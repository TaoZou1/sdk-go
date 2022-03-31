
#! /bin/bash
set -e

modulePath=${1:-}
echo "Module path: $modulePath"

mkdir /workspace/go-sdk-main
mkdir /workspace/go-sdk-tag
mkdir -p /workspace/results/$modulePath

# Clone repo
cd /workspace/go-sdk-main
git clone https://oauth2:${SDK_CLONE_ACCESS_TOKEN}@gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go.git
echo "Copy repo in tag dir..."
cp -R /workspace/go-sdk-main/vsphere-automation-sdk-go /workspace/go-sdk-tag/

# Checkout respective branch codes
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
git checkout aagrawal3/main/automate-sementic-versioning
git pull origin aagrawal3/main/automate-sementic-versioning
cd /workspace/go-sdk-tag/vsphere-automation-sdk-go/
git fetch --tags
latestTag=$(git describe --match "$modulePath*" --tags `git rev-list --tags --max-count=1`)
echo "Last tagged version for $modulePath was $latestTag"
git checkout $latestTag

# run diff check
cd /workspace/diff-checker
echo "Generating diff..."
go run cmd/module-diff-check.go generate-report --o /workspace/go-sdk-tag/vsphere-automation-sdk-go/$modulePath --n /workspace/go-sdk-main/vsphere-automation-sdk-go/$modulePath --result-dir /workspace/results/$modulePath --lang go

# find release type
apt-get install -y jq
RELEASE_TYPE=cat /workspace/results/$modulePath/go-mod-final-report.json | jq '.ReleaseType'
echo "Detected release type: $RELEASE_TYPE"

# Update version file in main branch...
cd /workspace/go-sdk-main/vsphere-automation-sdk-go/
lastReleaseVersion="${latestTag##*/}"
versionArray=($(echo $lastReleaseVersion | tr '.' "\n"))
nextRelease=lastReleaseVersion

if [[ $RELEASE_TYPE  ==  "MAJOR" ]]
then
  # remove 'v' from beginning
  majorVersion=${versionArray[1]:1}
  majorVersion=$(( majorVersion + 1 ))
  nextRelease="v$majorVersion.$versionArray[2].$versionArray[3]"
  echo "Next release version: $nextRelease"
elif [[ $RELEASE_TYPE  ==  "MINOR" ]]
then
  minorVersion=$versionArray[2]
  minorVersion=$(( minorVersion + 1 ))
  nextRelease="$versionArray[1].$minorVersion.$versionArray[3]"
  echo "Next release version: $nextRelease"
elif [[ $RELEASE_TYPE  ==  "PATCH" ]]
then
  patchVersion=$versionArray[3]
  patchVersion=$(( patchVersion + 1 ))
  nextRelease="$versionArray[1].$versionArray[2].$patchVersion"
  echo "Next release version: $nextRelease"
else
  echo "no change detected..."
fi

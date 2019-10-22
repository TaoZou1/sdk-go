/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package storage

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``VmHomePolicySpec`` class provides a specification for the storage policy to be associated with the virtual machine home's directory.
 type PolicyVmHomePolicySpec struct {
    // Policy type to be used while performing update operation on the virtual machine home's directory.
    Type_ PolicyVmHomePolicySpec_PolicyType
    // Storage Policy identification.
    Policy *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``PolicyType`` enumeration class defines the choices for how to specify the policy to be associated with the virtual machine home's directory.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type PolicyVmHomePolicySpec_PolicyType string

    const (
        // Use the specified policy (see PolicyVmHomePolicySpec#policy).
         PolicyVmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY PolicyVmHomePolicySpec_PolicyType = "USE_SPECIFIED_POLICY"
        // Use the default storage policy of the datastore.
         PolicyVmHomePolicySpec_PolicyType_USE_DEFAULT_POLICY PolicyVmHomePolicySpec_PolicyType = "USE_DEFAULT_POLICY"
    )

    func (p PolicyVmHomePolicySpec_PolicyType) PolicyVmHomePolicySpec_PolicyType() bool {
        switch p {
            case PolicyVmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY:
                return true
            case PolicyVmHomePolicySpec_PolicyType_USE_DEFAULT_POLICY:
                return true
            default:
                return false
        }
    }



// The ``DiskPolicySpec`` class provides a specification for the storage policy to be associated with the virtual disks.
 type PolicyDiskPolicySpec struct {
    // Policy type to be used while performing update operation on the virtual disks.
    Type_ PolicyDiskPolicySpec_PolicyType
    // Storage Policy identification.
    Policy *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``DiskPolicySpec`` enumeration class defines the choices for how to specify the policy to be associated with a virtual disk.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type PolicyDiskPolicySpec_PolicyType string

    const (
        // Use the specified policy (see PolicyDiskPolicySpec#policy).
         PolicyDiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY PolicyDiskPolicySpec_PolicyType = "USE_SPECIFIED_POLICY"
        // Use the default storage policy of the datastore.
         PolicyDiskPolicySpec_PolicyType_USE_DEFAULT_POLICY PolicyDiskPolicySpec_PolicyType = "USE_DEFAULT_POLICY"
    )

    func (p PolicyDiskPolicySpec_PolicyType) PolicyDiskPolicySpec_PolicyType() bool {
        switch p {
            case PolicyDiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY:
                return true
            case PolicyDiskPolicySpec_PolicyType_USE_DEFAULT_POLICY:
                return true
            default:
                return false
        }
    }



// The ``UpdateSpec`` class describes the updates to be made to the storage policies associated with the virtual machine home and/or its virtual disks.
 type PolicyUpdateSpec struct {
    // Storage policy to be used when reconfiguring the virtual machine home.
    VmHome *PolicyVmHomePolicySpec
    // Storage policy or policies to be used when reconfiguring virtual machine diks.
    Disks map[string]PolicyDiskPolicySpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains information about the storage policies associated with virtual machine's home directory and virtual hard disks.
 type PolicyInfo struct {
    // Storage Policy associated with virtual machine home.
    VmHome *string
    // Storage policies associated with virtual disks. The values in this map are storage policy identifiers. They will be identifiers for the resource type:com.vmware.vcenter.StoragePolicy If the map is empty, the virtual machine does not have any disks or its disks are not associated with a storage policy.
    Disks map[string]string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func policyUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(PolicyUpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func policyUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500})
}


func policyGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PolicyInfoBindingType)
}

func policyGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func PolicyVmHomePolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec.policy_type", reflect.TypeOf(PolicyVmHomePolicySpec_PolicyType(PolicyVmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_DEFAULT_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec",fields, reflect.TypeOf(PolicyVmHomePolicySpec{}), fieldNameMap, validators)
}

func PolicyDiskPolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec.policy_type", reflect.TypeOf(PolicyDiskPolicySpec_PolicyType(PolicyDiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_DEFAULT_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec",fields, reflect.TypeOf(PolicyDiskPolicySpec{}), fieldNameMap, validators)
}

func PolicyUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewReferenceType(PolicyVmHomePolicySpecBindingType))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(PolicyDiskPolicySpecBindingType),reflect.TypeOf(map[string]PolicyDiskPolicySpec{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.update_spec",fields, reflect.TypeOf(PolicyUpdateSpec{}), fieldNameMap, validators)
}

func PolicyInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.info",fields, reflect.TypeOf(PolicyInfo{}), fieldNameMap, validators)
}



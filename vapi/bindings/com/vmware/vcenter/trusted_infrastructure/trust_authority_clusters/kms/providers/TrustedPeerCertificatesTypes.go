/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TrustedPeerCertificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package providers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Info`` class contains x509 certificate list.
type TrustedPeerCertificatesInfo struct {
    // List of certificate strings, PEM format
	Certificates []string
}

// The ``UpdateSpec`` class contains properties that describe the server certificate update for a Key Provider.
type TrustedPeerCertificatesUpdateSpec struct {
    // Public certificates of key server to trust.
	Certificates []string
}


func trustedPeerCertificatesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewReferenceType(TrustedPeerCertificatesUpdateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedPeerCertificatesUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func trustedPeerCertificatesUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(TrustedPeerCertificatesUpdateSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"PATCH",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/peer-certs/trusted",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func trustedPeerCertificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedPeerCertificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TrustedPeerCertificatesInfoBindingType)
}

func trustedPeerCertificatesGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/peer-certs/trusted",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}


func TrustedPeerCertificatesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificates"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["certificates"] = "Certificates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.trusted_peer_certificates.info", fields, reflect.TypeOf(TrustedPeerCertificatesInfo{}), fieldNameMap, validators)
}

func TrustedPeerCertificatesUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["certificates"] = "Certificates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.trusted_peer_certificates.update_spec", fields, reflect.TypeOf(TrustedPeerCertificatesUpdateSpec{}), fieldNameMap, validators)
}


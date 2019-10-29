/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: NamespaceResourceOptions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespace_management

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains the information about the objects used to set and update resource quota keys on a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespaceResourceOptionsInfo struct {
    // Identifier of the class used to set resource quotas on the namespace. See null and null. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CreateResourceQuotaType string
    // Identifier of the class used to update resource quotas on the namespace. See null. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	UpdateResourceQuotaType string
}



func namespaceResourceOptionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceResourceOptionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NamespaceResourceOptionsInfoBindingType)
}

func namespaceResourceOptionsGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
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
		"/vcenter/namespace-management/clusters/{cluster}/workload-resource-options",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unsupported": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func NamespaceResourceOptionsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["create_resource_quota_type"] = bindings.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["create_resource_quota_type"] = "CreateResourceQuotaType"
	fields["update_resource_quota_type"] = bindings.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["update_resource_quota_type"] = "UpdateResourceQuotaType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.namespace_resource_options.info", fields, reflect.TypeOf(NamespaceResourceOptionsInfo{}), fieldNameMap, validators)
}


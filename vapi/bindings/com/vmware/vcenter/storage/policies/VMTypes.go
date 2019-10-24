/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package policies

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information about a virtual machine and its virtual disks that are associated with the given storage policy.
type VMInfo struct {
    // Flag to indicate whether or not the virtual machine home is associated with the given storage policy.
	VmHome bool
    // List of the virtual disks that are associated with the given storage policy.
	Disks []string
}



func vMListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"VirtualMachine"}, ""), bindings.NewReferenceType(VMInfoBindingType),reflect.TypeOf(map[string]VMInfo{}))
}

func vMListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"UnableToAllocateResource": 500,"Unauthenticated": 401,"ServiceUnavailable": 503,"Error": 500,"Unauthorized": 403})
}


func VMInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_home"] = bindings.NewBooleanType()
	fieldNameMap["vm_home"] = "VmHome"
	fields["disks"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["disks"] = "Disks"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.VM.info", fields, reflect.TypeOf(VMInfo{}), fieldNameMap, validators)
}


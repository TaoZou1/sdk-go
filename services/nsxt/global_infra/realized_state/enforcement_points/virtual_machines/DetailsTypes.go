/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Details.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package virtual_machines

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func detailsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_name"] = bindings.NewStringType()
	fields["virtual_machine_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_name"] = "EnforcementPointName"
	fieldNameMap["virtual_machine_id"] = "VirtualMachineId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detailsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VirtualMachineDetailsBindingType)
}

func detailsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["enforcement_point_name"] = bindings.NewStringType()
	fields["virtual_machine_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_name"] = "EnforcementPointName"
	fieldNameMap["virtual_machine_id"] = "VirtualMachineId"
	paramsTypeMap["enforcement_point_name"] = bindings.NewStringType()
	paramsTypeMap["virtual_machine_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointName"] = bindings.NewStringType()
	paramsTypeMap["virtualMachineId"] = bindings.NewStringType()
	pathParams["virtual_machine_id"] = "virtualMachineId"
	pathParams["enforcement_point_name"] = "enforcementPointName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/global-infra/realized-state/enforcement-points/{enforcementPointName}/virtual-machines/{virtualMachineId}/details",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



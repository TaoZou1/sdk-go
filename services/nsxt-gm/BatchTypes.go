/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Batch.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx_global_policy

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func batchCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["batch_request"] = bindings.NewReferenceType(model.BatchRequestBindingType)
	fields["atomic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["batch_request"] = "BatchRequest"
	fieldNameMap["atomic"] = "Atomic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BatchResponseBindingType)
}

func batchCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["batch_request"] = bindings.NewReferenceType(model.BatchRequestBindingType)
	fields["atomic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["batch_request"] = "BatchRequest"
	fieldNameMap["atomic"] = "Atomic"
	paramsTypeMap["batch_request"] = bindings.NewReferenceType(model.BatchRequestBindingType)
	paramsTypeMap["atomic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["atomic"] = "atomic"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"batch_request",
		"POST",
		"/global-manager/api/v1/batch",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}



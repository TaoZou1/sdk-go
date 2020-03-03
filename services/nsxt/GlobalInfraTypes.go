/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: GlobalInfra.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx_policy

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func globalInfraGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalInfraGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InfraBindingType)
}

func globalInfraGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["filter"] = "filter"
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
		"/policy/api/v1/global-infra",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func globalInfraPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fields["enforce_revision_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["infra"] = "Infra"
	fieldNameMap["enforce_revision_check"] = "EnforceRevisionCheck"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalInfraPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func globalInfraPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fields["enforce_revision_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["infra"] = "Infra"
	fieldNameMap["enforce_revision_check"] = "EnforceRevisionCheck"
	paramsTypeMap["enforce_revision_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	queryParams["enforce_revision_check"] = "enforce_revision_check"
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
		"infra",
		"PATCH",
		"/policy/api/v1/global-infra",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func globalInfraUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fieldNameMap["infra"] = "Infra"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalInfraUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InfraBindingType)
}

func globalInfraUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fieldNameMap["infra"] = "Infra"
	paramsTypeMap["infra"] = bindings.NewReferenceType(model.InfraBindingType)
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
		"infra",
		"PUT",
		"/policy/api/v1/global-infra",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



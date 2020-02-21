/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ServiceReferences.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package global_infra

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func serviceReferencesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceReferencesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceReferencesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	queryParams["cascade"] = "cascade"
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
		"DELETE",
		"/policy/api/v1/global-infra/service-references/{serviceReferenceId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func serviceReferencesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceReferencesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceReferenceBindingType)
}

func serviceReferencesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
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
		"/policy/api/v1/global-infra/service-references/{serviceReferenceId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func serviceReferencesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceReferencesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceReferenceListResultBindingType)
}

func serviceReferencesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/global-infra/service-references",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func serviceReferencesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_reference"] = "ServiceReference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceReferencesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceReferencesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_reference"] = "ServiceReference"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
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
		"service_reference",
		"PATCH",
		"/policy/api/v1/global-infra/service-references/{serviceReferenceId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func serviceReferencesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_reference"] = "ServiceReference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceReferencesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceReferenceBindingType)
}

func serviceReferencesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_reference"] = "ServiceReference"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["service_reference"] = bindings.NewReferenceType(model.ServiceReferenceBindingType)
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
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
		"service_reference",
		"PUT",
		"/policy/api/v1/global-infra/service-references/{serviceReferenceId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



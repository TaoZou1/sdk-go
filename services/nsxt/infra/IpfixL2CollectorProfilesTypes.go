/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: IpfixL2CollectorProfiles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package infra

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func ipfixL2CollectorProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixL2CollectorProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipfixL2CollectorProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	paramsTypeMap["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixL2CollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_l2_collector_profile_id"] = "ipfixL2CollectorProfileId"
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
		"/policy/api/v1/infra/ipfix-l2-collector-profiles/{ipfixL2CollectorProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixL2CollectorProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixL2CollectorProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
}

func ipfixL2CollectorProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	paramsTypeMap["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixL2CollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_l2_collector_profile_id"] = "ipfixL2CollectorProfileId"
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
		"/policy/api/v1/infra/ipfix-l2-collector-profiles/{ipfixL2CollectorProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixL2CollectorProfilesListInputType() bindings.StructType {
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

func ipfixL2CollectorProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXL2CollectorProfileListResultBindingType)
}

func ipfixL2CollectorProfilesListRestMetadata() protocol.OperationRestMetadata {
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
		"/policy/api/v1/infra/ipfix-l2-collector-profiles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixL2CollectorProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fields["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	fieldNameMap["i_PFIX_l2_collector_profile"] = "IPFIXL2CollectorProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixL2CollectorProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipfixL2CollectorProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fields["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	fieldNameMap["i_PFIX_l2_collector_profile"] = "IPFIXL2CollectorProfile"
	paramsTypeMap["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	paramsTypeMap["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixL2CollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_l2_collector_profile_id"] = "ipfixL2CollectorProfileId"
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
		"i_PFIX_l2_collector_profile",
		"PATCH",
		"/policy/api/v1/infra/ipfix-l2-collector-profiles/{ipfixL2CollectorProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixL2CollectorProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fields["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	fieldNameMap["i_PFIX_l2_collector_profile"] = "IPFIXL2CollectorProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixL2CollectorProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
}

func ipfixL2CollectorProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	fields["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	fieldNameMap["ipfix_l2_collector_profile_id"] = "IpfixL2CollectorProfileId"
	fieldNameMap["i_PFIX_l2_collector_profile"] = "IPFIXL2CollectorProfile"
	paramsTypeMap["i_PFIX_l2_collector_profile"] = bindings.NewReferenceType(model.IPFIXL2CollectorProfileBindingType)
	paramsTypeMap["ipfix_l2_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixL2CollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_l2_collector_profile_id"] = "ipfixL2CollectorProfileId"
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
		"i_PFIX_l2_collector_profile",
		"PUT",
		"/policy/api/v1/infra/ipfix-l2-collector-profiles/{ipfixL2CollectorProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



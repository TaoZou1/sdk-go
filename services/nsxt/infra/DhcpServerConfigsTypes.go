/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: DhcpServerConfigs.
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





func dhcpServerConfigsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dhcpServerConfigsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func dhcpServerConfigsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	paramsTypeMap["dhcp_server_config_id"] = bindings.NewStringType()
	paramsTypeMap["dhcpServerConfigId"] = bindings.NewStringType()
	pathParams["dhcp_server_config_id"] = "dhcpServerConfigId"
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
		"/policy/api/v1/infra/dhcp-server-configs/{dhcpServerConfigId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func dhcpServerConfigsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dhcpServerConfigsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpServerConfigBindingType)
}

func dhcpServerConfigsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	paramsTypeMap["dhcp_server_config_id"] = bindings.NewStringType()
	paramsTypeMap["dhcpServerConfigId"] = bindings.NewStringType()
	pathParams["dhcp_server_config_id"] = "dhcpServerConfigId"
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
		"/policy/api/v1/infra/dhcp-server-configs/{dhcpServerConfigId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func dhcpServerConfigsListInputType() bindings.StructType {
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

func dhcpServerConfigsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpServerConfigListResultBindingType)
}

func dhcpServerConfigsListRestMetadata() protocol.OperationRestMetadata {
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
		"/policy/api/v1/infra/dhcp-server-configs",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func dhcpServerConfigsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fields["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	fieldNameMap["dhcp_server_config"] = "DhcpServerConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dhcpServerConfigsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func dhcpServerConfigsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fields["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	fieldNameMap["dhcp_server_config"] = "DhcpServerConfig"
	paramsTypeMap["dhcp_server_config_id"] = bindings.NewStringType()
	paramsTypeMap["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	paramsTypeMap["dhcpServerConfigId"] = bindings.NewStringType()
	pathParams["dhcp_server_config_id"] = "dhcpServerConfigId"
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
		"dhcp_server_config",
		"PATCH",
		"/policy/api/v1/infra/dhcp-server-configs/{dhcpServerConfigId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func dhcpServerConfigsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fields["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	fieldNameMap["dhcp_server_config"] = "DhcpServerConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dhcpServerConfigsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpServerConfigBindingType)
}

func dhcpServerConfigsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["dhcp_server_config_id"] = bindings.NewStringType()
	fields["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	fieldNameMap["dhcp_server_config_id"] = "DhcpServerConfigId"
	fieldNameMap["dhcp_server_config"] = "DhcpServerConfig"
	paramsTypeMap["dhcp_server_config_id"] = bindings.NewStringType()
	paramsTypeMap["dhcp_server_config"] = bindings.NewReferenceType(model.DhcpServerConfigBindingType)
	paramsTypeMap["dhcpServerConfigId"] = bindings.NewStringType()
	pathParams["dhcp_server_config_id"] = "dhcpServerConfigId"
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
		"dhcp_server_config",
		"PUT",
		"/policy/api/v1/infra/dhcp-server-configs/{dhcpServerConfigId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



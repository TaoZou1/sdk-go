// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Ipv6StaticBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package servers

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func ipv6StaticBindingsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6StaticBindingsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"dhcp_v6_static_binding",
		"POST",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6StaticBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipv6StaticBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"",
		"DELETE",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6StaticBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"",
		"GET",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6StaticBindingsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpV6StaticBindingListResultBindingType)
}

func ipv6StaticBindingsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"",
		"GET",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6StaticBindingsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["dhcp_v6_static_binding"] = bindings.NewReferenceType(model.DhcpV6StaticBindingBindingType)
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"dhcp_v6_static_binding",
		"PUT",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

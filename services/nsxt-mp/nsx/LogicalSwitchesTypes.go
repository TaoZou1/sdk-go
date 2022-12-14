// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LogicalSwitches.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_DEFAULT = "DEFAULT"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_SERVICE_PLANE = "SERVICE_PLANE"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_DHCP_RELAY = "DHCP_RELAY"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_GLOBAL = "GLOBAL"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_INTER_ROUTER = "INTER_ROUTER"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_EVPN = "EVPN"

// Possible value for ``switchType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_SWITCH_TYPE_DVPG = "DVPG"

// Possible value for ``transportType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_TRANSPORT_TYPE_OVERLAY = "OVERLAY"

// Possible value for ``transportType`` of method LogicalSwitches#list.
const LogicalSwitches_LIST_TRANSPORT_TYPE_VLAN = "VLAN"

func logicalSwitchesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalSwitchesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func logicalSwitchesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	paramsTypeMap["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
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
		"logical_switch",
		"POST",
		"/api/v1/logical-switches",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalSwitchesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cascade"] = "Cascade"
	fieldNameMap["detach"] = "Detach"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalSwitchesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func logicalSwitchesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cascade"] = "Cascade"
	fieldNameMap["detach"] = "Detach"
	paramsTypeMap["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["cascade"] = "cascade"
	queryParams["detach"] = "detach"
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
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalSwitchesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalSwitchesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func logicalSwitchesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalSwitchesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switch_type"] = "SwitchType"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalSwitchesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchListResultBindingType)
}

func logicalSwitchesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switch_type"] = "SwitchType"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	paramsTypeMap["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["switching_profile_id"] = "switching_profile_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["switch_type"] = "switch_type"
	queryParams["uplink_teaming_policy_name"] = "uplink_teaming_policy_name"
	queryParams["vni"] = "vni"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["vlan"] = "vlan"
	queryParams["included_fields"] = "included_fields"
	queryParams["transport_type"] = "transport_type"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["diagnostic"] = "diagnostic"
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
		"/api/v1/logical-switches",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalSwitchesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalSwitchesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func logicalSwitchesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"logical_switch",
		"PUT",
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

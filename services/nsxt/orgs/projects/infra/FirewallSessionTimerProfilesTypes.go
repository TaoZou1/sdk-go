// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FirewallSessionTimerProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func firewallSessionTimerProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallSessionTimerProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["firewall_session_timer_profile_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileId"] = bindings.NewStringType()
	pathParams["firewall_session_timer_profile_id"] = "firewallSessionTimerProfileId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["override"] = "override"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/firewall-session-timer-profiles/{firewallSessionTimerProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
}

func firewallSessionTimerProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_session_timer_profile_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileId"] = bindings.NewStringType()
	pathParams["firewall_session_timer_profile_id"] = "firewallSessionTimerProfileId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/firewall-session-timer-profiles/{firewallSessionTimerProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileListResultBindingType)
}

func firewallSessionTimerProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/firewall-session-timer-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["policy_firewall_session_timer_profile"] = "PolicyFirewallSessionTimerProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallSessionTimerProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["policy_firewall_session_timer_profile"] = "PolicyFirewallSessionTimerProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["firewall_session_timer_profile_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileId"] = bindings.NewStringType()
	pathParams["firewall_session_timer_profile_id"] = "firewallSessionTimerProfileId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["override"] = "override"
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
		"policy_firewall_session_timer_profile",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/firewall-session-timer-profiles/{firewallSessionTimerProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["policy_firewall_session_timer_profile"] = "PolicyFirewallSessionTimerProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
}

func firewallSessionTimerProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_session_timer_profile_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["firewall_session_timer_profile_id"] = "FirewallSessionTimerProfileId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["policy_firewall_session_timer_profile"] = "PolicyFirewallSessionTimerProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["firewall_session_timer_profile_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_session_timer_profile"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileId"] = bindings.NewStringType()
	pathParams["firewall_session_timer_profile_id"] = "firewallSessionTimerProfileId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["override"] = "override"
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
		"policy_firewall_session_timer_profile",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/firewall-session-timer-profiles/{firewallSessionTimerProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

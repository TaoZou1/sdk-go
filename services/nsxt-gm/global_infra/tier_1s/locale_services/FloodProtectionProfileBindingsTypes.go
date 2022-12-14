// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FloodProtectionProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package locale_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func floodProtectionProfileBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floodProtectionProfileBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floodProtectionProfileBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["locale_services_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServicesId"] = bindings.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = bindings.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floodProtectionProfileBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
}

func floodProtectionProfileBindingsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["locale_services_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServicesId"] = bindings.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = bindings.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floodProtectionProfileBindingsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floodProtectionProfileBindingsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	paramsTypeMap["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["locale_services_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServicesId"] = bindings.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = bindings.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"flood_protection_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floodProtectionProfileBindingsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
}

func floodProtectionProfileBindingsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_services_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_id"] = bindings.NewStringType()
	fields["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	paramsTypeMap["flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.FloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["locale_services_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServicesId"] = bindings.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = bindings.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"flood_protection_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

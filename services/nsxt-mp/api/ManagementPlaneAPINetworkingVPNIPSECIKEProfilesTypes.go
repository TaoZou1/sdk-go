// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingVPNIPSECIKEProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func managementPlaneAPINetworkingVPNIPSECIKEProfilesCreateIPSecVPNIKEProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
	fieldNameMap["ip_sec_VPNIKE_profile"] = "IpSecVPNIKEProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesCreateIPSecVPNIKEProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesCreateIPSecVPNIKEProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
	fieldNameMap["ip_sec_VPNIKE_profile"] = "IpSecVPNIKEProfile"
	paramsTypeMap["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
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
		"ip_sec_VPNIKE_profile",
		"POST",
		"/api/v1/vpn/ipsec/ike-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesDeleteIPSecVPNIKEProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesDeleteIPSecVPNIKEProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesDeleteIPSecVPNIKEProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnIkeProfileId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_ike_profile_id"] = "ipsecVpnIkeProfileId"
	queryParams["force"] = "force"
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
		"/api/v1/vpn/ipsec/ike-profiles/{ipsecVpnIkeProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesGetIPSecVPNIKEProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesGetIPSecVPNIKEProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesGetIPSecVPNIKEProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	paramsTypeMap["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnIkeProfileId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_ike_profile_id"] = "ipsecVpnIkeProfileId"
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
		"/api/v1/vpn/ipsec/ike-profiles/{ipsecVpnIkeProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesListIPSecVPNIKEProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesListIPSecVPNIKEProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNIKEProfileListResultBindingType)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesListIPSecVPNIKEProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		"/api/v1/vpn/ipsec/ike-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesUpdateIPSecVPNIKEProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fields["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	fieldNameMap["ip_sec_VPNIKE_profile"] = "IpSecVPNIKEProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesUpdateIPSecVPNIKEProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
}

func managementPlaneAPINetworkingVPNIPSECIKEProfilesUpdateIPSecVPNIKEProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	fields["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
	fieldNameMap["ipsec_vpn_ike_profile_id"] = "IpsecVpnIkeProfileId"
	fieldNameMap["ip_sec_VPNIKE_profile"] = "IpSecVPNIKEProfile"
	paramsTypeMap["ip_sec_VPNIKE_profile"] = bindings.NewReferenceType(model.IPSecVPNIKEProfileBindingType)
	paramsTypeMap["ipsec_vpn_ike_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnIkeProfileId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_ike_profile_id"] = "ipsecVpnIkeProfileId"
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
		"ip_sec_VPNIKE_profile",
		"PUT",
		"/api/v1/vpn/ipsec/ike-profiles/{ipsecVpnIkeProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

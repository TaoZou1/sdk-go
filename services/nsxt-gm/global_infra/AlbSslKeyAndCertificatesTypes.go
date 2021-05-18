// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AlbSslKeyAndCertificates.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package global_infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func albSslKeyAndCertificatesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func albSslKeyAndCertificatesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func albSslKeyAndCertificatesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["albSslkeyandcertificateId"] = bindings.NewStringType()
	pathParams["alb_sslkeyandcertificate_id"] = "albSslkeyandcertificateId"
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
		"/global-manager/api/v1/global-infra/alb-ssl-key-and-certificates/{albSslkeyandcertificateId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func albSslKeyAndCertificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func albSslKeyAndCertificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
}

func albSslKeyAndCertificatesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	paramsTypeMap["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	paramsTypeMap["albSslkeyandcertificateId"] = bindings.NewStringType()
	pathParams["alb_sslkeyandcertificate_id"] = "albSslkeyandcertificateId"
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
		"/global-manager/api/v1/global-infra/alb-ssl-key-and-certificates/{albSslkeyandcertificateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func albSslKeyAndCertificatesListInputType() bindings.StructType {
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

func albSslKeyAndCertificatesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBSSLKeyAndCertificateApiResponseBindingType)
}

func albSslKeyAndCertificatesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/global-manager/api/v1/global-infra/alb-ssl-key-and-certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func albSslKeyAndCertificatesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["a_LBSSL_key_and_certificate"] = "ALBSSLKeyAndCertificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func albSslKeyAndCertificatesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func albSslKeyAndCertificatesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["a_LBSSL_key_and_certificate"] = "ALBSSLKeyAndCertificate"
	paramsTypeMap["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	paramsTypeMap["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	paramsTypeMap["albSslkeyandcertificateId"] = bindings.NewStringType()
	pathParams["alb_sslkeyandcertificate_id"] = "albSslkeyandcertificateId"
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
		"a_LBSSL_key_and_certificate",
		"PATCH",
		"/global-manager/api/v1/global-infra/alb-ssl-key-and-certificates/{albSslkeyandcertificateId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func albSslKeyAndCertificatesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["a_LBSSL_key_and_certificate"] = "ALBSSLKeyAndCertificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func albSslKeyAndCertificatesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
}

func albSslKeyAndCertificatesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	fields["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	fieldNameMap["alb_sslkeyandcertificate_id"] = "AlbSslkeyandcertificateId"
	fieldNameMap["a_LBSSL_key_and_certificate"] = "ALBSSLKeyAndCertificate"
	paramsTypeMap["alb_sslkeyandcertificate_id"] = bindings.NewStringType()
	paramsTypeMap["a_LBSSL_key_and_certificate"] = bindings.NewReferenceType(model.ALBSSLKeyAndCertificateBindingType)
	paramsTypeMap["albSslkeyandcertificateId"] = bindings.NewStringType()
	pathParams["alb_sslkeyandcertificate_id"] = "albSslkeyandcertificateId"
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
		"a_LBSSL_key_and_certificate",
		"PUT",
		"/global-manager/api/v1/global-infra/alb-ssl-key-and-certificates/{albSslkeyandcertificateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

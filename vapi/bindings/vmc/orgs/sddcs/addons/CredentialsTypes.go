/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Credentials.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package addons

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Possible value for ``addonType`` of method Credentials#create.
const Credentials_CREATE_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#get.
const Credentials_GET_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#list.
const Credentials_LIST_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#update.
const Credentials_UPDATE_ADDON_TYPE_HCX = "HCX"




func credentialsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc_id"] = bindings.NewStringType()
	fields["addon_type"] = bindings.NewStringType()
	fields["credentials"] = bindings.NewReferenceType(model.NewCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["credentials"] = "Credentials"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["addon_type"] = bindings.NewStringType()
	paramsTypeMap["credentials"] = bindings.NewReferenceType(model.NewCredentialsBindingType)
	paramsTypeMap["sddc_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddcId"] = bindings.NewStringType()
	paramsTypeMap["addonType"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"credentials",
		"POST",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"ConcurrentChange": 409,"InvalidRequest": 400,"Unauthorized": 403})
}

func credentialsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc_id"] = bindings.NewStringType()
	fields["addon_type"] = bindings.NewStringType()
	fields["name"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["addon_type"] = bindings.NewStringType()
	paramsTypeMap["name"] = bindings.NewStringType()
	paramsTypeMap["sddc_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddcId"] = bindings.NewStringType()
	paramsTypeMap["addonType"] = bindings.NewStringType()
	paramsTypeMap["name"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthorized": 403})
}

func credentialsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc_id"] = bindings.NewStringType()
	fields["addon_type"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(model.NewCredentialsBindingType), reflect.TypeOf([]model.NewCredentials{}))
}

func credentialsListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["addon_type"] = bindings.NewStringType()
	paramsTypeMap["sddc_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddcId"] = bindings.NewStringType()
	paramsTypeMap["addonType"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthorized": 403})
}

func credentialsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc_id"] = bindings.NewStringType()
	fields["addon_type"] = bindings.NewStringType()
	fields["name"] = bindings.NewStringType()
	fields["credentials"] = bindings.NewReferenceType(model.UpdateCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	fieldNameMap["credentials"] = "Credentials"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["addon_type"] = bindings.NewStringType()
	paramsTypeMap["name"] = bindings.NewStringType()
	paramsTypeMap["credentials"] = bindings.NewReferenceType(model.UpdateCredentialsBindingType)
	paramsTypeMap["sddc_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddcId"] = bindings.NewStringType()
	paramsTypeMap["addonType"] = bindings.NewStringType()
	paramsTypeMap["name"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"credentials",
		"PUT",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403})
}



/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Search.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ldap_identity_sources

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func searchCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fields["filter_value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	fieldNameMap["filter_value"] = "FilterValue"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func searchCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LdapIdentitySourceSearchResultListBindingType)
}

func searchCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fields["filter_value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	fieldNameMap["filter_value"] = "FilterValue"
	paramsTypeMap["ldap_identity_source_id"] = bindings.NewStringType()
	paramsTypeMap["filter_value"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ldapIdentitySourceId"] = bindings.NewStringType()
	pathParams["ldap_identity_source_id"] = "ldapIdentitySourceId"
	queryParams["filter_value"] = "filter_value"
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
		"POST",
		"/policy/api/v1/aaa/ldap-identity-sources/{ldapIdentitySourceId}/search",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPIServicesMetadataProxy.
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

// Possible value for ``source`` of method ManagementPlaneAPIServicesMetadataProxy#getMetadataProxyStatistics.
const ManagementPlaneAPIServicesMetadataProxy_GET_METADATA_PROXY_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPIServicesMetadataProxy#getMetadataProxyStatistics.
const ManagementPlaneAPIServicesMetadataProxy_GET_METADATA_PROXY_STATISTICS_SOURCE_CACHED = "cached"

func managementPlaneAPIServicesMetadataProxyGetMetadataProxyStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["proxy_id"] = bindings.NewStringType()
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["proxy_id"] = "ProxyId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPIServicesMetadataProxyGetMetadataProxyStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MetadataProxyStatisticsBindingType)
}

func managementPlaneAPIServicesMetadataProxyGetMetadataProxyStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["proxy_id"] = bindings.NewStringType()
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["proxy_id"] = "ProxyId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["proxy_id"] = bindings.NewStringType()
	paramsTypeMap["proxyId"] = bindings.NewStringType()
	pathParams["proxy_id"] = "proxyId"
	queryParams["source"] = "source"
	queryParams["logical_switch_id"] = "logical_switch_id"
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
		"/api/v1/md-proxies/{proxyId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

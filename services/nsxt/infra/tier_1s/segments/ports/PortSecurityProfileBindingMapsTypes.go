/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: PortSecurityProfileBindingMaps.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ports

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func portSecurityProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portSecurityProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portSecurityProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portSecurityProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portSecurityProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
}

func portSecurityProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portSecurityProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portSecurityProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortSecurityProfileBindingMapListResultBindingType)
}

func portSecurityProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portSecurityProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portSecurityProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portSecurityProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	paramsTypeMap["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
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
		"port_security_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portSecurityProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portSecurityProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
}

func portSecurityProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	paramsTypeMap["port_security_profile_binding_map"] = bindings.NewReferenceType(model.PortSecurityProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
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
		"port_security_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



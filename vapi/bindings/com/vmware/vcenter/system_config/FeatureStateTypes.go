/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: FeatureState.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package system_config

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// The ``Status`` enumeration class defines the possible states of a feature state switch.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FeatureStateStatus string

const (
    // Feature state switch is enabled.
	FeatureStateStatus_ENABLED FeatureStateStatus = "ENABLED"
    // Feature state switch is disabled.
	FeatureStateStatus_DISABLED FeatureStateStatus = "DISABLED"
)

func (s FeatureStateStatus) FeatureStateStatus() bool {
	switch s {
	case FeatureStateStatus_ENABLED:
		return true
	case FeatureStateStatus_DISABLED:
		return true
	default:
		return false
	}
}


func featureStateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["features"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["features"] = "Features"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func featureStateGetOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewStringType(), bindings.NewEnumType("com.vmware.vcenter.system_config.feature_state.status", reflect.TypeOf(FeatureStateStatus(FeatureStateStatus_ENABLED))),reflect.TypeOf(map[string]FeatureStateStatus{}))
}

func featureStateGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}



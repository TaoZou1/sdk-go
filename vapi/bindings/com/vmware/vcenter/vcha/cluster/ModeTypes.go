/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Mode.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package cluster

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// The ``ClusterMode`` enumeration class defines the possible modes for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ModeClusterMode string

const (
    // VCHA Cluster is enabled. State replication between the Active and Passive node is enabled and automatic failover is allowed.
	ModeClusterMode_ENABLED ModeClusterMode = "ENABLED"
    // VCHA Cluster is disabled. State replication between the Active and Passive node is disabled and automatic failover is not allowed.
	ModeClusterMode_DISABLED ModeClusterMode = "DISABLED"
    // VCHA Cluster is in maintenance mode. State replication between the and Passive node is enabled but automatic failover is not allowed.
	ModeClusterMode_MAINTENANCE ModeClusterMode = "MAINTENANCE"
)

func (c ModeClusterMode) ModeClusterMode() bool {
	switch c {
	case ModeClusterMode_ENABLED:
		return true
	case ModeClusterMode_DISABLED:
		return true
	case ModeClusterMode_MAINTENANCE:
		return true
	default:
		return false
	}
}

// The ``Info`` class contains the mode of the VCHA Cluster.
type ModeInfo struct {
    // Identifies the mode of the VCHA cluster.
	Mode ModeClusterMode
}


func modeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ModeInfoBindingType)
}

func modeGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"NotAllowedInCurrentState": 400,"Unauthorized": 403,"Error": 500})
}

func modeSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.mode.cluster_mode", reflect.TypeOf(ModeClusterMode(ModeClusterMode_ENABLED)))
	fieldNameMap["mode"] = "Mode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modeSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func modeSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}


func ModeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.mode.cluster_mode", reflect.TypeOf(ModeClusterMode(ModeClusterMode_ENABLED)))
	fieldNameMap["mode"] = "Mode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.mode.info", fields, reflect.TypeOf(ModeInfo{}), fieldNameMap, validators)
}


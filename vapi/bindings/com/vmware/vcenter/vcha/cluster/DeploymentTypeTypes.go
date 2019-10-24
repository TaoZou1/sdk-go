/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: DeploymentType.
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


// The ``Type`` enumeration class defines the possible deployment types for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DeploymentTypeType string

const (
    // VCHA Cluster is not configured.
	DeploymentTypeType_NONE DeploymentTypeType = "NONE"
    // VCHA Cluster was deployed automatically.
	DeploymentTypeType_AUTO DeploymentTypeType = "AUTO"
    // VCHA Cluster was deployed manually.
	DeploymentTypeType_MANUAL DeploymentTypeType = "MANUAL"
)

func (t DeploymentTypeType) DeploymentTypeType() bool {
	switch t {
	case DeploymentTypeType_NONE:
		return true
	case DeploymentTypeType_AUTO:
		return true
	case DeploymentTypeType_MANUAL:
		return true
	default:
		return false
	}
}

// The ``Info`` class contains the deployment type of the VCHA Cluster.
type DeploymentTypeInfo struct {
    // Identifies the deployment type of the VCHA cluster.
	DeploymentType DeploymentTypeType
}


func deploymentTypeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentTypeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DeploymentTypeInfoBindingType)
}

func deploymentTypeGetRestMetadata() protocol.OperationRestMetadata {
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


func DeploymentTypeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deployment_type"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.deployment_type.type", reflect.TypeOf(DeploymentTypeType(DeploymentTypeType_NONE)))
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.deployment_type.info", fields, reflect.TypeOf(DeploymentTypeInfo{}), fieldNameMap, validators)
}


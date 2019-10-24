/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Active.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package cluster

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vcha"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Info`` class contains the network and placement information of the active node of a VCHA Cluster.
type ActiveInfo struct {
    // IP specification for the Management network.
	Management vcha.IpSpec
    // IP specification for the HA network.
	Ha *vcha.IpSpec
    // Contains the placement information of the active node.
	Placement *vcha.PlacementInfo
}


func activeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
	fields["partial"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["vc_spec"] = "VcSpec"
	fieldNameMap["partial"] = "Partial"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func activeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ActiveInfoBindingType)
}

func activeGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"InvalidElementConfiguration": 400,"NotFound": 404,"Error": 500})
}


func ActiveInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management"] = bindings.NewReferenceType(vcha.IpSpecBindingType)
	fieldNameMap["management"] = "Management"
	fields["ha"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
	fieldNameMap["ha"] = "Ha"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.PlacementInfoBindingType))
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.active.info", fields, reflect.TypeOf(ActiveInfo{}), fieldNameMap, validators)
}


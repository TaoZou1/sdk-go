/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Status.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package attestation

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information that describes the status of the service.
type StatusInfo struct {
    // The service health status.
	Health StatusInfoHealth
    // Details regarding the health of the service. 
    //
    //  When the service ``Health`` is not StatusInfoHealth#Health_OK or StatusInfoHealth#Health_NONE, this member will provide an actionable description of the issues present.
	Details []std.LocalizableMessage
}

// The ``Health`` enumeration class defines the possible service health states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type StatusInfoHealth string

const (
    // No status available.
	StatusInfoHealth_NONE StatusInfoHealth = "NONE"
    // Attestation is functioning normally.
	StatusInfoHealth_OK StatusInfoHealth = "OK"
    // Attestation is functioning, however there is an issue that requires attention.
	StatusInfoHealth_WARNING StatusInfoHealth = "WARNING"
    // Attestation is not functioning.
	StatusInfoHealth_ERROR StatusInfoHealth = "ERROR"
)

func (h StatusInfoHealth) StatusInfoHealth() bool {
	switch h {
	case StatusInfoHealth_NONE:
		return true
	case StatusInfoHealth_OK:
		return true
	case StatusInfoHealth_WARNING:
		return true
	case StatusInfoHealth_ERROR:
		return true
	default:
		return false
	}
}




func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(StatusInfoBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
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
		"GET",
		"/esx/attestation/status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func StatusInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.esx.attestation.status.info.health", reflect.TypeOf(StatusInfoHealth(StatusInfoHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.status.info", fields, reflect.TypeOf(StatusInfo{}), fieldNameMap, validators)
}


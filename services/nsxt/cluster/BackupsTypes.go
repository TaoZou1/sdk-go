/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Backups.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func backupsRetrievesshfingerprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	fieldNameMap["remote_server_fingerprint_request"] = "RemoteServerFingerprintRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func backupsRetrievesshfingerprintOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RemoteServerFingerprintBindingType)
}

func backupsRetrievesshfingerprintRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	fieldNameMap["remote_server_fingerprint_request"] = "RemoteServerFingerprintRequest"
	paramsTypeMap["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=retrieve_ssh_fingerprint",
		"remote_server_fingerprint_request",
		"POST",
		"/policy/api/v1/cluster/backups",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



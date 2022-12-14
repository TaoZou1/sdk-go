// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PrecheckByClusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nvds_urt

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func precheckByClustersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
	fields["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func precheckByClustersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func precheckByClustersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
	fields["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	paramsTypeMap["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
	queryParams["tolerate_different_configurations"] = "tolerate_different_configurations"
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
		"precheck_parameters",
		"PUT",
		"/policy/api/v1/infra/nvds-urt/precheck-by-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

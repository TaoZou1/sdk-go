/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Statistics.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nat_rules

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func statisticsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["nat_id"] = bindings.NewStringType()
	fields["nat_rule_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statisticsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyNatRuleStatisticsListResultBindingType)
}

func statisticsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["nat_id"] = bindings.NewStringType()
	fields["nat_rule_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["nat_rule_id"] = bindings.NewStringType()
	paramsTypeMap["nat_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["natId"] = bindings.NewStringType()
	paramsTypeMap["natRuleId"] = bindings.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tier1_id"] = "tier1Id"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/global-infra/tier-1s/{tier1Id}/nat/{natId}/nat-rules/{natRuleId}/statistics",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



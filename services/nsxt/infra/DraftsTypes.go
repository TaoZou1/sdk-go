/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Drafts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package infra

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func draftsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fieldNameMap["draft_id"] = "DraftId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func draftsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fieldNameMap["draft_id"] = "DraftId"
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
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
		"/policy/api/v1/infra/drafts/{draftId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func draftsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fieldNameMap["draft_id"] = "DraftId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyDraftBindingType)
}

func draftsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fieldNameMap["draft_id"] = "DraftId"
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
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
		"/policy/api/v1/infra/drafts/{draftId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func draftsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["auto_drafts"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["auto_drafts"] = "AutoDrafts"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyDraftListResultBindingType)
}

func draftsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["auto_drafts"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["auto_drafts"] = "AutoDrafts"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["auto_drafts"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["auto_drafts"] = "auto_drafts"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/infra/drafts",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func draftsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fields["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["policy_draft"] = "PolicyDraft"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func draftsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fields["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["policy_draft"] = "PolicyDraft"
	paramsTypeMap["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
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
		"policy_draft",
		"PATCH",
		"/policy/api/v1/infra/drafts/{draftId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func draftsPublishInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["infra"] = "Infra"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsPublishOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func draftsPublishRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fields["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["infra"] = "Infra"
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["infra"] = bindings.NewReferenceType(model.InfraBindingType)
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=publish",
		"infra",
		"POST",
		"/policy/api/v1/infra/drafts/{draftId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func draftsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fields["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["policy_draft"] = "PolicyDraft"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func draftsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyDraftBindingType)
}

func draftsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fields["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["policy_draft"] = "PolicyDraft"
	paramsTypeMap["policy_draft"] = bindings.NewReferenceType(model.PolicyDraftBindingType)
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
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
		"policy_draft",
		"PUT",
		"/policy/api/v1/infra/drafts/{draftId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}



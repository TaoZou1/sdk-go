// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationLifecycleManagementMigrationVmgroup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func systemAdministrationLifecycleManagementMigrationVmgroupPostVmGroupMigratePostMigrateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["post_vm_group_migration_spec"] = bindings.NewReferenceType(model.PostVmGroupMigrationSpecBindingType)
	fieldNameMap["post_vm_group_migration_spec"] = "PostVmGroupMigrationSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementMigrationVmgroupPostVmGroupMigratePostMigrateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementMigrationVmgroupPostVmGroupMigratePostMigrateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["post_vm_group_migration_spec"] = bindings.NewReferenceType(model.PostVmGroupMigrationSpecBindingType)
	fieldNameMap["post_vm_group_migration_spec"] = "PostVmGroupMigrationSpec"
	paramsTypeMap["post_vm_group_migration_spec"] = bindings.NewReferenceType(model.PostVmGroupMigrationSpecBindingType)
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
		"action=post_migrate",
		"post_vm_group_migration_spec",
		"POST",
		"/api/v1/migration/vmgroup",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementMigrationVmgroupPreVmGroupMigratePreMigrateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pre_vm_group_migration_spec"] = bindings.NewReferenceType(model.PreVmGroupMigrationSpecBindingType)
	fieldNameMap["pre_vm_group_migration_spec"] = "PreVmGroupMigrationSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementMigrationVmgroupPreVmGroupMigratePreMigrateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementMigrationVmgroupPreVmGroupMigratePreMigrateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["pre_vm_group_migration_spec"] = bindings.NewReferenceType(model.PreVmGroupMigrationSpecBindingType)
	fieldNameMap["pre_vm_group_migration_spec"] = "PreVmGroupMigrationSpec"
	paramsTypeMap["pre_vm_group_migration_spec"] = bindings.NewReferenceType(model.PreVmGroupMigrationSpecBindingType)
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
		"action=pre_migrate",
		"pre_vm_group_migration_spec",
		"POST",
		"/api/v1/migration/vmgroup",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

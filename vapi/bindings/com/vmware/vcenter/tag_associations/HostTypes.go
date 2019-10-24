/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Host.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package tag_associations

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// The last status for the iterator. A field of this type is returned as part of the result and indicates to the caller of the API whether it can continue to make requests for more data. The last status only reports on the state of the iteration at the time data was last returned. As a result, it not does guarantee if the next call will succeed in getting more data or not. Failures to retrieve results will be returned as Error responses. These last statuses are only returned when the iterator is operating as expected. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostLastIterationStatus string

const (
    // Iterator has more data pending and is ready to provide it. The caller can request the next page of data at any time. The number of results returned may be less than the requested size. In other words, the iterator may not fill the page. The iterator has returned at least 1 result. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	HostLastIterationStatus_READY HostLastIterationStatus = "READY"
    // Iterator has finished iterating through its inventory. There are currently no more entities to return and the caller can terminate iteration. If the iterator returned some data, the marker may be set to allow the iterator to continue from where it left off when additional data does become available. This value is used to indicate that all available data has been returned by the iterator. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	HostLastIterationStatus_END_OF_DATA HostLastIterationStatus = "END_OF_DATA"
)

func (l HostLastIterationStatus) HostLastIterationStatus() bool {
	switch l {
	case HostLastIterationStatus_READY:
		return true
	case HostLastIterationStatus_END_OF_DATA:
		return true
	default:
		return false
	}
}

// The ``IterationSpec`` class contains properties used to break results into pages when listing tags associated to hosts, see Host#list). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostIterationSpec struct {
    // Size specifies the maximum number of results to return. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Size *int64
    // Marker is an opaque data structure that allows the ``Host`` interface to find where the previously returned page of tag associations has ended and where the new page should begin. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Marker *string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing tags associated to hosts, see Host#list). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostFilterSpec struct {
    // Tag associations need to have one of the host identifiers in this set to match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Hosts map[string]bool
    // Tag associations need to have one of the tag identifiers in this set to match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Tags map[string]bool
}

// The ``Summary`` describes a tag association. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostSummary struct {
    // The identifier of a tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Tag string
    // The identifier of a host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Host string
}

// The ``ListResult`` class contains the list of tag associations in a page, as well as related metadata fields. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostListResult struct {
    // List of tag associations. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Associations []HostSummary
    // Marker is an opaque data structure that allows the caller to request the next page of tag associations. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Marker *string
    // The last status for the iterator that indicates whether any more results can be expected if the caller continues to make requests for more data using the iterator. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Status HostLastIterationStatus
}


func hostListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["iterate"] = bindings.NewOptionalType(bindings.NewReferenceType(HostIterationSpecBindingType))
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(HostFilterSpecBindingType))
	fieldNameMap["iterate"] = "Iterate"
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(HostListResultBindingType)
}

func hostListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"Unauthorized": 403})
}


func HostIterationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["marker"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.tag_associations.Marker"}, ""))
	fieldNameMap["marker"] = "Marker"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tag_associations.host.iteration_spec", fields, reflect.TypeOf(HostIterationSpec{}), fieldNameMap, validators)
}

func HostFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tag_associations.host.filter_spec", fields, reflect.TypeOf(HostFilterSpec{}), fieldNameMap, validators)
}

func HostSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, "")
	fieldNameMap["tag"] = "Tag"
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tag_associations.host.summary", fields, reflect.TypeOf(HostSummary{}), fieldNameMap, validators)
}

func HostListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["associations"] = bindings.NewListType(bindings.NewReferenceType(HostSummaryBindingType), reflect.TypeOf([]HostSummary{}))
	fieldNameMap["associations"] = "Associations"
	fields["marker"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.tag_associations.Marker"}, ""))
	fieldNameMap["marker"] = "Marker"
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.tag_associations.host.last_iteration_status", reflect.TypeOf(HostLastIterationStatus(HostLastIterationStatus_READY)))
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tag_associations.host.list_result", fields, reflect.TypeOf(HostListResult{}), fieldNameMap, validators)
}


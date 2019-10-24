/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ReplicationStatus.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package topology

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Summary`` class contains replication information of partner vCenter or Platform Services Controller node of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type).
type ReplicationStatusSummary struct {
    // Identifier for the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node.
	Node string
    // Identifier for the vCenter or Platform Services Controller replication partner. Identifier can be either IP address or DNS resolvable name of the replication partner.
	ReplicationPartner string
    // Indicates if the VMware Directory Service on partner is reachable or not.
	PartnerAvailable bool
    // Indicates if the replication status for the node with respect to replication partner can be retrieved or not.
	StatusAvailable bool
    // Indicates if node is processing replication changes from the replication partner.
	Replicating *bool
    // Number of replication changes node is behind the replication partner.
	ChangeLag *int64
}

// The ``FilterSpec`` class contains property used to filter the results when listing replication status for the vCenter and Platform Services Controller nodes (see ReplicationStatus#list) of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type).
type ReplicationStatusFilterSpec struct {
    // Identifier that a vCenter and Platform Services Controller node must have to match the filter. (see ReplicationStatusSummary#node).
	Nodes map[string]bool
}


func replicationStatusListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ReplicationStatusFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func replicationStatusListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ReplicationStatusSummaryBindingType), reflect.TypeOf([]ReplicationStatusSummary{}))
}

func replicationStatusListRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400})
}


func ReplicationStatusSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node"] = bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, "")
	fieldNameMap["node"] = "Node"
	fields["replication_partner"] = bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, "")
	fieldNameMap["replication_partner"] = "ReplicationPartner"
	fields["partner_available"] = bindings.NewBooleanType()
	fieldNameMap["partner_available"] = "PartnerAvailable"
	fields["status_available"] = bindings.NewBooleanType()
	fieldNameMap["status_available"] = "StatusAvailable"
	fields["replicating"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["replicating"] = "Replicating"
	fields["change_lag"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["change_lag"] = "ChangeLag"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.topology.replication_status.summary", fields, reflect.TypeOf(ReplicationStatusSummary{}), fieldNameMap, validators)
}

func ReplicationStatusFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nodes"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["nodes"] = "Nodes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.topology.replication_status.filter_spec", fields, reflect.TypeOf(ReplicationStatusFilterSpec{}), fieldNameMap, validators)
}


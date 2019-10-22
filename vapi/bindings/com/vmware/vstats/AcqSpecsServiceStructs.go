/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: AcqSpecs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for acquisition specifications.
const AcqSpecs_RESOURCE_TYPE = "com.vmware.vstats.model.AcqSpec"


// Describes the status of an Acquisition Specification.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type AcqSpecs_Status string

const (
    // The acquisition specification is enabled.
     AcqSpecs_Status_ENABLED AcqSpecs_Status = "ENABLED"
    // The acquisition specification is disabled.
     AcqSpecs_Status_DISABLED AcqSpecs_Status = "DISABLED"
    // The acquisition specification is expired.
     AcqSpecs_Status_EXPIRED AcqSpecs_Status = "EXPIRED"
)

func (s AcqSpecs_Status) AcqSpecs_Status() bool {
    switch s {
        case AcqSpecs_Status_ENABLED:
            return true
        case AcqSpecs_Status_DISABLED:
            return true
        case AcqSpecs_Status_EXPIRED:
            return true
        default:
            return false
    }
}





// The ``CounterSpec`` class designates a counter or counter set in an acquisition specification.
 type AcqSpecsCounterSpec struct {
    // Counter and optional meatadata identifier.
    CidMid *CidMid
    // Counter set identifier
    SetId *string
}



//


// The ``CreateSpec`` class contains information for a new data acquisition specification.
 type AcqSpecsCreateSpec struct {
    // Designates the counter(s) to be sampled.
    Counters AcqSpecsCounterSpec
    // A set of resource identifiers representing a single resource to be monitored.
    Resources []RsrcId
    // Desired sampling rate in seconds.
    Interval *int64
    // Expiration time for this acquisition specification in Unix UTC number of seconds (since epoch).
    Expiration *int64
    // Consumer provided text about this acquisition specification.
    Memo_ *string
}



//


// The ``Info`` class is the information about an acquisition specification. It specifies the statistical data that should be collected at desired sampling rates. It designates the resources and their counters which should be sampled, and a desired sampling rate.
 type AcqSpecsInfo struct {
    // Acquisition specification identifier.
    Id string
    // Designates the counter(s) to be sampled.
    Counters AcqSpecsCounterSpec
    // A set of resource identifiers representing a single resource to be monitored.
    Resources []RsrcId
    // Desired sampling rate in seconds.
    Interval *int64
    // Expiration time for this acquisition specification represented as Unix UTC number of seconds (since epoch).
    Expiration *int64
    // Consumer provided text about this Acquisition Specification.
    Memo_ string
    // Acquisition Specification status.
    Status AcqSpecs_Status
}



//


// The ``ListResult`` class contains properties used to return the acquisition specifications.
 type AcqSpecsListResult struct {
    // List of acquisition specifications received.
    AcqSpecs []AcqSpecsInfo
    // The ``next`` property is a token used to retrieve paged data for larger result sets. This is opaque token generated by the server. It is to be sent in the AcqSpecsFilterSpec#page property to issue a subsequent call to the list method for retrieving results that did not fit the current page.
    Next *string
}



//


// The ``FilterSpec`` class contains properties used to filter the results when listing acquisition specifications.
 type AcqSpecsFilterSpec struct {
    // Used to retrieve paged data for larger result sets. The value of this token is generated by server and returned as ``next`` property in the result of ``list()`` methods.
    Page *string
}



//


// The ``UpdateSpec`` class contains properties that can be updated in an acquisition specification.
 type AcqSpecsUpdateSpec struct {
    // Designates the counter(s) to be sampled.
    Counters *AcqSpecsCounterSpec
    // A set of resource identifiers representing a single resource to be monitored.
    Resources []RsrcId
    // Desired sampling rate in seconds.
    Interval *int64
    // Expiration time for this acquisition specification in Unix UTC number of seconds (since epoch). When this property is 0, the default expiration duration is 100 minutes.
    Expiration *int64
    // Consumer provided text about this Acquisition Specification.
    Memo_ *string
}



//





func acqSpecsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["acq_spec"] = bindings.NewReferenceType(AcqSpecsCreateSpecBindingType)
    fieldNameMap["acq_spec"] = "AcqSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func acqSpecsCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
}

func acqSpecsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["acq_spec"] = bindings.NewReferenceType(AcqSpecsCreateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "acq_spec",
      "POST",
      "/stats/acq-specs",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func acqSpecsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func acqSpecsDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func acqSpecsDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/stats/acq-specs/{id}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func acqSpecsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(AcqSpecsFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func acqSpecsListOutputType() bindings.BindingType {
    return bindings.NewReferenceType(AcqSpecsListResultBindingType)
}

func acqSpecsListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.page"] = bindings.NewOptionalType(bindings.NewStringType())
    queryParams["filter.page"] = "page"
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
      "/stats/acq-specs",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func acqSpecsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func acqSpecsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(AcqSpecsInfoBindingType)
}

func acqSpecsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
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
      "/stats/acq-specs/{id}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func acqSpecsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fields["acq_spec"] = bindings.NewReferenceType(AcqSpecsUpdateSpecBindingType)
    fieldNameMap["id"] = "Id"
    fieldNameMap["acq_spec"] = "AcqSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func acqSpecsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func acqSpecsUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["acq_spec"] = bindings.NewReferenceType(AcqSpecsUpdateSpecBindingType)
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "acq_spec",
      "PATCH",
      "/stats/acq-specs/{id}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func AcqSpecsCounterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid_mid"] = bindings.NewOptionalType(bindings.NewReferenceType(CidMidBindingType))
    fieldNameMap["cid_mid"] = "CidMid"
    fields["set_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.CounterSet"}, ""))
    fieldNameMap["set_id"] = "SetId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.counter_spec",fields, reflect.TypeOf(AcqSpecsCounterSpec{}), fieldNameMap, validators)
}

func AcqSpecsCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewReferenceType(AcqSpecsCounterSpecBindingType)
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewListType(bindings.NewReferenceType(RsrcIdBindingType), reflect.TypeOf([]RsrcId{}))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["memo_"] = "Memo_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.create_spec",fields, reflect.TypeOf(AcqSpecsCreateSpec{}), fieldNameMap, validators)
}

func AcqSpecsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
    fields["counters"] = bindings.NewReferenceType(AcqSpecsCounterSpecBindingType)
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewListType(bindings.NewReferenceType(RsrcIdBindingType), reflect.TypeOf([]RsrcId{}))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewStringType()
    fieldNameMap["memo_"] = "Memo_"
    fields["status"] = bindings.NewEnumType("com.vmware.vstats.acq_specs.status", reflect.TypeOf(AcqSpecs_Status(AcqSpecs_Status_ENABLED)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.info",fields, reflect.TypeOf(AcqSpecsInfo{}), fieldNameMap, validators)
}

func AcqSpecsListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["acq_specs"] = bindings.NewListType(bindings.NewReferenceType(AcqSpecsInfoBindingType), reflect.TypeOf([]AcqSpecsInfo{}))
    fieldNameMap["acq_specs"] = "AcqSpecs"
    fields["next"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["next"] = "Next"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.list_result",fields, reflect.TypeOf(AcqSpecsListResult{}), fieldNameMap, validators)
}

func AcqSpecsFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["page"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["page"] = "Page"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.filter_spec",fields, reflect.TypeOf(AcqSpecsFilterSpec{}), fieldNameMap, validators)
}

func AcqSpecsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewOptionalType(bindings.NewReferenceType(AcqSpecsCounterSpecBindingType))
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(RsrcIdBindingType), reflect.TypeOf([]RsrcId{})))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["memo_"] = "Memo_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.update_spec",fields, reflect.TypeOf(AcqSpecsUpdateSpec{}), fieldNameMap, validators)
}



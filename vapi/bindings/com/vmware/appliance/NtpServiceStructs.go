/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ntp.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``ServerStatus`` enumeration class Status of server during test. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ntp_ServerStatus string

const (
    // Server is reachable. This constant field was added in vSphere API 6.7.
     Ntp_ServerStatus_SERVER_REACHABLE Ntp_ServerStatus = "SERVER_REACHABLE"
    // Server is unreachable. This constant field was added in vSphere API 6.7.
     Ntp_ServerStatus_SERVER_UNREACHABLE Ntp_ServerStatus = "SERVER_UNREACHABLE"
)

func (s Ntp_ServerStatus) Ntp_ServerStatus() bool {
    switch s {
        case Ntp_ServerStatus_SERVER_REACHABLE:
            return true
        case Ntp_ServerStatus_SERVER_UNREACHABLE:
            return true
        default:
            return false
    }
}





// ``LocalizableMessage`` class Structure representing message. This class was added in vSphere API 6.7.
 type NtpLocalizableMessage struct {
    // id in message bundle. This property was added in vSphere API 6.7.
    Id string
    // text in english. This property was added in vSphere API 6.7.
    DefaultMessage string
    // nested data. This property was added in vSphere API 6.7.
    Args []string
}



//


// ``TestRunStatus`` class Status of the test. This class was added in vSphere API 6.7.
 type NtpTestRunStatus struct {
    // Server name associated with the test run. This property was added in vSphere API 6.7.
    Server string
    // Server status. This property was added in vSphere API 6.7.
    Status Ntp_ServerStatus
    // Message associated with status. This property was added in vSphere API 6.7.
    Message NtpLocalizableMessage
}



//





func ntpTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ntpTestOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(NtpTestRunStatusBindingType), reflect.TypeOf([]NtpTestRunStatus{}))
}

func ntpTestRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func ntpSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ntpSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func ntpSetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func ntpGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ntpGetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
}

func ntpGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}



func NtpLocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.ntp.localizable_message",fields, reflect.TypeOf(NtpLocalizableMessage{}), fieldNameMap, validators)
}

func NtpTestRunStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server"] = bindings.NewStringType()
    fieldNameMap["server"] = "Server"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.ntp.server_status", reflect.TypeOf(Ntp_ServerStatus(Ntp_ServerStatus_SERVER_REACHABLE)))
    fieldNameMap["status"] = "Status"
    fields["message"] = bindings.NewReferenceType(NtpLocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.ntp.test_run_status",fields, reflect.TypeOf(NtpTestRunStatus{}), fieldNameMap, validators)
}



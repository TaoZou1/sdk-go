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

package kms

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains properties that describe the status of the key management service.
 type StatusInfo struct {
    // The service health status.
    Health StatusInfo_Health
    // Details regarding the health of the service. 
    //
    //  When the service ``Health`` is not StatusInfo_Health#Health_OK, this property will provide an actionable description of the issue.
    Details []std.LocalizableMessage
}



func (StatusInfo StatusInfo) Error() string {
    return "com.vmware.esx.kms.info"
}

    
    // The ``Health`` enumeration class defines the possible service health states.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type StatusInfo_Health string

    const (
        // None. 
        //
        //  No status available.
         StatusInfo_Health_NONE StatusInfo_Health = "NONE"
        // OK. 
        //
        //  Key management service is functioning normally.
         StatusInfo_Health_OK StatusInfo_Health = "OK"
        // Warning. 
        //
        //  Key management service is functioning, however there is an issue that requires attention.
         StatusInfo_Health_WARNING StatusInfo_Health = "WARNING"
        // Error. 
        //
        //  Key management service is not functioning.
         StatusInfo_Health_ERROR StatusInfo_Health = "ERROR"
    )

    func (h StatusInfo_Health) StatusInfo_Health() bool {
        switch h {
            case StatusInfo_Health_NONE:
                return true
            case StatusInfo_Health_OK:
                return true
            case StatusInfo_Health_WARNING:
                return true
            case StatusInfo_Health_ERROR:
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
      "/esx/kms/status",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403})
}



func StatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health"] = bindings.NewEnumType("com.vmware.esx.kms.status.info.health", reflect.TypeOf(StatusInfo_Health(StatusInfo_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.status.info",fields, reflect.TypeOf(StatusInfo{}), fieldNameMap, validators)
}



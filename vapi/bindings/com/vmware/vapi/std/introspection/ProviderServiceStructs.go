/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Provider.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package introspection

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// Information about a vAPI provider
 type ProviderInfo struct {
    // Identifier of the provider
    Id string
    // Checksum of the information present in the provider. 
//
//  Clients can use this information to check if the service information has changed. When a new service is added or removed (or) one of the existing service information is modified, the value of the checksum changes. 
//
//  The information used to calculate the checksum includes: 
//
// * service identifiers
// * operation identifiers inside the service
// * input, output and error definitions of an operation
    Checksum string
}









func providerGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providerGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProviderInfoBindingType)
}

func providerGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func ProviderInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vapi.provider"}, "")
    fieldNameMap["id"] = "Id"
    fields["checksum"] = bindings.NewStringType()
    fieldNameMap["checksum"] = "Checksum"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.introspection.provider.info",fields, reflect.TypeOf(ProviderInfo{}), fieldNameMap, validators)
}



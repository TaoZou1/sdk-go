/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Network.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package network

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class contains information about a vCenter Server network.
 type Info struct {
    Type_ string
}









func findInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["networks"] = bindings.NewListType(bindings.NewIdType([]string {"Network"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["networks"] = "Networks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func findOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"Network"}, ""), bindings.NewOptionalType(bindings.NewReferenceType(InfoBindingType)),reflect.TypeOf(map[string]*Info{}))
}

func findRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    isv1 := bindings.NewIsOneOfValidator(
        "type",
        []string {
             "Network",
             "DistributedVirtualPortgroup",
             "OpaqueNetwork",
        },
    )
    validators = append(validators, isv1)
    return bindings.NewStructType("com.vmware.vcenter.inventory.network.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



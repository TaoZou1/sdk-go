/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Namespace.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cli

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Identity`` class uniquely identifies a namespace in the CLI namespace tree.
 type NamespaceIdentity struct {
    // The dot-separated path of the namespace containing the namespace in the CLI node tree. For top-level namespace this will be empty.
    Path string
    // The name displayed to the user for this namespace.
    Name string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains information about a namespace. It includes the identity of the namespace, a description, information children namespaces.
 type NamespaceInfo struct {
    // Basic namespace identity.
    Identity NamespaceIdentity
    // The text description displayed to the user in help output.
    Description string
    // The children of this namespace in the tree of CLI namespaces.
    Children []NamespaceIdentity
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func namespaceListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
}

func namespaceListRestMetadata() protocol.OperationRestMetadata {
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


func namespaceGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["identity"] = bindings.NewReferenceType(NamespaceIdentityBindingType)
    fieldNameMap["identity"] = "Identity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(NamespaceInfoBindingType)
}

func namespaceGetRestMetadata() protocol.OperationRestMetadata {
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


func namespaceFingerprintInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceFingerprintOutputType() bindings.BindingType {
    return bindings.NewStringType()
}

func namespaceFingerprintRestMetadata() protocol.OperationRestMetadata {
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



func NamespaceIdentityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewStringType()
    fieldNameMap["path"] = "Path"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.namespace.identity",fields, reflect.TypeOf(NamespaceIdentity{}), fieldNameMap, validators)
}

func NamespaceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["identity"] = bindings.NewReferenceType(NamespaceIdentityBindingType)
    fieldNameMap["identity"] = "Identity"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["children"] = bindings.NewListType(bindings.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
    fieldNameMap["children"] = "Children"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.namespace.info",fields, reflect.TypeOf(NamespaceInfo{}), fieldNameMap, validators)
}



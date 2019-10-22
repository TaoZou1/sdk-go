/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vapi

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/authentication"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/cli"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/metamodel"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/privilege"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/routing"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)



// The ``ComponentInfo`` class holds component metadata of the different metadata types for an API component. The class allows any combination of metadata types to be aggregated into one instance.
type ComponentInfo struct {
    // The metamodel component data
    Metamodel metamodel.ComponentInfo
    // The CLI component data
    Cli *cli.ComponentInfo
    // The authentication component data
    Authentication *authentication.ComponentInfo
    // The routing component data
    Routing *routing.ComponentInfo
    // The privilege component data
    Privilege *privilege.ComponentInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``MetadataInfo`` is a class which holds a map of the available metadata aggregated in a ComponentInfo class.
type MetadataInfo struct {
    // Version of the current ``MetadataInfo`` class. Property value changes when the content of the ``MetadataInfo`` or referenced classes changes. This enables class processing adjustments.
    Version string
    // Component information of all available components. The key in the map is the identifier of the component and the value is the aggregated ComponentInfo.
    Metadata map[string]ComponentInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["metamodel"] = bindings.NewReferenceType(metamodel.ComponentInfoBindingType)
    fieldNameMap["metamodel"] = "Metamodel"
    fields["cli"] = bindings.NewOptionalType(bindings.NewReferenceType(cli.ComponentInfoBindingType))
    fieldNameMap["cli"] = "Cli"
    fields["authentication"] = bindings.NewOptionalType(bindings.NewReferenceType(authentication.ComponentInfoBindingType))
    fieldNameMap["authentication"] = "Authentication"
    fields["routing"] = bindings.NewOptionalType(bindings.NewReferenceType(routing.ComponentInfoBindingType))
    fieldNameMap["routing"] = "Routing"
    fields["privilege"] = bindings.NewOptionalType(bindings.NewReferenceType(privilege.ComponentInfoBindingType))
    fieldNameMap["privilege"] = "Privilege"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func MetadataInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["metadata"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.component"}, ""), bindings.NewReferenceType(ComponentInfoBindingType),reflect.TypeOf(map[string]ComponentInfo{}))
    fieldNameMap["metadata"] = "Metadata"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata_info",fields, reflect.TypeOf(MetadataInfo{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for a Key Provider.
const Providers_RESOURCE_TYPE = "com.vmware.esx.kms.providers"



// The ``NetworkAddress`` class contains an IP address or hostname and a port on which a connection can be established.
 type NetworkAddress struct {
    Hostname string
    Port *int64
}






// The ``Server`` class contains properties that describe a connection endpoint.
 type Server struct {
    Name string
    Address NetworkAddress
}






// The ``KmipServerCreateSpec`` class contains properties that describe the desired configuration for KMIP based key server.
 type KmipServerCreateSpec struct {
    Servers []Server
    Username *string
}






// The ``KeyServerCreateSpec`` class contains properties that describe the desired configuration for key server.
 type KeyServerCreateSpec struct {
    Type_ KeyServerCreateSpec_Type
    Description *string
    ProxyServer *NetworkAddress
    ConnectionTimeout *int64
    KmipServer *KmipServerCreateSpec
}




    
    // The ``Type`` enumeration class list the key server types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type KeyServerCreateSpec_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         KeyServerCreateSpec_Type_KMIP KeyServerCreateSpec_Type = "KMIP"
    )

    func (t KeyServerCreateSpec_Type) KeyServerCreateSpec_Type() bool {
        switch t {
            case KeyServerCreateSpec_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``CreateSpec`` class contains properties that describe the desired configuration for a new provider.
 type CreateSpec struct {
    Provider string
    MasterKeyId string
    KeyServer KeyServerCreateSpec
}






// The ``KmipServerUpdateSpec`` class contains properties that describe new configuration for KMIP based key server.
 type KmipServerUpdateSpec struct {
    Servers []Server
    Username *string
}






// The ``KeyServerUpdateSpec`` class contains properties that describe new configuration for an existing key server.
 type KeyServerUpdateSpec struct {
    Type_ *KeyServerUpdateSpec_Type
    Description *string
    ProxyServer *NetworkAddress
    ConnectionTimeout *int64
    KmipServer *KmipServerUpdateSpec
}




    
    // The ``Type`` enumeration class list the {\\\\@name KeyServer) types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type KeyServerUpdateSpec_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         KeyServerUpdateSpec_Type_KMIP KeyServerUpdateSpec_Type = "KMIP"
    )

    func (t KeyServerUpdateSpec_Type) KeyServerUpdateSpec_Type() bool {
        switch t {
            case KeyServerUpdateSpec_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``UpdateSpec`` class contains properties that describe the new configuration for an existing provider.
 type UpdateSpec struct {
    MasterKeyId *string
    KeyServer *KeyServerUpdateSpec
}






// The ``Summary`` class contains properties that summarize a provider.
 type Summary struct {
    Provider string
}






// The ``KmipServerInfo`` class contains properties that describe the current configuration of a KMIP based key server.
 type KmipServerInfo struct {
    Servers []Server
    Username *string
}






// The ``KeyServerInfo`` class contains properties that describe the current configuration of a key server.
 type KeyServerInfo struct {
    Type_ KeyServerInfo_Type
    Description string
    ProxyServer *NetworkAddress
    ConnectionTimeout *int64
    KmipServer *KmipServerInfo
}




    
    // The ``Type`` enumeration class list the {\\\\@name KeyServer) types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type KeyServerInfo_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         KeyServerInfo_Type_KMIP KeyServerInfo_Type = "KMIP"
    )

    func (t KeyServerInfo_Type) KeyServerInfo_Type() bool {
        switch t {
            case KeyServerInfo_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``Info`` class contains properties that describe the current configuration of a provider.
 type Info struct {
    MasterKeyId string
    KeyServer KeyServerInfo
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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
      "/esx/kms/providers",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/esx/kms/providers",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/esx/kms/providers/{provider}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func NetworkAddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.network_address",fields, reflect.TypeOf(NetworkAddress{}), fieldNameMap, validators)
}

func ServerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["address"] = bindings.NewReferenceType(NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.server",fields, reflect.TypeOf(Server{}), fieldNameMap, validators)
}

func KmipServerCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ServerBindingType), reflect.TypeOf([]Server{}))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_create_spec",fields, reflect.TypeOf(KmipServerCreateSpec{}), fieldNameMap, validators)
}

func KeyServerCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_create_spec.type", reflect.TypeOf(KeyServerCreateSpec_Type(KeyServerCreateSpec_Type_KMIP)))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(KmipServerCreateSpecBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_create_spec",fields, reflect.TypeOf(KeyServerCreateSpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["master_key_id"] = bindings.NewStringType()
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewReferenceType(KeyServerCreateSpecBindingType)
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func KmipServerUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ServerBindingType), reflect.TypeOf([]Server{})))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_update_spec",fields, reflect.TypeOf(KmipServerUpdateSpec{}), fieldNameMap, validators)
}

func KeyServerUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_update_spec.type", reflect.TypeOf(KeyServerUpdateSpec_Type(KeyServerUpdateSpec_Type_KMIP))))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(KmipServerUpdateSpecBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_update_spec",fields, reflect.TypeOf(KeyServerUpdateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["master_key_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewOptionalType(bindings.NewReferenceType(KeyServerUpdateSpecBindingType))
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func KmipServerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ServerBindingType), reflect.TypeOf([]Server{}))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_info",fields, reflect.TypeOf(KmipServerInfo{}), fieldNameMap, validators)
}

func KeyServerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_info.type", reflect.TypeOf(KeyServerInfo_Type(KeyServerInfo_Type_KMIP)))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(KmipServerInfoBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_info",fields, reflect.TypeOf(KeyServerInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["master_key_id"] = bindings.NewStringType()
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewReferenceType(KeyServerInfoBindingType)
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



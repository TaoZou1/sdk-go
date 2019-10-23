/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package kms

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/trusted_infrastructure"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the ``Services`` instances.
const Services_RESOURCE_TYPE = "com.vmware.vcenter.trusted_platform.kms.Service"



// The ``Summary`` class contains a summary of an ``Services`` instance.
 type ServicesSummary struct {
    // The service's unique identifier.
    Service string
    // The service's address.
    Address trusted_infrastructure.NetworkAddress
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to.
    Group string
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to.
    TrustAuthorityCluster string
}



func (ServicesSummary ServicesSummary) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.kms.summary"
}



// The ``Info`` class contains all the stored information about a ``Services`` instance.
 type ServicesInfo struct {
    // The service's address.
    Address trusted_infrastructure.NetworkAddress
    // The service's TLS certificate chain.
    TrustedCA trusted_infrastructure.X509CertChain
    // The group determines reports issued by which Attestation Service instances this ``Services`` instance can accept.
    Group string
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to.
    TrustAuthorityCluster string
}



func (ServicesInfo ServicesInfo) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.kms.info"
}



// The ``CreateSpec`` class contains the data necessary for adding a ``Services`` instance to the environment
 type ServicesCreateSpec struct {
    // The service's address.
    Address trusted_infrastructure.NetworkAddress
    // The service's TLS certificate chain.
    TrustedCA trusted_infrastructure.X509CertChain
    // The group determines reports issued by which Attestation Service instances this ``Services`` instance can accept.
    Group string
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to.
    TrustAuthorityCluster string
}



func (ServicesCreateSpec ServicesCreateSpec) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.kms.create_spec"
}



// The ``FilterSpec`` class contains the data necessary for identifying a Key Provider Service instance
 type ServicesFilterSpec struct {
    // A set of IDs by which to filter the services.
    Services map[string]bool
    // A set of address by which to filter.
    Address []trusted_infrastructure.NetworkAddress
    // The group determines reports issued by which Attestation Service instances this ``Services`` instance can accept.
    Group map[string]bool
    // The cluster determines which Trust Authority Cluster this ``Services`` belongs to.
    TrustAuthorityCluster map[string]bool
}



func (ServicesFilterSpec ServicesFilterSpec) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.kms.filter_spec"
}






func servicesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ServicesSummaryBindingType), reflect.TypeOf([]ServicesSummary{}))
}

func servicesListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
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
      "/vcenter/trusted-infrastructure/kms/services",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}


func servicesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ServicesInfoBindingType)
}

func servicesGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    pathParams["service"] = "service"
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
      "/vcenter/trusted-infrastructure/kms/services/{service}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func servicesCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
}

func servicesCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
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
      "/vcenter/trusted-infrastructure/kms/services",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}


func servicesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func servicesDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    pathParams["service"] = "service"
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
      "/vcenter/trusted-infrastructure/kms/services/{service}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ResourceBusy": 500,"Unauthenticated": 401})
}



func ServicesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, "")
    fieldNameMap["service"] = "Service"
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    fields["trust_authority_cluster"] = bindings.NewStringType()
    fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.kms.services.summary",fields, reflect.TypeOf(ServicesSummary{}), fieldNameMap, validators)
}

func ServicesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
    fieldNameMap["trusted_CA"] = "TrustedCA"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    fields["trust_authority_cluster"] = bindings.NewStringType()
    fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.kms.services.info",fields, reflect.TypeOf(ServicesInfo{}), fieldNameMap, validators)
}

func ServicesCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
    fieldNameMap["trusted_CA"] = "TrustedCA"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    fields["trust_authority_cluster"] = bindings.NewStringType()
    fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.kms.services.create_spec",fields, reflect.TypeOf(ServicesCreateSpec{}), fieldNameMap, validators)
}

func ServicesFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.kms.Service"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["services"] = "Services"
    fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["group"] = "Group"
    fields["trust_authority_cluster"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["trust_authority_cluster"] = "TrustAuthorityCluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.kms.services.filter_spec",fields, reflect.TypeOf(ServicesFilterSpec{}), fieldNameMap, validators)
}



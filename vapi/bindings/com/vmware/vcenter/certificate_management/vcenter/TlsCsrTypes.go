/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TlsCsr.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Info`` class contains information for a Certificate signing request.
type TlsCsrInfo struct {
    // Certificate Signing Request in PEM format.
	Csr string
}

// The ``Spec`` class contains information to generate a Private Key and CSR.
type TlsCsrSpec struct {
	KeySize *int64
	CommonName *string
    // Organization field in certificate subject
	Organization string
    // Organization unit field in certificate subject
	OrganizationUnit string
    // Locality field in certificate subject
	Locality string
    // State field in certificate subject
	StateOrProvince string
    // Country field in certificate subject
	Country string
    // Email field in Certificate extensions
	EmailAddress string
	SubjectAltName []string
}


func tlsCsrCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(TlsCsrSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsCsrCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TlsCsrInfoBindingType)
}

func tlsCsrCreateRestMetadata() protocol.OperationRestMetadata {
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


func TlsCsrInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["csr"] = bindings.NewStringType()
	fieldNameMap["csr"] = "Csr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls_csr.info", fields, reflect.TypeOf(TlsCsrInfo{}), fieldNameMap, validators)
}

func TlsCsrSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["key_size"] = "KeySize"
	fields["common_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["common_name"] = "CommonName"
	fields["organization"] = bindings.NewStringType()
	fieldNameMap["organization"] = "Organization"
	fields["organization_unit"] = bindings.NewStringType()
	fieldNameMap["organization_unit"] = "OrganizationUnit"
	fields["locality"] = bindings.NewStringType()
	fieldNameMap["locality"] = "Locality"
	fields["state_or_province"] = bindings.NewStringType()
	fieldNameMap["state_or_province"] = "StateOrProvince"
	fields["country"] = bindings.NewStringType()
	fieldNameMap["country"] = "Country"
	fields["email_address"] = bindings.NewStringType()
	fieldNameMap["email_address"] = "EmailAddress"
	fields["subject_alt_name"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["subject_alt_name"] = "SubjectAltName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls_csr.spec", fields, reflect.TypeOf(TlsCsrSpec{}), fieldNameMap, validators)
}


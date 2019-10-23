/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ImportSession.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
)



// The ``State`` enumeration class defines the different states of a transfer. 
//
//  For pull imports files are pulled by the server using HTTP GET, for push imports the client uses HTTP PUT to upload the data. For a content library import, the server retrieves the files from the content library. 
//
//  The transition diagram is as follows for a pull or content library import: 
//
// * ImportSession_State#ImportSessionState_IMPORT_OVF_TRANSFER: Initial state, where server is retrieving OVF descriptor, message bundles, certificates, and manifests.
// * ImportSession_State#ImportSessionState_IMPORT_SELECTING_OVF_PARAMS: Waiting for user to call Instantiate to begin the instantiation.
// * ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER: Transfering files from the source.
// * ImportSession_State#ImportSessionState_IMPORT_INSTANTIATING: Instantiating the virtual machines/virtual appliance.
// * ImportSession_State#ImportSessionState_IMPORT_COMPLETED: Import completed successfully.
//
//  
//
//  The transition diagram is as follows for a push import: 
//
// * ImportSession_State#ImportSessionState_IMPORT_OVF_TRANSFER: Initial state where server is waiting for the OVF descriptor to be uploaded.
// * ImportSession_State#ImportSessionState_IMPORT_MSG_BUNDLES_TRANSFER: Server is waiting for message bundles to be uploaded.
// * ImportSession_State#ImportSessionState_IMPORT_SELECTING_OVF_PARAMS: Waiting for user to begin the instantiation.
// * ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER: Server is waiting for all files to be uploaded.
// * ImportSession_State#ImportSessionState_IMPORT_INSTANTIATING: Instantiating the virtual machines/virtual appliance.
// * ImportSession_State#ImportSessionState_IMPORT_COMPLETED: Import completed successfully.
//
//  
//
//  The client is allowed to upload files as soon as the appear in the import session object. Thus upload can begin before reaching the ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER state.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ImportSession_State string

const (
    // State of an import transfer that does not have any files available. The transfer needs the OVF descriptor to continue. If this is a push transfer, the client must upload the OVF descriptor, and the transfer file list has one file info entry with a URL to which the client must upload the OVF descriptor using HTTP PUT. For pull transfers (including content library), the server is in the process of retrieving the OVF descriptor. 
    //
    //  Transition to the next state is done when the server has retrieved the complete OVF content and parsed it.
     ImportSession_State_IMPORT_OVF_TRANSFER ImportSession_State = "IMPORT_OVF_TRANSFER"
    // The file list contains a number of message bundles that need to be transferred to the server. If this is a push transfer, the client must PUT the requested files to the server. 
    //
    //  In case the OVF descriptor does not specify any bundles this state is skipped. 
    //
    //  Transition to next state is done when the complete content of all message bundles has been retrieved by the server.
     ImportSession_State_IMPORT_MSG_BUNDLES_TRANSFER ImportSession_State = "IMPORT_MSG_BUNDLES_TRANSFER"
    // The server can be queried for OVF parameters, and the client can specify instantiation parameters. 
    //
    //  Specifying an OVF instantiation parameter might affect other OVF instantiation parameters and change the set of files that needs to get transferred. 
    //
    //  During this state the client is allowed to push other files to the server with HTTP PUT. 
    //
    //  Transition to the next state is done by calling the instantiate method. If all needed files have been transferred to the server a transition is made to ImportSession_State#ImportSessionState_IMPORT_INSTANTIATING, otherwise a transition is made to ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER.
     ImportSession_State_IMPORT_SELECTING_OVF_PARAMS ImportSession_State = "IMPORT_SELECTING_OVF_PARAMS"
    // State for transferring remaining files that have not been transferred during an earlier state. 
    //
    //  When the content of all files has been transferred the state is changed to ImportSession_State#ImportSessionState_IMPORT_INSTANTIATING: 
    //
    // * Pull transfer: All needed files have been completely transferred, and the manifest and certificate have been transferred if they are available.
    // * Push transfer: All files marked as required (attribute optional == false) have been transferred and all initiated transfers of optional files are complete. In particular the manifest and certificate are optional so the upload of these has to at least be initiated before other files are completely uploaded.
     ImportSession_State_IMPORT_FILE_TRANSFER ImportSession_State = "IMPORT_FILE_TRANSFER"
    // The virtual machine or virtual appliance is being instantiated.
     ImportSession_State_IMPORT_INSTANTIATING ImportSession_State = "IMPORT_INSTANTIATING"
    // The virtual machine or virtual appliance is instantiated, and the upload transfer is done.
     ImportSession_State_IMPORT_COMPLETED ImportSession_State = "IMPORT_COMPLETED"
    // The transfer failed.
     ImportSession_State_IMPORT_ERROR ImportSession_State = "IMPORT_ERROR"
)

func (s ImportSession_State) ImportSession_State() bool {
    switch s {
        case ImportSession_State_IMPORT_OVF_TRANSFER:
            return true
        case ImportSession_State_IMPORT_MSG_BUNDLES_TRANSFER:
            return true
        case ImportSession_State_IMPORT_SELECTING_OVF_PARAMS:
            return true
        case ImportSession_State_IMPORT_FILE_TRANSFER:
            return true
        case ImportSession_State_IMPORT_INSTANTIATING:
            return true
        case ImportSession_State_IMPORT_COMPLETED:
            return true
        case ImportSession_State_IMPORT_ERROR:
            return true
        default:
            return false
    }
}




// The ``SourceType`` enumeration class defines the type of source.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ImportSession_SourceType string

const (
    // The client is uploading content using HTTP(S) PUT requests.
     ImportSession_SourceType_PUSH_SOURCE ImportSession_SourceType = "PUSH_SOURCE"
    // The server is pulling content from a URL.
     ImportSession_SourceType_PULL_SOURCE ImportSession_SourceType = "PULL_SOURCE"
    // The server is pulling content from a library item.
     ImportSession_SourceType_CONTENT_LIBRARY_SOURCE ImportSession_SourceType = "CONTENT_LIBRARY_SOURCE"
)

func (s ImportSession_SourceType) ImportSession_SourceType() bool {
    switch s {
        case ImportSession_SourceType_PUSH_SOURCE:
            return true
        case ImportSession_SourceType_PULL_SOURCE:
            return true
        case ImportSession_SourceType_CONTENT_LIBRARY_SOURCE:
            return true
        default:
            return false
    }
}




// The ``PushSourceContentType`` enumeration class defines the content type of source by push.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ImportSession_PushSourceContentType string

const (
    // The client is uploading an OVF template.
     ImportSession_PushSourceContentType_OVF_SOURCE ImportSession_PushSourceContentType = "OVF_SOURCE"
    // The client is uploading an OVA template.
     ImportSession_PushSourceContentType_OVA_SOURCE ImportSession_PushSourceContentType = "OVA_SOURCE"
)

func (p ImportSession_PushSourceContentType) ImportSession_PushSourceContentType() bool {
    switch p {
        case ImportSession_PushSourceContentType_OVF_SOURCE:
            return true
        case ImportSession_PushSourceContentType_OVA_SOURCE:
            return true
        default:
            return false
    }
}




// The ``PushSourceOvfOption`` class specifies the optional information that the OVF template provides. This only applies to ImportSession_PushSourceContentType#ImportSessionPushSourceContentType_OVF_SOURCE.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ImportSession_PushSourceOvfOption string

const (
    // Indicates that manifest file will be provided when pushing the OVF template.
     ImportSession_PushSourceOvfOption_MANIFEST ImportSession_PushSourceOvfOption = "MANIFEST"
    // Indicates that manifest and certificate file will be provided when pushing the OVF template.
     ImportSession_PushSourceOvfOption_MANIFEST_CERTIFICATE ImportSession_PushSourceOvfOption = "MANIFEST_CERTIFICATE"
    // Indicates that neither manifest nor certificate file will be provided when pushing the OVF template.
     ImportSession_PushSourceOvfOption_NONE ImportSession_PushSourceOvfOption = "NONE"
)

func (p ImportSession_PushSourceOvfOption) ImportSession_PushSourceOvfOption() bool {
    switch p {
        case ImportSession_PushSourceOvfOption_MANIFEST:
            return true
        case ImportSession_PushSourceOvfOption_MANIFEST_CERTIFICATE:
            return true
        case ImportSession_PushSourceOvfOption_NONE:
            return true
        default:
            return false
    }
}




// The ``FileTransferMode`` class specifies the file transfer mode for disk files during the OVF template deployment.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ImportSession_FileTransferMode string

const (
    // Indicates that the server will directly handle all file transfers. This is the default transfer mode.
     ImportSession_FileTransferMode_SERVER_TRANSFER_MODE ImportSession_FileTransferMode = "SERVER_TRANSFER_MODE"
    // Indicates that the server delegates files transfer to external services. In this mode, if the ImportSessionCreateSpec#sourceType is {ImportSession_SourceType#ImportSessionSourceType_PUSH_SOURCE, client should call ImportSession#progress to update file transfer progress, besides pushing files to OvfFileInfo#fileUrl.
     ImportSession_FileTransferMode_EXTERNAL_TRANSFER_MODE ImportSession_FileTransferMode = "EXTERNAL_TRANSFER_MODE"
)

func (f ImportSession_FileTransferMode) ImportSession_FileTransferMode() bool {
    switch f {
        case ImportSession_FileTransferMode_SERVER_TRANSFER_MODE:
            return true
        case ImportSession_FileTransferMode_EXTERNAL_TRANSFER_MODE:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class contains information about the import parameters.
 type ImportSessionCreateSpec struct {
    // The locale that should be used for parsing the OVF descriptor. This defaults to the server locale if not specified.
    Locale *string
    // Array of selected import flags. The supported flags are: 
    //
    // * LAX: Lax mode parsing of the OVF descriptor.
    ImportFlags []string
    // Type of import source.
    SourceType ImportSession_SourceType
    // If type is ImportSession_SourceType#ImportSessionSourceType_PULL_SOURCE, this is the source URL.
    PullSource *url.URL
    // Thumbprint of the expected SSL certificate for ImportSession_SourceType#ImportSessionSourceType_PULL_SOURCE. It is only used for HTTPS connections, otherwise it will be ignored. The thumbprint is the SHA-1 hash of the DER encoding of the remote endpoint's SSL certificate. If set, the remote endpoint's SSL certificate is only accepted if it matches this thumbprint, and no other certificate validation is performed.
    SslCertificateThumbprint *string
    // This is the source content type for ImportSession_SourceType#ImportSessionSourceType_PUSH_SOURCE, . The default source content type is ImportSession_PushSourceContentType#ImportSessionPushSourceContentType_OVF_SOURCE.
    PushSourceContentType *ImportSession_PushSourceContentType
    // This specifies the ImportSession_PushSourceOvfOption for ImportSession_PushSourceContentType#ImportSessionPushSourceContentType_OVF_SOURCE. This defaults to ImportSession_PushSourceOvfOption#ImportSessionPushSourceOvfOption_NONE if not specified.
    PushSourceOvfOption *ImportSession_PushSourceOvfOption
    // If type is ImportSession_SourceType#ImportSessionSourceType_CONTENT_LIBRARY_SOURCE, this is the identifier of source content library item.
    ContentLibraryItem *string
}



func (ImportSessionCreateSpec ImportSessionCreateSpec) Error() string {
    return "com.vmware.vcenter.ovf.create_spec"
}



// The ``OvfValidationResult`` class contains information about the result of calling tryInstantiate.
 type ImportSessionOvfValidationResult struct {
    // A array of OVF parameter classes providing information about the OVF package and the target environment.
    Parameters []*data.StructValue
    // Array of errors. Errors will prevent the import from proceeding.
    Errors []OvfError
    // Array of warnings. Warnings will not prevent the import from proceeding, but there might be issues that warrant attention.
    Warnings []OvfWarning
    // Array of information. A typical case is that a non-required element will be skipped for some reason.
    Information []OvfInfo
}



func (ImportSessionOvfValidationResult ImportSessionOvfValidationResult) Error() string {
    return "com.vmware.vcenter.ovf.ovf_validation_result"
}



// The ``Info`` class represents an import session.
 type ImportSessionInfo struct {
    // The state for the current transfer.
    State ImportSession_State
    // Progress of import.
    Progress int64
    // The transfer mode for disk files. 
    //
    //  It is unset initially until state ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER is reached.
    FileTransferMode *ImportSession_FileTransferMode
    // Array of files that are part of this transfer. 
    //
    //  There is an entry for the OVF descriptor, files defined in the file reference section of the OVF descriptor, and optional manifest and certificate. 
    //
    //  The entries are not fixed during the lifetime of the import. It can change when instantiation parameters are selected. It is fixed from state ImportSession_State#ImportSessionState_IMPORT_FILE_TRANSFER and onwards.
    Files []OvfFileInfo
    // Array of errors. Errors will prevent the import from proceeding.
    Errors []OvfError
    // Array of warnings. Warnings will not prevent the import from proceeding, but there might be issues that warrant attention.
    Warnings []OvfWarning
    // Array of information. A typical case is that a non-required element will be skipped for some reason.
    Information []OvfInfo
    // The identifier of the virtual appliance imported in case the OVF was a virtual appliance. This is not set until the virtual appliance is created. Only one of ImportSessionInfo#vappId and ImportSessionInfo#vmId will be set.
    VappId *string
    // The identifier of the virtual machine imported in case the OVF was a virtual machine. This is not set until a virtual machine is created. Only one of ImportSessionInfo#vappId and ImportSessionInfo#vmId will be set.
    VmId *string
    // Results of all the customizations applied to deployed resource entity or sub-entities as part of the ``instantiate()`` method. The key in the map is the identifier of the resource entity and the value is the success indicator of the application of the customization to the resource entity.
    CustomizationResults map[string]bool
}



func (ImportSessionInfo ImportSessionInfo) Error() string {
    return "com.vmware.vcenter.ovf.info"
}



// The ``PreviewFile`` class contains information about a file that is referenced in the OVF descriptor and will be imported according to an import preview.
 type ImportSessionPreviewFile struct {
    // The name of the file.
    Name string
}



func (ImportSessionPreviewFile ImportSessionPreviewFile) Error() string {
    return "com.vmware.vcenter.ovf.preview_file"
}



// The ``Preview`` class contains information about the result of an OVF import preview based on OVF descriptor, which includes referenced files and OVF validation information.
 type ImportSessionPreview struct {
    // The files that must be transferred during the import of an OVF template. This includes disks, CD images, resource bundles, etc.
    Files []ImportSessionPreviewFile
    // Array of errors. Errors will prevent the import from proceeding.
    Errors []OvfError
    // Array of warnings. Warnings will not prevent the import from proceeding, but there might be issues that warrant attention.
    Warnings []OvfWarning
    // Array of information. A typical case is that a non-required element will be skipped for some reason.
    Information []OvfInfo
}



func (ImportSessionPreview ImportSessionPreview) Error() string {
    return "com.vmware.vcenter.ovf.preview"
}



// The ``ProbeResult`` class contains information about the accessibility of a pull source URI.
 type ImportSessionProbeResult struct {
    // Status of the attempt to access the pull source URI.
    Status ImportSessionProbeResult_Status
    // The SSL thumbprint for a pull source URI with the ``https`` scheme.
    SslThumbprint *string
    // Detailed error messages if the the attempt to access the pull source URI failed.
    ErrorMessages []std.LocalizableMessage
}



func (ImportSessionProbeResult ImportSessionProbeResult) Error() string {
    return "com.vmware.vcenter.ovf.probe_result"
}

    
    // The ``Status`` enumeration class defines the possible status values from an attempt to access a pull source URI.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ImportSessionProbeResult_Status string

    const (
        // Indicates that the probe was successful.
         ImportSessionProbeResult_Status_SUCCESS ImportSessionProbeResult_Status = "SUCCESS"
        // Indicates that the supplied URL was not valid.
         ImportSessionProbeResult_Status_INVALID_URL ImportSessionProbeResult_Status = "INVALID_URL"
        // Indicates that the probe timed out while attempting to connect to the URL.
         ImportSessionProbeResult_Status_TIMED_OUT ImportSessionProbeResult_Status = "TIMED_OUT"
        // Indicates that the host in the URL could not be found.
         ImportSessionProbeResult_Status_HOST_NOT_FOUND ImportSessionProbeResult_Status = "HOST_NOT_FOUND"
        // Indicates that the provided server certificate thumbprint is invalid. In this case, the returned certificate thumbprint should be provided.
         ImportSessionProbeResult_Status_CERTIFICATE_ERROR ImportSessionProbeResult_Status = "CERTIFICATE_ERROR"
        // Indicates an unspecified error different from the other error cases defined in ImportSessionProbeResult_Status.
         ImportSessionProbeResult_Status_UNKNOWN_ERROR ImportSessionProbeResult_Status = "UNKNOWN_ERROR"
    )

    func (s ImportSessionProbeResult_Status) ImportSessionProbeResult_Status() bool {
        switch s {
            case ImportSessionProbeResult_Status_SUCCESS:
                return true
            case ImportSessionProbeResult_Status_INVALID_URL:
                return true
            case ImportSessionProbeResult_Status_TIMED_OUT:
                return true
            case ImportSessionProbeResult_Status_HOST_NOT_FOUND:
                return true
            case ImportSessionProbeResult_Status_CERTIFICATE_ERROR:
                return true
            case ImportSessionProbeResult_Status_UNKNOWN_ERROR:
                return true
            default:
                return false
        }
    }






func importSessionCreateForResourcePoolInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fields["host_system"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ImportSessionCreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["resource_pool"] = "ResourcePool"
    fieldNameMap["host_system"] = "HostSystem"
    fieldNameMap["folder"] = "Folder"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionCreateForResourcePoolOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
}

func importSessionCreateForResourcePoolRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}


func importSessionGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ImportSessionInfoBindingType)
}

func importSessionGetRestMetadata() protocol.OperationRestMetadata {
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


func importSessionTryInstantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionTryInstantiateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ImportSessionOvfValidationResultBindingType)
}

func importSessionTryInstantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthorized": 403})
}


func importSessionInstantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionInstantiateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func importSessionInstantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"Unsupported": 400,"Unauthorized": 403})
}


func importSessionProgressInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fields["percent"] = bindings.NewIntegerType()
    fieldNameMap["id"] = "Id"
    fieldNameMap["percent"] = "Percent"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionProgressOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func importSessionProgressRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}


func importSessionDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfImportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func importSessionDeleteRestMetadata() protocol.OperationRestMetadata {
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


func importSessionPreviewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ovf_descriptor"] = bindings.NewStringType()
    fieldNameMap["ovf_descriptor"] = "OvfDescriptor"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionPreviewOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ImportSessionPreviewBindingType)
}

func importSessionPreviewRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400})
}


func importSessionProbeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uri"] = bindings.NewUriType()
    fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["uri"] = "Uri"
    fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSessionProbeOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ImportSessionProbeResultBindingType)
}

func importSessionProbeRestMetadata() protocol.OperationRestMetadata {
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



func ImportSessionCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["locale"] = "Locale"
    fields["import_flags"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["import_flags"] = "ImportFlags"
    fields["source_type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.source_type", reflect.TypeOf(ImportSession_SourceType(ImportSession_SourceType_PUSH_SOURCE)))
    fieldNameMap["source_type"] = "SourceType"
    fields["pull_source"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["pull_source"] = "PullSource"
    fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
    fields["push_source_content_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.push_source_content_type", reflect.TypeOf(ImportSession_PushSourceContentType(ImportSession_PushSourceContentType_OVF_SOURCE))))
    fieldNameMap["push_source_content_type"] = "PushSourceContentType"
    fields["push_source_ovf_option"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.push_source_ovf_option", reflect.TypeOf(ImportSession_PushSourceOvfOption(ImportSession_PushSourceOvfOption_MANIFEST))))
    fieldNameMap["push_source_ovf_option"] = "PushSourceOvfOption"
    fields["content_library_item"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["content_library_item"] = "ContentLibraryItem"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("source_type",
        map[string][]bindings.FieldData {
            "PULL_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("pull_source", true),
                 bindings.NewFieldData("ssl_certificate_thumbprint", false),
            },
            "PUSH_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("push_source_content_type", false),
            },
            "CONTENT_LIBRARY_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("content_library_item", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("push_source_content_type",
        map[string][]bindings.FieldData {
            "OVF_SOURCE": []bindings.FieldData {
                 bindings.NewFieldData("push_source_ovf_option", false),
            },
            "OVA_SOURCE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.create_spec",fields, reflect.TypeOf(ImportSessionCreateSpec{}), fieldNameMap, validators)
}

func ImportSessionOvfValidationResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["parameters"] = "Parameters"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.ovf_validation_result",fields, reflect.TypeOf(ImportSessionOvfValidationResult{}), fieldNameMap, validators)
}

func ImportSessionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.state", reflect.TypeOf(ImportSession_State(ImportSession_State_IMPORT_OVF_TRANSFER)))
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    fields["file_transfer_mode"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.file_transfer_mode", reflect.TypeOf(ImportSession_FileTransferMode(ImportSession_FileTransferMode_SERVER_TRANSFER_MODE))))
    fieldNameMap["file_transfer_mode"] = "FileTransferMode"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(OvfFileInfoBindingType), reflect.TypeOf([]OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
    fieldNameMap["information"] = "Information"
    fields["vapp_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualApp"}, ""))
    fieldNameMap["vapp_id"] = "VappId"
    fields["vm_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualMachine"}, ""))
    fieldNameMap["vm_id"] = "VmId"
    fields["customization_results"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewBooleanType(),reflect.TypeOf(map[string]bool{})))
    fieldNameMap["customization_results"] = "CustomizationResults"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "IMPORT_FILE_TRANSFER": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_INSTANTIATING": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_COMPLETED": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("file_transfer_mode", false),
            },
            "IMPORT_OVF_TRANSFER": []bindings.FieldData {},
            "IMPORT_MSG_BUNDLES_TRANSFER": []bindings.FieldData {},
            "IMPORT_SELECTING_OVF_PARAMS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.info",fields, reflect.TypeOf(ImportSessionInfo{}), fieldNameMap, validators)
}

func ImportSessionPreviewFileBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.preview_file",fields, reflect.TypeOf(ImportSessionPreviewFile{}), fieldNameMap, validators)
}

func ImportSessionPreviewBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(ImportSessionPreviewFileBindingType), reflect.TypeOf([]ImportSessionPreviewFile{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.preview",fields, reflect.TypeOf(ImportSessionPreview{}), fieldNameMap, validators)
}

func ImportSessionProbeResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.ovf.import_session.probe_result.status", reflect.TypeOf(ImportSessionProbeResult_Status(ImportSessionProbeResult_Status_SUCCESS)))
    fieldNameMap["status"] = "Status"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["error_messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "CERTIFICATE_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", false),
            },
            "SUCCESS": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
            },
            "UNKNOWN_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", false),
            },
            "INVALID_URL": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
            "TIMED_OUT": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
            "HOST_NOT_FOUND": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.import_session.probe_result",fields, reflect.TypeOf(ImportSessionProbeResult{}), fieldNameMap, validators)
}



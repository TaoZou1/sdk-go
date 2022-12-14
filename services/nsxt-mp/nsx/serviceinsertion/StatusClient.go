// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Status
// Used by client-side stubs.

package serviceinsertion

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type StatusClient interface {

	// Fetches service insertion global status for supported contexts types e.g. 'east_west'
	//  Note- GET service insertion status API is deprecated. Please use the policy service insertion configuration status API.
	//
	// @param contextTypeParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(contextTypeParam string) (model.ServiceInsertionStatus, error)

	// Fetches service insertion status for all supported context types. Supported context types are 'east_west' and 'north_south'.
	//  Note- GET service insertion status API is deprecated. Please use the policy service insertion configuration status API.
	// @return com.vmware.nsx.model.ServiceInsertionStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.ServiceInsertionStatusListResult, error)

	// Update global serviceinsertion status for supported context types e.g. 'north_south' Please use 'north_south' context to update status of SI NFV usecase (rule applied on T0/T1 uplink with redirect_to service chain).
	//  The PUT service insertion status API is deprecated. Please use the policy service insertion configuration status API.
	//
	// @param contextTypeParam (required)
	// @param serviceInsertionStatusParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(contextTypeParam string, serviceInsertionStatusParam model.ServiceInsertionStatus) (model.ServiceInsertionStatus, error)
}

type statusClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewStatusClient(connector client.Connector) *statusClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.status")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := statusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statusClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statusClient) Get(contextTypeParam string) (model.ServiceInsertionStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(statusGetInputType(), typeConverter)
	sv.AddStructField("ContextType", contextTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := statusGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.status", "get", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), statusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *statusClient) List() (model.ServiceInsertionStatusListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(statusListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := statusListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.status", "list", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), statusListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *statusClient) Update(contextTypeParam string, serviceInsertionStatusParam model.ServiceInsertionStatus) (model.ServiceInsertionStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(statusUpdateInputType(), typeConverter)
	sv.AddStructField("ContextType", contextTypeParam)
	sv.AddStructField("ServiceInsertionStatus", serviceInsertionStatusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := statusUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.status", "update", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), statusUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

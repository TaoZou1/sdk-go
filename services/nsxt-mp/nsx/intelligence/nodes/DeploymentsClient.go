// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Deployments
// Used by client-side stubs.

package nodes

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type DeploymentsClient interface {

	// Deploys a Intelligence cluster node VM as specified by the deployment config.
	//
	// @param addIntelligenceClusterNodeVMInfoParam (required)
	// @return com.vmware.nsx.model.IntelligenceClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(addIntelligenceClusterNodeVMInfoParam model.AddIntelligenceClusterNodeVMInfo) (model.IntelligenceClusterNodeVMDeploymentRequestList, error)

	// Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted.
	//
	// @param nodeIdParam (required)
	// @param forceDeleteParam Delete by force (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(nodeIdParam string, forceDeleteParam *bool) error

	// Returns deployment request information for a specific attempted deployment of a cluster node VM.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.nsx.model.IntelligenceClusterNodeVMDeploymentRequest
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentRequest, error)

	// Returns request information for every attempted deployment of a cluster node VM.
	// @return com.vmware.nsx.model.IntelligenceClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.IntelligenceClusterNodeVMDeploymentRequestList, error)
}

type deploymentsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDeploymentsClient(connector client.Connector) *deploymentsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.intelligence.nodes.deployments")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := deploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentsClient) Create(addIntelligenceClusterNodeVMInfoParam model.AddIntelligenceClusterNodeVMInfo) (model.IntelligenceClusterNodeVMDeploymentRequestList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentsCreateInputType(), typeConverter)
	sv.AddStructField("AddIntelligenceClusterNodeVMInfo", addIntelligenceClusterNodeVMInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.nodes.deployments", "create", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) Delete(nodeIdParam string, forceDeleteParam *bool) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentsDeleteInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("ForceDelete", forceDeleteParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.nodes.deployments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *deploymentsClient) Get(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentRequest, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentsGetInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequest
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.nodes.deployments", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) List() (model.IntelligenceClusterNodeVMDeploymentRequestList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentsListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.nodes.deployments", "list", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

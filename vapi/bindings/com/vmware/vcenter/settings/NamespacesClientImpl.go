/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Namespaces
 * Functions that implement the generated NamespacesClient interface
 */

package settings

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type NamespacesClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewNamespacesClientImpl(connector client.Connector) *NamespacesClientImpl {
	interfaceName := "com.vmware.vcenter.settings.namespaces"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	nIface := NamespacesClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	nIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	nIface.methodNameToDefMap["list"] = nIface.listMethodDefinition()
	nIface.methodNameToDefMap["get"] = nIface.getMethodDefinition()
	nIface.methodNameToDefMap["create"] = nIface.createMethodDefinition()
	nIface.methodNameToDefMap["delete"] = nIface.deleteMethodDefinition()
	nIface.methodNameToDefMap["update"] = nIface.updateMethodDefinition()
	return &nIface
}


func (nIface *NamespacesClientImpl) List() ([]NamespacesInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(namespacesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []NamespacesInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := namespacesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []NamespacesInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), namespacesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]NamespacesInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *NamespacesClientImpl) Get(namespaceParam string) (NamespacesInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(namespacesGetInputType(), typeConverter)
	sv.AddStructField("Namespace", namespaceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput NamespacesInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := namespacesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput NamespacesInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), namespacesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(NamespacesInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *NamespacesClientImpl) Create(specParam NamespacesCreateSpec) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(namespacesCreateInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := namespacesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *NamespacesClientImpl) Delete(namespaceParam string) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(namespacesDeleteInputType(), typeConverter)
	sv.AddStructField("Namespace", namespaceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := namespacesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *NamespacesClientImpl) Update(namespaceParam string, specParam NamespacesUpdateSpec) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(namespacesUpdateInputType(), typeConverter)
	sv.AddStructField("Namespace", namespaceParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := namespacesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *NamespacesClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := nIface.connector.GetApiProvider().Invoke(nIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (nIface *NamespacesClientImpl) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(namespacesListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(namespacesListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NamespacesClientImpl) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(namespacesGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(namespacesGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NamespacesClientImpl) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(namespacesCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(namespacesCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's errors.AlreadyExists error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NamespacesClientImpl) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(namespacesDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(namespacesDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.delete method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NamespacesClientImpl) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(namespacesUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(namespacesUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NamespacesClientImpl.update method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

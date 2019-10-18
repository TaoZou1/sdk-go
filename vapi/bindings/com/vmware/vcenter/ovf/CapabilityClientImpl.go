
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Capability
 * Functions that implement the generated CapabilityClient interface
 */


package ovf
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type CapabilityClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewCapabilityClientImpl(connector client.Connector) *CapabilityClientImpl {
      interfaceName := "com.vmware.vcenter.ovf.capability"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      cIface := CapabilityClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      cIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      cIface.methodNameToDefMap["get"] = cIface.getMethodDefinition()
      return &cIface
}

func (cIface *CapabilityClientImpl) Get(serverGuidParam string) (CapabilityCapabilityInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(capabilityGetInputType(), typeConverter)
	sv.AddStructField("ServerGuid", serverGuidParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput CapabilityCapabilityInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := capabilityGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput CapabilityCapabilityInfo
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), capabilityGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(CapabilityCapabilityInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *CapabilityClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := cIface.connector.GetApiProvider().Invoke(cIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (cIface *CapabilityClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(capabilityGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(capabilityGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CapabilityClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CapabilityClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CapabilityClientImpl.get method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}




/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: HealthCheckSettings
 * Functions that implement the generated HealthCheckSettingsClient interface
 */


package appliance
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type HealthCheckSettingsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewHealthCheckSettingsClientImpl(connector client.Connector) *HealthCheckSettingsClientImpl {
      interfaceName := "com.vmware.appliance.health_check_settings"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
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
      hIface := HealthCheckSettingsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      hIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      hIface.methodNameToDefMap["get"] = hIface.getMethodDefinition()
      hIface.methodNameToDefMap["update"] = hIface.updateMethodDefinition()
      return &hIface
}

func (hIface *HealthCheckSettingsClientImpl) Get() (HealthCheckSettingsSettingSpec, error) {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(healthCheckSettingsGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput HealthCheckSettingsSettingSpec
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := healthCheckSettingsGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput HealthCheckSettingsSettingSpec
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), healthCheckSettingsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(HealthCheckSettingsSettingSpec), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (hIface *HealthCheckSettingsClientImpl) Update(healthSettingsParam HealthCheckSettingsUpdateSpec) error {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(healthCheckSettingsUpdateInputType(), typeConverter)
	sv.AddStructField("HealthSettings", healthSettingsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := healthCheckSettingsUpdateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *HealthCheckSettingsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := hIface.connector.GetApiProvider().Invoke(hIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (hIface *HealthCheckSettingsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
      typeConverter := hIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(healthCheckSettingsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(healthCheckSettingsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      hIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (hIface *HealthCheckSettingsClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
      typeConverter := hIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(healthCheckSettingsUpdateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(healthCheckSettingsUpdateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      hIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.update method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      hIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for HealthCheckSettingsClientImpl.update method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



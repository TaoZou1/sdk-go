
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Floppy
 * Functions that implement the generated FloppyClient interface
 */


package floppy
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type FloppyClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewFloppyClientImpl(connector client.Connector) *FloppyClientImpl {
      interfaceName := "com.vmware.vcenter.vm.hardware.floppy"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "update"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
          core.NewMethodIdentifier(interfaceIdentifier, "connect"),
          core.NewMethodIdentifier(interfaceIdentifier, "disconnect"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      fIface := FloppyClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      fIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      fIface.methodNameToDefMap["list"] = fIface.listMethodDefinition()
      fIface.methodNameToDefMap["get"] = fIface.getMethodDefinition()
      fIface.methodNameToDefMap["create"] = fIface.createMethodDefinition()
      fIface.methodNameToDefMap["update"] = fIface.updateMethodDefinition()
      fIface.methodNameToDefMap["delete"] = fIface.deleteMethodDefinition()
      fIface.methodNameToDefMap["connect"] = fIface.connectMethodDefinition()
      fIface.methodNameToDefMap["disconnect"] = fIface.disconnectMethodDefinition()
      return &fIface
}

func (fIface *FloppyClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(fIface.interfaceName)
}

func (fIface *FloppyClientImpl) Definition() core.InterfaceDefinition {
    return fIface.interfaceDefinition
}

func (fIface *FloppyClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := fIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (fIface *FloppyClientImpl) List(vmParam string) ([]Summary, error) {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(listInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []Summary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []Summary
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]Summary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Get(vmParam string, floppyParam string) (Info, error) {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Floppy", floppyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Create(vmParam string, specParam CreateSpec) (string, error) {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(createInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := createRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), createOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Update(vmParam string, floppyParam string, specParam UpdateSpec) error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(updateInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Floppy", floppyParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := updateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Delete(vmParam string, floppyParam string) error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(deleteInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Floppy", floppyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Connect(vmParam string, floppyParam string) error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "connect")
	sv := bindings.NewStructValueBuilder(connectInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Floppy", floppyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := connectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (fIface *FloppyClientImpl) Disconnect(vmParam string, floppyParam string) error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "disconnect")
	sv := bindings.NewStructValueBuilder(disconnectInputType(), typeConverter)
	sv.AddStructField("Vm", vmParam)
	sv.AddStructField("Floppy", floppyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := disconnectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (fIface *FloppyClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := fIface.connector.GetApiProvider().Invoke(fIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (fIface *FloppyClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(createInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(createOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.UnableToAllocateResource error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.ResourceBusy error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError9 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.create method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError9).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(updateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(updateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.ResourceBusy error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.update method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(deleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(deleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.ResourceBusy error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.delete method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) connectMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(connectInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(connectOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "connect")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.ResourceBusy error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError9 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.connect method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError9).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FloppyClientImpl) disconnectMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(disconnectInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(disconnectOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "disconnect")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ResourceBusyBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.ResourceBusy error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
      fIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError9 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FloppyClientImpl.disconnect method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError9).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}


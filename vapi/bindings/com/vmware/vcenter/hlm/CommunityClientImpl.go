
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Community
 * Functions that implement the generated CommunityClient interface
 */


package hlm
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type CommunityClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewCommunityClientImpl(connector client.Connector) *CommunityClientImpl {
      interfaceName := "com.vmware.vcenter.hlm.community"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "add"),
          core.NewMethodIdentifier(interfaceIdentifier, "remove"),
          core.NewMethodIdentifier(interfaceIdentifier, "check"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      cIface := CommunityClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      cIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      cIface.methodNameToDefMap["get"] = cIface.getMethodDefinition()
      cIface.methodNameToDefMap["add"] = cIface.addMethodDefinition()
      cIface.methodNameToDefMap["remove"] = cIface.removeMethodDefinition()
      cIface.methodNameToDefMap["check"] = cIface.checkMethodDefinition()
      return &cIface
}

func (cIface *CommunityClientImpl) Get() (CommunityInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(communityGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput CommunityInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := communityGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput CommunityInfo
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), communityGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(CommunityInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CommunityClientImpl) Add(specParam CommunityAddSpec) error {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "add")
	sv := bindings.NewStructValueBuilder(communityAddInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := communityAddRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (cIface *CommunityClientImpl) Remove(hostnameParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "remove")
	sv := bindings.NewStructValueBuilder(communityRemoveInputType(), typeConverter)
	sv.AddStructField("Hostname", hostnameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := communityRemoveRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (cIface *CommunityClientImpl) Check(specParam *CommunityAddCheckSpec) ([]CommunityCheckInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "check")
	sv := bindings.NewStructValueBuilder(communityCheckInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []CommunityCheckInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := communityCheckRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []CommunityCheckInfo
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), communityCheckOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]CommunityCheckInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *CommunityClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := cIface.connector.GetApiProvider().Invoke(cIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (cIface *CommunityClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(communityGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(communityGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CommunityClientImpl) addMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(communityAddInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(communityAddOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "add")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.AlreadyExists error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.UnverifiedPeer error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.add method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CommunityClientImpl) removeMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(communityRemoveInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(communityRemoveOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.remove method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.remove method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "remove")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.remove method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.remove method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CommunityClientImpl) checkMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(communityCheckInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(communityCheckOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "check")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's errors.UnverifiedPeer error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommunityClientImpl.check method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



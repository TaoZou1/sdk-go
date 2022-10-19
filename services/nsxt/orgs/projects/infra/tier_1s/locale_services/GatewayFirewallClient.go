// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GatewayFirewall
// Used by client-side stubs.

package locale_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type GatewayFirewallClient interface {

	// Get filtered view of Gateway Firewall rules associated with the Tier-1 Locale Services. The gateway policies are returned in the order of category and sequence number.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param tier1IdParam (required)
	// @param localeServicesIdParam (required)
	// @return com.vmware.nsx_policy.model.GatewayPolicyListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, tier1IdParam string, localeServicesIdParam string) (model.GatewayPolicyListResult, error)
}

type gatewayFirewallClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewGatewayFirewallClient(connector client.Connector) *gatewayFirewallClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.tier_1s.locale_services.gateway_firewall")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	gIface := gatewayFirewallClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *gatewayFirewallClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *gatewayFirewallClient) List(orgIdParam string, projectIdParam string, tier1IdParam string, localeServicesIdParam string) (model.GatewayPolicyListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(gatewayFirewallListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("LocaleServicesId", localeServicesIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GatewayPolicyListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := gatewayFirewallListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.tier_1s.locale_services.gateway_firewall", "list", inputDataValue, executionContext)
	var emptyOutput model.GatewayPolicyListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), gatewayFirewallListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GatewayPolicyListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
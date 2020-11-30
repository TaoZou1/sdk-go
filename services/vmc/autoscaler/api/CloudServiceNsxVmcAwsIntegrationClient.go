/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceNsxVmcAwsIntegration
 * Used by client-side stubs.
 */

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type CloudServiceNsxVmcAwsIntegrationClient interface {

    // Get the user-level SDDC configuration parameters
    // @return com.vmware.model.SddcUserConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetSddcUserConfig() (model.SddcUserConfiguration, error)
}

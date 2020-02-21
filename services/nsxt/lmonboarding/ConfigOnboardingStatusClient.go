/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ConfigOnboardingStatus
 * Used by client-side stubs.
 */

package lmonboarding

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ConfigOnboardingStatusClient interface {

    // Get on-boarding status.
    // @return com.vmware.nsx_policy.model.PolicyLMOnboardingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PolicyLMOnboardingStatus, error)
}

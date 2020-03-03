/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DeploymentState
 * Used by client-side stubs.
 */

package global_managers

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type DeploymentStateClient interface {

    // Read Global Manager deployment state.
    //
    // @param gmIdParam (required)
    // @return com.vmware.nsx_policy.model.GlobalManagerDeploymentState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(gmIdParam string) (model.GlobalManagerDeploymentState, error)
}

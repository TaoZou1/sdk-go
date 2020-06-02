/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Run
 * Used by client-side stubs.
 */

package xlb

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type RunClient interface {

    // Execute cross-cluster load balancing operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string) (model.Task, error)
}

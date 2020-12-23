/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package crl_distribution_points

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StatusClient interface {

    // Return the status of the CrlDistributionPoint
    //
    // @param crlDistributionPointIdParam (required)
    // @return com.vmware.nsx_global_policy.model.CrlDistributionPointStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(crlDistributionPointIdParam string) (model.CrlDistributionPointStatus, error)
}

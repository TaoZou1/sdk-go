/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package l2vpns

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param l2vpnIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.AggregateL2VpnStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, l2vpnIdParam string, enforcementPointPathParam *string) (model.AggregateL2VpnStatistics, error)
}

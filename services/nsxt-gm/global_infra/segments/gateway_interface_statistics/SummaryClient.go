/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Summary
 * Used by client-side stubs.
 */

package gateway_interface_statistics

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type SummaryClient interface {

    // Segment ID is the ID of the segment that is connected to the the tier-0
    //
    // @param segmentIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_global_policy.model.PolicyInterfaceStatisticsSummary
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(segmentIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.PolicyInterfaceStatisticsSummary, error)
}

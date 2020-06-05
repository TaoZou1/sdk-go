/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DetailedStatus
 * Used by client-side stubs.
 */

package lb_virtual_servers

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type DetailedStatusClient interface {

    // Get LBVirtualServer detailed status information. - no enforcement point path specified: Information will be aggregated from each enforcement point. - {enforcement_point_path}: Information will be retrieved only from the given enforcement point.
    //
    // @param lbServiceIdParam LBService id (required)
    // @param lbVirtualServerIdParam LBVirtualServer id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.AggregateLBVirtualServerStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(lbServiceIdParam string, lbVirtualServerIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.AggregateLBVirtualServerStatus, error)
}

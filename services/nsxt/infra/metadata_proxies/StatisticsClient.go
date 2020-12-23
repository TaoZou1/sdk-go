/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package metadata_proxies

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // Get metadata proxy status
    //
    // @param metadataProxyIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param segmentPathParam String Path of the segment which is associated with this metadata proxy (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.PolicyMetadataProxyStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(metadataProxyIdParam string, enforcementPointPathParam *string, segmentPathParam *string, sourceParam *string) (model.PolicyMetadataProxyStatistics, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package tag_operations

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StatusClient interface {

    // Get status of tag bulk operation with details of tag operation on each virtual machine.
    //
    // @param operationIdParam (required)
    // @return com.vmware.nsx_global_policy.model.TagBulkOperationStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(operationIdParam string) (model.TagBulkOperationStatus, error)
}

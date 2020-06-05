/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: NsxPartialPatchConfig
 * Used by client-side stubs.
 */

package system_config

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type NsxPartialPatchConfigClient interface {

    // Get Configuration values for nsx-partial-patch. By default partial patch is disbaled (i.e false).
    // @return com.vmware.nsx_global_policy.model.PartialPatchConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PartialPatchConfig, error)

    // Update partial patch configuration values. Only boolean value is allowed for enable_partial_patch
    //
    // @param partialPatchConfigParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(partialPatchConfigParam model.PartialPatchConfig) error
}

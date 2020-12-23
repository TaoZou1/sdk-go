/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: FeaturesWithProperties
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type FeaturesWithPropertiesClient interface {

    // List features
    // @return com.vmware.nsx_policy.model.FeaturePermissionListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.FeaturePermissionListResult, error)
}

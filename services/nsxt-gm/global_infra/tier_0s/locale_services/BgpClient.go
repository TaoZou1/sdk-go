/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Bgp
 * Used by client-side stubs.
 */

package locale_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type BgpClient interface {

    // Read BGP routing config
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @return com.vmware.nsx_global_policy.model.BgpRoutingConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string) (model.BgpRoutingConfig, error)

    // If an BGP routing config not present, create BGP routing config. If it already exists, update the routing config.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param bgpRoutingConfigParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, bgpRoutingConfigParam model.BgpRoutingConfig, overrideParam *bool) error

    // If BGP routing config is not already present, create BGP routing config. If it already exists, replace the BGP routing config with this object.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param bgpRoutingConfigParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @return com.vmware.nsx_global_policy.model.BgpRoutingConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, bgpRoutingConfigParam model.BgpRoutingConfig, overrideParam *bool) (model.BgpRoutingConfig, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MemberTypes
 * Used by client-side stubs.
 */

package groups

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type MemberTypesClient interface {

    // It retrieves member types for a given group. In case of nested groups, it calculates member types of child groups as well. Considers member type for members added via static members and dynamic membership criteria.
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @return com.vmware.nsx_global_policy.model.GroupMemberTypeListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string) (model.GroupMemberTypeListResult, error)
}

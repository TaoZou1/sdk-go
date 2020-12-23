/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbIpAddrGroups
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type AlbIpAddrGroupsClient interface {

    // Delete the ALBIpAddrGroup along with all the entities contained by this ALBIpAddrGroup.
    //
    // @param albIpaddrgroupIdParam ALBIpAddrGroup ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albIpaddrgroupIdParam string, forceParam *bool) error

    // Read a ALBIpAddrGroup.
    //
    // @param albIpaddrgroupIdParam ALBIpAddrGroup ID (required)
    // @return com.vmware.nsx_global_policy.model.ALBIpAddrGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albIpaddrgroupIdParam string) (model.ALBIpAddrGroup, error)

    // Paginated list of all ALBIpAddrGroup for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.ALBIpAddrGroupApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBIpAddrGroupApiResponse, error)

    // If a ALBipaddrgroup with the alb-ipaddrgroup-id is not already present, create a new ALBipaddrgroup. If it already exists, update the ALBipaddrgroup. This is a full replace.
    //
    // @param albIpaddrgroupIdParam ALBipaddrgroup ID (required)
    // @param aLBIpAddrGroupParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albIpaddrgroupIdParam string, aLBIpAddrGroupParam model.ALBIpAddrGroup) error

    // If a ALBIpAddrGroup with the alb-IpAddrGroup-id is not already present, create a new ALBIpAddrGroup. If it already exists, update the ALBIpAddrGroup. This is a full replace.
    //
    // @param albIpaddrgroupIdParam ALBIpAddrGroup ID (required)
    // @param aLBIpAddrGroupParam (required)
    // @return com.vmware.nsx_global_policy.model.ALBIpAddrGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albIpaddrgroupIdParam string, aLBIpAddrGroupParam model.ALBIpAddrGroup) (model.ALBIpAddrGroup, error)
}

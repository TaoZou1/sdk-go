/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbApplicationProfiles
 * Used by client-side stubs.
 */

package global_infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbApplicationProfilesClient interface {

    // Delete the ALBApplicationProfile along with all the entities contained by this ALBApplicationProfile.
    //
    // @param albApplicationprofileIdParam ALBApplicationProfile ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albApplicationprofileIdParam string, forceParam *bool) error

    // Read a ALBApplicationProfile.
    //
    // @param albApplicationprofileIdParam ALBApplicationProfile ID (required)
    // @return com.vmware.nsx_policy.model.ALBApplicationProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albApplicationprofileIdParam string) (model.ALBApplicationProfile, error)

    // Paginated list of all ALBApplicationProfile for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBApplicationProfileApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBApplicationProfileApiResponse, error)

    // If a ALBapplicationprofile with the alb-applicationprofile-id is not already present, create a new ALBapplicationprofile. If it already exists, update the ALBapplicationprofile. This is a full replace.
    //
    // @param albApplicationprofileIdParam ALBapplicationprofile ID (required)
    // @param aLBApplicationProfileParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albApplicationprofileIdParam string, aLBApplicationProfileParam model.ALBApplicationProfile) error

    // If a ALBApplicationProfile with the alb-ApplicationProfile-id is not already present, create a new ALBApplicationProfile. If it already exists, update the ALBApplicationProfile. This is a full replace.
    //
    // @param albApplicationprofileIdParam ALBApplicationProfile ID (required)
    // @param aLBApplicationProfileParam (required)
    // @return com.vmware.nsx_policy.model.ALBApplicationProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albApplicationprofileIdParam string, aLBApplicationProfileParam model.ALBApplicationProfile) (model.ALBApplicationProfile, error)
}

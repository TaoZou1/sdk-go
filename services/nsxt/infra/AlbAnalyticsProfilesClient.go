/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbAnalyticsProfiles
 * Used by client-side stubs.
 */

package infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbAnalyticsProfilesClient interface {

    // Delete the ALBAnalyticsProfile along with all the entities contained by this ALBAnalyticsProfile.
    //
    // @param albAnalyticsprofileIdParam ALBAnalyticsProfile ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albAnalyticsprofileIdParam string, forceParam *bool) error

    // Read a ALBAnalyticsProfile.
    //
    // @param albAnalyticsprofileIdParam ALBAnalyticsProfile ID (required)
    // @return com.vmware.nsx_policy.model.ALBAnalyticsProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albAnalyticsprofileIdParam string) (model.ALBAnalyticsProfile, error)

    // Paginated list of all ALBAnalyticsProfile for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBAnalyticsProfileApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBAnalyticsProfileApiResponse, error)

    // If a ALBanalyticsprofile with the alb-analyticsprofile-id is not already present, create a new ALBanalyticsprofile. If it already exists, update the ALBanalyticsprofile. This is a full replace.
    //
    // @param albAnalyticsprofileIdParam ALBanalyticsprofile ID (required)
    // @param aLBAnalyticsProfileParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albAnalyticsprofileIdParam string, aLBAnalyticsProfileParam model.ALBAnalyticsProfile) error

    // If a ALBAnalyticsProfile with the alb-AnalyticsProfile-id is not already present, create a new ALBAnalyticsProfile. If it already exists, update the ALBAnalyticsProfile. This is a full replace.
    //
    // @param albAnalyticsprofileIdParam ALBAnalyticsProfile ID (required)
    // @param aLBAnalyticsProfileParam (required)
    // @return com.vmware.nsx_policy.model.ALBAnalyticsProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albAnalyticsprofileIdParam string, aLBAnalyticsProfileParam model.ALBAnalyticsProfile) (model.ALBAnalyticsProfile, error)
}

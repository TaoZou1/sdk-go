/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbPkiProfiles
 * Used by client-side stubs.
 */

package global_infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbPkiProfilesClient interface {

    // Delete the ALBPKIProfile along with all the entities contained by this ALBPKIProfile.
    //
    // @param albPkiprofileIdParam ALBPKIProfile ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albPkiprofileIdParam string, forceParam *bool) error

    // Read a ALBPKIProfile.
    //
    // @param albPkiprofileIdParam ALBPKIProfile ID (required)
    // @return com.vmware.nsx_policy.model.ALBPKIProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albPkiprofileIdParam string) (model.ALBPKIProfile, error)

    // Paginated list of all ALBPKIProfile for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBPKIProfileApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBPKIProfileApiResponse, error)

    // If a ALBpkiprofile with the alb-pkiprofile-id is not already present, create a new ALBpkiprofile. If it already exists, update the ALBpkiprofile. This is a full replace.
    //
    // @param albPkiprofileIdParam ALBpkiprofile ID (required)
    // @param aLBPKIProfileParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albPkiprofileIdParam string, aLBPKIProfileParam model.ALBPKIProfile) error

    // If a ALBPKIProfile with the alb-PKIProfile-id is not already present, create a new ALBPKIProfile. If it already exists, update the ALBPKIProfile. This is a full replace.
    //
    // @param albPkiprofileIdParam ALBPKIProfile ID (required)
    // @param aLBPKIProfileParam (required)
    // @return com.vmware.nsx_policy.model.ALBPKIProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albPkiprofileIdParam string, aLBPKIProfileParam model.ALBPKIProfile) (model.ALBPKIProfile, error)
}

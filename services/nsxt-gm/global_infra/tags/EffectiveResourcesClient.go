/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EffectiveResources
 * Used by client-side stubs.
 */

package tags

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type EffectiveResourcesClient interface {

    // Paginated list of all objects assigned with matching scope and tag values. Objects are represented in form of resource reference.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param filterTextParam Filter text to restrict tagged objects list with matching filter text. (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param scopeParam Tag scope (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param tagParam Tag value (optional)
    // @return com.vmware.nsx_global_policy.model.PolicyResourceReferenceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, filterTextParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, scopeParam *string, sortAscendingParam *bool, sortByParam *string, tagParam *string) (model.PolicyResourceReferenceListResult, error)
}

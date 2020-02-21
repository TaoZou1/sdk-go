/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UrlReputationSeverities
 * Used by client-side stubs.
 */

package global_infra

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type UrlReputationSeveritiesClient interface {

    // Gets a URL reputation severity. The min_reputation and max_reputation specify the range of the reputations which belong to a particular severity. For instance, any reputation between 1 to 20 belongs to the severity 'High Risk'. Similary a reputation between 81 to 100 belong to the severity 'Trustworthy'.
    //
    // @param reputationSeverityIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyUrlReputationSeverity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(reputationSeverityIdParam string) (model.PolicyUrlReputationSeverity, error)

    // Gets the list of reputation severities. This will provide all the supported severities along with their ids, min and max reputaitons. The min_reputation and max_reputation specify the range of the reputations which belong to a particular severity. For instance, any reputation between 1 to 20 belongs to the severity 'High Risk'. Similary a reputation between 81 to 100 belong to the severity 'Trustworthy'.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyUrlReputationSeverityListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyUrlReputationSeverityListResult, error)
}

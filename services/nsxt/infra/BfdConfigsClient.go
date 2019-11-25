/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: BfdConfigs
 * Used by client-side stubs.
 */

package infra

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type BfdConfigsClient interface {

    // Delete BFD Config and all the entities contained by this BfdConfiguration.
    //
    // @param bfdConfigIdParam BfdConfiguration ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(bfdConfigIdParam string) error

    // Read a BfdConfiguration.
    //
    // @param bfdConfigIdParam BfdConfiguration ID (required)
    // @return com.vmware.nsx_policy.model.BfdConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(bfdConfigIdParam string) (model.BfdConfiguration, error)

    // Paginated list of all BfdConfigurations.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.BfdConfigurationListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BfdConfigurationListResult, error)

    // If a BfdConfiguration with the bfd-config-id is not already present, create a new BfdConfiguration. If it already exists, update the BfdConfiguration. This operation will fully replace the object.
    //
    // @param bfdConfigIdParam BfdConfiguration ID (required)
    // @param bfdConfigurationParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(bfdConfigIdParam string, bfdConfigurationParam model.BfdConfiguration) error

    // If a BfdConfiguration with the bfd-config-id is not already present, create a new BfdConfiguration. If it already exists, update the BfdConfiguration. This operation will fully replace the object.
    //
    // @param bfdConfigIdParam BfdConfiguration ID (required)
    // @param bfdConfigurationParam (required)
    // @return com.vmware.nsx_policy.model.BfdConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(bfdConfigIdParam string, bfdConfigurationParam model.BfdConfiguration) (model.BfdConfiguration, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortMonitoringProfileBindingMaps
 * Used by client-side stubs.
 */

package ports

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortMonitoringProfileBindingMapsClient interface {

    // API will delete Port Monitoring Profile Binding Profile.
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier1IdParam string, segmentIdParam string, portIdParam string, portMonitoringProfileBindingMapIdParam string) error

    // API will get Port Monitoring Profile Binding Map.
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, portIdParam string, portMonitoringProfileBindingMapIdParam string) (model.PortMonitoringProfileBindingMap, error)

    // API will list all Port Monitoring Profile Binding Maps in current port id.
    //
    // @param tier1IdParam (required)
    // @param segmentIdParam (required)
    // @param portIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortMonitoringProfileBindingMapListResult, error)

    // API will create Port Monitoring Profile Binding Map.
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @param portMonitoringProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier1IdParam string, segmentIdParam string, portIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam model.PortMonitoringProfileBindingMap) error

    // API will update Port Monitoring Profile Binding Map.
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @param portMonitoringProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier1IdParam string, segmentIdParam string, portIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam model.PortMonitoringProfileBindingMap) (model.PortMonitoringProfileBindingMap, error)
}

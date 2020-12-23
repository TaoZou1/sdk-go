/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortQosProfileBindingMaps
 * Used by client-side stubs.
 */

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortQosProfileBindingMapsClient interface {

    // API will delete Port QoS Profile Binding Profile.
    //
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portQosProfileBindingMapIdParam Port QoS Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(segmentIdParam string, portIdParam string, portQosProfileBindingMapIdParam string) error

    // API will get Port QoS Profile Binding Map.
    //
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portQosProfileBindingMapIdParam Port QoS Profile Binding Map ID (required)
    // @return com.vmware.nsx_policy.model.PortQoSProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(segmentIdParam string, portIdParam string, portQosProfileBindingMapIdParam string) (model.PortQosProfileBindingMap, error)

    // API will list all Port QoS Profile Binding Maps in current port id.
    //
    // @param segmentIdParam (required)
    // @param portIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortQoSProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortQosProfileBindingMapListResult, error)

    // API will create Port QoS Profile Binding Map. For objects with no binding maps, default profile is applied.
    //
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portQosProfileBindingMapIdParam Port QoS Profile Binding Map ID (required)
    // @param portQosProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(segmentIdParam string, portIdParam string, portQosProfileBindingMapIdParam string, portQosProfileBindingMapParam model.PortQosProfileBindingMap) error

    // API will update Port QoS Profile Binding Map. For objects with no binding maps, default profile is applied.
    //
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portQosProfileBindingMapIdParam Port QoS Profile Binding Map ID (required)
    // @param portQosProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PortQoSProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(segmentIdParam string, portIdParam string, portQosProfileBindingMapIdParam string, portQosProfileBindingMapParam model.PortQosProfileBindingMap) (model.PortQosProfileBindingMap, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: BfdPeers
 * Used by client-side stubs.
 */

package static_routes

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type BfdPeersClient interface {

    // Delete this StaticRouteBfdPeer and all the entities contained by it.
    //
    // @param tier0IdParam Tier-0 ID (required)
    // @param bfdPeerIdParam BFD peer ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, bfdPeerIdParam string) error

    // Read a StaticRouteBfdPeer with the bfd-peer-id.
    //
    // @param tier0IdParam Tier-0 ID (required)
    // @param bfdPeerIdParam BFD peer ID (required)
    // @return com.vmware.nsx_policy.model.StaticRouteBfdPeer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, bfdPeerIdParam string) (model.StaticRouteBfdPeer, error)

    // Paginated list of all StaticRouteBfdPeers.
    //
    // @param tier0IdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.StaticRouteBfdPeerListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.StaticRouteBfdPeerListResult, error)

    // If a StaticRouteBfdPeer with the bfd-peer-id is not already present, create a new StaticRouteBfdPeer. If it already exists, update the StaticRouteBfdPeer. This is a full replace.
    //
    // @param tier0IdParam Tier-0 ID (required)
    // @param bfdPeerIdParam BFD peer ID (required)
    // @param staticRouteBfdPeerParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, bfdPeerIdParam string, staticRouteBfdPeerParam model.StaticRouteBfdPeer) error

    // If a StaticRouteBfdPeer with the bfd-peer-id is not already present, create a new StaticRouteBfdPeer. If it already exists, update the StaticRouteBfdPeer. This operation will fully replace the object.
    //
    // @param tier0IdParam Tier-0 ID (required)
    // @param bfdPeerIdParam BFD peer ID (required)
    // @param staticRouteBfdPeerParam (required)
    // @return com.vmware.nsx_policy.model.StaticRouteBfdPeer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, bfdPeerIdParam string, staticRouteBfdPeerParam model.StaticRouteBfdPeer) (model.StaticRouteBfdPeer, error)
}

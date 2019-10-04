/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Leaseinfo
 * Used by client-side stubs.
 */

package leaseinfo

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type LeaseinfoClient interface {


    // Retrieve DHCP leaseinfo of a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @return com.vmware.vmc.model.DhcpLeases
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, edgeIdParam string) (model.DhcpLeases, error) 

}
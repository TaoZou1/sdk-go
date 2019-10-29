/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VcenterAppliance
 * Used by client-side stubs.
 */

package nsx


// The ``VcenterAppliance`` interface provides methods to query the configuration of the vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcenterApplianceClient interface {

    // Gets the configuration details of current vCenter Server Appliance. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param managementVcenterParam vCenter connection details for Uber management vCenter.
    // If map with bool value the vCenter is an Uber managed
    // @return ApplianceInfo for the vCenter appliance.
    // @throws NotFound if vCenter appliance could not be located.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
	Get(managementVcenterParam *Connection) (VcenterApplianceApplianceInfo, error)
}

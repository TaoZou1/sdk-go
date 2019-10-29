/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: License
 * Used by client-side stubs.
 */

package vcha


// The ``License`` interface provides methods to get the license of a vCenter High Availability (VCHA) feature. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type LicenseClient interface {

    // Gets the license of the active node of a VCHA cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Info structure containing the VCHA license.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
	Get() (LicenseInfo, error)
}

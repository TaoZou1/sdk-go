/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Identity
 * Used by client-side stubs.
 */

package guest

import (
)

// The ``Identity`` interface provides methods for retrieving guest operating system identification information.
type IdentityClient interface {


    // Return information about the guest.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return guest identification information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws ServiceUnavailable if VMware Tools has not provided any data.
    Get(vmParam string) (IdentityInfo, error) 

}

/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package VM

import (
)

// The ``VM`` interface provides methods managing the storage policy association for a virtual machine and its virtual disks.
type VMClient interface {


    // Returns information about the virtual machines and/or their virtual disks that are associated with the given storage policy.
    //
    // @param policyParam storage policy identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.StoragePolicy``.
    // @return Information about the virtual machines and/or their virtual disks that are associated with the given storage policy.
    // The key in the return value map will be an identifier for the resource type: ``VirtualMachine``.
    // @throws NotFound if there is no policy associated with ``policy`` in the system.
    // @throws UnableToAllocateResource if more than 1000 virtual machines are associated with the specified policy.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws ServiceUnavailable if the system is unable to communicate with a service necessary to complete the request.
    // @throws Error if the backend server encounters some an error while processing the request.
    // @throws Unauthorized if the user does not have the required priveleges.
    List(policyParam string) (map[string]Info, error) 

}

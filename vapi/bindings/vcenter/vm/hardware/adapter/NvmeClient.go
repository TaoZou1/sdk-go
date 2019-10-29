/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Nvme
 * Used by client-side stubs.
 */

package adapter


// The ``Nvme`` interface provides methods for configuring the virtual NVMe adapters of a virtual machine. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeClient interface {

    // Returns commonly used information about the virtual NVMe adapters belonging to the virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual NVMe adapters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(vmParam string) ([]NvmeSummary, error)

    // Returns information about a virtual NVMe adapter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual NVMe adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.NvmeAdapter``.
    // @return Information about the specified virtual NVMe adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual NVMe adapter is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string, adapterParam string) (NvmeInfo, error)

    // Adds a virtual NVMe adapter to the virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual NVMe adapter.
    // @return Virtual NVMe adapter identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.NvmeAdapter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Error if the system reported that the NVMe adapter was created but was unable to confirm the creation because the identifier of the new adapter could not be determined.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine is not found.
    // @throws UnableToAllocateResource if there are no more available NVMe buses on the virtual machine.
    // @throws ResourceInUse if the specified NVMe bus or PCI address is in use.
    // @throws InvalidArgument if the specified NVMe bus or PCI address is out of bounds.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
	Create(vmParam string, specParam NvmeCreateSpec) (string, error)

    // Removes a virtual NVMe adapter from the virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual NVMe adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.NvmeAdapter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine or virtual NVMe adapter is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(vmParam string, adapterParam string) error
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Service
 * Used by client-side stubs.
 */

package authentication


// The ``Service`` interface provides methods to retrieve authentication information of a service element. 
//
//  A service element is said to contain authentication information if there is a default authentication assigned to all operation elements contained in a service element or if one of the operation elements contained in this service element has authentication information.
type ServiceClient interface {

    // Returns the identifiers for the service elements that have authentication information.
    // @return The list of identifiers for the service elements that have authentication information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
	List() ([]string, error)

    // Retrieves authentication information about the service element corresponding to ``service_id``.
    //
    // @param serviceIdParam Identifier of the service element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return The ServiceInfo instance that corresponds to ``service_id``
    // @throws NotFound if the service element associated with ``service_id`` does not have any authentication information.
	Get(serviceIdParam string) (ServiceInfo, error)
}

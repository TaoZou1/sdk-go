/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Events
 * Used by client-side stubs.
 */

package events


// The ``Events`` interface provides methods to get Kubernetes events related to a particular cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EventsClient interface {

    // Returns Kubernetes events related to a specific cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which Namespaces are enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return List of latest Kubernetes events by 'lastTimestamp', with a limit of 1000 events.
    // @throws NotFound if cluster could not be located.
    // @throws Unsupported if the specified cluster does not have Namespaces enabled.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
	Get(clusterParam string) ([]EventsEvent, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Compliance
 * Used by client-side stubs.
 */
package policies


// The Compliance interface provides methods related to all the associated entities of given compliance statuses.
type ComplianceClient interface {

    // Returns compliance information about entities matching the filter ComplianceFilterSpec. Entities without storage policy association are not returned.
    //
    // @param filterParam compliance status of matching entities for which information should be returned.
    // @return compliance information about entities matching the filter ComplianceFilterSpec.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the ComplianceFilterSpec#status property contains a value that is not supported by the server.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam ComplianceFilterSpec) ([]ComplianceSummary, error)
}

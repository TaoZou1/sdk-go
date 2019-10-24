/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */
package vstats


// The ``Providers`` interface manages list of statistical data provider services that are currently used.
type ProvidersClient interface {

    // Returns a list of stats providers.
    // @return List of stats providers.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List() ([]ProvidersSummary, error)
}

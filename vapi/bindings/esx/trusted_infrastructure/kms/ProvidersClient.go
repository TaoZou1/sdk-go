/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package kms


// The ``Providers`` interface provides methods to retrieve available key providers.
type ProvidersClient interface {

    // Return the available providers. 
    //
    //  Contacts the KMS service to get the list of providers available for key operation. 
    //
    // @param filterSpecParam 
    // @return Summary of providers.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Unauthorized If the caller is not authorized.
    // @throws Error For any other error.
	List(filterSpecParam *ProvidersFilterSpec) ([]ProvidersSummary, error)
}

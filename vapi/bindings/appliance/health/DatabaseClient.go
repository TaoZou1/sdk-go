/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Database
 * Used by client-side stubs.
 */

package health


// The ``Database`` interface provides methods to retrieve the health status of the vcdb. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DatabaseClient interface {

    // Returns the health status of the database. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Health status of the database
    // @throws Error if issue in retrieving health of the database
	Get() (DatabaseInfo, error)
}

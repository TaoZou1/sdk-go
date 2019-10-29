/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: RetentionPlans
 * Used by client-side stubs.
 */

package vstats


// The ``RetentionPlans`` interface manages the retention settings of the product.
type RetentionPlansClient interface {

    // Returns the default retention plan.
    // @return Retention plan information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if no default retention plan is found.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	GetDefault() (RetentionPlansInfo, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Time
 * Used by client-side stubs.
 */

package system

import (
)

// ``Time`` interface provides methods Gets system time.
type TimeClient interface {


    // Get system time.
    // @return System time
    // @throws Error Generic error
    Get() (TimeSystemTimeStruct, error) 

}

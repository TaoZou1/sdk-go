/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Mem
 * Used by client-side stubs.
 */

package health

import (
)

// ``Mem`` interface provides methods Get memory health.
type MemClient interface {


    // Get memory health.
    // @return Memory health.
    // @throws Error Generic error
    Get() (MemHealthLevel, error) 

}

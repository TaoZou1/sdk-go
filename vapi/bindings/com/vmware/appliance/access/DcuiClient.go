/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Dcui
 * Used by client-side stubs.
 */

package access

import (
)

// ``Dcui`` interface provides methods Get/Set enabled state of DCUI.
type DcuiClient interface {


    // Set enabled state of Direct Console User Interface (DCUI TTY2).
    //
    // @param enabledParam DCUI is enabled.
    // @throws Error Generic error
    Set(enabledParam bool) error 


    // Get enabled state of Direct Console User Interface (DCUI TTY2).
    // @return DCUI is enabled.
    // @throws Error Generic error
    Get() (bool, error) 

}

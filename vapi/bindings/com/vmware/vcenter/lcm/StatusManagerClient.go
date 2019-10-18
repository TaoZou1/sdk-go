/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: StatusManager
 * Used by client-side stubs.
 */

package lcm

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
)

// The service to install Embedded VCSA, PSC, Management VCSA, VMC gateway. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type StatusManagerClient interface {


    // Retrieves the status of the appliance. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Check() (std.LocalizableMessage, error) 

}

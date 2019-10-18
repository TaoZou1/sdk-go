/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VmcaRoot
 * Used by client-side stubs.
 */

package vcenter

import (
)

// The ``VmcaRoot`` interface provides methods to replace VMware Certificate Authority (VMCA) root certificate. This interface was added in vSphere API 6.9.1.
type VmcaRootClient interface {


    // Replace Root Certificate with VMCA signed one using the given Spec. 
    //
    // After this method completes, the services using the certificate will be restarted for the new certificate to take effect.. This method was added in vSphere API 6.9.1.
    //
    // @param specParam The information needed to generate VMCA signed Root Certificate.
    // Default values will be set for all null parameters.
    // @throws Error If the system failed to renew the TLS certificate.
    Create(specParam *VmcaRootCreateSpec) error 

}

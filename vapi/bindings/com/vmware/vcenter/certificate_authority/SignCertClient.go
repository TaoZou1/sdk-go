/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SignCert
 * Used by client-side stubs.
 */

package certificate_authority

import (
)

// The ``SignCert`` interface provides methods to generate certificate using CSR (certificate signing request). **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SignCertClient interface {


    // Sign the provided CSR and generate a certificate. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param csrParam The information needed to create certificate
    // @param durationParam The duration (in days) of the new TLS certificate. The duration should be less than or equal to 730 days..
    // If null, the duration will be 730 days (two years).
    // @return certificate in String format.
    // @throws InvalidArgument if CSR is not valid or duration is more than 10 years.
    // @throws Error If the system failed to sign the CSR.
    SignCertFromCSR(csrParam string, durationParam *int64) (string, error) 

}

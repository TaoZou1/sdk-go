/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Job
 * Used by client-side stubs.
 */

package reconciliation

import (
)

// The ``Job`` interface provides methods to create and get the status of reconciliation job.
type JobClient interface {


    // Initiate reconciliation.
    //
    // @param specParam CreateSpec Structure
    // @return Info Structure
    // @throws FeatureInUse A backup or restore is already in progress.
    // @throws NotAllowedInCurrentState Reconciliation is allowed only after restore has finished successfully.
    // @throws Error if any error occurs during the execution of the operation.
    Create(specParam JobCreateSpec) (JobInfo, error) 


    // Get reconciliation job progress/result.
    // @return Info Structure
    // @throws NotFound if there is no running reconciliation job.
    // @throws Error if any error occurs during the execution of the operation.
    Get() (JobInfo, error) 

}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Metrics
 * Used by client-side stubs.
 */

package vstats

import (
)

// The ``Metrics`` interface provides method to list metrics. A metric is a fundamental concept. It means a measurable aspect or property. For instance, temperature, count, velocity, data size, bandwidth.
type MetricsClient interface {


    // Returns list of available Metrics as supplied by the discovered providers.
    // @return List of Metrics.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List() ([]MetricsSummary, error) 

}

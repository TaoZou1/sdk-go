/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ConnectivityTests
 * Used by client-side stubs.
 */

package connectivityTests

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ConnectivityTestsClient interface {


    // Retrieve metadata for connectivity tests.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.ConnectivityValidationGroups
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    Get(orgParam string, sddcParam string) (model.ConnectivityValidationGroups, error) 


    // ConnectivityValidationGroupResultWrapper will be available at task.params['test_result'].
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param requestInfoParam request information (required)
    // @param actionParam If = 'start', start connectivity tests. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Bad Request
    // @throws Unauthorized  Forbidden
    Post(orgParam string, sddcParam string, requestInfoParam model.ConnectivityValidationGroup, actionParam string) (model.Task, error) 

}

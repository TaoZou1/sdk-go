/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Roles
 * Used by client-side stubs.
 */

package aaa

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type RolesClient interface {

    // The role with id role-id is cloned and the new id, name and description are the ones provided in the request body.
    //
    // @param roleIdParam Role id (required)
    // @param newRoleParam (required)
    // @return com.vmware.nsx_policy.model.NewRole
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Clone(roleIdParam string, newRoleParam model.NewRole) (model.NewRole, error)

    // If a role is assigned to a role binding then the deletion of the role is not allowed. Precanned roles cannot be deleted.
    //
    // @param roleIdParam Custom role id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(roleIdParam string) error

    // Get role information
    //
    // @param roleParam Role Name (required)
    // @return com.vmware.nsx_policy.model.RoleWithFeatures
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(roleParam string) (model.RoleWithFeatures, error)

    // Get information about all roles
    // @return com.vmware.nsx_policy.model.RoleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.RoleListResult, error)

    // Creates a new role with id as role-id if there does not exist any role with id role-id, else updates the existing role.
    //
    // @param roleIdParam Custom role id (required)
    // @param roleWithFeaturesParam (required)
    // @return com.vmware.nsx_policy.model.RoleWithFeatures
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(roleIdParam string, roleWithFeaturesParam model.RoleWithFeatures) (model.RoleWithFeatures, error)

    // Validate the permissions of an incoming role. Also, recommend the permissions which need to be corrected.
    //
    // @param featurePermissionArrayParam (required)
    // @return com.vmware.nsx_policy.model.RecommendedFeaturePermissionListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Validate(featurePermissionArrayParam model.FeaturePermissionArray) (model.RecommendedFeaturePermissionListResult, error)
}

/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Enumeration
 * Used by client-side stubs.
 */

package enumeration

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
)

// The ``Enumeration`` interface provides methods to retrieve metamodel information about an enumeration element in the interface definition language. 
//
//  The ``Enumeration`` has a list of enumeration value elements.
type EnumerationClient interface {


    // Returns the identifiers for the enumeration elements that are contained in all the package elements, service elements and structure elements.
    // @return The list of identifiers for the enumeration elements.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.enumeration``.
    List() ([]string, error) 


    // Retrieves information about the enumeration element corresponding to ``enumeration_id``. 
    //
    //  The metamodel.EnumerationInfo contains the metamodel information about the enumeration value element contained in the enumeration element.
    //
    // @param enumerationIdParam Identifier of the enumeration element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.enumeration``.
    // @return The metamodel.EnumerationInfo instance that corresponds to ``enumeration_id``
    // @throws NotFound if the enumeration element associated with ``enumeration_id`` is not contained in any of the package elements, service elements and structure elements.
    Get(enumerationIdParam string) (metamodel.EnumerationInfo, error) 

}

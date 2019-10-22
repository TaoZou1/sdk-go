/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster_scale_in_ignore_vm_capabilities

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)



// The ``CreateSpec`` class contains information used to create a new policy to ignore virtual machine capabilities when scaling-in a cluster, see Policies#create. When considering scaling-in a cluster, policies that have been created with one of the listed CreateSpec#vmCapabilities are ignored for virtual machines that have the tag indicated by CreateSpec#vmTag in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CreateSpec struct {
    // When considering scaling-in a cluster, policies that have been created with one of the listed CreateSpec#vmCapabilities are ignored for virtual machines that have this tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    VmTag string
    // When considering scaling-in a cluster, policies that have been created with one of these capabilities are ignored for virtual machines that have the tag indicated by CreateSpec#vmTag. This map with bool value must contain at least one item. Currently, the only allowed capability identifier is ``com.vmware.vcenter.compute.policies.capabilities.disable_drs_vmotion``. In the future, other capabilities may be specified. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    VmCapabilities map[string]bool
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Capability string
    // Name of the policy. The name needs to be unique within this vCenter server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Description string
}



//


// The ``Info`` class contains information about a policy to ignore virtual machine capabilities when scaling-in a cluster, see Policies#get. When considering scaling-in a cluster, policies that have been created with one of the listed Info#vmCapabilities are ignored for virtual machines that have the tag indicated by Info#vmTag in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type Info struct {
    // When considering scaling-in a cluster, policies that have been created with one of the listed Info#vmCapabilities are ignored for virtual machines that have this tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    VmTag string
    // When considering scaling-in a cluster, policies that have been created with one of these capabilities are ignored for virtual machines that have the tag indicated by Info#vmTag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    VmCapabilities map[string]bool
    // Name of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Description string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Capability string
}



//






func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["vm_capabilities"] = bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability:VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["vm_capabilities"] = "VmCapabilities"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["vm_capabilities"] = bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability:VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["vm_capabilities"] = "VmCapabilities"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



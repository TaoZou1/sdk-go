/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.guest.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``HostnameGenerator`` class specifies various mechanisms of generating the hostname for the guest inside the virtual machine while customization. This class was added in vSphere API 7.0.0.
type HostnameGenerator struct {
    // The type of the Name Generator. This property was added in vSphere API 7.0.0.
	Type_ HostnameGeneratorType
    // The virtual machine name specified by the client. This property was added in vSphere API 7.0.0.
	FixedName *string
    // Base prefix, to which a unique number is appended. This property was added in vSphere API 7.0.0.
	Prefix *string
}

// The ``Type`` enumeration class specifies different types of Name Generators. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostnameGeneratorType string

const (
    // Specifies that a fixed name should be used for the hostname for the virtual machine during the customization. This constant field was added in vSphere API 7.0.0.
	HostnameGeneratorType_FIXED HostnameGeneratorType = "FIXED"
    // Specifies that a unique name should be generated by concatenating the base string with a number including the network name of the guest. 
    //
    //  Virtual machine names are unique across the set of hosts and virtual machines known to the VirtualCenter instance. VirtualCenter tracks the network names of virtual machines as well as hosts. VMware Tools runs in a guest operating system and reports information to VirtualCenter.. This constant field was added in vSphere API 7.0.0.
	HostnameGeneratorType_PREFIX HostnameGeneratorType = "PREFIX"
    // Specifies that the VirtualCenter should generate a virtual machine name. 
    //
    //  The name is generated from a base prefix comprising the virtual machine entity name. A number is appended, if necessary, to make it unique. 
    //
    //  Virtual machine names are unique across the set of hosts and virtual machines known to the VirtualCenter instance. VirtualCenter tracks the network names of the virtual machines as well as hosts. VMware Tools runs in a guest operating system and reports information to VirtualCenter, including the network name of the guest.. This constant field was added in vSphere API 7.0.0.
	HostnameGeneratorType_VIRTUAL_MACHINE HostnameGeneratorType = "VIRTUAL_MACHINE"
    // Indicates that the host name is not specified in advance. The user must provide a valid value when the actual customization spec is applied to a virtual machine. This constant field was added in vSphere API 7.0.0.
	HostnameGeneratorType_USER_INPUT_REQUIRED HostnameGeneratorType = "USER_INPUT_REQUIRED"
)

func (t HostnameGeneratorType) HostnameGeneratorType() bool {
	switch t {
	case HostnameGeneratorType_FIXED:
		return true
	case HostnameGeneratorType_PREFIX:
		return true
	case HostnameGeneratorType_VIRTUAL_MACHINE:
		return true
	case HostnameGeneratorType_USER_INPUT_REQUIRED:
		return true
	default:
		return false
	}
}


// The ``UserData`` class specifies the personal data pertaining to the user of the Windows guest operating system. This class maps to the UserData key in the sysprep.xml answer file. These values are transferred directly into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more detailed information about performing unattended installation, check https://technet.microsoft.com/en-us/library/cc771830(v=ws.10).aspx. This class was added in vSphere API 7.0.0.
type UserData struct {
    // The computer name of the (Windows) virtual machine. A computer name may contain letters (A-Z), numbers(0-9) and hyphens (-) but no spaces or periods (.). The name may not consist entirely of digits. A computer name is restricted to 15 characters in length. If the computer name is longer than 15 characters, it will be truncated to 15 characters. Check HostnameGenerator for various options. This property was added in vSphere API 7.0.0.
	ComputerName HostnameGenerator
    // Full name of the end user. Note that this is not the username but full name specified in "Firstname Lastname" format. This property was added in vSphere API 7.0.0.
	FullName string
    // Name of the organization that owns the computer. This property was added in vSphere API 7.0.0.
	Organization string
    // The product Key to use for activating Windows guest operating system. This property was added in vSphere API 7.0.0.
	ProductKey string
}

// The ``Domain`` class specifies the information needed to join a workgroup or domain. This structure maps to the Identification key in the sysprep.xml answer file. These values are transferred into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more information about performing unattended installation, check https://technet.microsoft.com/en-us/library/cc771830(v=ws.10).aspx. This class was added in vSphere API 7.0.0.
type Domain struct {
    // The type of network to join after the customization. This property was added in vSphere API 7.0.0.
	Type_ DomainType
    // The workgroup that the virtual machine should join. This property was added in vSphere API 7.0.0.
	Workgroup *string
    // The domain to which the virtual machine should be joined. This property was added in vSphere API 7.0.0.
	Domain *string
    // The domain user that has permission to join the domain after virtual machine is joined. This property was added in vSphere API 7.0.0.
	DomainUsername *string
    // The domain user password that has permission to join the Domain#domainUsername after customization. This property was added in vSphere API 7.0.0.
	DomainPassword *string
}

// The ``Type`` enumeration class defines the types of network the virtual machine should join to after the customization is completed. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DomainType string

const (
    // The virtual machine should be be joined to a workgroup. This constant field was added in vSphere API 7.0.0.
	DomainType_WORKGROUP DomainType = "WORKGROUP"
    // The virtual machine should be be joined to a domain. This constant field was added in vSphere API 7.0.0.
	DomainType_DOMAIN DomainType = "DOMAIN"
)

func (t DomainType) DomainType() bool {
	switch t {
	case DomainType_WORKGROUP:
		return true
	case DomainType_DOMAIN:
		return true
	default:
		return false
	}
}


// The ``GuiUnattended`` class maps to the GuiUnattended key in the sysprep.xml answer file. These values are plugged directly into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more detailed information about performing unattended installations, check https://technet.microsoft.com/en-us/library/cc771830(v=ws.10).aspx. This class was added in vSphere API 7.0.0.
type GuiUnattended struct {
    // Flag to determine whether or not the machine automatically logs on as Administrator. See the GuiUnattended#password property. If GuiUnattended#autoLogon flag is set, then GuiUnattended#password must not be null or the guest customization will fail. This property was added in vSphere API 7.0.0.
	AutoLogon bool
    // If the GuiUnattended#autoLogon flag is map with bool value, then this property specifies the number of times the machine should automatically log on as Administrator. Generally it should be 1, but if the setup requires a number of reboots, you may want to increase it. This number may be determined by the list of commands. This property was added in vSphere API 7.0.0.
	AutoLogonCount int64
    // The new administrator password for the machine. To specify that the password should be set to blank (that is, no password), leave it unset. This property was added in vSphere API 7.0.0.
	Password *string
    // The time zone index for the virtual machine. Numbers correspond to time zones at https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values. This property was added in vSphere API 7.0.0.
	TimeZone int64
}

// The ``WindowsSysprep`` class provides all the settings like user details, administrator details, etc that need to applied for a windows guest operating system during customization. This class was added in vSphere API 7.0.0.
type WindowsSysprep struct {
    // A list of commands to run at first user logon, after customizing the guest. These commands are directly mapped to the GuiRunOnce key in the sysprep.xml answer file. These commands are transferred into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more information about performing unattended installation, check https://technet.microsoft.com/en-us/library/cc771830(v=ws.10).aspx The commands listed here ar executed when a user logs on the first time after customization completes. The logon may be driven by GuiUnattended#autoLogon setting. These commands are directly mapped to the GuiRunOnce key in the. This property was added in vSphere API 7.0.0.
	GuiRunOnceCommands []string
    // Personal data pertaining to the owner of the virtual machine. This property was added in vSphere API 7.0.0.
	UserData UserData
    // Information needed to join a workgroup or domain. This property was added in vSphere API 7.0.0.
	Domain *Domain
    // Information related to unattended installation. This property was added in vSphere API 7.0.0.
	GuiUnattended GuiUnattended
}

// The ``WindowsConfiguration`` class specifies the settings for customizing a windows guest operating system. This class was added in vSphere API 7.0.0.
type WindowsConfiguration struct {
    // A value specifying the action that should be taken after customization. This property was added in vSphere API 7.0.0.
	Reboot *WindowsConfigurationRebootOption
    // Customization settings like user details, administrator details, etc for the windows guest operating system. Exactly one of ``sysprep`` or ``sysprepXml`` must be specified. This property was added in vSphere API 7.0.0.
	Sysprep *WindowsSysprep
    // All settings specified in a XML format. This is the content of a typical answer.xml file that is used by System administrators during the Windows image customization. Check https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/update-windows-settings-and-scripts-create-your-own-answer-file-sxs Exactly one of ``sysprep`` or ``sysprepXml`` must be specified. This property was added in vSphere API 7.0.0.
	SysprepXml *string
}

// The ``RebootOption`` enumeration class specifies what should be done to the guest after the customization. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type WindowsConfigurationRebootOption string

const (
    // Reboot the guest after customization. This constant field was added in vSphere API 7.0.0.
	WindowsConfigurationRebootOption_REBOOT WindowsConfigurationRebootOption = "REBOOT"
    // Take no action. Leave the guest OS running after customization. This option can be used to look at values for debugging purposes. This constant field was added in vSphere API 7.0.0.
	WindowsConfigurationRebootOption_NO_REBOOT WindowsConfigurationRebootOption = "NO_REBOOT"
    // Shutdown the guest after customization. This constant field was added in vSphere API 7.0.0.
	WindowsConfigurationRebootOption_SHUTDOWN WindowsConfigurationRebootOption = "SHUTDOWN"
)

func (r WindowsConfigurationRebootOption) WindowsConfigurationRebootOption() bool {
	switch r {
	case WindowsConfigurationRebootOption_REBOOT:
		return true
	case WindowsConfigurationRebootOption_NO_REBOOT:
		return true
	case WindowsConfigurationRebootOption_SHUTDOWN:
		return true
	default:
		return false
	}
}


// Guest customization Settings to customize a Linux guest operating system. 
//
//  The ``LinuxConfiguration`` class contains settings that identify a Linux machine in the same way that the WindowsConfiguration class identifies a Windows machine.. This class was added in vSphere API 7.0.0.
type LinuxConfiguration struct {
    // The network host name of the Linux virtual machine. This property was added in vSphere API 7.0.0.
	Hostname HostnameGenerator
    // The fully qualified domain name. This property was added in vSphere API 7.0.0.
	Domain string
    // The case-sensitive time zone, such as Europe/Sofia. Valid time zone values are based on the tz (time zone) database used by Linux. The values are strings (string) in the form "Area/Location," in which Area is a continent or ocean name, and Location is the city, island, or other regional designation. 
    //
    //  See the https://kb.vmware.com/kb/2145518 for a list of supported time zones for different versions in Linux.. This property was added in vSphere API 7.0.0.
	TimeZone *string
    // The script to run before and after Linux guest customization.
    //  The max size of the script is 1500 bytes. As long as the script (shell, perl, python...) has the right "#!" in the header, it is supported. The caller should not assume any environment variables when the script is run. 
    //
    //  The script is invoked by the customization engine using the command line: 1) with argument "precustomization" before customization, 2) with argument "postcustomization" after customization. The script should parse this argument and implement pre-customization or post-customization task code details in the corresponding block. 
    //
    //  A Linux shell script example: 
    //
    //  ``#!/bin/sh
    //  if [ x$1 == x"precustomization" ]; then
    //  echo "Do Precustomization tasks"
    //  #code for pre-customization actions...
    //  elif [ x$1 == x"postcustomization" ]; then
    //  echo "Do Postcustomization tasks"
    //  #code for post-customization actions...
    //  fi``. This property was added in vSphere API 7.0.0.
	ScriptText *string
}

// The ``ConfigurationSpec`` class specifies the settings for customizing a guest operating system. This class was added in vSphere API 7.0.0.
type ConfigurationSpec struct {
    // Guest customization specification for a Windows guest operating system. This property was added in vSphere API 7.0.0.
	WindowsConfig *WindowsConfiguration
    // Guest customization specification for a linux guest operating system. This property was added in vSphere API 7.0.0.
	LinuxConfig *LinuxConfiguration
}

// The ``GlobalDNSSettings`` class specifies a collection of global IP settings in the guest operating system. In Linux, DNS server settings are global. The settings can either be statically set or supplied by a DHCP server. This class was added in vSphere API 7.0.0.
type GlobalDNSSettings struct {
    // List of name resolution suffixes for the virtual network adapter. This list applies to both Windows and Linux guest customization. For Linux, this setting is global, whereas in Windows, this setting is listed on a per-adapter basis. This property was added in vSphere API 7.0.0.
	DnsSuffixList []string
    // List of DNS servers, for a virtual network adapter with a static IP address. If this list is empty, then the guest operating system is expected to use a DHCP server to get its DNS server settings. These settings configure the virtual machine to use the specified DNS servers. These DNS server settings are listed in the order of preference. This property was added in vSphere API 7.0.0.
	DnsServers []string
}

// The ``Ipv4`` class specifies the IPv4 settings that need to be applied to a virtual network adapter. This class was added in vSphere API 7.0.0.
type Ipv4 struct {
    // The type of the IPv4 configuration. This property was added in vSphere API 7.0.0.
	Type_ Ipv4Type
    // The static IPv4 address. This property was added in vSphere API 7.0.0.
	IpAddress *string
    // The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. This property was added in vSphere API 7.0.0.
	Prefix *int64
    // Gateways for the IPv4 address. This property was added in vSphere API 7.0.0.
	Gateways []string
}

// The ``Type`` enumeration class specifies different types of the IPv4 configuration. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Ipv4Type string

const (
    // DHCP configuration. This constant field was added in vSphere API 7.0.0.
	Ipv4Type_DHCP Ipv4Type = "DHCP"
    // Static configuration. This constant field was added in vSphere API 7.0.0.
	Ipv4Type_STATIC Ipv4Type = "STATIC"
    // Indicates that the IPv4 settings are not specified in advance. The user must provide a valid value when the actual customization spec is applied to a virtual machine. This constant field was added in vSphere API 7.0.0.
	Ipv4Type_USER_INPUT_REQUIRED Ipv4Type = "USER_INPUT_REQUIRED"
)

func (t Ipv4Type) Ipv4Type() bool {
	switch t {
	case Ipv4Type_DHCP:
		return true
	case Ipv4Type_STATIC:
		return true
	case Ipv4Type_USER_INPUT_REQUIRED:
		return true
	default:
		return false
	}
}


// The ``Ipv6Address`` class specifies the the settings for a Static IPv6 configuration. This class was added in vSphere API 7.0.0.
type Ipv6Address struct {
    // Static IPv6 Address. This property was added in vSphere API 7.0.0.
	IpAddress string
    // The CIDR prefix for the interface. This property was added in vSphere API 7.0.0.
	Prefix int64
}

// The ``Ipv6`` class specifies the IPv6 settings that need to be applied to a virtual network adapter. This class was added in vSphere API 7.0.0.
type Ipv6 struct {
    // The IPv6 configuration type. This property was added in vSphere API 7.0.0.
	Type_ Ipv6Type
    // IPv6 address. This property was added in vSphere API 7.0.0.
	Ipv6 []Ipv6Address
    // gateways for the IPv6 address. This property was added in vSphere API 7.0.0.
	Gateways []string
}

// The ``Type`` enumeration class specifies different types of the IPv6 configuration. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Ipv6Type string

const (
    // DHCP IPv6 configuration. This constant field was added in vSphere API 7.0.0.
	Ipv6Type_DHCP Ipv6Type = "DHCP"
    // Static IPv6 configuration. This constant field was added in vSphere API 7.0.0.
	Ipv6Type_STATIC Ipv6Type = "STATIC"
    // Indicates that the IPv6 settings are not specified in advance. The user must provide a valid value when the actual customization spec is applied to a virtual machine. This constant field was added in vSphere API 7.0.0.
	Ipv6Type_USER_INPUT_REQUIRED Ipv6Type = "USER_INPUT_REQUIRED"
)

func (t Ipv6Type) Ipv6Type() bool {
	switch t {
	case Ipv6Type_DHCP:
		return true
	case Ipv6Type_STATIC:
		return true
	case Ipv6Type_USER_INPUT_REQUIRED:
		return true
	default:
		return false
	}
}


// The ``WindowsNetworkAdapterSettings`` class specifies Windows specific Network settings for a virtual network adapter. This class was added in vSphere API 7.0.0.
type WindowsNetworkAdapterSettings struct {
    // A list of server IP addresses to use for DNS lookup in a Windows guest operating system. 
    //
    //  These servers should be specified in the order of preference. If this list is not empty, and if a DHCP IP address is used, then these settings override the DHCP settings.. This property was added in vSphere API 7.0.0.
	DnsServers []string
    // A DNS domain suffix such as vmware.com. This property was added in vSphere API 7.0.0.
	DnsDomain *string
    // List of WINS Servers to set for the Windows guest operating system. A Maximum of two IP addresses can be specified in this list. The first IP address will be set as the primary WINS server. The second IP address will be set as the secondary WINS server. This property was added in vSphere API 7.0.0.
	WinsServers []string
    // NetBIOS setting for Windows. This property was added in vSphere API 7.0.0.
	NetBIOSMode *WindowsNetworkAdapterSettingsNetBIOSMode
}

// The ``NetBIOSMode`` enumeration class specifies different NetBIOS settings for Windows guest operating systems. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type WindowsNetworkAdapterSettingsNetBIOSMode string

const (
    // DHCP server decides whether or not to use NetBIOS. This constant field was added in vSphere API 7.0.0.
	WindowsNetworkAdapterSettingsNetBIOSMode_USE_DHCP WindowsNetworkAdapterSettingsNetBIOSMode = "USE_DHCP"
    // Always use NetBIOS. This constant field was added in vSphere API 7.0.0.
	WindowsNetworkAdapterSettingsNetBIOSMode_ENABLE WindowsNetworkAdapterSettingsNetBIOSMode = "ENABLE"
    // Never use NetBIOS. This constant field was added in vSphere API 7.0.0.
	WindowsNetworkAdapterSettingsNetBIOSMode_DISABLE WindowsNetworkAdapterSettingsNetBIOSMode = "DISABLE"
)

func (n WindowsNetworkAdapterSettingsNetBIOSMode) WindowsNetworkAdapterSettingsNetBIOSMode() bool {
	switch n {
	case WindowsNetworkAdapterSettingsNetBIOSMode_USE_DHCP:
		return true
	case WindowsNetworkAdapterSettingsNetBIOSMode_ENABLE:
		return true
	case WindowsNetworkAdapterSettingsNetBIOSMode_DISABLE:
		return true
	default:
		return false
	}
}


// The ``IPSettings`` class specifies the IP settings for a virtual network adapter. This class was added in vSphere API 7.0.0.
type IPSettings struct {
    // Specification to configure IPv4 address, subnet mask and gateway info for this virtual network adapter. This property was added in vSphere API 7.0.0.
	Ipv4 *Ipv4
    // Specification to configure IPv6 address, subnet mask and gateway info for this virtual network adapter. This property was added in vSphere API 7.0.0.
	Ipv6 *Ipv6
    // Windows settings to be configured for this specific virtual Network adapter. This is valid only for Windows guest operating systems. This property was added in vSphere API 7.0.0.
	Windows *WindowsNetworkAdapterSettings
}

// The ``AdapterMapping`` class specifies the association between a virtual network adapter and its IP settings. This class was added in vSphere API 7.0.0.
type AdapterMapping struct {
    // The MAC address of a network adapter being customized. This property was added in vSphere API 7.0.0.
	MacAddress *string
    // The IP settings for the associated virtual network adapter. This property was added in vSphere API 7.0.0.
	Adapter IPSettings
}

// The ``CustomizationSpec`` class contains information required to customize a virtual machine when deploying it or migrating it to a new host. This class was added in vSphere API 7.0.0.
type CustomizationSpec struct {
    // Settings to be applied to the guest during the customization. This property was added in vSphere API 7.0.0.
	ConfigurationSpec ConfigurationSpec
    // Global DNS settings constitute the DNS settings that are not specific to a particular virtual network adapter. This property was added in vSphere API 7.0.0.
	GlobalDNSSettings GlobalDNSSettings
    // IP settings that are specific to a particular virtual network adapter. The AdapterMapping class maps a network adapter's MAC address to its IPSettings. May be empty if there are no network adapters, else should match number of network adapters configured for the VM. This property was added in vSphere API 7.0.0.
	Interfaces []AdapterMapping
}




func HostnameGeneratorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.guest.hostname_generator.type", reflect.TypeOf(HostnameGeneratorType(HostnameGeneratorType_FIXED)))
	fieldNameMap["type"] = "Type_"
	fields["fixed_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["fixed_name"] = "FixedName"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FIXED": []bindings.FieldData{
				bindings.NewFieldData("fixed_name", true),
			},
			"PREFIX": []bindings.FieldData{
				bindings.NewFieldData("prefix", true),
			},
			"VIRTUAL_MACHINE": []bindings.FieldData{},
			"USER_INPUT_REQUIRED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.guest.hostname_generator", fields, reflect.TypeOf(HostnameGenerator{}), fieldNameMap, validators)
}

func UserDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["computer_name"] = bindings.NewReferenceType(HostnameGeneratorBindingType)
	fieldNameMap["computer_name"] = "ComputerName"
	fields["full_name"] = bindings.NewStringType()
	fieldNameMap["full_name"] = "FullName"
	fields["organization"] = bindings.NewStringType()
	fieldNameMap["organization"] = "Organization"
	fields["product_key"] = bindings.NewStringType()
	fieldNameMap["product_key"] = "ProductKey"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.user_data", fields, reflect.TypeOf(UserData{}), fieldNameMap, validators)
}

func DomainBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.guest.domain.type", reflect.TypeOf(DomainType(DomainType_WORKGROUP)))
	fieldNameMap["type"] = "Type_"
	fields["workgroup"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["workgroup"] = "Workgroup"
	fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain"] = "Domain"
	fields["domain_username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_username"] = "DomainUsername"
	fields["domain_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["domain_password"] = "DomainPassword"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"WORKGROUP": []bindings.FieldData{
				bindings.NewFieldData("workgroup", true),
			},
			"DOMAIN": []bindings.FieldData{
				bindings.NewFieldData("domain", true),
				bindings.NewFieldData("domain_username", true),
				bindings.NewFieldData("domain_password", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.guest.domain", fields, reflect.TypeOf(Domain{}), fieldNameMap, validators)
}

func GuiUnattendedBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["auto_logon"] = bindings.NewBooleanType()
	fieldNameMap["auto_logon"] = "AutoLogon"
	fields["auto_logon_count"] = bindings.NewIntegerType()
	fieldNameMap["auto_logon_count"] = "AutoLogonCount"
	fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["password"] = "Password"
	fields["time_zone"] = bindings.NewIntegerType()
	fieldNameMap["time_zone"] = "TimeZone"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.gui_unattended", fields, reflect.TypeOf(GuiUnattended{}), fieldNameMap, validators)
}

func WindowsSysprepBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["gui_run_once_commands"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["gui_run_once_commands"] = "GuiRunOnceCommands"
	fields["user_data"] = bindings.NewReferenceType(UserDataBindingType)
	fieldNameMap["user_data"] = "UserData"
	fields["domain"] = bindings.NewOptionalType(bindings.NewReferenceType(DomainBindingType))
	fieldNameMap["domain"] = "Domain"
	fields["gui_unattended"] = bindings.NewReferenceType(GuiUnattendedBindingType)
	fieldNameMap["gui_unattended"] = "GuiUnattended"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.windows_sysprep", fields, reflect.TypeOf(WindowsSysprep{}), fieldNameMap, validators)
}

func WindowsConfigurationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reboot"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.windows_configuration.reboot_option", reflect.TypeOf(WindowsConfigurationRebootOption(WindowsConfigurationRebootOption_REBOOT))))
	fieldNameMap["reboot"] = "Reboot"
	fields["sysprep"] = bindings.NewOptionalType(bindings.NewReferenceType(WindowsSysprepBindingType))
	fieldNameMap["sysprep"] = "Sysprep"
	fields["sysprep_xml"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sysprep_xml"] = "SysprepXml"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.windows_configuration", fields, reflect.TypeOf(WindowsConfiguration{}), fieldNameMap, validators)
}

func LinuxConfigurationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewReferenceType(HostnameGeneratorBindingType)
	fieldNameMap["hostname"] = "Hostname"
	fields["domain"] = bindings.NewStringType()
	fieldNameMap["domain"] = "Domain"
	fields["time_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["time_zone"] = "TimeZone"
	fields["script_text"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["script_text"] = "ScriptText"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.linux_configuration", fields, reflect.TypeOf(LinuxConfiguration{}), fieldNameMap, validators)
}

func ConfigurationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["windows_config"] = bindings.NewOptionalType(bindings.NewReferenceType(WindowsConfigurationBindingType))
	fieldNameMap["windows_config"] = "WindowsConfig"
	fields["linux_config"] = bindings.NewOptionalType(bindings.NewReferenceType(LinuxConfigurationBindingType))
	fieldNameMap["linux_config"] = "LinuxConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.configuration_spec", fields, reflect.TypeOf(ConfigurationSpec{}), fieldNameMap, validators)
}

func GlobalDNSSettingsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dns_suffix_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns_suffix_list"] = "DnsSuffixList"
	fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns_servers"] = "DnsServers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.global_DNS_settings", fields, reflect.TypeOf(GlobalDNSSettings{}), fieldNameMap, validators)
}

func Ipv4BindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.guest.ipv4.type", reflect.TypeOf(Ipv4Type(Ipv4Type_DHCP)))
	fieldNameMap["type"] = "Type_"
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["prefix"] = "Prefix"
	fields["gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["gateways"] = "Gateways"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"STATIC": []bindings.FieldData{
				bindings.NewFieldData("ip_address", true),
				bindings.NewFieldData("prefix", true),
			},
			"DHCP": []bindings.FieldData{},
			"USER_INPUT_REQUIRED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.guest.ipv4", fields, reflect.TypeOf(Ipv4{}), fieldNameMap, validators)
}

func Ipv6AddressBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_address"] = bindings.NewStringType()
	fieldNameMap["ip_address"] = "IpAddress"
	fields["prefix"] = bindings.NewIntegerType()
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.ipv6_address", fields, reflect.TypeOf(Ipv6Address{}), fieldNameMap, validators)
}

func Ipv6BindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.guest.ipv6.type", reflect.TypeOf(Ipv6Type(Ipv6Type_DHCP)))
	fieldNameMap["type"] = "Type_"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(Ipv6AddressBindingType), reflect.TypeOf([]Ipv6Address{})))
	fieldNameMap["ipv6"] = "Ipv6"
	fields["gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["gateways"] = "Gateways"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"STATIC": []bindings.FieldData{
				bindings.NewFieldData("ipv6", true),
			},
			"DHCP": []bindings.FieldData{},
			"USER_INPUT_REQUIRED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.guest.ipv6", fields, reflect.TypeOf(Ipv6{}), fieldNameMap, validators)
}

func WindowsNetworkAdapterSettingsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns_servers"] = "DnsServers"
	fields["dns_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dns_domain"] = "DnsDomain"
	fields["wins_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["wins_servers"] = "WinsServers"
	fields["net_BIOS_mode"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.windows_network_adapter_settings.net_BIOS_mode", reflect.TypeOf(WindowsNetworkAdapterSettingsNetBIOSMode(WindowsNetworkAdapterSettingsNetBIOSMode_USE_DHCP))))
	fieldNameMap["net_BIOS_mode"] = "NetBIOSMode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.windows_network_adapter_settings", fields, reflect.TypeOf(WindowsNetworkAdapterSettings{}), fieldNameMap, validators)
}

func IPSettingsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4BindingType))
	fieldNameMap["ipv4"] = "Ipv4"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv6BindingType))
	fieldNameMap["ipv6"] = "Ipv6"
	fields["windows"] = bindings.NewOptionalType(bindings.NewReferenceType(WindowsNetworkAdapterSettingsBindingType))
	fieldNameMap["windows"] = "Windows"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.IP_settings", fields, reflect.TypeOf(IPSettings{}), fieldNameMap, validators)
}

func AdapterMappingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["adapter"] = bindings.NewReferenceType(IPSettingsBindingType)
	fieldNameMap["adapter"] = "Adapter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.adapter_mapping", fields, reflect.TypeOf(AdapterMapping{}), fieldNameMap, validators)
}

func CustomizationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["configuration_spec"] = bindings.NewReferenceType(ConfigurationSpecBindingType)
	fieldNameMap["configuration_spec"] = "ConfigurationSpec"
	fields["global_DNS_settings"] = bindings.NewReferenceType(GlobalDNSSettingsBindingType)
	fieldNameMap["global_DNS_settings"] = "GlobalDNSSettings"
	fields["interfaces"] = bindings.NewListType(bindings.NewReferenceType(AdapterMappingBindingType), reflect.TypeOf([]AdapterMapping{}))
	fieldNameMap["interfaces"] = "Interfaces"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_spec", fields, reflect.TypeOf(CustomizationSpec{}), fieldNameMap, validators)
}



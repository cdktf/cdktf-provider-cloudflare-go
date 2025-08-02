// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceCustomProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#account_id ZeroTrustDeviceCustomProfile#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The wirefilter expression to match devices. Available values: "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name", "os.version".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#match ZeroTrustDeviceCustomProfile#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// The name of the device settings profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#name ZeroTrustDeviceCustomProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The precedence of the policy.
	//
	// Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#precedence ZeroTrustDeviceCustomProfile#precedence}
	Precedence *float64 `field:"required" json:"precedence" yaml:"precedence"`
	// Whether to allow devices to leave the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#allowed_to_leave ZeroTrustDeviceCustomProfile#allowed_to_leave}
	AllowedToLeave interface{} `field:"optional" json:"allowedToLeave" yaml:"allowedToLeave"`
	// Whether to allow the user to switch WARP between modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#allow_mode_switch ZeroTrustDeviceCustomProfile#allow_mode_switch}
	AllowModeSwitch interface{} `field:"optional" json:"allowModeSwitch" yaml:"allowModeSwitch"`
	// Whether to receive update notifications when a new version of the client is available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#allow_updates ZeroTrustDeviceCustomProfile#allow_updates}
	AllowUpdates interface{} `field:"optional" json:"allowUpdates" yaml:"allowUpdates"`
	// The amount of time in seconds to reconnect after having been disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#auto_connect ZeroTrustDeviceCustomProfile#auto_connect}
	AutoConnect *float64 `field:"optional" json:"autoConnect" yaml:"autoConnect"`
	// Turn on the captive portal after the specified amount of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#captive_portal ZeroTrustDeviceCustomProfile#captive_portal}
	CaptivePortal *float64 `field:"optional" json:"captivePortal" yaml:"captivePortal"`
	// A description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#description ZeroTrustDeviceCustomProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#disable_auto_fallback ZeroTrustDeviceCustomProfile#disable_auto_fallback}
	DisableAutoFallback interface{} `field:"optional" json:"disableAutoFallback" yaml:"disableAutoFallback"`
	// Whether the policy will be applied to matching devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#enabled ZeroTrustDeviceCustomProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// List of routes excluded in the WARP client's tunnel.
	//
	// Both 'exclude' and 'include' cannot be set in the same request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#exclude ZeroTrustDeviceCustomProfile#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#exclude_office_ips ZeroTrustDeviceCustomProfile#exclude_office_ips}
	ExcludeOfficeIps interface{} `field:"optional" json:"excludeOfficeIps" yaml:"excludeOfficeIps"`
	// List of routes included in the WARP client's tunnel.
	//
	// Both 'exclude' and 'include' cannot be set in the same request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#include ZeroTrustDeviceCustomProfile#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// The amount of time in minutes a user is allowed access to their LAN.
	//
	// A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#lan_allow_minutes ZeroTrustDeviceCustomProfile#lan_allow_minutes}
	LanAllowMinutes *float64 `field:"optional" json:"lanAllowMinutes" yaml:"lanAllowMinutes"`
	// The size of the subnet for the local access network.
	//
	// Note that this field is omitted from the response if null or unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#lan_allow_subnet_size ZeroTrustDeviceCustomProfile#lan_allow_subnet_size}
	LanAllowSubnetSize *float64 `field:"optional" json:"lanAllowSubnetSize" yaml:"lanAllowSubnetSize"`
	// Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#register_interface_ip_with_dns ZeroTrustDeviceCustomProfile#register_interface_ip_with_dns}
	RegisterInterfaceIpWithDns interface{} `field:"optional" json:"registerInterfaceIpWithDns" yaml:"registerInterfaceIpWithDns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN boundary. (Windows only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#sccm_vpn_boundary_support ZeroTrustDeviceCustomProfile#sccm_vpn_boundary_support}
	SccmVpnBoundarySupport interface{} `field:"optional" json:"sccmVpnBoundarySupport" yaml:"sccmVpnBoundarySupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#service_mode_v2 ZeroTrustDeviceCustomProfile#service_mode_v2}.
	ServiceModeV2 *ZeroTrustDeviceCustomProfileServiceModeV2 `field:"optional" json:"serviceModeV2" yaml:"serviceModeV2"`
	// The URL to launch when the Send Feedback button is clicked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#support_url ZeroTrustDeviceCustomProfile#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#switch_locked ZeroTrustDeviceCustomProfile#switch_locked}
	SwitchLocked interface{} `field:"optional" json:"switchLocked" yaml:"switchLocked"`
	// Determines which tunnel protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_custom_profile#tunnel_protocol ZeroTrustDeviceCustomProfile#tunnel_protocol}
	TunnelProtocol *string `field:"optional" json:"tunnelProtocol" yaml:"tunnelProtocol"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceDefaultProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#account_id ZeroTrustDeviceDefaultProfile#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether to allow devices to leave the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#allowed_to_leave ZeroTrustDeviceDefaultProfile#allowed_to_leave}
	AllowedToLeave interface{} `field:"optional" json:"allowedToLeave" yaml:"allowedToLeave"`
	// Whether to allow the user to switch WARP between modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#allow_mode_switch ZeroTrustDeviceDefaultProfile#allow_mode_switch}
	AllowModeSwitch interface{} `field:"optional" json:"allowModeSwitch" yaml:"allowModeSwitch"`
	// Whether to receive update notifications when a new version of the client is available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#allow_updates ZeroTrustDeviceDefaultProfile#allow_updates}
	AllowUpdates interface{} `field:"optional" json:"allowUpdates" yaml:"allowUpdates"`
	// The amount of time in seconds to reconnect after having been disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#auto_connect ZeroTrustDeviceDefaultProfile#auto_connect}
	AutoConnect *float64 `field:"optional" json:"autoConnect" yaml:"autoConnect"`
	// Turn on the captive portal after the specified amount of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#captive_portal ZeroTrustDeviceDefaultProfile#captive_portal}
	CaptivePortal *float64 `field:"optional" json:"captivePortal" yaml:"captivePortal"`
	// If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#disable_auto_fallback ZeroTrustDeviceDefaultProfile#disable_auto_fallback}
	DisableAutoFallback interface{} `field:"optional" json:"disableAutoFallback" yaml:"disableAutoFallback"`
	// List of routes excluded in the WARP client's tunnel.
	//
	// Both 'exclude' and 'include' cannot be set in the same request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#exclude ZeroTrustDeviceDefaultProfile#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#exclude_office_ips ZeroTrustDeviceDefaultProfile#exclude_office_ips}
	ExcludeOfficeIps interface{} `field:"optional" json:"excludeOfficeIps" yaml:"excludeOfficeIps"`
	// List of routes included in the WARP client's tunnel.
	//
	// Both 'exclude' and 'include' cannot be set in the same request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#include ZeroTrustDeviceDefaultProfile#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// The amount of time in minutes a user is allowed access to their LAN.
	//
	// A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#lan_allow_minutes ZeroTrustDeviceDefaultProfile#lan_allow_minutes}
	LanAllowMinutes *float64 `field:"optional" json:"lanAllowMinutes" yaml:"lanAllowMinutes"`
	// The size of the subnet for the local access network.
	//
	// Note that this field is omitted from the response if null or unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#lan_allow_subnet_size ZeroTrustDeviceDefaultProfile#lan_allow_subnet_size}
	LanAllowSubnetSize *float64 `field:"optional" json:"lanAllowSubnetSize" yaml:"lanAllowSubnetSize"`
	// Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#register_interface_ip_with_dns ZeroTrustDeviceDefaultProfile#register_interface_ip_with_dns}
	RegisterInterfaceIpWithDns interface{} `field:"optional" json:"registerInterfaceIpWithDns" yaml:"registerInterfaceIpWithDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#service_mode_v2 ZeroTrustDeviceDefaultProfile#service_mode_v2}.
	ServiceModeV2 *ZeroTrustDeviceDefaultProfileServiceModeV2 `field:"optional" json:"serviceModeV2" yaml:"serviceModeV2"`
	// The URL to launch when the Send Feedback button is clicked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#support_url ZeroTrustDeviceDefaultProfile#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#switch_locked ZeroTrustDeviceDefaultProfile#switch_locked}
	SwitchLocked interface{} `field:"optional" json:"switchLocked" yaml:"switchLocked"`
	// Determines which tunnel protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_device_default_profile#tunnel_protocol ZeroTrustDeviceDefaultProfile#tunnel_protocol}
	TunnelProtocol *string `field:"optional" json:"tunnelProtocol" yaml:"tunnelProtocol"`
}


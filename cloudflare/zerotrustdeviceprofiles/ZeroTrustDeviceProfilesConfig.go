// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceprofiles

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceProfilesConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#account_id ZeroTrustDeviceProfiles#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Description of Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#description ZeroTrustDeviceProfiles#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#name ZeroTrustDeviceProfiles#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to allow devices to leave the organization. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#allowed_to_leave ZeroTrustDeviceProfiles#allowed_to_leave}
	AllowedToLeave interface{} `field:"optional" json:"allowedToLeave" yaml:"allowedToLeave"`
	// Whether to allow mode switch for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#allow_mode_switch ZeroTrustDeviceProfiles#allow_mode_switch}
	AllowModeSwitch interface{} `field:"optional" json:"allowModeSwitch" yaml:"allowModeSwitch"`
	// Whether to allow updates under this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#allow_updates ZeroTrustDeviceProfiles#allow_updates}
	AllowUpdates interface{} `field:"optional" json:"allowUpdates" yaml:"allowUpdates"`
	// The amount of time in seconds to reconnect after having been disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#auto_connect ZeroTrustDeviceProfiles#auto_connect}
	AutoConnect *float64 `field:"optional" json:"autoConnect" yaml:"autoConnect"`
	// The captive portal value for this policy. Defaults to `180`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#captive_portal ZeroTrustDeviceProfiles#captive_portal}
	CaptivePortal *float64 `field:"optional" json:"captivePortal" yaml:"captivePortal"`
	// Whether the policy refers to the default account policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#default ZeroTrustDeviceProfiles#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Whether to disable auto fallback for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#disable_auto_fallback ZeroTrustDeviceProfiles#disable_auto_fallback}
	DisableAutoFallback interface{} `field:"optional" json:"disableAutoFallback" yaml:"disableAutoFallback"`
	// Whether the policy is enabled (cannot be set for default policies). Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#enabled ZeroTrustDeviceProfiles#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#exclude_office_ips ZeroTrustDeviceProfiles#exclude_office_ips}
	ExcludeOfficeIps interface{} `field:"optional" json:"excludeOfficeIps" yaml:"excludeOfficeIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#id ZeroTrustDeviceProfiles#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Wirefilter expression to match a device against when evaluating whether this policy should take effect for that device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#match ZeroTrustDeviceProfiles#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// The precedence of the policy. Lower values indicate higher precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#precedence ZeroTrustDeviceProfiles#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The service mode. Available values: `1dot1`, `warp`, `proxy`, `posture_only`, `warp_tunnel_only`. Defaults to `warp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#service_mode_v2_mode ZeroTrustDeviceProfiles#service_mode_v2_mode}
	ServiceModeV2Mode *string `field:"optional" json:"serviceModeV2Mode" yaml:"serviceModeV2Mode"`
	// The port to use for the proxy service mode. Required when using `service_mode_v2_mode`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#service_mode_v2_port ZeroTrustDeviceProfiles#service_mode_v2_port}
	ServiceModeV2Port *float64 `field:"optional" json:"serviceModeV2Port" yaml:"serviceModeV2Port"`
	// The support URL that will be opened when sending feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#support_url ZeroTrustDeviceProfiles#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// Enablement of the ZT client switch lock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#switch_locked ZeroTrustDeviceProfiles#switch_locked}
	SwitchLocked interface{} `field:"optional" json:"switchLocked" yaml:"switchLocked"`
	// Determines which tunnel protocol to use. Available values: `""`, `wireguard`, `masque`. Defaults to `wireguard`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_device_profiles#tunnel_protocol ZeroTrustDeviceProfiles#tunnel_protocol}
	TunnelProtocol *string `field:"optional" json:"tunnelProtocol" yaml:"tunnelProtocol"`
}


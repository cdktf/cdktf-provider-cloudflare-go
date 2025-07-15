// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofile


type ZeroTrustDeviceCustomProfileServiceModeV2 struct {
	// The mode to run the WARP client under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_custom_profile#mode ZeroTrustDeviceCustomProfile#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The port number when used with proxy mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_custom_profile#port ZeroTrustDeviceCustomProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


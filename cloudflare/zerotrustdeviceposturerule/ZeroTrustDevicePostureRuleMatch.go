// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleMatch struct {
	// The platform of the device. Available values: `windows`, `mac`, `linux`, `android`, `ios`, `chromeos`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#platform ZeroTrustDevicePostureRule#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}


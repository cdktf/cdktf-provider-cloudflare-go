// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleMatch struct {
	// Available values: "windows", "mac", "linux", "android", "ios", "chromeos".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_device_posture_rule#platform ZeroTrustDevicePostureRule#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}


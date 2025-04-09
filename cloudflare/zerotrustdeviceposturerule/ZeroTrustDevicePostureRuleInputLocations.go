// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleInputLocations struct {
	// List of paths to check for client certificate on linux.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_device_posture_rule#paths ZeroTrustDevicePostureRule#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// List of trust stores to check for client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_device_posture_rule#trust_stores ZeroTrustDevicePostureRule#trust_stores}
	TrustStores *[]*string `field:"optional" json:"trustStores" yaml:"trustStores"`
}


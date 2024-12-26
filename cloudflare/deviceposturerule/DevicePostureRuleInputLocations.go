// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deviceposturerule


type DevicePostureRuleInputLocations struct {
	// List of paths to check for client certificate rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/device_posture_rule#paths DevicePostureRule#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// List of trust stores to check for client certificate rule. Available values: `system`, `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/device_posture_rule#trust_stores DevicePostureRule#trust_stores}
	TrustStores *[]*string `field:"optional" json:"trustStores" yaml:"trustStores"`
}


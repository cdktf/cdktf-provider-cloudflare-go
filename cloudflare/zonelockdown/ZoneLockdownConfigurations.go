// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonelockdown


type ZoneLockdownConfigurations struct {
	// The configuration target.
	//
	// You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.
	// Available values: "ip", "ip_range".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_lockdown#target ZoneLockdown#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The IP address to match. This address will be compared to the IP address of incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_lockdown#value ZoneLockdown#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


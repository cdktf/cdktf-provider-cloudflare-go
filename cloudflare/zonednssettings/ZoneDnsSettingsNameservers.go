// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonednssettings


type ZoneDnsSettingsNameservers struct {
	// Nameserver type Available values: "cloudflare.standard", "custom.account", "custom.tenant", "custom.zone".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zone_dns_settings#type ZoneDnsSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configured nameserver set to be used for this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zone_dns_settings#ns_set ZoneDnsSettings#ns_set}
	NsSet *float64 `field:"optional" json:"nsSet" yaml:"nsSet"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings


type DnsSettingsZoneDefaultsInternalDns struct {
	// The ID of the zone to fallback to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#reference_zone_id DnsSettings#reference_zone_id}
	ReferenceZoneId *string `field:"optional" json:"referenceZoneId" yaml:"referenceZoneId"`
}


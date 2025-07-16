// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonednssettings


type ZoneDnsSettingsInternalDns struct {
	// The ID of the zone to fallback to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zone_dns_settings#reference_zone_id ZoneDnsSettings#reference_zone_id}
	ReferenceZoneId *string `field:"optional" json:"referenceZoneId" yaml:"referenceZoneId"`
}


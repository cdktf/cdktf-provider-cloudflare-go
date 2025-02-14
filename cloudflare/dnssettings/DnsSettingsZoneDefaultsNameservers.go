// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings


type DnsSettingsZoneDefaultsNameservers struct {
	// Nameserver type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#type DnsSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}


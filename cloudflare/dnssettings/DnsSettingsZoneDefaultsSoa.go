// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings


type DnsSettingsZoneDefaultsSoa struct {
	// Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#expire DnsSettings#expire}
	Expire *float64 `field:"required" json:"expire" yaml:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#min_ttl DnsSettings#min_ttl}
	MinTtl *float64 `field:"required" json:"minTtl" yaml:"minTtl"`
	// The primary nameserver, which may be used for outbound zone transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#mname DnsSettings#mname}
	Mname *string `field:"required" json:"mname" yaml:"mname"`
	// Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#refresh DnsSettings#refresh}
	Refresh *float64 `field:"required" json:"refresh" yaml:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#retry DnsSettings#retry}
	Retry *float64 `field:"required" json:"retry" yaml:"retry"`
	// The email address of the zone administrator, with the first label representing the local part of the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#rname DnsSettings#rname}
	Rname *string `field:"required" json:"rname" yaml:"rname"`
	// The time to live (TTL) of the SOA record itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#ttl DnsSettings#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}


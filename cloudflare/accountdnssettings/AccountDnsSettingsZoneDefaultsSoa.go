// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountdnssettings


type AccountDnsSettingsZoneDefaultsSoa struct {
	// Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#expire AccountDnsSettings#expire}
	Expire *float64 `field:"required" json:"expire" yaml:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#min_ttl AccountDnsSettings#min_ttl}
	MinTtl *float64 `field:"required" json:"minTtl" yaml:"minTtl"`
	// The primary nameserver, which may be used for outbound zone transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#mname AccountDnsSettings#mname}
	Mname *string `field:"required" json:"mname" yaml:"mname"`
	// Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#refresh AccountDnsSettings#refresh}
	Refresh *float64 `field:"required" json:"refresh" yaml:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#retry AccountDnsSettings#retry}
	Retry *float64 `field:"required" json:"retry" yaml:"retry"`
	// The email address of the zone administrator, with the first label representing the local part of the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#rname AccountDnsSettings#rname}
	Rname *string `field:"required" json:"rname" yaml:"rname"`
	// The time to live (TTL) of the SOA record itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/account_dns_settings#ttl AccountDnsSettings#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}


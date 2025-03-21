// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountdnssettings


type AccountDnsSettingsZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone.
	//
	// Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#flatten_all_cnames AccountDnsSettings#flatten_all_cnames}
	FlattenAllCnames interface{} `field:"optional" json:"flattenAllCnames" yaml:"flattenAllCnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#foundation_dns AccountDnsSettings#foundation_dns}
	FoundationDns interface{} `field:"optional" json:"foundationDns" yaml:"foundationDns"`
	// Settings for this internal zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#internal_dns AccountDnsSettings#internal_dns}
	InternalDns *AccountDnsSettingsZoneDefaultsInternalDns `field:"optional" json:"internalDns" yaml:"internalDns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#multi_provider AccountDnsSettings#multi_provider}
	MultiProvider interface{} `field:"optional" json:"multiProvider" yaml:"multiProvider"`
	// Settings determining the nameservers through which the zone should be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#nameservers AccountDnsSettings#nameservers}
	Nameservers *AccountDnsSettingsZoneDefaultsNameservers `field:"optional" json:"nameservers" yaml:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#ns_ttl AccountDnsSettings#ns_ttl}
	NsTtl *float64 `field:"optional" json:"nsTtl" yaml:"nsTtl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#secondary_overrides AccountDnsSettings#secondary_overrides}
	SecondaryOverrides interface{} `field:"optional" json:"secondaryOverrides" yaml:"secondaryOverrides"`
	// Components of the zone's SOA record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#soa AccountDnsSettings#soa}
	Soa *AccountDnsSettingsZoneDefaultsSoa `field:"optional" json:"soa" yaml:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone. Available values: "standard", "cdn_only", "dns_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/account_dns_settings#zone_mode AccountDnsSettings#zone_mode}
	ZoneMode *string `field:"optional" json:"zoneMode" yaml:"zoneMode"`
}


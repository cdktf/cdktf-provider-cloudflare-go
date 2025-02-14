// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings


type DnsSettingsZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone.
	//
	// Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#flatten_all_cnames DnsSettings#flatten_all_cnames}
	FlattenAllCnames interface{} `field:"optional" json:"flattenAllCnames" yaml:"flattenAllCnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#foundation_dns DnsSettings#foundation_dns}
	FoundationDns interface{} `field:"optional" json:"foundationDns" yaml:"foundationDns"`
	// Settings for this internal zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#internal_dns DnsSettings#internal_dns}
	InternalDns *DnsSettingsZoneDefaultsInternalDns `field:"optional" json:"internalDns" yaml:"internalDns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#multi_provider DnsSettings#multi_provider}
	MultiProvider interface{} `field:"optional" json:"multiProvider" yaml:"multiProvider"`
	// Settings determining the nameservers through which the zone should be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#nameservers DnsSettings#nameservers}
	Nameservers *DnsSettingsZoneDefaultsNameservers `field:"optional" json:"nameservers" yaml:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#ns_ttl DnsSettings#ns_ttl}
	NsTtl *float64 `field:"optional" json:"nsTtl" yaml:"nsTtl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#secondary_overrides DnsSettings#secondary_overrides}
	SecondaryOverrides interface{} `field:"optional" json:"secondaryOverrides" yaml:"secondaryOverrides"`
	// Components of the zone's SOA record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#soa DnsSettings#soa}
	Soa *DnsSettingsZoneDefaultsSoa `field:"optional" json:"soa" yaml:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_settings#zone_mode DnsSettings#zone_mode}
	ZoneMode *string `field:"optional" json:"zoneMode" yaml:"zoneMode"`
}


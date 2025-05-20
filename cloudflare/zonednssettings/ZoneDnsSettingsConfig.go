// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonednssettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneDnsSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#zone_id ZoneDnsSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Whether to flatten all CNAME records in the zone.
	//
	// Note that, due to DNS limitations, a CNAME record at the zone apex will always be flattened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#flatten_all_cnames ZoneDnsSettings#flatten_all_cnames}
	FlattenAllCnames interface{} `field:"optional" json:"flattenAllCnames" yaml:"flattenAllCnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#foundation_dns ZoneDnsSettings#foundation_dns}
	FoundationDns interface{} `field:"optional" json:"foundationDns" yaml:"foundationDns"`
	// Settings for this internal zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#internal_dns ZoneDnsSettings#internal_dns}
	InternalDns *ZoneDnsSettingsInternalDns `field:"optional" json:"internalDns" yaml:"internalDns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the zone even when non-Cloudflare NS records exist, and to respect NS records at the zone apex during outbound zone transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#multi_provider ZoneDnsSettings#multi_provider}
	MultiProvider interface{} `field:"optional" json:"multiProvider" yaml:"multiProvider"`
	// Settings determining the nameservers through which the zone should be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#nameservers ZoneDnsSettings#nameservers}
	Nameservers *ZoneDnsSettingsNameservers `field:"optional" json:"nameservers" yaml:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#ns_ttl ZoneDnsSettings#ns_ttl}
	NsTtl *float64 `field:"optional" json:"nsTtl" yaml:"nsTtl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME flattening at the zone apex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#secondary_overrides ZoneDnsSettings#secondary_overrides}
	SecondaryOverrides interface{} `field:"optional" json:"secondaryOverrides" yaml:"secondaryOverrides"`
	// Components of the zone's SOA record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#soa ZoneDnsSettings#soa}
	Soa *ZoneDnsSettingsSoa `field:"optional" json:"soa" yaml:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone. Available values: "standard", "cdn_only", "dns_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zone_dns_settings#zone_mode ZoneDnsSettings#zone_mode}
	ZoneMode *string `field:"optional" json:"zoneMode" yaml:"zoneMode"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountdnssettingsinternalview


type DataCloudflareAccountDnsSettingsInternalViewFilter struct {
	// Direction to order DNS views in. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#direction DataCloudflareAccountDnsSettingsInternalView#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Whether to match all search requirements or at least one (any).
	//
	// If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#match DataCloudflareAccountDnsSettingsInternalView#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#name DataCloudflareAccountDnsSettingsInternalView#name}.
	Name *DataCloudflareAccountDnsSettingsInternalViewFilterName `field:"optional" json:"name" yaml:"name"`
	// Field to order DNS views by. Available values: "name", "created_on", "modified_on".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#order DataCloudflareAccountDnsSettingsInternalView#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A zone ID that exists in the zones list for the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#zone_id DataCloudflareAccountDnsSettingsInternalView#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
	// A zone name that exists in the zones list for the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/account_dns_settings_internal_view#zone_name DataCloudflareAccountDnsSettingsInternalView#zone_name}
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}


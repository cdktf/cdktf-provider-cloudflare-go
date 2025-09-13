// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountdnssettingsinternalview


type DataCloudflareAccountDnsSettingsInternalViewFilterName struct {
	// Substring of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_dns_settings_internal_view#contains DataCloudflareAccountDnsSettingsInternalView#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Suffix of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_dns_settings_internal_view#endswith DataCloudflareAccountDnsSettingsInternalView#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// Exact value of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_dns_settings_internal_view#exact DataCloudflareAccountDnsSettingsInternalView#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Prefix of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_dns_settings_internal_view#startswith DataCloudflareAccountDnsSettingsInternalView#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}


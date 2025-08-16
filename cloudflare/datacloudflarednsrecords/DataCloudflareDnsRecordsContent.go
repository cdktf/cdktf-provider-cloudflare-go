// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords


type DataCloudflareDnsRecordsContent struct {
	// Substring of the DNS record content. Content filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/dns_records#contains DataCloudflareDnsRecords#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Suffix of the DNS record content. Content filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/dns_records#endswith DataCloudflareDnsRecords#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// Exact value of the DNS record content. Content filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/dns_records#exact DataCloudflareDnsRecords#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Prefix of the DNS record content. Content filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/dns_records#startswith DataCloudflareDnsRecords#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}


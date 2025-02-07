// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords


type DataCloudflareDnsRecordsComment struct {
	// If this parameter is present, only records *without* a comment are returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#absent DataCloudflareDnsRecords#absent}
	Absent *string `field:"optional" json:"absent" yaml:"absent"`
	// Substring of the DNS record comment. Comment filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#contains DataCloudflareDnsRecords#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Suffix of the DNS record comment. Comment filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#endswith DataCloudflareDnsRecords#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// Exact value of the DNS record comment. Comment filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#exact DataCloudflareDnsRecords#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// If this parameter is present, only records *with* a comment are returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#present DataCloudflareDnsRecords#present}
	Present *string `field:"optional" json:"present" yaml:"present"`
	// Prefix of the DNS record comment. Comment filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/dns_records#startswith DataCloudflareDnsRecords#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}


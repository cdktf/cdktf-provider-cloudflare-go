// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords


type DataCloudflareDnsRecordsTag struct {
	// Name of a tag which must *not* be present on the DNS record. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#absent DataCloudflareDnsRecords#absent}
	Absent *string `field:"optional" json:"absent" yaml:"absent"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value contains `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#contains DataCloudflareDnsRecords#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value ends with `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#endswith DataCloudflareDnsRecords#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#exact DataCloudflareDnsRecords#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Name of a tag which must be present on the DNS record. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#present DataCloudflareDnsRecords#present}
	Present *string `field:"optional" json:"present" yaml:"present"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value starts with `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#startswith DataCloudflareDnsRecords#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}


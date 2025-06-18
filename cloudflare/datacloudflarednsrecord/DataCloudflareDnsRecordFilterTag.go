// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecord


type DataCloudflareDnsRecordFilterTag struct {
	// Name of a tag which must *not* be present on the DNS record. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#absent DataCloudflareDnsRecord#absent}
	Absent *string `field:"optional" json:"absent" yaml:"absent"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value contains `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#contains DataCloudflareDnsRecord#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value ends with `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#endswith DataCloudflareDnsRecord#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#exact DataCloudflareDnsRecord#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Name of a tag which must be present on the DNS record. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#present DataCloudflareDnsRecord#present}
	Present *string `field:"optional" json:"present" yaml:"present"`
	// A tag and value, of the form `<tag-name>:<tag-value>`.
	//
	// The API will only return DNS records that have a tag named `<tag-name>` whose value starts with `<tag-value>`. Tag filters are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/dns_record#startswith DataCloudflareDnsRecord#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}


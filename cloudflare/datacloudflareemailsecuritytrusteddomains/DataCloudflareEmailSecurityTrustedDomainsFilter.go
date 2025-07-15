// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecuritytrusteddomains


type DataCloudflareEmailSecurityTrustedDomainsFilter struct {
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_trusted_domains#direction DataCloudflareEmailSecurityTrustedDomains#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_trusted_domains#is_recent DataCloudflareEmailSecurityTrustedDomains#is_recent}.
	IsRecent interface{} `field:"optional" json:"isRecent" yaml:"isRecent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_trusted_domains#is_similarity DataCloudflareEmailSecurityTrustedDomains#is_similarity}.
	IsSimilarity interface{} `field:"optional" json:"isSimilarity" yaml:"isSimilarity"`
	// The field to sort by. Available values: "pattern", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_trusted_domains#order DataCloudflareEmailSecurityTrustedDomains#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Allows searching in multiple properties of a record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact
	// behavior is intentionally left unspecified and is subject to change
	// in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_trusted_domains#search DataCloudflareEmailSecurityTrustedDomains#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}


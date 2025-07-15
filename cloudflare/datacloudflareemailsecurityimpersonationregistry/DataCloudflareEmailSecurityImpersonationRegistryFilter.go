// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityimpersonationregistry


type DataCloudflareEmailSecurityImpersonationRegistryFilter struct {
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_impersonation_registry#direction DataCloudflareEmailSecurityImpersonationRegistry#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The field to sort by. Available values: "name", "email", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_impersonation_registry#order DataCloudflareEmailSecurityImpersonationRegistry#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Available values: "A1S_INTERNAL", "SNOOPY-CASB_OFFICE_365", "SNOOPY-OFFICE_365", "SNOOPY-GOOGLE_DIRECTORY".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_impersonation_registry#provenance DataCloudflareEmailSecurityImpersonationRegistry#provenance}
	Provenance *string `field:"optional" json:"provenance" yaml:"provenance"`
	// Allows searching in multiple properties of a record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact
	// behavior is intentionally left unspecified and is subject to change
	// in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_impersonation_registry#search DataCloudflareEmailSecurityImpersonationRegistry#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}


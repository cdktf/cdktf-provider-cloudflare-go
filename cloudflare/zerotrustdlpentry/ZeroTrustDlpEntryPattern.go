// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpentry


type ZeroTrustDlpEntryPattern struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dlp_entry#regex ZeroTrustDlpEntry#regex}.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Available values: "luhn".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dlp_entry#validation ZeroTrustDlpEntry#validation}
	Validation *string `field:"optional" json:"validation" yaml:"validation"`
}


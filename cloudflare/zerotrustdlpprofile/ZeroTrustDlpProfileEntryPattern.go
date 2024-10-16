// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpprofile


type ZeroTrustDlpProfileEntryPattern struct {
	// The regex that defines the pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_dlp_profile#regex ZeroTrustDlpProfile#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// The validation algorithm to apply with this pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_dlp_profile#validation ZeroTrustDlpProfile#validation}
	Validation *string `field:"optional" json:"validation" yaml:"validation"`
}


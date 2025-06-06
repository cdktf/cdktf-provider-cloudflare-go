// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpcustomprofile


type ZeroTrustDlpCustomProfileProfilesContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches surrounded by keywords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_dlp_custom_profile#enabled ZeroTrustDlpCustomProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Content types to exclude from context analysis and return all matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_dlp_custom_profile#skip ZeroTrustDlpCustomProfile#skip}
	Skip *ZeroTrustDlpCustomProfileProfilesContextAwarenessSkip `field:"required" json:"skip" yaml:"skip"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlppredefinedprofile


type ZeroTrustDlpPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches surrounded by keywords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_dlp_predefined_profile#enabled ZeroTrustDlpPredefinedProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Content types to exclude from context analysis and return all matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_dlp_predefined_profile#skip ZeroTrustDlpPredefinedProfile#skip}
	Skip *ZeroTrustDlpPredefinedProfileContextAwarenessSkip `field:"required" json:"skip" yaml:"skip"`
}


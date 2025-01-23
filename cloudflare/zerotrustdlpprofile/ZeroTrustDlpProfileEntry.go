// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpprofile


type ZeroTrustDlpProfileEntry struct {
	// Name of the entry to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_dlp_profile#name ZeroTrustDlpProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the entry is active. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_dlp_profile#enabled ZeroTrustDlpProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Unique entry identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_dlp_profile#id ZeroTrustDlpProfile#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_dlp_profile#pattern ZeroTrustDlpProfile#pattern}
	Pattern *ZeroTrustDlpProfileEntryPattern `field:"optional" json:"pattern" yaml:"pattern"`
}


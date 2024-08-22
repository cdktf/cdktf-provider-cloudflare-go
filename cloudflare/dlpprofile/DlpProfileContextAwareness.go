// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlpprofile


type DlpProfileContextAwareness struct {
	// Scan the context of predefined entries to only return matches surrounded by keywords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/dlp_profile#enabled DlpProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// skip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/dlp_profile#skip DlpProfile#skip}
	Skip *DlpProfileContextAwarenessSkip `field:"required" json:"skip" yaml:"skip"`
}


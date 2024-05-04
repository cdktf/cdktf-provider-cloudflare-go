// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlpprofile


type DlpProfileEntryPattern struct {
	// The regex that defines the pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.31.0/docs/resources/dlp_profile#regex DlpProfile#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// The validation algorithm to apply with this pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.31.0/docs/resources/dlp_profile#validation DlpProfile#validation}
	Validation *string `field:"optional" json:"validation" yaml:"validation"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization


type ZeroTrustOrganizationLoginDesign struct {
	// The background color on your login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_organization#background_color ZeroTrustOrganization#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The text at the bottom of your login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_organization#footer_text ZeroTrustOrganization#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// The text at the top of your login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_organization#header_text ZeroTrustOrganization#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// The URL of the logo on your login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_organization#logo_path ZeroTrustOrganization#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// The text color on your login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_organization#text_color ZeroTrustOrganization#text_color}
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}


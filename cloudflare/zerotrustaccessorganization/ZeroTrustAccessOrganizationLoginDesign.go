// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessorganization


type ZeroTrustAccessOrganizationLoginDesign struct {
	// The background color on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_organization#background_color ZeroTrustAccessOrganization#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The text at the bottom of the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_organization#footer_text ZeroTrustAccessOrganization#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// The text at the top of the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_organization#header_text ZeroTrustAccessOrganization#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// The URL of the logo on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_organization#logo_path ZeroTrustAccessOrganization#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// The text color on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_organization#text_color ZeroTrustAccessOrganization#text_color}
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#background_color ZeroTrustGatewaySettings#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Block page footer text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#footer_text ZeroTrustGatewaySettings#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// Block page header text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#header_text ZeroTrustGatewaySettings#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// Full URL to the logo file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#logo_path ZeroTrustGatewaySettings#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// Admin email for users to contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#mailto_address ZeroTrustGatewaySettings#mailto_address}
	MailtoAddress *string `field:"optional" json:"mailtoAddress" yaml:"mailtoAddress"`
	// Subject line for emails created from block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#mailto_subject ZeroTrustGatewaySettings#mailto_subject}
	MailtoSubject *string `field:"optional" json:"mailtoSubject" yaml:"mailtoSubject"`
	// Block page title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#name ZeroTrustGatewaySettings#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Suppress detailed info at the bottom of the block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_gateway_settings#suppress_footer ZeroTrustGatewaySettings#suppress_footer}
	SuppressFooter interface{} `field:"optional" json:"suppressFooter" yaml:"suppressFooter"`
}


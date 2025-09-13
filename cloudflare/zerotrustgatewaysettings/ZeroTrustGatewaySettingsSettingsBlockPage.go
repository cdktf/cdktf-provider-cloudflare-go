// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBlockPage struct {
	// If mode is customized_block_page: block page background color in #rrggbb format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#background_color ZeroTrustGatewaySettings#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Enable only cipher suites and TLS versions compliant with FIPS. 140-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If mode is customized_block_page: block page footer text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#footer_text ZeroTrustGatewaySettings#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// If mode is customized_block_page: block page header text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#header_text ZeroTrustGatewaySettings#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// If mode is redirect_uri: when enabled, context will be appended to target_uri as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#include_context ZeroTrustGatewaySettings#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// If mode is customized_block_page: full URL to the logo file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#logo_path ZeroTrustGatewaySettings#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// If mode is customized_block_page: admin email for users to contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#mailto_address ZeroTrustGatewaySettings#mailto_address}
	MailtoAddress *string `field:"optional" json:"mailtoAddress" yaml:"mailtoAddress"`
	// If mode is customized_block_page: subject line for emails created from block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#mailto_subject ZeroTrustGatewaySettings#mailto_subject}
	MailtoSubject *string `field:"optional" json:"mailtoSubject" yaml:"mailtoSubject"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to a customer-provided URI.
	//
	// Available values: "", "customized_block_page", "redirect_uri".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#mode ZeroTrustGatewaySettings#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// If mode is customized_block_page: block page title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#name ZeroTrustGatewaySettings#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// This setting was shared via the Orgs API and cannot be edited by the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#read_only ZeroTrustGatewaySettings#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Account tag of account that shared this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#source_account ZeroTrustGatewaySettings#source_account}
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#suppress_footer ZeroTrustGatewaySettings#suppress_footer}
	SuppressFooter interface{} `field:"optional" json:"suppressFooter" yaml:"suppressFooter"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#target_uri ZeroTrustGatewaySettings#target_uri}
	TargetUri *string `field:"optional" json:"targetUri" yaml:"targetUri"`
	// Version number of the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_gateway_settings#version ZeroTrustGatewaySettings#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}


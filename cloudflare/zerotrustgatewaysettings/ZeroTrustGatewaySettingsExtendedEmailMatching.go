// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsExtendedEmailMatching struct {
	// Whether e-mails should be matched on all variants of user emails (with + or . modifiers) in Firewall policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


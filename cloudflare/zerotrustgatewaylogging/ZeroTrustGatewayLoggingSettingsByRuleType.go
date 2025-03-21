// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaylogging


type ZeroTrustGatewayLoggingSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/zero_trust_gateway_logging#dns ZeroTrustGatewayLogging#dns}
	Dns *string `field:"optional" json:"dns" yaml:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/zero_trust_gateway_logging#http ZeroTrustGatewayLogging#http}
	Http *string `field:"optional" json:"http" yaml:"http"`
	// Logging settings for Network firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/zero_trust_gateway_logging#l4 ZeroTrustGatewayLogging#l4}
	L4 *string `field:"optional" json:"l4" yaml:"l4"`
}


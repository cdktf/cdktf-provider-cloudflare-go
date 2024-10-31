// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsLoggingSettingsByRuleType struct {
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_gateway_settings#dns ZeroTrustGatewaySettings#dns}
	Dns *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns `field:"required" json:"dns" yaml:"dns"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_gateway_settings#http ZeroTrustGatewaySettings#http}
	Http *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp `field:"required" json:"http" yaml:"http"`
	// l4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_gateway_settings#l4 ZeroTrustGatewaySettings#l4}
	L4 *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4 `field:"required" json:"l4" yaml:"l4"`
}


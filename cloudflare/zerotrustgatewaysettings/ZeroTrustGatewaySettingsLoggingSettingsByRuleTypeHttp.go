// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp struct {
	// Whether to log all activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_settings#log_all ZeroTrustGatewaySettings#log_all}
	LogAll interface{} `field:"required" json:"logAll" yaml:"logAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_settings#log_blocks ZeroTrustGatewaySettings#log_blocks}.
	LogBlocks interface{} `field:"required" json:"logBlocks" yaml:"logBlocks"`
}


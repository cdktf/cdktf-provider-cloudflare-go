// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsL4Override struct {
	// Override IP to forward traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#ip ZeroTrustGatewayPolicy#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Override Port to forward traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#port ZeroTrustGatewayPolicy#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}


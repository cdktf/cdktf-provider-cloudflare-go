// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsDnsResolvers struct {
	// ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_gateway_policy#ipv4 ZeroTrustGatewayPolicy#ipv4}
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_gateway_policy#ipv6 ZeroTrustGatewayPolicy#ipv6}
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
}


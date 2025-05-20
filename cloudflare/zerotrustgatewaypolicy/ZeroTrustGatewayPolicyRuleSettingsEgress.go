// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_gateway_policy#ipv4 ZeroTrustGatewayPolicy#ipv4}
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error egressing with the primary IPv4.
	//
	// Can be '0.0.0.0' to indicate local egress via WARP IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_gateway_policy#ipv4_fallback ZeroTrustGatewayPolicy#ipv4_fallback}
	Ipv4Fallback *string `field:"optional" json:"ipv4Fallback" yaml:"ipv4Fallback"`
	// The IPv6 range to be used for egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_gateway_policy#ipv6 ZeroTrustGatewayPolicy#ipv6}
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
}


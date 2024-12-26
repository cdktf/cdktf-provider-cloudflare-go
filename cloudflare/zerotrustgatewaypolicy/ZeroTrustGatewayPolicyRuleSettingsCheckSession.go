// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Enable session enforcement for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_gateway_policy#enforce ZeroTrustGatewayPolicy#enforce}
	Enforce interface{} `field:"required" json:"enforce" yaml:"enforce"`
}


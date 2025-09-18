// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	//
	// The API automatically formats and sanitizes this expression. This returns a normalized version that may differ from your input and cause Terraform state drift.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Set to true to enable session enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_gateway_policy#enforce ZeroTrustGatewayPolicy#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
}


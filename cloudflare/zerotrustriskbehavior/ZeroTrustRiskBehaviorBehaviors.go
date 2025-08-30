// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustriskbehavior


type ZeroTrustRiskBehaviorBehaviors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_risk_behavior#enabled ZeroTrustRiskBehavior#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Available values: "low", "medium", "high".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_risk_behavior#risk_level ZeroTrustRiskBehavior#risk_level}
	RiskLevel *string `field:"required" json:"riskLevel" yaml:"riskLevel"`
}


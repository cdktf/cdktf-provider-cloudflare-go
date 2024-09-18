// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustriskbehavior


type ZeroTrustRiskBehaviorBehavior struct {
	// Whether this risk behavior type is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_risk_behavior#enabled ZeroTrustRiskBehavior#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Name of this risk behavior type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_risk_behavior#name ZeroTrustRiskBehavior#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Risk level. Available values: `low`, `medium`, `high`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_risk_behavior#risk_level ZeroTrustRiskBehavior#risk_level}
	RiskLevel *string `field:"required" json:"riskLevel" yaml:"riskLevel"`
}


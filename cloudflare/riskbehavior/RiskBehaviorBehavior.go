// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package riskbehavior


type RiskBehaviorBehavior struct {
	// Whether this risk behavior type is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/risk_behavior#enabled RiskBehavior#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Name of this risk behavior type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/risk_behavior#name RiskBehavior#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Risk level. Available values: `low`, `medium`, `high`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/risk_behavior#risk_level RiskBehavior#risk_level}
	RiskLevel *string `field:"required" json:"riskLevel" yaml:"riskLevel"`
}


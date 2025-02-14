// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersOrigin struct {
	// Override the resolved hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/ruleset#host Ruleset#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Override the destination port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/ruleset#port Ruleset#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersOrigin struct {
	// Origin Hostname where request is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#host Ruleset#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Origin Port where request is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#port Ruleset#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


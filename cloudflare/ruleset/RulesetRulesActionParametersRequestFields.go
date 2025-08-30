// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersRequestFields struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}


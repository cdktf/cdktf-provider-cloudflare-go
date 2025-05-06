// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersResponseFields struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to log duplicate values of the same header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#preserve_duplicates Ruleset#preserve_duplicates}
	PreserveDuplicates interface{} `field:"optional" json:"preserveDuplicates" yaml:"preserveDuplicates"`
}


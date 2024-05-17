// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersAlgorithms struct {
	// Name of the compression algorithm to use. Available values: `gzip`, `brotli`, `auto`, `default`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.33.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}


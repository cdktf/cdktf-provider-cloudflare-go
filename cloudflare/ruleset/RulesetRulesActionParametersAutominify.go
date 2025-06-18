// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersAutominify struct {
	// Minify CSS files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#css Ruleset#css}
	Css interface{} `field:"optional" json:"css" yaml:"css"`
	// Minify HTML files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#html Ruleset#html}
	Html interface{} `field:"optional" json:"html" yaml:"html"`
	// Minify JS files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#js Ruleset#js}
	Js interface{} `field:"optional" json:"js" yaml:"js"`
}


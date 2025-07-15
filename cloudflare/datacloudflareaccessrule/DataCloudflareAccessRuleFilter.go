// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccessrule


type DataCloudflareAccessRuleFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#configuration DataCloudflareAccessRule#configuration}.
	Configuration *DataCloudflareAccessRuleFilterConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Defines the direction used to sort returned rules. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#direction DataCloudflareAccessRule#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Defines the search requirements.
	//
	// When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#match DataCloudflareAccessRule#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// The action to apply to a matched request. Available values: "block", "challenge", "whitelist", "js_challenge", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#mode DataCloudflareAccessRule#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Defines the string to search for in the notes of existing IP Access rules.
	//
	// Notes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#notes DataCloudflareAccessRule#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Defines the field used to sort returned rules. Available values: "configuration.target", "configuration.value", "mode".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/access_rule#order DataCloudflareAccessRule#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}


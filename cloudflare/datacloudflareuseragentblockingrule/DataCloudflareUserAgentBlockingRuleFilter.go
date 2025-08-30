// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareuseragentblockingrule


type DataCloudflareUserAgentBlockingRuleFilter struct {
	// A string to search for in the description of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/user_agent_blocking_rule#description DataCloudflareUserAgentBlockingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When true, indicates that the rule is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/user_agent_blocking_rule#paused DataCloudflareUserAgentBlockingRule#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// A string to search for in the user agent values of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/user_agent_blocking_rule#user_agent DataCloudflareUserAgentBlockingRule#user_agent}
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}


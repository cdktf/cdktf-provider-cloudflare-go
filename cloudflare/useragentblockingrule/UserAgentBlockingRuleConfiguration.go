// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useragentblockingrule


type UserAgentBlockingRuleConfiguration struct {
	// The configuration target.
	//
	// You must set the target to `ua` when specifying a user agent in the rule.
	// Available values: "ua".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/user_agent_blocking_rule#target UserAgentBlockingRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// the user agent to exactly match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/user_agent_blocking_rule#value UserAgentBlockingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


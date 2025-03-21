// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useragentblockingrule


type UserAgentBlockingRuleConfiguration struct {
	// The configuration target.
	//
	// You must set the target to `ip` when specifying an IP address in the rule.
	// Available values: "ip".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/user_agent_blocking_rule#target UserAgentBlockingRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The IP address to match. This address will be compared to the IP address of incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/user_agent_blocking_rule#value UserAgentBlockingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


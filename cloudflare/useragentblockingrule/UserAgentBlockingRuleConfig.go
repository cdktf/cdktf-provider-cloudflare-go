// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useragentblockingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserAgentBlockingRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/user_agent_blocking_rule#configuration UserAgentBlockingRule#configuration}.
	Configuration *UserAgentBlockingRuleConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The action to apply to a matched request. Available values: "block", "challenge", "whitelist", "js_challenge", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/user_agent_blocking_rule#mode UserAgentBlockingRule#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/user_agent_blocking_rule#zone_id UserAgentBlockingRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// An informative summary of the rule. This value is sanitized and any tags will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/user_agent_blocking_rule#description UserAgentBlockingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When true, indicates that the rule is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/user_agent_blocking_rule#paused UserAgentBlockingRule#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
}


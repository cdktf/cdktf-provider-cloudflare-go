// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webanalyticsrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebAnalyticsRuleConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#account_id WebAnalyticsRule#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The Web Analytics ruleset identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#ruleset_id WebAnalyticsRule#ruleset_id}
	RulesetId *string `field:"required" json:"rulesetId" yaml:"rulesetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#host WebAnalyticsRule#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#inclusive WebAnalyticsRule#inclusive}
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
	// Whether the rule is paused or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#is_paused WebAnalyticsRule#is_paused}
	IsPaused interface{} `field:"optional" json:"isPaused" yaml:"isPaused"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/web_analytics_rule#paths WebAnalyticsRule#paths}.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}


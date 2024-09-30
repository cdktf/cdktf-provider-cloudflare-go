// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailRoutingRuleConfig struct {
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
	// Routing rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#name EmailRoutingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#zone_id EmailRoutingRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#action EmailRoutingRule#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Whether the email routing rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#enabled EmailRoutingRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#matcher EmailRoutingRule#matcher}
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// The priority of the email routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/email_routing_rule#priority EmailRoutingRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}


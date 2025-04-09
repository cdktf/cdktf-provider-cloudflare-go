// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PageRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#actions PageRule#actions}.
	Actions *PageRuleActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#target PageRule#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#zone_id PageRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The priority of the rule, used to define which Page Rule is processed over another.
	//
	// A higher number indicates a higher priority. For example,
	// if you have a catch-all Page Rule (rule A: `/images/*`) but want a more
	// specific Page Rule to take precedence (rule B: `/images/special/*`),
	// specify a higher priority for rule B so it overrides rule A.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#priority PageRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The status of the Page Rule. Available values: "active", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#status PageRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


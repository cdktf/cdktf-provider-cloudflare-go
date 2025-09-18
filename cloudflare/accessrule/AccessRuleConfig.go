// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessRuleConfig struct {
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
	// The rule configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/access_rule#configuration AccessRule#configuration}
	Configuration *AccessRuleConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The action to apply to a matched request. Available values: "block", "challenge", "whitelist", "js_challenge", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/access_rule#mode AccessRule#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/access_rule#account_id AccessRule#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/access_rule#notes AccessRule#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/access_rule#zone_id AccessRule#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


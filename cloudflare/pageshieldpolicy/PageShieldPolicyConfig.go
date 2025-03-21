// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pageshieldpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PageShieldPolicyConfig struct {
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
	// The action to take if the expression matches Available values: "allow", "log".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#action PageShieldPolicy#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// A description for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#description PageShieldPolicy#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Whether the policy is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#enabled PageShieldPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#expression PageShieldPolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The policy which will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#value PageShieldPolicy#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/page_shield_policy#zone_id PageShieldPolicy#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}


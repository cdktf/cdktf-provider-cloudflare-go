// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustriskbehavior

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustRiskBehaviorConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_risk_behavior#account_id ZeroTrustRiskBehavior#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_risk_behavior#behavior ZeroTrustRiskBehavior#behavior}
	Behavior interface{} `field:"optional" json:"behavior" yaml:"behavior"`
}

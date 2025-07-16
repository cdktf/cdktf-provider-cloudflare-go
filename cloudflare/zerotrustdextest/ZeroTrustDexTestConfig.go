// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdextest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDexTestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#account_id ZeroTrustDexTest#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The configuration object which contains the details for the WARP client to conduct the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#data ZeroTrustDexTest#data}
	Data *ZeroTrustDexTestData `field:"required" json:"data" yaml:"data"`
	// Determines whether or not the test is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#enabled ZeroTrustDexTest#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// How often the test will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#interval ZeroTrustDexTest#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The name of the DEX test. Must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#name ZeroTrustDexTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional details about the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#description ZeroTrustDexTest#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#targeted ZeroTrustDexTest#targeted}.
	Targeted interface{} `field:"optional" json:"targeted" yaml:"targeted"`
	// DEX rules targeted by this test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dex_test#target_policies ZeroTrustDexTest#target_policies}
	TargetPolicies interface{} `field:"optional" json:"targetPolicies" yaml:"targetPolicies"`
}


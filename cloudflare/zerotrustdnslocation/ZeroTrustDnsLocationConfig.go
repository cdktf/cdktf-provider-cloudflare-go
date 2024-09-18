// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDnsLocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#account_id ZeroTrustDnsLocation#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the teams location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#name ZeroTrustDnsLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicator that this is the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#client_default ZeroTrustDnsLocation#client_default}
	ClientDefault interface{} `field:"optional" json:"clientDefault" yaml:"clientDefault"`
	// Indicator that this location needs to resolve EDNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#ecs_support ZeroTrustDnsLocation#ecs_support}
	EcsSupport interface{} `field:"optional" json:"ecsSupport" yaml:"ecsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#id ZeroTrustDnsLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}


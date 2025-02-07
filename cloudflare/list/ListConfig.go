// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/list#account_id List#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/list#kind List#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// An informative name for the list. Use this name in filter and rule expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/list#name List#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An informative summary of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/list#description List#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package callsturnapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CallsTurnAppConfig struct {
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
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/calls_turn_app#account_id CallsTurnApp#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Cloudflare-generated unique identifier for a item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/calls_turn_app#key_id CallsTurnApp#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// A short description of a TURN key, not shown to end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/calls_turn_app#name CallsTurnApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


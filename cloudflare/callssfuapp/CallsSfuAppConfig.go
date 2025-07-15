// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package callssfuapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CallsSfuAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/calls_sfu_app#account_id CallsSfuApp#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Cloudflare-generated unique identifier for a item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/calls_sfu_app#app_id CallsSfuApp#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// A short description of Calls app, not shown to end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/calls_sfu_app#name CallsSfuApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


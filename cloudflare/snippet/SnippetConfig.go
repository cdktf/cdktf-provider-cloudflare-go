// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnippetConfig struct {
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
	// Main module file name of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/snippet#main_module Snippet#main_module}
	MainModule *string `field:"required" json:"mainModule" yaml:"mainModule"`
	// Name of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/snippet#name Snippet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/snippet#zone_id Snippet#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/snippet#files Snippet#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
}


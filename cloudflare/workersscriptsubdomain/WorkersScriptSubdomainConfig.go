// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscriptsubdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersScriptSubdomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script_subdomain#account_id WorkersScriptSubdomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether the Worker should be available on the workers.dev subdomain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script_subdomain#enabled WorkersScriptSubdomain#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Name of the script, used in URLs and route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script_subdomain#script_name WorkersScriptSubdomain#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// Whether the Worker's Preview URLs should be available on the workers.dev subdomain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script_subdomain#previews_enabled WorkersScriptSubdomain#previews_enabled}
	PreviewsEnabled interface{} `field:"optional" json:"previewsEnabled" yaml:"previewsEnabled"`
}


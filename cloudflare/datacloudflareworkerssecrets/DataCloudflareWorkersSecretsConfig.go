// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkerssecrets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWorkersSecretsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secrets#account_id DataCloudflareWorkersSecrets#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the Workers for Platforms dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secrets#dispatch_namespace DataCloudflareWorkersSecrets#dispatch_namespace}
	DispatchNamespace *string `field:"required" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// Name of the script, used in URLs and route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secrets#script_name DataCloudflareWorkersSecrets#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secrets#max_items DataCloudflareWorkersSecrets#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkerssecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWorkersSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secret#account_id DataCloudflareWorkersSecret#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the Workers for Platforms dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secret#dispatch_namespace DataCloudflareWorkersSecret#dispatch_namespace}
	DispatchNamespace *string `field:"required" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// Name of the script, used in URLs and route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secret#script_name DataCloudflareWorkersSecret#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// A JavaScript variable name for the secret binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/workers_secret#secret_name DataCloudflareWorkersSecret#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerssecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#account_id WorkersSecret#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the Workers for Platforms dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#dispatch_namespace WorkersSecret#dispatch_namespace}
	DispatchNamespace *string `field:"required" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// The name of this secret, this is what will be used to access it inside the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#name WorkersSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the script, used in URLs and route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#script_name WorkersSecret#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// The value of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#text WorkersSecret#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The type of secret to put. Available values: "secret_text".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_secret#type WorkersSecret#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


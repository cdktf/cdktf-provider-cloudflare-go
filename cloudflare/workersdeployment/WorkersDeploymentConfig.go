// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_deployment#account_id WorkersDeployment#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_deployment#script_name WorkersDeployment#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// Available values: "percentage".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_deployment#strategy WorkersDeployment#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_deployment#versions WorkersDeployment#versions}.
	Versions interface{} `field:"required" json:"versions" yaml:"versions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/workers_deployment#annotations WorkersDeployment#annotations}.
	Annotations *WorkersDeploymentAnnotations `field:"optional" json:"annotations" yaml:"annotations"`
}


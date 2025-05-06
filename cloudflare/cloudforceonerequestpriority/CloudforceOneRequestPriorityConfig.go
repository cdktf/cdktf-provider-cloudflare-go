// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudforceonerequestpriority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudforceOneRequestPriorityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request_priority#account_identifier CloudforceOneRequestPriority#account_identifier}
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// List of labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request_priority#labels CloudforceOneRequestPriority#labels}
	Labels *[]*string `field:"required" json:"labels" yaml:"labels"`
	// Priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request_priority#priority CloudforceOneRequestPriority#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request_priority#requirement CloudforceOneRequestPriority#requirement}
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
	// The CISA defined Traffic Light Protocol (TLP) Available values: "clear", "amber", "amber-strict", "green", "red".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request_priority#tlp CloudforceOneRequestPriority#tlp}
	Tlp *string `field:"required" json:"tlp" yaml:"tlp"`
}


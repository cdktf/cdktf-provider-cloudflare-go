// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudforceonerequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudforceOneRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#account_identifier CloudforceOneRequest#account_identifier}
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// Request content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#content CloudforceOneRequest#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Priority for analyzing the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#priority CloudforceOneRequest#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Requested information from request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#request_type CloudforceOneRequest#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// Brief description of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#summary CloudforceOneRequest#summary}
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
	// The CISA defined Traffic Light Protocol (TLP) Available values: "clear", "amber", "amber-strict", "green", "red".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/cloudforce_one_request#tlp CloudforceOneRequest#tlp}
	Tlp *string `field:"optional" json:"tlp" yaml:"tlp"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/cloudforce_one_request#account_id DataCloudflareCloudforceOneRequest#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/cloudforce_one_request#filter DataCloudflareCloudforceOneRequest#filter}.
	Filter *DataCloudflareCloudforceOneRequestFilter `field:"optional" json:"filter" yaml:"filter"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/cloudforce_one_request#request_id DataCloudflareCloudforceOneRequest#request_id}
	RequestId *string `field:"optional" json:"requestId" yaml:"requestId"`
}


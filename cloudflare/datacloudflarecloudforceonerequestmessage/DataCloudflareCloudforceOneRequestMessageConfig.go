// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequestmessage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestMessageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/cloudforce_one_request_message#account_identifier DataCloudflareCloudforceOneRequestMessage#account_identifier}
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/cloudforce_one_request_message#request_identifier DataCloudflareCloudforceOneRequestMessage#request_identifier}
	RequestIdentifier *string `field:"required" json:"requestIdentifier" yaml:"requestIdentifier"`
}


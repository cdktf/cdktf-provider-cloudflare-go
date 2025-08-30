// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package image

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#account_id Image#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// An optional custom unique identifier for your image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#id Image#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Can set the creator field with an internal user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#creator Image#creator}
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// An image binary data. Only needed when type is uploading a file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#file Image#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// User modifiable key-value store. Can use used for keeping references to another system of record for managing images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#metadata Image#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether the image requires a signature token for the access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#require_signed_urls Image#require_signed_urls}
	RequireSignedUrls interface{} `field:"optional" json:"requireSignedUrls" yaml:"requireSignedUrls"`
	// A URL to fetch an image from origin. Only needed when type is uploading from a URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/image#url Image#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}


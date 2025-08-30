// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type R2BucketConfig struct {
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
	// Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket#account_id R2Bucket#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket#name R2Bucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored. Available values: "default", "eu", "fedramp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket#jurisdiction R2Bucket#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Location of the bucket. Available values: "apac", "eeur", "enam", "weur", "wnam", "oc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket#location R2Bucket#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Storage class for newly uploaded objects, unless specified otherwise. Available values: "Standard", "InfrequentAccess".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket#storage_class R2Bucket#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}


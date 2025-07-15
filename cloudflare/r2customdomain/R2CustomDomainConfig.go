// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2customdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type R2CustomDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#account_id R2CustomDomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#bucket_name R2CustomDomain#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Name of the custom domain to be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#domain R2CustomDomain#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Whether to enable public bucket access at the custom domain. If undefined, the domain will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#enabled R2CustomDomain#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Zone ID of the custom domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#zone_id R2CustomDomain#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Jurisdiction of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#jurisdiction R2CustomDomain#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Minimum TLS Version the custom domain will accept for incoming connections.
	//
	// If not set, defaults to 1.0.
	// Available values: "1.0", "1.1", "1.2", "1.3".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/r2_custom_domain#min_tls R2CustomDomain#min_tls}
	MinTls *string `field:"optional" json:"minTls" yaml:"minTls"`
}


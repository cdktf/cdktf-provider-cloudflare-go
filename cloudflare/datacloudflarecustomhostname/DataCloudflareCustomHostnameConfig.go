// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomhostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCustomHostnameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_hostname#zone_id DataCloudflareCustomHostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_hostname#custom_hostname_id DataCloudflareCustomHostname#custom_hostname_id}
	CustomHostnameId *string `field:"optional" json:"customHostnameId" yaml:"customHostnameId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_hostname#filter DataCloudflareCustomHostname#filter}.
	Filter *DataCloudflareCustomHostnameFilter `field:"optional" json:"filter" yaml:"filter"`
}


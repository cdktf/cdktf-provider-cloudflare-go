// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequestasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestAssetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/cloudforce_one_request_asset#account_identifier DataCloudflareCloudforceOneRequestAsset#account_identifier}
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/cloudforce_one_request_asset#asset_identifer DataCloudflareCloudforceOneRequestAsset#asset_identifer}
	AssetIdentifer *string `field:"required" json:"assetIdentifer" yaml:"assetIdentifer"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/cloudforce_one_request_asset#request_identifier DataCloudflareCloudforceOneRequestAsset#request_identifier}
	RequestIdentifier *string `field:"required" json:"requestIdentifier" yaml:"requestIdentifier"`
}


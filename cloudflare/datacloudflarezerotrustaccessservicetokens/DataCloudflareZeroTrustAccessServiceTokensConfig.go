// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessservicetokens

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessServiceTokensConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/zero_trust_access_service_tokens#account_id DataCloudflareZeroTrustAccessServiceTokens#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/zero_trust_access_service_tokens#max_items DataCloudflareZeroTrustAccessServiceTokens#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The name of the service token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/zero_trust_access_service_tokens#name DataCloudflareZeroTrustAccessServiceTokens#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Search for service tokens by other listed query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/zero_trust_access_service_tokens#search DataCloudflareZeroTrustAccessServiceTokens#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/zero_trust_access_service_tokens#zone_id DataCloudflareZeroTrustAccessServiceTokens#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


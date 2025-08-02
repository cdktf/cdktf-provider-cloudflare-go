// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apitoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiTokenConfig struct {
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
	// Token name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#name ApiToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of access policies assigned to the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#policies ApiToken#policies}
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#condition ApiToken#condition}.
	Condition *ApiTokenCondition `field:"optional" json:"condition" yaml:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#expires_on ApiToken#expires_on}
	ExpiresOn *string `field:"optional" json:"expiresOn" yaml:"expiresOn"`
	// The time before which the token MUST NOT be accepted for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#not_before ApiToken#not_before}
	NotBefore *string `field:"optional" json:"notBefore" yaml:"notBefore"`
	// Status of the token. Available values: "active", "disabled", "expired".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/api_token#status ApiToken#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


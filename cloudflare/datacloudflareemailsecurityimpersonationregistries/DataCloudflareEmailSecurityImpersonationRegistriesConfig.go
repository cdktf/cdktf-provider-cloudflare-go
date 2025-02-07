// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityimpersonationregistries

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareEmailSecurityImpersonationRegistriesConfig struct {
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
	// Account Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#account_id DataCloudflareEmailSecurityImpersonationRegistries#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The sorting direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#direction DataCloudflareEmailSecurityImpersonationRegistries#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#max_items DataCloudflareEmailSecurityImpersonationRegistries#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The field to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#order DataCloudflareEmailSecurityImpersonationRegistries#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#provenance DataCloudflareEmailSecurityImpersonationRegistries#provenance}.
	Provenance *string `field:"optional" json:"provenance" yaml:"provenance"`
	// Allows searching in multiple properties of a record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact
	// behavior is intentionally left unspecified and is subject to change
	// in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_security_impersonation_registries#search DataCloudflareEmailSecurityImpersonationRegistries#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}


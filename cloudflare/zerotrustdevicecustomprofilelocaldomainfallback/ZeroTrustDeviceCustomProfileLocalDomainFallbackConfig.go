// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofilelocaldomainfallback

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceCustomProfileLocalDomainFallbackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#account_id ZeroTrustDeviceCustomProfileLocalDomainFallback#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#domains ZeroTrustDeviceCustomProfileLocalDomainFallback#domains}.
	Domains interface{} `field:"required" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#policy_id ZeroTrustDeviceCustomProfileLocalDomainFallback#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}


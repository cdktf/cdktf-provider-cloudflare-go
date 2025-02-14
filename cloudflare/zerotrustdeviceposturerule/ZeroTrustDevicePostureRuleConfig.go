// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDevicePostureRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#account_id ZeroTrustDevicePostureRule#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the device posture rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#name ZeroTrustDevicePostureRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of device posture rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#type ZeroTrustDevicePostureRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the device posture rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#description ZeroTrustDevicePostureRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Sets the expiration time for a posture check result.
	//
	// If empty, the result remains valid until it is overwritten by new data from the WARP client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#expiration ZeroTrustDevicePostureRule#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// The value to be checked against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#input ZeroTrustDevicePostureRule#input}
	Input *ZeroTrustDevicePostureRuleInput `field:"optional" json:"input" yaml:"input"`
	// The conditions that the client must match to run the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#match ZeroTrustDevicePostureRule#match}
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every five minutes). Minimum: `1m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#schedule ZeroTrustDevicePostureRule#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}


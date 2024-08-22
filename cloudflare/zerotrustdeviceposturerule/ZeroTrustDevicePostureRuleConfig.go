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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#account_id ZeroTrustDevicePostureRule#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The device posture rule type.
	//
	// Available values: `serial_number`, `file`, `application`, `gateway`, `warp`, `domain_joined`, `os_version`, `disk_encryption`, `firewall`, `client_certificate`, `client_certificate_v2`, `workspace_one`, `unique_client_id`, `crowdstrike_s2s`, `sentinelone`, `kolide`, `tanium_s2s`, `intune`, `sentinelone_s2s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#type ZeroTrustDevicePostureRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#description ZeroTrustDevicePostureRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Expire posture results after the specified amount of time.
	//
	// Must be in the format `1h` or `30m`. Valid units are `h` and `m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#expiration ZeroTrustDevicePostureRule#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#id ZeroTrustDevicePostureRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#input ZeroTrustDevicePostureRule#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#match ZeroTrustDevicePostureRule#match}
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// Name of the device posture rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#name ZeroTrustDevicePostureRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tells the client when to run the device posture check.
	//
	// Must be in the format `1h` or `30m`. Valid units are `h` and `m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_device_posture_rule#schedule ZeroTrustDevicePostureRule#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}


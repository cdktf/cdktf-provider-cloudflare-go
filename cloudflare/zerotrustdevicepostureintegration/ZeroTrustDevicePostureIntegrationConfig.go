// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicepostureintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDevicePostureIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_posture_integration#account_id ZeroTrustDevicePostureIntegration#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The configuration object containing third-party integration information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_posture_integration#config ZeroTrustDevicePostureIntegration#config}
	Config *ZeroTrustDevicePostureIntegrationConfigA `field:"required" json:"config" yaml:"config"`
	// The interval between each posture check with the third-party API.
	//
	// Use `m` for minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_posture_integration#interval ZeroTrustDevicePostureIntegration#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The name of the device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_posture_integration#name ZeroTrustDevicePostureIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of device posture integration. Available values: "workspace_one", "crowdstrike_s2s", "uptycs", "intune", "kolide", "tanium_s2s", "sentinelone_s2s", "custom_s2s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_posture_integration#type ZeroTrustDevicePostureIntegration#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}


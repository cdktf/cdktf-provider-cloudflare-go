// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infrastructureaccesstarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InfrastructureAccessTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/infrastructure_access_target#account_id InfrastructureAccessTarget#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A non-unique field that refers to a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/infrastructure_access_target#hostname InfrastructureAccessTarget#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The IPv4/IPv6 address that identifies where to reach a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/infrastructure_access_target#ip InfrastructureAccessTarget#ip}
	Ip *InfrastructureAccessTargetIp `field:"required" json:"ip" yaml:"ip"`
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apishielddiscoveryoperation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiShieldDiscoveryOperationConfig struct {
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
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/api_shield_discovery_operation#operation_id ApiShieldDiscoveryOperation#operation_id}
	OperationId *string `field:"required" json:"operationId" yaml:"operationId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/api_shield_discovery_operation#zone_id ApiShieldDiscoveryOperation#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Mark state of operation in API Discovery   * `review` - Mark operation as for review   * `ignored` - Mark operation as ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/api_shield_discovery_operation#state ApiShieldDiscoveryOperation#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}


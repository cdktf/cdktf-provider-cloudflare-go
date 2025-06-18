// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schemavalidationoperationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchemaValidationOperationSettingsConfig struct {
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
	// When set, this applies a mitigation action to this operation.
	//
	// - `"log"` - log request when request does not conform to schema for this operation
	//   - `"block"` - deny access to the site when request does not conform to schema for this operation
	//   - `"none"` - will skip mitigation for this operation
	//   - `null` - clears any mitigation action
	// Available values: "log", "block", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/schema_validation_operation_settings#mitigation_action SchemaValidationOperationSettings#mitigation_action}
	MitigationAction *string `field:"required" json:"mitigationAction" yaml:"mitigationAction"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/schema_validation_operation_settings#operation_id SchemaValidationOperationSettings#operation_id}
	OperationId *string `field:"required" json:"operationId" yaml:"operationId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/schema_validation_operation_settings#zone_id SchemaValidationOperationSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}


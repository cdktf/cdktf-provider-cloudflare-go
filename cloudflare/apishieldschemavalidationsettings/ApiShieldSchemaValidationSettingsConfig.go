// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apishieldschemavalidationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiShieldSchemaValidationSettingsConfig struct {
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
	// The default mitigation action used when there is no mitigation action defined on the operation.
	//
	// Mitigation actions are as follows:
	//
	//   * `log` - log request when request does not conform to schema
	//   * `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation
	// Available values: "none", "log", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/api_shield_schema_validation_settings#validation_default_mitigation_action ApiShieldSchemaValidationSettings#validation_default_mitigation_action}
	ValidationDefaultMitigationAction *string `field:"required" json:"validationDefaultMitigationAction" yaml:"validationDefaultMitigationAction"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/api_shield_schema_validation_settings#zone_id ApiShieldSchemaValidationSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	//   - `null` indicates that no override is in place
	//
	// To clear any override, use the special value `disable_override` or `null`
	// Available values: "none", "disable_override".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/api_shield_schema_validation_settings#validation_override_mitigation_action ApiShieldSchemaValidationSettings#validation_override_mitigation_action}
	ValidationOverrideMitigationAction *string `field:"optional" json:"validationOverrideMitigationAction" yaml:"validationOverrideMitigationAction"`
}


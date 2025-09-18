// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schemavalidationschemas

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchemaValidationSchemasConfig struct {
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
	// The kind of the schema Available values: "openapi_v3".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/schema_validation_schemas#kind SchemaValidationSchemas#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// A human-readable name for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/schema_validation_schemas#name SchemaValidationSchemas#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/schema_validation_schemas#source SchemaValidationSchemas#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// An indicator if this schema is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/schema_validation_schemas#validation_enabled SchemaValidationSchemas#validation_enabled}
	ValidationEnabled interface{} `field:"required" json:"validationEnabled" yaml:"validationEnabled"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/schema_validation_schemas#zone_id SchemaValidationSchemas#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}


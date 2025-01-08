// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationScimConfigMappings struct {
	// Which SCIM resource type this mapping applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#schema AccessApplication#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Whether or not this mapping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#enabled AccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#filter AccessApplication#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#operations AccessApplication#operations}
	Operations *AccessApplicationScimConfigMappingsOperations `field:"optional" json:"operations" yaml:"operations"`
	// How strictly to adhere to outbound resource schemas when provisioning to this mapping.
	//
	// "strict" will remove unknown values when provisioning, while "passthrough" will pass unknown values to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#strictness AccessApplication#strictness}
	Strictness *string `field:"optional" json:"strictness" yaml:"strictness"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#transform_jsonata AccessApplication#transform_jsonata}
	TransformJsonata *string `field:"optional" json:"transformJsonata" yaml:"transformJsonata"`
}


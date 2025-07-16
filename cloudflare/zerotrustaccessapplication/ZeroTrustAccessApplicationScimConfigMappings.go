// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfigMappings struct {
	// Which SCIM resource type this mapping applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#schema ZeroTrustAccessApplication#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Whether or not this mapping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#filter ZeroTrustAccessApplication#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Whether or not this mapping applies to creates, updates, or deletes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#operations ZeroTrustAccessApplication#operations}
	Operations *ZeroTrustAccessApplicationScimConfigMappingsOperations `field:"optional" json:"operations" yaml:"operations"`
	// The level of adherence to outbound resource schemas when provisioning to this mapping.
	//
	// ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown values to the target.
	// Available values: "strict", "passthrough".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#strictness ZeroTrustAccessApplication#strictness}
	Strictness *string `field:"optional" json:"strictness" yaml:"strictness"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#transform_jsonata ZeroTrustAccessApplication#transform_jsonata}
	TransformJsonata *string `field:"optional" json:"transformJsonata" yaml:"transformJsonata"`
}


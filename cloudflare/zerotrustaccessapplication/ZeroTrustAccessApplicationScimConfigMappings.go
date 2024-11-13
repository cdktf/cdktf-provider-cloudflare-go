// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfigMappings struct {
	// Which SCIM resource type this mapping applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#schema ZeroTrustAccessApplication#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Whether or not this mapping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#filter ZeroTrustAccessApplication#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#operations ZeroTrustAccessApplication#operations}
	Operations *ZeroTrustAccessApplicationScimConfigMappingsOperations `field:"optional" json:"operations" yaml:"operations"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#transform_jsonata ZeroTrustAccessApplication#transform_jsonata}
	TransformJsonata *string `field:"optional" json:"transformJsonata" yaml:"transformJsonata"`
}


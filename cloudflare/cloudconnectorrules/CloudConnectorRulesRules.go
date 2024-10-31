// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudconnectorrules


type CloudConnectorRulesRules struct {
	// Criteria for an HTTP request to trigger the cloud connector rule.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#expression CloudConnectorRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#parameters CloudConnectorRules#parameters}
	Parameters *CloudConnectorRulesRulesParameters `field:"required" json:"parameters" yaml:"parameters"`
	// Type of provider. Available values: `aws_s3`, `cloudflare_r2`, `azure_storage`, `gcp_storage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#provider CloudConnectorRules#provider}
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Brief summary of the cloud connector rule and its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#description CloudConnectorRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the headers rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#enabled CloudConnectorRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


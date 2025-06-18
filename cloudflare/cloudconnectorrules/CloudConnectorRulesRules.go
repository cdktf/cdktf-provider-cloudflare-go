// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudconnectorrules


type CloudConnectorRulesRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/cloud_connector_rules#description CloudConnectorRules#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/cloud_connector_rules#enabled CloudConnectorRules#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/cloud_connector_rules#expression CloudConnectorRules#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Parameters of Cloud Connector Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/cloud_connector_rules#parameters CloudConnectorRules#parameters}
	Parameters *CloudConnectorRulesRulesParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// Cloud Provider type Available values: "aws_s3", "cloudflare_r2", "gcp_storage", "azure_storage".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/cloud_connector_rules#provider CloudConnectorRules#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}


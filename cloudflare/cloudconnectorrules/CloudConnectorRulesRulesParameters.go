// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudconnectorrules


type CloudConnectorRulesRulesParameters struct {
	// Host parameter for cloud connector rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/cloud_connector_rules#host CloudConnectorRules#host}
	Host *string `field:"required" json:"host" yaml:"host"`
}


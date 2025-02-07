// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailroutingrule


type DataCloudflareEmailRoutingRuleFilter struct {
	// Filter by enabled routing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/email_routing_rule#enabled DataCloudflareEmailRoutingRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


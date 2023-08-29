// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActionsCacheKeyFieldsHost struct {
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.13.0/docs/resources/page_rule#resolved PageRule#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}


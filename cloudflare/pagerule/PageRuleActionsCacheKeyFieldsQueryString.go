// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActionsCacheKeyFieldsQueryString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/page_rule#exclude PageRule#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/page_rule#ignore PageRule#ignore}.
	Ignore interface{} `field:"optional" json:"ignore" yaml:"ignore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/page_rule#include PageRule#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}


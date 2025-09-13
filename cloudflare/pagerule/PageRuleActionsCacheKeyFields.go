// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActionsCacheKeyFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/page_rule#cookie PageRule#cookie}.
	Cookie *PageRuleActionsCacheKeyFieldsCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/page_rule#header PageRule#header}.
	Header *PageRuleActionsCacheKeyFieldsHeader `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/page_rule#host PageRule#host}.
	Host *PageRuleActionsCacheKeyFieldsHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/page_rule#query_string PageRule#query_string}.
	QueryString *PageRuleActionsCacheKeyFieldsQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/page_rule#user PageRule#user}.
	User *PageRuleActionsCacheKeyFieldsUser `field:"optional" json:"user" yaml:"user"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActionsCacheTtlByStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/page_rule#codes PageRule#codes}.
	Codes *string `field:"required" json:"codes" yaml:"codes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/page_rule#ttl PageRule#ttl}.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}


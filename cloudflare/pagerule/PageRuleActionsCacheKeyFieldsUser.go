// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActionsCacheKeyFieldsUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/page_rule#device_type PageRule#device_type}.
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/page_rule#geo PageRule#geo}.
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/page_rule#lang PageRule#lang}.
	Lang interface{} `field:"optional" json:"lang" yaml:"lang"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ratelimit


type RateLimitMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#headers RateLimit#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#request RateLimit#request}.
	Request *RateLimitMatchRequest `field:"optional" json:"request" yaml:"request"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#response RateLimit#response}.
	Response *RateLimitMatchResponse `field:"optional" json:"response" yaml:"response"`
}


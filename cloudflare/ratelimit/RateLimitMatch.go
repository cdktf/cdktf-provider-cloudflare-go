// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ratelimit


type RateLimitMatch struct {
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/rate_limit#request RateLimit#request}
	Request *RateLimitMatchRequest `field:"optional" json:"request" yaml:"request"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/rate_limit#response RateLimit#response}
	Response *RateLimitMatchResponse `field:"optional" json:"response" yaml:"response"`
}


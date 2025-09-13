// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ratelimit


type RateLimitMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count towards rate limiting.
	//
	// In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.
	// Notes: This field is deprecated. Instead, use response headers and set "origin_traffic" to "false" to avoid legacy behaviour interacting with the "response_headers" property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/rate_limit#origin_traffic RateLimit#origin_traffic}
	OriginTraffic interface{} `field:"optional" json:"originTraffic" yaml:"originTraffic"`
}


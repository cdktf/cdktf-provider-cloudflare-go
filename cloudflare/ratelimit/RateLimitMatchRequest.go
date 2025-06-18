// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ratelimit


type RateLimitMatchRequest struct {
	// The HTTP methods to match.
	//
	// You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/rate_limit#methods RateLimit#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// The HTTP schemes to match.
	//
	// You can specify one scheme (`['HTTPS']`), both schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/rate_limit#schemes RateLimit#schemes}
	Schemes *[]*string `field:"optional" json:"schemes" yaml:"schemes"`
	// The URL pattern to match, composed of a host and a path such as `example.org/path*`. Normalization is applied before the pattern is matched. `*` wildcards are expanded to match applicable traffic. Query strings are not matched. Set the value to `*` to match all traffic to your zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/rate_limit#url RateLimit#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}


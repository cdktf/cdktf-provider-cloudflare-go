// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketcors


type R2BucketCorsRules struct {
	// Object specifying allowed origins, methods and headers for this CORS rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_cors#allowed R2BucketCors#allowed}
	Allowed *R2BucketCorsRulesAllowed `field:"required" json:"allowed" yaml:"allowed"`
	// Specifies the headers that can be exposed back, and accessed by, the JavaScript making the cross-origin request.
	//
	// If you need to access headers beyond the safelisted response headers, such as Content-Encoding or cf-cache-status, you must specify it here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_cors#expose_headers R2BucketCors#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Identifier for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_cors#id R2BucketCors#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the amount of time (in seconds) browsers are allowed to cache CORS preflight responses.
	//
	// Browsers may limit this to 2 hours or less, even if the maximum value (86400) is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_cors#max_age_seconds R2BucketCors#max_age_seconds}
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}


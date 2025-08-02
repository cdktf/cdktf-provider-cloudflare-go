// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketcors


type R2BucketCorsRulesAllowed struct {
	// Specifies the value for the Access-Control-Allow-Methods header R2 sets when requesting objects in a bucket from a browser.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/r2_bucket_cors#methods R2BucketCors#methods}
	Methods *[]*string `field:"required" json:"methods" yaml:"methods"`
	// Specifies the value for the Access-Control-Allow-Origin header R2 sets when requesting objects in a bucket from a browser.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/r2_bucket_cors#origins R2BucketCors#origins}
	Origins *[]*string `field:"required" json:"origins" yaml:"origins"`
	// Specifies the value for the Access-Control-Allow-Headers header R2 sets when requesting objects in this bucket from a browser.
	//
	// Cross-origin requests that include custom headers (e.g. x-user-id) should specify these headers as AllowedHeaders.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/r2_bucket_cors#headers R2BucketCors#headers}
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}


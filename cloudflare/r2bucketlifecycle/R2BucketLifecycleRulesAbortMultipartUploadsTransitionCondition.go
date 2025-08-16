// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle


type R2BucketLifecycleRulesAbortMultipartUploadsTransitionCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/r2_bucket_lifecycle#max_age R2BucketLifecycle#max_age}.
	MaxAge *float64 `field:"required" json:"maxAge" yaml:"maxAge"`
	// Available values: "Age".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/r2_bucket_lifecycle#type R2BucketLifecycle#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle


type R2BucketLifecycleRulesStorageClassTransitionsCondition struct {
	// Available values: "Age", "Date".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#type R2BucketLifecycle#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#date R2BucketLifecycle#date}.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#max_age R2BucketLifecycle#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}


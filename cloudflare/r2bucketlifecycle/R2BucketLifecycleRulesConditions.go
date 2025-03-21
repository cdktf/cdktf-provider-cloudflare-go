// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle


type R2BucketLifecycleRulesConditions struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/r2_bucket_lifecycle#prefix R2BucketLifecycle#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle


type R2BucketLifecycleRulesStorageClassTransitions struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket_lifecycle#condition R2BucketLifecycle#condition}
	Condition *R2BucketLifecycleRulesStorageClassTransitionsCondition `field:"required" json:"condition" yaml:"condition"`
	// Available values: "InfrequentAccess".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/r2_bucket_lifecycle#storage_class R2BucketLifecycle#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}


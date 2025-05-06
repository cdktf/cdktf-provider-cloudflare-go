// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlock


type R2BucketLockRules struct {
	// Condition to apply a lock rule to an object for how long in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_lock#condition R2BucketLock#condition}
	Condition *R2BucketLockRulesCondition `field:"required" json:"condition" yaml:"condition"`
	// Whether or not this rule is in effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_lock#enabled R2BucketLock#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Unique identifier for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_lock#id R2BucketLock#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Rule will only apply to objects/uploads in the bucket that start with the given prefix, an empty prefix can be provided to scope rule to all objects/uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_lock#prefix R2BucketLock#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


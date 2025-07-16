// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle


type R2BucketLifecycleRules struct {
	// Conditions that apply to all transitions of this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#conditions R2BucketLifecycle#conditions}
	Conditions *R2BucketLifecycleRulesConditions `field:"required" json:"conditions" yaml:"conditions"`
	// Whether or not this rule is in effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#enabled R2BucketLifecycle#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Unique identifier for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#id R2BucketLifecycle#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Transition to abort ongoing multipart uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#abort_multipart_uploads_transition R2BucketLifecycle#abort_multipart_uploads_transition}
	AbortMultipartUploadsTransition *R2BucketLifecycleRulesAbortMultipartUploadsTransition `field:"optional" json:"abortMultipartUploadsTransition" yaml:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#delete_objects_transition R2BucketLifecycle#delete_objects_transition}
	DeleteObjectsTransition *R2BucketLifecycleRulesDeleteObjectsTransition `field:"optional" json:"deleteObjectsTransition" yaml:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#storage_class_transitions R2BucketLifecycle#storage_class_transitions}
	StorageClassTransitions interface{} `field:"optional" json:"storageClassTransitions" yaml:"storageClassTransitions"`
}


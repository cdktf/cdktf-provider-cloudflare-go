// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketsippy


type R2BucketSippySource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#access_key_id R2BucketSippy#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Name of the AWS S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#bucket R2BucketSippy#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#client_email R2BucketSippy#client_email}
	ClientEmail *string `field:"optional" json:"clientEmail" yaml:"clientEmail"`
	// Available values: "aws", "gcs".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#cloud_provider R2BucketSippy#cloud_provider}
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#private_key R2BucketSippy#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Name of the AWS availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#region R2BucketSippy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/r2_bucket_sippy#secret_access_key R2BucketSippy#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}


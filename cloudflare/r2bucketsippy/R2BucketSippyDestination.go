// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketsippy


type R2BucketSippyDestination struct {
	// ID of a Cloudflare API token.
	//
	// This is the value labelled "Access Key ID" when creating an API.
	// token from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is
	// best to scope this token to the bucket you're enabling Sippy for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/r2_bucket_sippy#access_key_id R2BucketSippy#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Available values: "r2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/r2_bucket_sippy#cloud_provider R2BucketSippy#cloud_provider}
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Value of a Cloudflare API token.
	//
	// This is the value labelled "Secret Access Key" when creating an API.
	// token from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is
	// best to scope this token to the bucket you're enabling Sippy for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/r2_bucket_sippy#secret_access_key R2BucketSippy#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}


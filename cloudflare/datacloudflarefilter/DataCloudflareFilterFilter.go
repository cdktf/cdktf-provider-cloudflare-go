// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarefilter


type DataCloudflareFilterFilter struct {
	// A case-insensitive string to find in the description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/filter#description DataCloudflareFilter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A case-insensitive string to find in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/filter#expression DataCloudflareFilter#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The unique identifier of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/filter#id DataCloudflareFilter#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, indicates that the filter is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/filter#paused DataCloudflareFilter#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The filter ref (a short reference tag) to search for. Must be an exact match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/data-sources/filter#ref DataCloudflareFilter#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}


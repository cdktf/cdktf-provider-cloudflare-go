// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccount


type DataCloudflareAccountFilter struct {
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/account#direction DataCloudflareAccount#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Name of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/account#name DataCloudflareAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


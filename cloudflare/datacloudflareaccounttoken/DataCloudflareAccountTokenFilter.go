// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccounttoken


type DataCloudflareAccountTokenFilter struct {
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_token#direction DataCloudflareAccountToken#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}


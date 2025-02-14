// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountmember


type DataCloudflareAccountMemberFilter struct {
	// Direction to order results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/account_member#direction DataCloudflareAccountMember#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/account_member#order DataCloudflareAccountMember#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A member's status in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/account_member#status DataCloudflareAccountMember#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


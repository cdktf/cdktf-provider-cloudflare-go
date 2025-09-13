// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountmember


type DataCloudflareAccountMemberFilter struct {
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_member#direction DataCloudflareAccountMember#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order results by. Available values: "user.first_name", "user.last_name", "user.email", "status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_member#order DataCloudflareAccountMember#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A member's status in the account. Available values: "accepted", "pending", "rejected".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/account_member#status DataCloudflareAccountMember#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


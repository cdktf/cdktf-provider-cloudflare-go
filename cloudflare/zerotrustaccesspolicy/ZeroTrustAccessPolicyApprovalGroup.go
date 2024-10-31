// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyApprovalGroup struct {
	// Number of approvals needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_access_policy#approvals_needed ZeroTrustAccessPolicy#approvals_needed}
	ApprovalsNeeded *float64 `field:"required" json:"approvalsNeeded" yaml:"approvalsNeeded"`
	// List of emails to request approval from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_access_policy#email_addresses ZeroTrustAccessPolicy#email_addresses}
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_access_policy#email_list_uuid ZeroTrustAccessPolicy#email_list_uuid}.
	EmailListUuid *string `field:"optional" json:"emailListUuid" yaml:"emailListUuid"`
}


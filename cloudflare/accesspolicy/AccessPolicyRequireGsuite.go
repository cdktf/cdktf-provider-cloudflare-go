// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyRequireGsuite struct {
	// The email of the Google Workspace group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/access_policy#email AccessPolicy#email}
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// The ID of your Google Workspace identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}


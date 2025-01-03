// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyIncludeGsuite struct {
	// The email of the Google Workspace group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/access_policy#email AccessPolicy#email}
	Email *[]*string `field:"required" json:"email" yaml:"email"`
	// The ID of your Google Workspace identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessgroup


type AccessGroupIncludeGsuite struct {
	// The email of the Google Workspace group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/access_group#email AccessGroup#email}
	Email *[]*string `field:"required" json:"email" yaml:"email"`
	// The ID of your Google Workspace identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/access_group#identity_provider_id AccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}


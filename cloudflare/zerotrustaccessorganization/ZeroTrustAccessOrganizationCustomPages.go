// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessorganization


type ZeroTrustAccessOrganizationCustomPages struct {
	// The id of the forbidden page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_organization#forbidden ZeroTrustAccessOrganization#forbidden}
	Forbidden *string `field:"optional" json:"forbidden" yaml:"forbidden"`
	// The id of the identity denied page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_organization#identity_denied ZeroTrustAccessOrganization#identity_denied}
	IdentityDenied *string `field:"optional" json:"identityDenied" yaml:"identityDenied"`
}


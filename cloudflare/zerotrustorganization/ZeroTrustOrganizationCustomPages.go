// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization


type ZeroTrustOrganizationCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a non-identity rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_organization#forbidden ZeroTrustOrganization#forbidden}
	Forbidden *string `field:"optional" json:"forbidden" yaml:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_organization#identity_denied ZeroTrustOrganization#identity_denied}
	IdentityDenied *string `field:"optional" json:"identityDenied" yaml:"identityDenied"`
}


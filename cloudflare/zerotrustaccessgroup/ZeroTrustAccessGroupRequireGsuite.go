// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireGsuite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_access_group#email ZeroTrustAccessGroup#email}.
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}

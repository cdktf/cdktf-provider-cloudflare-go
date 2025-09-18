// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeOidc struct {
	// The name of the OIDC claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_access_group#claim_name ZeroTrustAccessGroup#claim_name}
	ClaimName *string `field:"required" json:"claimName" yaml:"claimName"`
	// The OIDC claim value to look for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_access_group#claim_value ZeroTrustAccessGroup#claim_value}
	ClaimValue *string `field:"required" json:"claimValue" yaml:"claimValue"`
	// The ID of your OIDC identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_access_group#identity_provider_id ZeroTrustAccessGroup#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}


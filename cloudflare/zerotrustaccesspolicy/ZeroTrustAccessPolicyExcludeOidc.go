// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeOidc struct {
	// The name of the OIDC claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_policy#claim_name ZeroTrustAccessPolicy#claim_name}
	ClaimName *string `field:"required" json:"claimName" yaml:"claimName"`
	// The OIDC claim value to look for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_policy#claim_value ZeroTrustAccessPolicy#claim_value}
	ClaimValue *string `field:"required" json:"claimValue" yaml:"claimValue"`
	// The ID of your OIDC identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_policy#identity_provider_id ZeroTrustAccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}


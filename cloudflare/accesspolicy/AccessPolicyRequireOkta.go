// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyRequireOkta struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.39.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.39.0/docs/resources/access_policy#name AccessPolicy#name}.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessgroup


type AccessGroupRequireOkta struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.19.0/docs/resources/access_group#identity_provider_id AccessGroup#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.19.0/docs/resources/access_group#name AccessGroup#name}.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyRequireAuthContext struct {
	// The ACID of the Authentication Context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.38.0/docs/resources/access_policy#ac_id AccessPolicy#ac_id}
	AcId *string `field:"required" json:"acId" yaml:"acId"`
	// The ID of the Authentication Context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.38.0/docs/resources/access_policy#id AccessPolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the Azure Identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.38.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
}


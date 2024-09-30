// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyRequireSaml struct {
	// The name of the SAML attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/access_policy#attribute_name AccessPolicy#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// The SAML attribute value to look for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/access_policy#attribute_value AccessPolicy#attribute_value}
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// The ID of your SAML identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}


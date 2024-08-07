// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessgroup


type AccessGroupExcludeSaml struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.39.0/docs/resources/access_group#attribute_name AccessGroup#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.39.0/docs/resources/access_group#attribute_value AccessGroup#attribute_value}.
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.39.0/docs/resources/access_group#identity_provider_id AccessGroup#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}


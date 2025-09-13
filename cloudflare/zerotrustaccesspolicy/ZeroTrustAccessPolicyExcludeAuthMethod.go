// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_policy#auth_method ZeroTrustAccessPolicy#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
}


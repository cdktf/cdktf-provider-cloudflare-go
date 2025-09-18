// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesIncludeAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_access_application#auth_method ZeroTrustAccessApplication#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
}


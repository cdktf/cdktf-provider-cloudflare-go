// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_application#return_access_token_from_authorization_endpoint ZeroTrustAccessApplication#return_access_token_from_authorization_endpoint}
	ReturnAccessTokenFromAuthorizationEndpoint interface{} `field:"optional" json:"returnAccessTokenFromAuthorizationEndpoint" yaml:"returnAccessTokenFromAuthorizationEndpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_application#return_id_token_from_authorization_endpoint ZeroTrustAccessApplication#return_id_token_from_authorization_endpoint}
	ReturnIdTokenFromAuthorizationEndpoint interface{} `field:"optional" json:"returnIdTokenFromAuthorizationEndpoint" yaml:"returnIdTokenFromAuthorizationEndpoint"`
}


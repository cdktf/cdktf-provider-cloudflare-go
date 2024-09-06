// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasAppHybridAndImplicitOptions struct {
	// If true, the authorization endpoint will return an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/access_application#return_access_token_from_authorization_endpoint AccessApplication#return_access_token_from_authorization_endpoint}
	ReturnAccessTokenFromAuthorizationEndpoint interface{} `field:"optional" json:"returnAccessTokenFromAuthorizationEndpoint" yaml:"returnAccessTokenFromAuthorizationEndpoint"`
	// If true, the authorization endpoint will return an id token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/access_application#return_id_token_from_authorization_endpoint AccessApplication#return_id_token_from_authorization_endpoint}
	ReturnIdTokenFromAuthorizationEndpoint interface{} `field:"optional" json:"returnIdTokenFromAuthorizationEndpoint" yaml:"returnIdTokenFromAuthorizationEndpoint"`
}


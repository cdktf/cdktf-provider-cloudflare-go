// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfigAuthentication struct {
	// The authentication scheme to use when making SCIM requests to this application. Available values: "httpbasic", "oauthbearertoken", "oauth2", "access_service_token".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#scheme ZeroTrustAccessApplication#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// URL used to generate the auth code used during token generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#authorization_url ZeroTrustAccessApplication#authorization_url}
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#client_id ZeroTrustAccessApplication#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Secret used to authenticate when generating a token for authenticating with the remove SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#client_secret ZeroTrustAccessApplication#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Password used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#password ZeroTrustAccessApplication#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#scopes ZeroTrustAccessApplication#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Token used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#token ZeroTrustAccessApplication#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// URL used to generate the token used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#token_url ZeroTrustAccessApplication#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// User name used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_application#user ZeroTrustAccessApplication#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider


type ZeroTrustAccessIdentityProviderConfigA struct {
	// Your companies TLD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#apps_domain ZeroTrustAccessIdentityProvider#apps_domain}
	AppsDomain *string `field:"optional" json:"appsDomain" yaml:"appsDomain"`
	// A list of SAML attribute names that will be added to your signed JWT token and can be used in SAML policy rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#attributes ZeroTrustAccessIdentityProvider#attributes}
	Attributes *[]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Your okta authorization server id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#authorization_server_id ZeroTrustAccessIdentityProvider#authorization_server_id}
	AuthorizationServerId *string `field:"optional" json:"authorizationServerId" yaml:"authorizationServerId"`
	// The authorization_endpoint URL of your IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#auth_url ZeroTrustAccessIdentityProvider#auth_url}
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Your centrify account url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#centrify_account ZeroTrustAccessIdentityProvider#centrify_account}
	CentrifyAccount *string `field:"optional" json:"centrifyAccount" yaml:"centrifyAccount"`
	// Your centrify app id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#centrify_app_id ZeroTrustAccessIdentityProvider#centrify_app_id}
	CentrifyAppId *string `field:"optional" json:"centrifyAppId" yaml:"centrifyAppId"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#certs_url ZeroTrustAccessIdentityProvider#certs_url}
	CertsUrl *string `field:"optional" json:"certsUrl" yaml:"certsUrl"`
	// Custom claims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#claims ZeroTrustAccessIdentityProvider#claims}
	Claims *[]*string `field:"optional" json:"claims" yaml:"claims"`
	// Your OAuth Client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#client_id ZeroTrustAccessIdentityProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Your OAuth Client Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#client_secret ZeroTrustAccessIdentityProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Should Cloudflare try to load authentication contexts from your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#conditional_access_enabled ZeroTrustAccessIdentityProvider#conditional_access_enabled}
	ConditionalAccessEnabled interface{} `field:"optional" json:"conditionalAccessEnabled" yaml:"conditionalAccessEnabled"`
	// Your Azure directory uuid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#directory_id ZeroTrustAccessIdentityProvider#directory_id}
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The attribute name for email in the SAML response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#email_attribute_name ZeroTrustAccessIdentityProvider#email_attribute_name}
	EmailAttributeName *string `field:"optional" json:"emailAttributeName" yaml:"emailAttributeName"`
	// The claim name for email in the id_token response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#email_claim_name ZeroTrustAccessIdentityProvider#email_claim_name}
	EmailClaimName *string `field:"optional" json:"emailClaimName" yaml:"emailClaimName"`
	// Add a list of attribute names that will be returned in the response header from the Access callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#header_attributes ZeroTrustAccessIdentityProvider#header_attributes}
	HeaderAttributes interface{} `field:"optional" json:"headerAttributes" yaml:"headerAttributes"`
	// X509 certificate to verify the signature in the SAML authentication response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#idp_public_certs ZeroTrustAccessIdentityProvider#idp_public_certs}
	IdpPublicCerts *[]*string `field:"optional" json:"idpPublicCerts" yaml:"idpPublicCerts"`
	// IdP Entity ID or Issuer URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#issuer_url ZeroTrustAccessIdentityProvider#issuer_url}
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// Your okta account url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#okta_account ZeroTrustAccessIdentityProvider#okta_account}
	OktaAccount *string `field:"optional" json:"oktaAccount" yaml:"oktaAccount"`
	// Your OneLogin account url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#onelogin_account ZeroTrustAccessIdentityProvider#onelogin_account}
	OneloginAccount *string `field:"optional" json:"oneloginAccount" yaml:"oneloginAccount"`
	// Your PingOne environment identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#ping_env_id ZeroTrustAccessIdentityProvider#ping_env_id}
	PingEnvId *string `field:"optional" json:"pingEnvId" yaml:"pingEnvId"`
	// Enable Proof Key for Code Exchange (PKCE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#pkce_enabled ZeroTrustAccessIdentityProvider#pkce_enabled}
	PkceEnabled interface{} `field:"optional" json:"pkceEnabled" yaml:"pkceEnabled"`
	// Indicates the type of user interaction that is required.
	//
	// prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.
	// Available values: "login", "select_account", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#prompt ZeroTrustAccessIdentityProvider#prompt}
	Prompt *string `field:"optional" json:"prompt" yaml:"prompt"`
	// OAuth scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#scopes ZeroTrustAccessIdentityProvider#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Sign the SAML authentication request with Access credentials.
	//
	// To verify the signature, use the public key from the Access certs endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#sign_request ZeroTrustAccessIdentityProvider#sign_request}
	SignRequest interface{} `field:"optional" json:"signRequest" yaml:"signRequest"`
	// URL to send the SAML authentication requests to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#sso_target_url ZeroTrustAccessIdentityProvider#sso_target_url}
	SsoTargetUrl *string `field:"optional" json:"ssoTargetUrl" yaml:"ssoTargetUrl"`
	// Should Cloudflare try to load groups from your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#support_groups ZeroTrustAccessIdentityProvider#support_groups}
	SupportGroups interface{} `field:"optional" json:"supportGroups" yaml:"supportGroups"`
	// The token_endpoint URL of your IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_access_identity_provider#token_url ZeroTrustAccessIdentityProvider#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}


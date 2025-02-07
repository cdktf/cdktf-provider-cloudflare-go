// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasApp struct {
	// The lifetime of the OIDC Access Token after creation.
	//
	// Valid units are m,h. Must be greater than or equal to 1m and less than or equal to 24h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#access_token_lifetime ZeroTrustAccessApplication#access_token_lifetime}
	AccessTokenLifetime *string `field:"optional" json:"accessTokenLifetime" yaml:"accessTokenLifetime"`
	// If client secret should be required on the token endpoint when authorization_code_with_pkce grant is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#allow_pkce_without_client_secret ZeroTrustAccessApplication#allow_pkce_without_client_secret}
	AllowPkceWithoutClientSecret interface{} `field:"optional" json:"allowPkceWithoutClientSecret" yaml:"allowPkceWithoutClientSecret"`
	// The URL where this applications tile redirects users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#app_launcher_url ZeroTrustAccessApplication#app_launcher_url}
	AppLauncherUrl *string `field:"optional" json:"appLauncherUrl" yaml:"appLauncherUrl"`
	// Optional identifier indicating the authentication protocol used for the saas app. Required for OIDC. Default if unset is "saml".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#auth_type ZeroTrustAccessApplication#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The application client id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#client_id ZeroTrustAccessApplication#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The application client secret, only returned on POST request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#client_secret ZeroTrustAccessApplication#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#consumer_service_url ZeroTrustAccessApplication#consumer_service_url}
	ConsumerServiceUrl *string `field:"optional" json:"consumerServiceUrl" yaml:"consumerServiceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#custom_attributes ZeroTrustAccessApplication#custom_attributes}.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#custom_claims ZeroTrustAccessApplication#custom_claims}.
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// The URL that the user will be redirected to after a successful login for IDP initiated logins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#default_relay_state ZeroTrustAccessApplication#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// The OIDC flows supported by this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#grant_types ZeroTrustAccessApplication#grant_types}
	GrantTypes *[]*string `field:"optional" json:"grantTypes" yaml:"grantTypes"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#group_filter_regex ZeroTrustAccessApplication#group_filter_regex}
	GroupFilterRegex *string `field:"optional" json:"groupFilterRegex" yaml:"groupFilterRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#hybrid_and_implicit_options ZeroTrustAccessApplication#hybrid_and_implicit_options}.
	HybridAndImplicitOptions *ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions `field:"optional" json:"hybridAndImplicitOptions" yaml:"hybridAndImplicitOptions"`
	// The unique identifier for your SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#idp_entity_id ZeroTrustAccessApplication#idp_entity_id}
	IdpEntityId *string `field:"optional" json:"idpEntityId" yaml:"idpEntityId"`
	// The format of the name identifier sent to the SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#name_id_format ZeroTrustAccessApplication#name_id_format}
	NameIdFormat *string `field:"optional" json:"nameIdFormat" yaml:"nameIdFormat"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#name_id_transform_jsonata ZeroTrustAccessApplication#name_id_transform_jsonata}
	NameIdTransformJsonata *string `field:"optional" json:"nameIdTransformJsonata" yaml:"nameIdTransformJsonata"`
	// The Access public certificate that will be used to verify your identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#public_key ZeroTrustAccessApplication#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#redirect_uris ZeroTrustAccessApplication#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#refresh_token_options ZeroTrustAccessApplication#refresh_token_options}.
	RefreshTokenOptions *ZeroTrustAccessApplicationSaasAppRefreshTokenOptions `field:"optional" json:"refreshTokenOptions" yaml:"refreshTokenOptions"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#saml_attribute_transform_jsonata ZeroTrustAccessApplication#saml_attribute_transform_jsonata}
	SamlAttributeTransformJsonata *string `field:"optional" json:"samlAttributeTransformJsonata" yaml:"samlAttributeTransformJsonata"`
	// Define the user information shared with access, "offline_access" scope will be automatically enabled if refresh tokens are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#scopes ZeroTrustAccessApplication#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#sp_entity_id ZeroTrustAccessApplication#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
	// The endpoint where your SaaS application will send login requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#sso_endpoint ZeroTrustAccessApplication#sso_endpoint}
	SsoEndpoint *string `field:"optional" json:"ssoEndpoint" yaml:"ssoEndpoint"`
}


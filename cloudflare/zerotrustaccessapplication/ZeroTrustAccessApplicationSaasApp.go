// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasApp struct {
	// The lifetime of the Access Token after creation.
	//
	// Valid units are `m` and `h`. Must be greater than or equal to 1m and less than or equal to 24h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#access_token_lifetime ZeroTrustAccessApplication#access_token_lifetime}
	AccessTokenLifetime *string `field:"optional" json:"accessTokenLifetime" yaml:"accessTokenLifetime"`
	// Allow PKCE flow without a client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#allow_pkce_without_client_secret ZeroTrustAccessApplication#allow_pkce_without_client_secret}
	AllowPkceWithoutClientSecret interface{} `field:"optional" json:"allowPkceWithoutClientSecret" yaml:"allowPkceWithoutClientSecret"`
	// The URL where this applications tile redirects users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#app_launcher_url ZeroTrustAccessApplication#app_launcher_url}
	AppLauncherUrl *string `field:"optional" json:"appLauncherUrl" yaml:"appLauncherUrl"`
	// **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#auth_type ZeroTrustAccessApplication#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#consumer_service_url ZeroTrustAccessApplication#consumer_service_url}
	ConsumerServiceUrl *string `field:"optional" json:"consumerServiceUrl" yaml:"consumerServiceUrl"`
	// custom_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#custom_attribute ZeroTrustAccessApplication#custom_attribute}
	CustomAttribute interface{} `field:"optional" json:"customAttribute" yaml:"customAttribute"`
	// custom_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#custom_claim ZeroTrustAccessApplication#custom_claim}
	CustomClaim interface{} `field:"optional" json:"customClaim" yaml:"customClaim"`
	// The relay state used if not provided by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#default_relay_state ZeroTrustAccessApplication#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// The OIDC flows supported by this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#grant_types ZeroTrustAccessApplication#grant_types}
	GrantTypes *[]*string `field:"optional" json:"grantTypes" yaml:"grantTypes"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#group_filter_regex ZeroTrustAccessApplication#group_filter_regex}
	GroupFilterRegex *string `field:"optional" json:"groupFilterRegex" yaml:"groupFilterRegex"`
	// hybrid_and_implicit_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#hybrid_and_implicit_options ZeroTrustAccessApplication#hybrid_and_implicit_options}
	HybridAndImplicitOptions *ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions `field:"optional" json:"hybridAndImplicitOptions" yaml:"hybridAndImplicitOptions"`
	// The format of the name identifier sent to the SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#name_id_format ZeroTrustAccessApplication#name_id_format}
	NameIdFormat *string `field:"optional" json:"nameIdFormat" yaml:"nameIdFormat"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#name_id_transform_jsonata ZeroTrustAccessApplication#name_id_transform_jsonata}
	NameIdTransformJsonata *string `field:"optional" json:"nameIdTransformJsonata" yaml:"nameIdTransformJsonata"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#redirect_uris ZeroTrustAccessApplication#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
	// refresh_token_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#refresh_token_options ZeroTrustAccessApplication#refresh_token_options}
	RefreshTokenOptions interface{} `field:"optional" json:"refreshTokenOptions" yaml:"refreshTokenOptions"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#saml_attribute_transform_jsonata ZeroTrustAccessApplication#saml_attribute_transform_jsonata}
	SamlAttributeTransformJsonata *string `field:"optional" json:"samlAttributeTransformJsonata" yaml:"samlAttributeTransformJsonata"`
	// Define the user information shared with access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#scopes ZeroTrustAccessApplication#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_access_application#sp_entity_id ZeroTrustAccessApplication#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
}


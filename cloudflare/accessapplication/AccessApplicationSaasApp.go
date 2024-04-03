// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasApp struct {
	// The URL where this applications tile redirects users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#app_launcher_url AccessApplication#app_launcher_url}
	AppLauncherUrl *string `field:"optional" json:"appLauncherUrl" yaml:"appLauncherUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#auth_type AccessApplication#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#consumer_service_url AccessApplication#consumer_service_url}
	ConsumerServiceUrl *string `field:"optional" json:"consumerServiceUrl" yaml:"consumerServiceUrl"`
	// custom_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#custom_attribute AccessApplication#custom_attribute}
	CustomAttribute interface{} `field:"optional" json:"customAttribute" yaml:"customAttribute"`
	// The relay state used if not provided by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#default_relay_state AccessApplication#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// The OIDC flows supported by this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#grant_types AccessApplication#grant_types}
	GrantTypes *[]*string `field:"optional" json:"grantTypes" yaml:"grantTypes"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#group_filter_regex AccessApplication#group_filter_regex}
	GroupFilterRegex *string `field:"optional" json:"groupFilterRegex" yaml:"groupFilterRegex"`
	// The format of the name identifier sent to the SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#name_id_format AccessApplication#name_id_format}
	NameIdFormat *string `field:"optional" json:"nameIdFormat" yaml:"nameIdFormat"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#name_id_transform_jsonata AccessApplication#name_id_transform_jsonata}
	NameIdTransformJsonata *string `field:"optional" json:"nameIdTransformJsonata" yaml:"nameIdTransformJsonata"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#redirect_uris AccessApplication#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#saml_attribute_transform_jsonata AccessApplication#saml_attribute_transform_jsonata}
	SamlAttributeTransformJsonata *string `field:"optional" json:"samlAttributeTransformJsonata" yaml:"samlAttributeTransformJsonata"`
	// Define the user information shared with access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#scopes AccessApplication#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.29.0/docs/resources/access_application#sp_entity_id AccessApplication#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
}


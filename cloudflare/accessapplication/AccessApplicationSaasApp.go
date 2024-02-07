// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasApp struct {
	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/access_application#consumer_service_url AccessApplication#consumer_service_url}
	ConsumerServiceUrl *string `field:"required" json:"consumerServiceUrl" yaml:"consumerServiceUrl"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/access_application#sp_entity_id AccessApplication#sp_entity_id}
	SpEntityId *string `field:"required" json:"spEntityId" yaml:"spEntityId"`
	// custom_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/access_application#custom_attribute AccessApplication#custom_attribute}
	CustomAttribute interface{} `field:"optional" json:"customAttribute" yaml:"customAttribute"`
	// The relay state used if not provided by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/access_application#default_relay_state AccessApplication#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// The format of the name identifier sent to the SaaS application. Defaults to `email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/access_application#name_id_format AccessApplication#name_id_format}
	NameIdFormat *string `field:"optional" json:"nameIdFormat" yaml:"nameIdFormat"`
}


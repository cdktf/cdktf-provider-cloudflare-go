// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfig struct {
	// The UIDs of the IdP to use as the source for SCIM resources to provision to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#idp_uid ZeroTrustAccessApplication#idp_uid}
	IdpUid *string `field:"required" json:"idpUid" yaml:"idpUid"`
	// The base URI for the application's SCIM-compatible API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#remote_uri ZeroTrustAccessApplication#remote_uri}
	RemoteUri *string `field:"required" json:"remoteUri" yaml:"remoteUri"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#authentication ZeroTrustAccessApplication#authentication}
	Authentication *ZeroTrustAccessApplicationScimConfigAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM resources.
	//
	// If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#deactivate_on_delete ZeroTrustAccessApplication#deactivate_on_delete}
	DeactivateOnDelete interface{} `field:"optional" json:"deactivateOnDelete" yaml:"deactivateOnDelete"`
	// Whether SCIM provisioning is turned on for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#mappings ZeroTrustAccessApplication#mappings}
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
}


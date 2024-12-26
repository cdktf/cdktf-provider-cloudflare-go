// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationScimConfig struct {
	// The UIDs of the IdP to use as the source for SCIM resources to provision to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#idp_uid AccessApplication#idp_uid}
	IdpUid *string `field:"required" json:"idpUid" yaml:"idpUid"`
	// The base URI for the application's SCIM-compatible API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#remote_uri AccessApplication#remote_uri}
	RemoteUri *string `field:"required" json:"remoteUri" yaml:"remoteUri"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#authentication AccessApplication#authentication}
	Authentication interface{} `field:"optional" json:"authentication" yaml:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM resources.
	//
	// If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#deactivate_on_delete AccessApplication#deactivate_on_delete}
	DeactivateOnDelete interface{} `field:"optional" json:"deactivateOnDelete" yaml:"deactivateOnDelete"`
	// Whether SCIM provisioning is turned on for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#enabled AccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#mappings AccessApplication#mappings}
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
}


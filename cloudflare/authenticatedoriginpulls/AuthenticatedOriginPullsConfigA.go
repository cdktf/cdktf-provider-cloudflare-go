// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticatedoriginpulls


type AuthenticatedOriginPullsConfigA struct {
	// Certificate identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/authenticated_origin_pulls#cert_id AuthenticatedOriginPulls#cert_id}
	CertId *string `field:"optional" json:"certId" yaml:"certId"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null value voids the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/authenticated_origin_pulls#enabled AuthenticatedOriginPulls#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/authenticated_origin_pulls#hostname AuthenticatedOriginPulls#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}


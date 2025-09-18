// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplication


type DataCloudflareZeroTrustAccessApplicationFilter struct {
	// The aud of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_access_application#aud DataCloudflareZeroTrustAccessApplication#aud}
	Aud *string `field:"optional" json:"aud" yaml:"aud"`
	// The domain of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_access_application#domain DataCloudflareZeroTrustAccessApplication#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// True for only exact string matches against passed name/domain query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_access_application#exact DataCloudflareZeroTrustAccessApplication#exact}
	Exact interface{} `field:"optional" json:"exact" yaml:"exact"`
	// The name of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_access_application#name DataCloudflareZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Search for apps by other listed query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_access_application#search DataCloudflareZeroTrustAccessApplication#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}


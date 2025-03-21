// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomssl


type DataCloudflareCustomSslFilter struct {
	// Whether to match all search requirements or at least one (any). Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/custom_ssl#match DataCloudflareCustomSsl#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Status of the zone's custom SSL. Available values: "active", "expired", "deleted", "pending", "initializing".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/custom_ssl#status DataCloudflareCustomSsl#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


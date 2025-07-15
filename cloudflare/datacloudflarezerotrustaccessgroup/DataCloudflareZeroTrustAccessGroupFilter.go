// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessgroup


type DataCloudflareZeroTrustAccessGroupFilter struct {
	// The name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_access_group#name DataCloudflareZeroTrustAccessGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Search for groups by other listed query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_access_group#search DataCloudflareZeroTrustAccessGroup#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}


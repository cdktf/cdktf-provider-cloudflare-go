// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustlist


type DataCloudflareZeroTrustListFilter struct {
	// The type of list. Available values: "SERIAL", "URL", "DOMAIN", "EMAIL", "IP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/zero_trust_list#type DataCloudflareZeroTrustList#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


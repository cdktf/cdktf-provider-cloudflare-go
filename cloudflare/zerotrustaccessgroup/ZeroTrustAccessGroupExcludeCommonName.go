// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeCommonName struct {
	// The common name to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_group#common_name ZeroTrustAccessGroup#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
}


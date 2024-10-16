// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppRefreshTokenOptions struct {
	// How long a refresh token will be valid for after creation.
	//
	// Valid units are `m`, `h` and `d`. Must be longer than 1m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#lifetime ZeroTrustAccessApplication#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
}


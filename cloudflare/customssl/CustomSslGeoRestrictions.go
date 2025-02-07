// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customssl


type CustomSslGeoRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/custom_ssl#label CustomSsl#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareorigincacertificate


type DataCloudflareOriginCaCertificateFilter struct {
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/origin_ca_certificate#zone_id DataCloudflareOriginCaCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Limit to the number of records returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/origin_ca_certificate#limit DataCloudflareOriginCaCertificate#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/origin_ca_certificate#offset DataCloudflareOriginCaCertificate#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}


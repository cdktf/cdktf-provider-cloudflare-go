// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareorigincacertificate


type DataCloudflareOriginCaCertificateFilter struct {
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/origin_ca_certificate#zone_id DataCloudflareOriginCaCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}


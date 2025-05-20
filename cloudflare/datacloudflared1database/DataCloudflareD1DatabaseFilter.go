// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflared1database


type DataCloudflareD1DatabaseFilter struct {
	// a database name to search for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/d1_database#name DataCloudflareD1Database#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


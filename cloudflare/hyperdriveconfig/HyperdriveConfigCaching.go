// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigCaching struct {
	// Set to true to disable caching of SQL responses. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/hyperdrive_config#disabled HyperdriveConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if set to the default (60).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/hyperdrive_config#max_age HyperdriveConfig#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if set to the default (15).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/hyperdrive_config#stale_while_revalidate HyperdriveConfig#stale_while_revalidate}
	StaleWhileRevalidate *float64 `field:"optional" json:"staleWhileRevalidate" yaml:"staleWhileRevalidate"`
}


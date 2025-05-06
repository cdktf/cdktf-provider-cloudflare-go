// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#disabled HyperdriveConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// When present, specifies max duration for which items should persist in the cache.
	//
	// Not returned if set to default. (Default: 60)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#max_age HyperdriveConfig#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// When present, indicates the number of seconds cache may serve the response after it becomes stale.
	//
	// Not returned if set to default. (Default: 15)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#stale_while_revalidate HyperdriveConfig#stale_while_revalidate}
	StaleWhileRevalidate *float64 `field:"optional" json:"staleWhileRevalidate" yaml:"staleWhileRevalidate"`
}


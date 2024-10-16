// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigCaching struct {
	// Disable caching for this Hyperdrive configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/hyperdrive_config#disabled HyperdriveConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Configure the `max_age` value of this Hyperdrive configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/hyperdrive_config#max_age HyperdriveConfig#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Disable caching for this Hyperdrive configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/hyperdrive_config#stale_while_revalidate HyperdriveConfig#stale_while_revalidate}
	StaleWhileRevalidate *float64 `field:"optional" json:"staleWhileRevalidate" yaml:"staleWhileRevalidate"`
}


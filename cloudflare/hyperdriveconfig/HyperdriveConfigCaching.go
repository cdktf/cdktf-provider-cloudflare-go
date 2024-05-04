// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigCaching struct {
	// Disable caching for this Hyperdrive configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.31.0/docs/resources/hyperdrive_config#disabled HyperdriveConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}


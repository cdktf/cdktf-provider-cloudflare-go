// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsite


type MagicTransitSiteLocation struct {
	// Latitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/magic_transit_site#lat MagicTransitSite#lat}
	Lat *string `field:"optional" json:"lat" yaml:"lat"`
	// Longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/magic_transit_site#lon MagicTransitSite#lon}
	Lon *string `field:"optional" json:"lon" yaml:"lon"`
}


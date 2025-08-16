// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanstaticroute


type MagicWanStaticRouteScope struct {
	// List of colo names for the ECMP scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_wan_static_route#colo_names MagicWanStaticRoute#colo_names}
	ColoNames *[]*string `field:"optional" json:"coloNames" yaml:"coloNames"`
	// List of colo regions for the ECMP scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_wan_static_route#colo_regions MagicWanStaticRoute#colo_regions}
	ColoRegions *[]*string `field:"optional" json:"coloRegions" yaml:"coloRegions"`
}


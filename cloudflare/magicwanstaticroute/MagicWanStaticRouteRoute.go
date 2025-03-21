// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanstaticroute


type MagicWanStaticRouteRoute struct {
	// An optional human provided description of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#description MagicWanStaticRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The next-hop IP Address for the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#nexthop MagicWanStaticRoute#nexthop}
	Nexthop *string `field:"optional" json:"nexthop" yaml:"nexthop"`
	// IP Prefix in Classless Inter-Domain Routing format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#prefix MagicWanStaticRoute#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Priority of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#priority MagicWanStaticRoute#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Used only for ECMP routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#scope MagicWanStaticRoute#scope}
	Scope *MagicWanStaticRouteRouteScope `field:"optional" json:"scope" yaml:"scope"`
	// Optional weight of the ECMP scope - if provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/magic_wan_static_route#weight MagicWanStaticRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


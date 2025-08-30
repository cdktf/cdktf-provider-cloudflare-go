// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanstaticroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicWanStaticRouteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#account_id MagicWanStaticRoute#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The next-hop IP Address for the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#nexthop MagicWanStaticRoute#nexthop}
	Nexthop *string `field:"required" json:"nexthop" yaml:"nexthop"`
	// IP Prefix in Classless Inter-Domain Routing format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#prefix MagicWanStaticRoute#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Priority of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#priority MagicWanStaticRoute#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// An optional human provided description of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#description MagicWanStaticRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Used only for ECMP routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#scope MagicWanStaticRoute#scope}
	Scope *MagicWanStaticRouteScope `field:"optional" json:"scope" yaml:"scope"`
	// Optional weight of the ECMP scope - if provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_static_route#weight MagicWanStaticRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


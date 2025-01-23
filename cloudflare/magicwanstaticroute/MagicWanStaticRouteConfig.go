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
	// The nexthop IP address where traffic will be routed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#nexthop MagicWanStaticRoute#nexthop}
	Nexthop *string `field:"required" json:"nexthop" yaml:"nexthop"`
	// Your network prefix using CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#prefix MagicWanStaticRoute#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The priority for the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#priority MagicWanStaticRoute#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#account_id MagicWanStaticRoute#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// List of Cloudflare colocation regions for this static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#colo_names MagicWanStaticRoute#colo_names}
	ColoNames *[]*string `field:"optional" json:"coloNames" yaml:"coloNames"`
	// List of Cloudflare colocation names for this static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#colo_regions MagicWanStaticRoute#colo_regions}
	ColoRegions *[]*string `field:"optional" json:"coloRegions" yaml:"coloRegions"`
	// Description of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#description MagicWanStaticRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#id MagicWanStaticRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The optional weight for ECMP routes. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/magic_wan_static_route#weight MagicWanStaticRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


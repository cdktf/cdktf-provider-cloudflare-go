// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#account_id LoadBalancerPool#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#name LoadBalancerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of origins within this pool.
	//
	// Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#origins LoadBalancerPool#origins}
	Origins interface{} `field:"required" json:"origins" yaml:"origins"`
	// A list of regions from which to run health checks. Null means every Cloudflare data center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#check_regions LoadBalancerPool#check_regions}
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// A human-readable description of the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#description LoadBalancerPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable (the default) or disable this pool.
	//
	// Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#enabled LoadBalancerPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The latitude of the data center containing the origins used in this pool in decimal degrees.
	//
	// If this is set, longitude must also be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#latitude LoadBalancerPool#latitude}
	Latitude *float64 `field:"optional" json:"latitude" yaml:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#load_shedding LoadBalancerPool#load_shedding}
	LoadShedding *LoadBalancerPoolLoadShedding `field:"optional" json:"loadShedding" yaml:"loadShedding"`
	// The longitude of the data center containing the origins used in this pool in decimal degrees.
	//
	// If this is set, latitude must also be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#longitude LoadBalancerPool#longitude}
	Longitude *float64 `field:"optional" json:"longitude" yaml:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve traffic.
	//
	// If the number of healthy origins falls below this number, the pool will be marked unhealthy and will failover to the next available pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#minimum_origins LoadBalancerPool#minimum_origins}
	MinimumOrigins *float64 `field:"optional" json:"minimumOrigins" yaml:"minimumOrigins"`
	// The ID of the Monitor to use for checking the health of origins within this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#monitor LoadBalancerPool#monitor}
	Monitor *string `field:"optional" json:"monitor" yaml:"monitor"`
	// This field is now deprecated.
	//
	// It has been moved to Cloudflare's Centralized Notification service https://developers.cloudflare.com/fundamentals/notifications/. The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#notification_email LoadBalancerPool#notification_email}
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// Filter pool and origin health notifications by resource type or health status. Use null to reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#notification_filter LoadBalancerPool#notification_filter}
	NotificationFilter *LoadBalancerPoolNotificationFilter `field:"optional" json:"notificationFilter" yaml:"notificationFilter"`
	// Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer_pool#origin_steering LoadBalancerPool#origin_steering}
	OriginSteering *LoadBalancerPoolOriginSteering `field:"optional" json:"originSteering" yaml:"originSteering"`
}


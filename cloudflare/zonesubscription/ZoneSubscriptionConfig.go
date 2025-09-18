// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonesubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneSubscriptionConfig struct {
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
	// Subscription identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#zone_id ZoneSubscription#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// How often the subscription is renewed automatically. Available values: "weekly", "monthly", "quarterly", "yearly".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#frequency ZoneSubscription#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The rate plan applied to the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#rate_plan ZoneSubscription#rate_plan}
	RatePlan *ZoneSubscriptionRatePlan `field:"optional" json:"ratePlan" yaml:"ratePlan"`
}


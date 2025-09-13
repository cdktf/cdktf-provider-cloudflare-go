// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonelockdown

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneLockdownConfig struct {
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
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule.
	//
	// You can include any number of `ip` or `ip_range` configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#configurations ZoneLockdown#configurations}
	Configurations interface{} `field:"required" json:"configurations" yaml:"configurations"`
	// The URLs to include in the current WAF override.
	//
	// You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#urls ZoneLockdown#urls}
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#zone_id ZoneLockdown#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// An informative summary of the rule. This value is sanitized and any tags will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#description ZoneLockdown#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When true, indicates that the rule is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#paused ZoneLockdown#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The priority of the rule to control the processing order.
	//
	// A lower number indicates higher priority. If not provided, any rules with a configured priority will be processed before rules without a priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zone_lockdown#priority ZoneLockdown#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}


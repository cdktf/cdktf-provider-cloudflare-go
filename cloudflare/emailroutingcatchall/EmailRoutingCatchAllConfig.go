// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingcatchall

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailRoutingCatchAllConfig struct {
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
	// List actions for the catch-all routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/email_routing_catch_all#actions EmailRoutingCatchAll#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// List of matchers for the catch-all routing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/email_routing_catch_all#matchers EmailRoutingCatchAll#matchers}
	Matchers interface{} `field:"required" json:"matchers" yaml:"matchers"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/email_routing_catch_all#zone_id EmailRoutingCatchAll#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Routing rule status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/email_routing_catch_all#enabled EmailRoutingCatchAll#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Routing rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/email_routing_catch_all#name EmailRoutingCatchAll#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observatoryscheduledtest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservatoryScheduledTestConfig struct {
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
	// A URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/observatory_scheduled_test#url ObservatoryScheduledTest#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/observatory_scheduled_test#zone_id ObservatoryScheduledTest#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}


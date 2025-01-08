// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webanalyticssite

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebAnalyticsSiteConfig struct {
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
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#account_id WebAnalyticsSite#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether Cloudflare will automatically inject the JavaScript snippet for orange-clouded sites.
	//
	// **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#auto_install WebAnalyticsSite#auto_install}
	AutoInstall interface{} `field:"required" json:"autoInstall" yaml:"autoInstall"`
	// The hostname to use for gray-clouded sites.
	//
	// Must provide only one of `zone_tag`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#host WebAnalyticsSite#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#id WebAnalyticsSite#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#timeouts WebAnalyticsSite#timeouts}
	Timeouts *WebAnalyticsSiteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone identifier for orange-clouded sites.
	//
	// Must provide only one of `host`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/web_analytics_site#zone_tag WebAnalyticsSite#zone_tag}
	ZoneTag *string `field:"optional" json:"zoneTag" yaml:"zoneTag"`
}


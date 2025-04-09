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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#account_id WebAnalyticsSite#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#auto_install WebAnalyticsSite#auto_install}
	AutoInstall interface{} `field:"optional" json:"autoInstall" yaml:"autoInstall"`
	// Enables or disables RUM. This option can be used only when auto_install is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#enabled WebAnalyticsSite#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The hostname to use for gray-clouded sites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#host WebAnalyticsSite#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// If enabled, the JavaScript snippet will not be injected for visitors from the EU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#lite WebAnalyticsSite#lite}
	Lite interface{} `field:"optional" json:"lite" yaml:"lite"`
	// The zone identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/web_analytics_site#zone_tag WebAnalyticsSite#zone_tag}
	ZoneTag *string `field:"optional" json:"zoneTag" yaml:"zoneTag"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botmanagement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotManagementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#zone_id BotManagement#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Enable rule to block AI Scrapers and Crawlers.
	//
	// Please note the value `only_on_ad_pages` is currently not available for Enterprise customers.
	// Available values: "block", "disabled", "only_on_ad_pages".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#ai_bots_protection BotManagement#ai_bots_protection}
	AiBotsProtection *string `field:"optional" json:"aiBotsProtection" yaml:"aiBotsProtection"`
	// Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#auto_update_model BotManagement#auto_update_model}
	AutoUpdateModel interface{} `field:"optional" json:"autoUpdateModel" yaml:"autoUpdateModel"`
	// Indicates that the bot management cookie can be placed on end user devices accessing the site. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#bm_cookie_enabled BotManagement#bm_cookie_enabled}
	BmCookieEnabled interface{} `field:"optional" json:"bmCookieEnabled" yaml:"bmCookieEnabled"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze. Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#crawler_protection BotManagement#crawler_protection}
	CrawlerProtection *string `field:"optional" json:"crawlerProtection" yaml:"crawlerProtection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#enable_js BotManagement#enable_js}
	EnableJs interface{} `field:"optional" json:"enableJs" yaml:"enableJs"`
	// Whether to enable Bot Fight Mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#fight_mode BotManagement#fight_mode}
	FightMode interface{} `field:"optional" json:"fightMode" yaml:"fightMode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected, then managed robots.txt will be prepended to the existing robots.txt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#is_robots_txt_managed BotManagement#is_robots_txt_managed}
	IsRobotsTxtManaged interface{} `field:"optional" json:"isRobotsTxtManaged" yaml:"isRobotsTxtManaged"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#optimize_wordpress BotManagement#optimize_wordpress}
	OptimizeWordpress interface{} `field:"optional" json:"optimizeWordpress" yaml:"optimizeWordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests. Available values: "allow", "block", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#sbfm_definitely_automated BotManagement#sbfm_definitely_automated}
	SbfmDefinitelyAutomated *string `field:"optional" json:"sbfmDefinitelyAutomated" yaml:"sbfmDefinitelyAutomated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests. Available values: "allow", "block", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#sbfm_likely_automated BotManagement#sbfm_likely_automated}
	SbfmLikelyAutomated *string `field:"optional" json:"sbfmLikelyAutomated" yaml:"sbfmLikelyAutomated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection.
	//
	// Enable if static resources on your application need bot protection.
	// Note: Static resource protection can also result in legitimate traffic being blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#sbfm_static_resource_protection BotManagement#sbfm_static_resource_protection}
	SbfmStaticResourceProtection interface{} `field:"optional" json:"sbfmStaticResourceProtection" yaml:"sbfmStaticResourceProtection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests. Available values: "allow", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#sbfm_verified_bots BotManagement#sbfm_verified_bots}
	SbfmVerifiedBots *string `field:"optional" json:"sbfmVerifiedBots" yaml:"sbfmVerifiedBots"`
	// Whether to disable tracking the highest bot score for a session in the Bot Management cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/bot_management#suppress_session_score BotManagement#suppress_session_score}
	SuppressSessionScore interface{} `field:"optional" json:"suppressSessionScore" yaml:"suppressSessionScore"`
}


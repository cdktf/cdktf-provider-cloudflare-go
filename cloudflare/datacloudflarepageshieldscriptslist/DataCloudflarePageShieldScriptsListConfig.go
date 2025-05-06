// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldscriptslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflarePageShieldScriptsListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#zone_id DataCloudflarePageShieldScriptsList#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The direction used to sort returned scripts. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#direction DataCloudflarePageShieldScriptsList#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#exclude_cdn_cgi DataCloudflarePageShieldScriptsList#exclude_cdn_cgi}
	ExcludeCdnCgi interface{} `field:"optional" json:"excludeCdnCgi" yaml:"excludeCdnCgi"`
	// When true, excludes duplicate scripts.
	//
	// We consider a script duplicate of another if their javascript
	// content matches and they share the same url host and zone hostname. In such case, we return the most
	// recent script for the URL host and zone hostname combination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#exclude_duplicates DataCloudflarePageShieldScriptsList#exclude_duplicates}
	ExcludeDuplicates interface{} `field:"optional" json:"excludeDuplicates" yaml:"excludeDuplicates"`
	// Excludes scripts whose URL contains one of the URL-encoded URLs separated by commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#exclude_urls DataCloudflarePageShieldScriptsList#exclude_urls}
	ExcludeUrls *string `field:"optional" json:"excludeUrls" yaml:"excludeUrls"`
	// Export the list of scripts as a file. Available values: "csv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#export DataCloudflarePageShieldScriptsList#export}
	Export *string `field:"optional" json:"export" yaml:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#hosts DataCloudflarePageShieldScriptsList#hosts}
	Hosts *string `field:"optional" json:"hosts" yaml:"hosts"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#max_items DataCloudflarePageShieldScriptsList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The field used to sort returned scripts. Available values: "first_seen_at", "last_seen_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#order_by DataCloudflarePageShieldScriptsList#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will return all the scripts
	// with the applied filters in a single page. This feature is best-effort and it may only work for zones with
	// a low number of scripts
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#page DataCloudflarePageShieldScriptsList#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// Includes scripts that match one or more page URLs (separated by commas) where they were last seen.
	//
	// Wildcards are supported at the start and end of each page URL to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#page_url DataCloudflarePageShieldScriptsList#page_url}
	PageUrl *string `field:"optional" json:"pageUrl" yaml:"pageUrl"`
	// The number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#per_page DataCloudflarePageShieldScriptsList#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
	// When true, malicious scripts appear first in the returned scripts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#prioritize_malicious DataCloudflarePageShieldScriptsList#prioritize_malicious}
	PrioritizeMalicious interface{} `field:"optional" json:"prioritizeMalicious" yaml:"prioritizeMalicious"`
	// Filters the returned scripts using a comma-separated list of scripts statuses.
	//
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is `active`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#status DataCloudflarePageShieldScriptsList#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Includes scripts whose URL contain one or more URL-encoded URLs separated by commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/data-sources/page_shield_scripts_list#urls DataCloudflarePageShieldScriptsList#urls}
	Urls *string `field:"optional" json:"urls" yaml:"urls"`
}


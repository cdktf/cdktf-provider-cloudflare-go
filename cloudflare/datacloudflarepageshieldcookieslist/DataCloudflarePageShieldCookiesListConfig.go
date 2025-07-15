// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldcookieslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflarePageShieldCookiesListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#zone_id DataCloudflarePageShieldCookiesList#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The direction used to sort returned cookies.' Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#direction DataCloudflarePageShieldCookiesList#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filters the returned cookies that match the specified domain attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#domain DataCloudflarePageShieldCookiesList#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Export the list of cookies as a file. Available values: "csv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#export DataCloudflarePageShieldCookiesList#export}
	Export *string `field:"optional" json:"export" yaml:"export"`
	// Includes cookies that match one or more URL-encoded hostnames separated by commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#hosts DataCloudflarePageShieldCookiesList#hosts}
	Hosts *string `field:"optional" json:"hosts" yaml:"hosts"`
	// Filters the returned cookies that are set with HttpOnly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#http_only DataCloudflarePageShieldCookiesList#http_only}
	HttpOnly interface{} `field:"optional" json:"httpOnly" yaml:"httpOnly"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#max_items DataCloudflarePageShieldCookiesList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Filters the returned cookies that match the specified name.
	//
	// Wildcards are supported at the start and end to support starts with, ends with
	// and contains. e.g. session*
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#name DataCloudflarePageShieldCookiesList#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The field used to sort returned cookies. Available values: "first_seen_at", "last_seen_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#order_by DataCloudflarePageShieldCookiesList#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will return all the cookies
	// with the applied filters in a single page. This feature is best-effort and it may only work for zones with
	// a low number of cookies
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#page DataCloudflarePageShieldCookiesList#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// Includes connections that match one or more page URLs (separated by commas) where they were last seen.
	//
	// Wildcards are supported at the start and end of each page URL to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#page_url DataCloudflarePageShieldCookiesList#page_url}
	PageUrl *string `field:"optional" json:"pageUrl" yaml:"pageUrl"`
	// Filters the returned cookies that match the specified path attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#path DataCloudflarePageShieldCookiesList#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#per_page DataCloudflarePageShieldCookiesList#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
	// Filters the returned cookies that match the specified same_site attribute Available values: "lax", "strict", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#same_site DataCloudflarePageShieldCookiesList#same_site}
	SameSite *string `field:"optional" json:"sameSite" yaml:"sameSite"`
	// Filters the returned cookies that are set with Secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#secure DataCloudflarePageShieldCookiesList#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
	// Filters the returned cookies that match the specified type attribute Available values: "first_party", "unknown".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/page_shield_cookies_list#type DataCloudflarePageShieldCookiesList#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


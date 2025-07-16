// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldconnectionslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflarePageShieldConnectionsListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#zone_id DataCloudflarePageShieldConnectionsList#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The direction used to sort returned connections. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#direction DataCloudflarePageShieldConnectionsList#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned connections. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#exclude_cdn_cgi DataCloudflarePageShieldConnectionsList#exclude_cdn_cgi}
	ExcludeCdnCgi interface{} `field:"optional" json:"excludeCdnCgi" yaml:"excludeCdnCgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#exclude_urls DataCloudflarePageShieldConnectionsList#exclude_urls}
	ExcludeUrls *string `field:"optional" json:"excludeUrls" yaml:"excludeUrls"`
	// Export the list of connections as a file. Available values: "csv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#export DataCloudflarePageShieldConnectionsList#export}
	Export *string `field:"optional" json:"export" yaml:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#hosts DataCloudflarePageShieldConnectionsList#hosts}
	Hosts *string `field:"optional" json:"hosts" yaml:"hosts"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#max_items DataCloudflarePageShieldConnectionsList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The field used to sort returned connections. Available values: "first_seen_at", "last_seen_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#order_by DataCloudflarePageShieldConnectionsList#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will return all the connections
	// with the applied filters in a single page. This feature is best-effort and it may only work for zones with
	// a low number of connections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#page DataCloudflarePageShieldConnectionsList#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// Includes connections that match one or more page URLs (separated by commas) where they were last seen.
	//
	// Wildcards are supported at the start and end of each page URL to support starts with, ends with
	// and contains. If no wildcards are used, results will be filtered by exact match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#page_url DataCloudflarePageShieldConnectionsList#page_url}
	PageUrl *string `field:"optional" json:"pageUrl" yaml:"pageUrl"`
	// The number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#per_page DataCloudflarePageShieldConnectionsList#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
	// When true, malicious connections appear first in the returned connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#prioritize_malicious DataCloudflarePageShieldConnectionsList#prioritize_malicious}
	PrioritizeMalicious interface{} `field:"optional" json:"prioritizeMalicious" yaml:"prioritizeMalicious"`
	// Filters the returned connections using a comma-separated list of connection statuses.
	//
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is `active`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#status DataCloudflarePageShieldConnectionsList#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Includes connections whose URL contain one or more URL-encoded URLs separated by commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/page_shield_connections_list#urls DataCloudflarePageShieldConnectionsList#urls}
	Urls *string `field:"optional" json:"urls" yaml:"urls"`
}


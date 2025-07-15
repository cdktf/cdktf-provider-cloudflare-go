// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequests

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#account_id DataCloudflareCloudforceOneRequests#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Page number of results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#page DataCloudflareCloudforceOneRequests#page}
	Page *float64 `field:"required" json:"page" yaml:"page"`
	// Number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#per_page DataCloudflareCloudforceOneRequests#per_page}
	PerPage *float64 `field:"required" json:"perPage" yaml:"perPage"`
	// Retrieve requests completed after this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#completed_after DataCloudflareCloudforceOneRequests#completed_after}
	CompletedAfter *string `field:"optional" json:"completedAfter" yaml:"completedAfter"`
	// Retrieve requests completed before this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#completed_before DataCloudflareCloudforceOneRequests#completed_before}
	CompletedBefore *string `field:"optional" json:"completedBefore" yaml:"completedBefore"`
	// Retrieve requests created after this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#created_after DataCloudflareCloudforceOneRequests#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Retrieve requests created before this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#created_before DataCloudflareCloudforceOneRequests#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#max_items DataCloudflareCloudforceOneRequests#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Requested information from request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#request_type DataCloudflareCloudforceOneRequests#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// Field to sort results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#sort_by DataCloudflareCloudforceOneRequests#sort_by}
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Sort order (asc or desc). Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#sort_order DataCloudflareCloudforceOneRequests#sort_order}
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// Request Status. Available values: "open", "accepted", "reported", "approved", "completed", "declined".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/cloudforce_one_requests#status DataCloudflareCloudforceOneRequests#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


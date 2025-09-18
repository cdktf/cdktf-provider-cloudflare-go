// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequest


type DataCloudflareCloudforceOneRequestFilter struct {
	// Page number of results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#page DataCloudflareCloudforceOneRequest#page}
	Page *float64 `field:"required" json:"page" yaml:"page"`
	// Number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#per_page DataCloudflareCloudforceOneRequest#per_page}
	PerPage *float64 `field:"required" json:"perPage" yaml:"perPage"`
	// Retrieve requests completed after this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#completed_after DataCloudflareCloudforceOneRequest#completed_after}
	CompletedAfter *string `field:"optional" json:"completedAfter" yaml:"completedAfter"`
	// Retrieve requests completed before this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#completed_before DataCloudflareCloudforceOneRequest#completed_before}
	CompletedBefore *string `field:"optional" json:"completedBefore" yaml:"completedBefore"`
	// Retrieve requests created after this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#created_after DataCloudflareCloudforceOneRequest#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Retrieve requests created before this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#created_before DataCloudflareCloudforceOneRequest#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Requested information from request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#request_type DataCloudflareCloudforceOneRequest#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// Field to sort results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#sort_by DataCloudflareCloudforceOneRequest#sort_by}
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Sort order (asc or desc). Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#sort_order DataCloudflareCloudforceOneRequest#sort_order}
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// Request Status. Available values: "open", "accepted", "reported", "approved", "completed", "declined".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/cloudforce_one_request#status DataCloudflareCloudforceOneRequest#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


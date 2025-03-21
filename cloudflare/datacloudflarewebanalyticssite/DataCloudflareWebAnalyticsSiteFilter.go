// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarewebanalyticssite


type DataCloudflareWebAnalyticsSiteFilter struct {
	// The property used to sort the list of results. Available values: "host", "created".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/web_analytics_site#order_by DataCloudflareWebAnalyticsSite#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}


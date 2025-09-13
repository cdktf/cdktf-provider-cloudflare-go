// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list


type ListItemsRedirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#source_url List#source_url}.
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#target_url List#target_url}.
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#include_subdomains List#include_subdomains}.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#preserve_path_suffix List#preserve_path_suffix}.
	PreservePathSuffix interface{} `field:"optional" json:"preservePathSuffix" yaml:"preservePathSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#preserve_query_string List#preserve_query_string}.
	PreserveQueryString interface{} `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// Available values: 301, 302, 307, 308.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#status_code List#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#subpath_matching List#subpath_matching}.
	SubpathMatching interface{} `field:"optional" json:"subpathMatching" yaml:"subpathMatching"`
}


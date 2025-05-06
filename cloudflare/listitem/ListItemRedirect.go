// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listitem


type ListItemRedirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#source_url ListItem#source_url}.
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#target_url ListItem#target_url}.
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#include_subdomains ListItem#include_subdomains}.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#preserve_path_suffix ListItem#preserve_path_suffix}.
	PreservePathSuffix interface{} `field:"optional" json:"preservePathSuffix" yaml:"preservePathSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#preserve_query_string ListItem#preserve_query_string}.
	PreserveQueryString interface{} `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// Available values: 301, 302, 307, 308.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#status_code ListItem#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/list_item#subpath_matching ListItem#subpath_matching}.
	SubpathMatching interface{} `field:"optional" json:"subpathMatching" yaml:"subpathMatching"`
}


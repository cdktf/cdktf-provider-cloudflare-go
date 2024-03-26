// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listitem


type ListItemRedirect struct {
	// The source url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#source_url ListItemA#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// The target url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#target_url ListItemA#target_url}
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Whether the redirect also matches subdomains of the source url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#include_subdomains ListItemA#include_subdomains}
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Whether the redirect target url should keep the query string of the request's url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#preserve_path_suffix ListItemA#preserve_path_suffix}
	PreservePathSuffix interface{} `field:"optional" json:"preservePathSuffix" yaml:"preservePathSuffix"`
	// Whether the redirect target url should keep the query string of the request's url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#preserve_query_string ListItemA#preserve_query_string}
	PreserveQueryString interface{} `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// The status code to be used when redirecting a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#status_code ListItemA#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Whether the redirect also matches subpaths of the source url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.27.0/docs/resources/list_item#subpath_matching ListItemA#subpath_matching}
	SubpathMatching interface{} `field:"optional" json:"subpathMatching" yaml:"subpathMatching"`
}


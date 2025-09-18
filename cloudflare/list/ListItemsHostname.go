// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list


type ListItemsHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list#url_hostname List#url_hostname}.
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
	// Only applies to wildcard hostnames (e.g., *.example.com). When true (default), only subdomains are blocked. When false, both the root domain and subdomains are blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list#exclude_exact_hostname List#exclude_exact_hostname}
	ExcludeExactHostname interface{} `field:"optional" json:"excludeExactHostname" yaml:"excludeExactHostname"`
}


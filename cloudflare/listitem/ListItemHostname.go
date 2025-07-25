// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listitem


type ListItemHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/list_item#url_hostname ListItem#url_hostname}.
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
}


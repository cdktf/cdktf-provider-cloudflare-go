// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listitem


type ListItemHostname struct {
	// The FQDN to match on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/list_item#url_hostname ListItemA#url_hostname}
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
}


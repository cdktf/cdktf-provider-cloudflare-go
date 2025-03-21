// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccessrules


type DataCloudflareAccessRulesConfiguration struct {
	// The target to search in existing rules. Available values: "ip", "ip_range", "asn", "country".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/access_rules#target DataCloudflareAccessRules#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`. Notes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/access_rules#value DataCloudflareAccessRules#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


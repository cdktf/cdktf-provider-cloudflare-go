// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezone


type DataCloudflareZoneFilterAccount struct {
	// An account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zone#id DataCloudflareZone#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An account Name.
	//
	// Optional filter operators can be provided to extend refine the search:
	//   * `equal` (default)
	//   * `not_equal`
	//   * `starts_with`
	//   * `ends_with`
	//   * `contains`
	//   * `starts_with_case_sensitive`
	//   * `ends_with_case_sensitive`
	//   * `contains_case_sensitive`
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zone#name DataCloudflareZone#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


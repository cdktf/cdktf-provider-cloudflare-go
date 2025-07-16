// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezone


type DataCloudflareZoneFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#account DataCloudflareZone#account}.
	Account *DataCloudflareZoneFilterAccount `field:"optional" json:"account" yaml:"account"`
	// Direction to order zones. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#direction DataCloudflareZone#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Whether to match all search requirements or at least one (any). Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#match DataCloudflareZone#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// A domain name.
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#name DataCloudflareZone#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Field to order zones by. Available values: "name", "status", "account.id", "account.name", "plan.id".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#order DataCloudflareZone#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Specify a zone status to filter by. Available values: "initializing", "pending", "active", "moved".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zone#status DataCloudflareZone#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


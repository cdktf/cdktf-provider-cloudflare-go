// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareturnstilewidget


type DataCloudflareTurnstileWidgetFilter struct {
	// Direction to order widgets. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/turnstile_widget#direction DataCloudflareTurnstileWidget#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order widgets by. Available values: "id", "sitekey", "name", "created_on", "modified_on".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/data-sources/turnstile_widget#order DataCloudflareTurnstileWidget#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}


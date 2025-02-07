// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareturnstilewidget


type DataCloudflareTurnstileWidgetFilter struct {
	// Direction to order widgets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/turnstile_widget#direction DataCloudflareTurnstileWidget#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order widgets by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/turnstile_widget#order DataCloudflareTurnstileWidget#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}


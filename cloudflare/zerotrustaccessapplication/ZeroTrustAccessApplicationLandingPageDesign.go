// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationLandingPageDesign struct {
	// The button color of the landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#button_color ZeroTrustAccessApplication#button_color}
	ButtonColor *string `field:"optional" json:"buttonColor" yaml:"buttonColor"`
	// The button text color of the landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#button_text_color ZeroTrustAccessApplication#button_text_color}
	ButtonTextColor *string `field:"optional" json:"buttonTextColor" yaml:"buttonTextColor"`
	// The URL of the image to be displayed in the landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#image_url ZeroTrustAccessApplication#image_url}
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The message of the landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#message ZeroTrustAccessApplication#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The title of the landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_access_application#title ZeroTrustAccessApplication#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}


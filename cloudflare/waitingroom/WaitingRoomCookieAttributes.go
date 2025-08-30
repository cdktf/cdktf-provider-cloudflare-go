// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroom


type WaitingRoomCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie.
	//
	// Value `auto` will be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled. Note that when using value `none`, the secure attribute cannot be set to `never`.
	// Available values: "auto", "lax", "none", "strict".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/waiting_room#samesite WaitingRoom#samesite}
	Samesite *string `field:"optional" json:"samesite" yaml:"samesite"`
	// Configures the Secure attribute on the waiting room cookie.
	//
	// Value `always` indicates that the Secure attribute will be set in the Set-Cookie header, `never` indicates that the Secure attribute will not be set, and `auto` will set the Secure attribute depending if **Always Use HTTPS** is enabled.
	// Available values: "auto", "always", "never".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/waiting_room#secure WaitingRoom#secure}
	Secure *string `field:"optional" json:"secure" yaml:"secure"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessOrganizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The unique subdomain assigned to your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#auth_domain AccessOrganization#auth_domain}
	AuthDomain *string `field:"required" json:"authDomain" yaml:"authDomain"`
	// The name of your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#name AccessOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#account_id AccessOrganization#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// When set to true, users can authenticate via WARP for any application in your organization.
	//
	// Application settings will take precedence over this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#allow_authenticate_via_warp AccessOrganization#allow_authenticate_via_warp}
	AllowAuthenticateViaWarp interface{} `field:"optional" json:"allowAuthenticateViaWarp" yaml:"allowAuthenticateViaWarp"`
	// When set to true, users skip the identity provider selection step during login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#auto_redirect_to_identity AccessOrganization#auto_redirect_to_identity}
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// custom_pages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#custom_pages AccessOrganization#custom_pages}
	CustomPages interface{} `field:"optional" json:"customPages" yaml:"customPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#id AccessOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#is_ui_read_only AccessOrganization#is_ui_read_only}
	IsUiReadOnly interface{} `field:"optional" json:"isUiReadOnly" yaml:"isUiReadOnly"`
	// login_design block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#login_design AccessOrganization#login_design}
	LoginDesign interface{} `field:"optional" json:"loginDesign" yaml:"loginDesign"`
	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#session_duration AccessOrganization#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// A description of the reason why the UI read only field is being toggled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#ui_read_only_toggle_reason AccessOrganization#ui_read_only_toggle_reason}
	UiReadOnlyToggleReason *string `field:"optional" json:"uiReadOnlyToggleReason" yaml:"uiReadOnlyToggleReason"`
	// The amount of time a user seat is inactive before it expires.
	//
	// When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#user_seat_expiration_inactive_time AccessOrganization#user_seat_expiration_inactive_time}
	UserSeatExpirationInactiveTime *string `field:"optional" json:"userSeatExpirationInactiveTime" yaml:"userSeatExpirationInactiveTime"`
	// The amount of time that tokens issued for applications will be valid.
	//
	// Must be in the format 30m or 2h45m. Valid time units are: m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#warp_auth_session_duration AccessOrganization#warp_auth_session_duration}
	WarpAuthSessionDuration *string `field:"optional" json:"warpAuthSessionDuration" yaml:"warpAuthSessionDuration"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/access_organization#zone_id AccessOrganization#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


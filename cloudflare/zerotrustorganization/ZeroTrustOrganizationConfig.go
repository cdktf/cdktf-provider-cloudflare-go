// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustOrganizationConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#account_id ZeroTrustOrganization#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// When set to true, users can authenticate via WARP for any application in your organization.
	//
	// Application settings will take precedence over this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#allow_authenticate_via_warp ZeroTrustOrganization#allow_authenticate_via_warp}
	AllowAuthenticateViaWarp interface{} `field:"optional" json:"allowAuthenticateViaWarp" yaml:"allowAuthenticateViaWarp"`
	// The unique subdomain assigned to your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#auth_domain ZeroTrustOrganization#auth_domain}
	AuthDomain *string `field:"optional" json:"authDomain" yaml:"authDomain"`
	// When set to `true`, users skip the identity provider selection step during login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#auto_redirect_to_identity ZeroTrustOrganization#auto_redirect_to_identity}
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#custom_pages ZeroTrustOrganization#custom_pages}.
	CustomPages *ZeroTrustOrganizationCustomPages `field:"optional" json:"customPages" yaml:"customPages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	//
	// Updates may only be made via the API or Terraform for this account when enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#is_ui_read_only ZeroTrustOrganization#is_ui_read_only}
	IsUiReadOnly interface{} `field:"optional" json:"isUiReadOnly" yaml:"isUiReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#login_design ZeroTrustOrganization#login_design}.
	LoginDesign *ZeroTrustOrganizationLoginDesign `field:"optional" json:"loginDesign" yaml:"loginDesign"`
	// The name of your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#name ZeroTrustOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The amount of time that tokens issued for applications will be valid.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#session_duration ZeroTrustOrganization#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// A description of the reason why the UI read only field is being toggled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#ui_read_only_toggle_reason ZeroTrustOrganization#ui_read_only_toggle_reason}
	UiReadOnlyToggleReason *string `field:"optional" json:"uiReadOnlyToggleReason" yaml:"uiReadOnlyToggleReason"`
	// The amount of time a user seat is inactive before it expires.
	//
	// When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#user_seat_expiration_inactive_time ZeroTrustOrganization#user_seat_expiration_inactive_time}
	UserSeatExpirationInactiveTime *string `field:"optional" json:"userSeatExpirationInactiveTime" yaml:"userSeatExpirationInactiveTime"`
	// The amount of time that tokens issued for applications will be valid.
	//
	// Must be in the format `30m` or `2h45m`. Valid time units are: m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#warp_auth_session_duration ZeroTrustOrganization#warp_auth_session_duration}
	WarpAuthSessionDuration *string `field:"optional" json:"warpAuthSessionDuration" yaml:"warpAuthSessionDuration"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_organization#zone_id ZeroTrustOrganization#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


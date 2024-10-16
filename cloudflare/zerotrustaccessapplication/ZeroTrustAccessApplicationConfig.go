// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessApplicationConfig struct {
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
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#account_id ZeroTrustAccessApplication#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// When set to true, users can authenticate to this application using their WARP session.
	//
	// When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#allow_authenticate_via_warp ZeroTrustAccessApplication#allow_authenticate_via_warp}
	AllowAuthenticateViaWarp interface{} `field:"optional" json:"allowAuthenticateViaWarp" yaml:"allowAuthenticateViaWarp"`
	// The identity providers selected for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#allowed_idps ZeroTrustAccessApplication#allowed_idps}
	AllowedIdps *[]*string `field:"optional" json:"allowedIdps" yaml:"allowedIdps"`
	// The logo URL of the app launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#app_launcher_logo_url ZeroTrustAccessApplication#app_launcher_logo_url}
	AppLauncherLogoUrl *string `field:"optional" json:"appLauncherLogoUrl" yaml:"appLauncherLogoUrl"`
	// Option to show/hide applications in App Launcher. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#app_launcher_visible ZeroTrustAccessApplication#app_launcher_visible}
	AppLauncherVisible interface{} `field:"optional" json:"appLauncherVisible" yaml:"appLauncherVisible"`
	// Option to skip identity provider selection if only one is configured in `allowed_idps`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#auto_redirect_to_identity ZeroTrustAccessApplication#auto_redirect_to_identity}
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// The background color of the app launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#bg_color ZeroTrustAccessApplication#bg_color}
	BgColor *string `field:"optional" json:"bgColor" yaml:"bgColor"`
	// cors_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#cors_headers ZeroTrustAccessApplication#cors_headers}
	CorsHeaders interface{} `field:"optional" json:"corsHeaders" yaml:"corsHeaders"`
	// Option that returns a custom error message when a user is denied access to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#custom_deny_message ZeroTrustAccessApplication#custom_deny_message}
	CustomDenyMessage *string `field:"optional" json:"customDenyMessage" yaml:"customDenyMessage"`
	// Option that redirects to a custom URL when a user is denied access to the application via identity based rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#custom_deny_url ZeroTrustAccessApplication#custom_deny_url}
	CustomDenyUrl *string `field:"optional" json:"customDenyUrl" yaml:"customDenyUrl"`
	// Option that redirects to a custom URL when a user is denied access to the application via non identity rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#custom_non_identity_deny_url ZeroTrustAccessApplication#custom_non_identity_deny_url}
	CustomNonIdentityDenyUrl *string `field:"optional" json:"customNonIdentityDenyUrl" yaml:"customNonIdentityDenyUrl"`
	// The custom pages selected for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#custom_pages ZeroTrustAccessApplication#custom_pages}
	CustomPages *[]*string `field:"optional" json:"customPages" yaml:"customPages"`
	// The primary hostname and path that Access will secure.
	//
	// If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#domain ZeroTrustAccessApplication#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests.
	//
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#enable_binding_cookie ZeroTrustAccessApplication#enable_binding_cookie}
	EnableBindingCookie interface{} `field:"optional" json:"enableBindingCookie" yaml:"enableBindingCookie"`
	// footer_links block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#footer_links ZeroTrustAccessApplication#footer_links}
	FooterLinks interface{} `field:"optional" json:"footerLinks" yaml:"footerLinks"`
	// The background color of the header bar in the app launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#header_bg_color ZeroTrustAccessApplication#header_bg_color}
	HeaderBgColor *string `field:"optional" json:"headerBgColor" yaml:"headerBgColor"`
	// Option to add the `HttpOnly` cookie flag to access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#http_only_cookie_attribute ZeroTrustAccessApplication#http_only_cookie_attribute}
	HttpOnlyCookieAttribute interface{} `field:"optional" json:"httpOnlyCookieAttribute" yaml:"httpOnlyCookieAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#id ZeroTrustAccessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// landing_page_design block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#landing_page_design ZeroTrustAccessApplication#landing_page_design}
	LandingPageDesign *ZeroTrustAccessApplicationLandingPageDesign `field:"optional" json:"landingPageDesign" yaml:"landingPageDesign"`
	// Image URL for the logo shown in the app launcher dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#logo_url ZeroTrustAccessApplication#logo_url}
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
	// Friendly name of the Access Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Allows options preflight requests to bypass Access authentication and go directly to the origin.
	//
	// Cannot turn on if cors_headers is set. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#options_preflight_bypass ZeroTrustAccessApplication#options_preflight_bypass}
	OptionsPreflightBypass interface{} `field:"optional" json:"optionsPreflightBypass" yaml:"optionsPreflightBypass"`
	// The policies associated with the application, in ascending order of precedence.
	//
	// Warning: Do not use this field while you still have this application ID referenced as `application_id` in any `cloudflare_access_policy` resource, as it can result in an inconsistent state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#policies ZeroTrustAccessApplication#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// saas_app block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#saas_app ZeroTrustAccessApplication#saas_app}
	SaasApp *ZeroTrustAccessApplicationSaasApp `field:"optional" json:"saasApp" yaml:"saasApp"`
	// Defines the same-site cookie setting for access tokens. Available values: `none`, `lax`, `strict`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#same_site_cookie_attribute ZeroTrustAccessApplication#same_site_cookie_attribute}
	SameSiteCookieAttribute *string `field:"optional" json:"sameSiteCookieAttribute" yaml:"sameSiteCookieAttribute"`
	// scim_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#scim_config ZeroTrustAccessApplication#scim_config}
	ScimConfig *ZeroTrustAccessApplicationScimConfig `field:"optional" json:"scimConfig" yaml:"scimConfig"`
	// List of domains that access will secure.
	//
	// Only present for self_hosted, vnc, and ssh applications. Always includes the value set as `domain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#self_hosted_domains ZeroTrustAccessApplication#self_hosted_domains}
	SelfHostedDomains *[]*string `field:"optional" json:"selfHostedDomains" yaml:"selfHostedDomains"`
	// Option to return a 401 status code in service authentication rules on failed requests. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#service_auth_401_redirect ZeroTrustAccessApplication#service_auth_401_redirect}
	ServiceAuth401Redirect interface{} `field:"optional" json:"serviceAuth401Redirect" yaml:"serviceAuth401Redirect"`
	// How often a user will be forced to re-authorise.
	//
	// Must be in the format `48h` or `2h45m`. Defaults to `24h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#session_duration ZeroTrustAccessApplication#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Option to skip the App Launcher landing page. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#skip_app_launcher_login_page ZeroTrustAccessApplication#skip_app_launcher_login_page}
	SkipAppLauncherLoginPage interface{} `field:"optional" json:"skipAppLauncherLoginPage" yaml:"skipAppLauncherLoginPage"`
	// Option to skip the authorization interstitial when using the CLI. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#skip_interstitial ZeroTrustAccessApplication#skip_interstitial}
	SkipInterstitial interface{} `field:"optional" json:"skipInterstitial" yaml:"skipInterstitial"`
	// The itags associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#tags ZeroTrustAccessApplication#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// target_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#target_criteria ZeroTrustAccessApplication#target_criteria}
	TargetCriteria interface{} `field:"optional" json:"targetCriteria" yaml:"targetCriteria"`
	// The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`, `infrastructure`. Defaults to `self_hosted`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#type ZeroTrustAccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_access_application#zone_id ZeroTrustAccessApplication#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


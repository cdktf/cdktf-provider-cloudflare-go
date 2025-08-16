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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#account_id ZeroTrustAccessApplication#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// When set to true, users can authenticate to this application using their WARP session.
	//
	// When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#allow_authenticate_via_warp ZeroTrustAccessApplication#allow_authenticate_via_warp}
	AllowAuthenticateViaWarp interface{} `field:"optional" json:"allowAuthenticateViaWarp" yaml:"allowAuthenticateViaWarp"`
	// The identity providers your users can select when connecting to this application.
	//
	// Defaults to all IdPs configured in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#allowed_idps ZeroTrustAccessApplication#allowed_idps}
	AllowedIdps *[]*string `field:"optional" json:"allowedIdps" yaml:"allowedIdps"`
	// Enables loading application content in an iFrame.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#allow_iframe ZeroTrustAccessApplication#allow_iframe}
	AllowIframe interface{} `field:"optional" json:"allowIframe" yaml:"allowIframe"`
	// The image URL of the logo shown in the App Launcher header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#app_launcher_logo_url ZeroTrustAccessApplication#app_launcher_logo_url}
	AppLauncherLogoUrl *string `field:"optional" json:"appLauncherLogoUrl" yaml:"appLauncherLogoUrl"`
	// Displays the application in the App Launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#app_launcher_visible ZeroTrustAccessApplication#app_launcher_visible}
	AppLauncherVisible interface{} `field:"optional" json:"appLauncherVisible" yaml:"appLauncherVisible"`
	// When set to `true`, users skip the identity provider selection step during login.
	//
	// You must specify only one identity provider in allowed_idps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#auto_redirect_to_identity ZeroTrustAccessApplication#auto_redirect_to_identity}
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// The background color of the App Launcher page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#bg_color ZeroTrustAccessApplication#bg_color}
	BgColor *string `field:"optional" json:"bgColor" yaml:"bgColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#cors_headers ZeroTrustAccessApplication#cors_headers}.
	CorsHeaders *ZeroTrustAccessApplicationCorsHeaders `field:"optional" json:"corsHeaders" yaml:"corsHeaders"`
	// The custom error message shown to a user when they are denied access to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#custom_deny_message ZeroTrustAccessApplication#custom_deny_message}
	CustomDenyMessage *string `field:"optional" json:"customDenyMessage" yaml:"customDenyMessage"`
	// The custom URL a user is redirected to when they are denied access to the application when failing identity-based rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#custom_deny_url ZeroTrustAccessApplication#custom_deny_url}
	CustomDenyUrl *string `field:"optional" json:"customDenyUrl" yaml:"customDenyUrl"`
	// The custom URL a user is redirected to when they are denied access to the application when failing non-identity rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#custom_non_identity_deny_url ZeroTrustAccessApplication#custom_non_identity_deny_url}
	CustomNonIdentityDenyUrl *string `field:"optional" json:"customNonIdentityDenyUrl" yaml:"customNonIdentityDenyUrl"`
	// The custom pages that will be displayed when applicable for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#custom_pages ZeroTrustAccessApplication#custom_pages}
	CustomPages *[]*string `field:"optional" json:"customPages" yaml:"customPages"`
	// List of destinations secured by Access.
	//
	// This supersedes `self_hosted_domains` to allow for more flexibility in defining different types of domains. If `destinations` are provided, then `self_hosted_domains` will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#destinations ZeroTrustAccessApplication#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The primary hostname and path secured by Access.
	//
	// This domain will be displayed if the app is visible in the App Launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#domain ZeroTrustAccessApplication#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Enables the binding cookie, which increases security against compromised authorization tokens and CSRF attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#enable_binding_cookie ZeroTrustAccessApplication#enable_binding_cookie}
	EnableBindingCookie interface{} `field:"optional" json:"enableBindingCookie" yaml:"enableBindingCookie"`
	// The links in the App Launcher footer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#footer_links ZeroTrustAccessApplication#footer_links}
	FooterLinks interface{} `field:"optional" json:"footerLinks" yaml:"footerLinks"`
	// The background color of the App Launcher header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#header_bg_color ZeroTrustAccessApplication#header_bg_color}
	HeaderBgColor *string `field:"optional" json:"headerBgColor" yaml:"headerBgColor"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#http_only_cookie_attribute ZeroTrustAccessApplication#http_only_cookie_attribute}
	HttpOnlyCookieAttribute interface{} `field:"optional" json:"httpOnlyCookieAttribute" yaml:"httpOnlyCookieAttribute"`
	// The design of the App Launcher landing page shown to users when they log in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#landing_page_design ZeroTrustAccessApplication#landing_page_design}
	LandingPageDesign *ZeroTrustAccessApplicationLandingPageDesign `field:"optional" json:"landingPageDesign" yaml:"landingPageDesign"`
	// The image URL for the logo shown in the App Launcher dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#logo_url ZeroTrustAccessApplication#logo_url}
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Allows options preflight requests to bypass Access authentication and go directly to the origin.
	//
	// Cannot turn on if cors_headers is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#options_preflight_bypass ZeroTrustAccessApplication#options_preflight_bypass}
	OptionsPreflightBypass interface{} `field:"optional" json:"optionsPreflightBypass" yaml:"optionsPreflightBypass"`
	// Enables cookie paths to scope an application's JWT to the application path.
	//
	// If disabled, the JWT will scope to the hostname by default
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#path_cookie_attribute ZeroTrustAccessApplication#path_cookie_attribute}
	PathCookieAttribute interface{} `field:"optional" json:"pathCookieAttribute" yaml:"pathCookieAttribute"`
	// The policies that Access applies to the application, in ascending order of precedence.
	//
	// Items can reference existing policies or create new policies exclusive to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#policies ZeroTrustAccessApplication#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// Allows matching Access Service Tokens passed HTTP in a single header with this name.
	//
	// This works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.
	// The header value will be interpreted as a json object similar to:
	//   {
	//     "cf-access-client-id": "88bf3b6d86161464f6509f7219099e57.access.example.com",
	//     "cf-access-client-secret": "bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5"
	//   }
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#read_service_tokens_from_header ZeroTrustAccessApplication#read_service_tokens_from_header}
	ReadServiceTokensFromHeader *string `field:"optional" json:"readServiceTokensFromHeader" yaml:"readServiceTokensFromHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#saas_app ZeroTrustAccessApplication#saas_app}.
	SaasApp *ZeroTrustAccessApplicationSaasApp `field:"optional" json:"saasApp" yaml:"saasApp"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#same_site_cookie_attribute ZeroTrustAccessApplication#same_site_cookie_attribute}
	SameSiteCookieAttribute *string `field:"optional" json:"sameSiteCookieAttribute" yaml:"sameSiteCookieAttribute"`
	// Configuration for provisioning to this application via SCIM. This is currently in closed beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#scim_config ZeroTrustAccessApplication#scim_config}
	ScimConfig *ZeroTrustAccessApplicationScimConfig `field:"optional" json:"scimConfig" yaml:"scimConfig"`
	// List of public domains that Access will secure.
	//
	// This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#self_hosted_domains ZeroTrustAccessApplication#self_hosted_domains}
	SelfHostedDomains *[]*string `field:"optional" json:"selfHostedDomains" yaml:"selfHostedDomains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#service_auth_401_redirect ZeroTrustAccessApplication#service_auth_401_redirect}
	ServiceAuth401Redirect interface{} `field:"optional" json:"serviceAuth401Redirect" yaml:"serviceAuth401Redirect"`
	// The amount of time that tokens issued for this application will be valid.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. Note: unsupported for infrastructure type applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#session_duration ZeroTrustAccessApplication#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Determines when to skip the App Launcher landing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#skip_app_launcher_login_page ZeroTrustAccessApplication#skip_app_launcher_login_page}
	SkipAppLauncherLoginPage interface{} `field:"optional" json:"skipAppLauncherLoginPage" yaml:"skipAppLauncherLoginPage"`
	// Enables automatic authentication through cloudflared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#skip_interstitial ZeroTrustAccessApplication#skip_interstitial}
	SkipInterstitial interface{} `field:"optional" json:"skipInterstitial" yaml:"skipInterstitial"`
	// The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#tags ZeroTrustAccessApplication#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#target_criteria ZeroTrustAccessApplication#target_criteria}.
	TargetCriteria interface{} `field:"optional" json:"targetCriteria" yaml:"targetCriteria"`
	// The application type. Available values: "self_hosted", "saas", "ssh", "vnc", "app_launcher", "warp", "biso", "bookmark", "dash_sso", "infrastructure", "rdp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#type ZeroTrustAccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_access_application#zone_id ZeroTrustAccessApplication#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}


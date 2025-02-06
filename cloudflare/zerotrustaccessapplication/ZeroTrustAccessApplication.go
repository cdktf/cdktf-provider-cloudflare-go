// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application}.
type ZeroTrustAccessApplication interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowAuthenticateViaWarp() interface{}
	SetAllowAuthenticateViaWarp(val interface{})
	AllowAuthenticateViaWarpInput() interface{}
	AllowedIdps() *[]*string
	SetAllowedIdps(val *[]*string)
	AllowedIdpsInput() *[]*string
	AppLauncherLogoUrl() *string
	SetAppLauncherLogoUrl(val *string)
	AppLauncherLogoUrlInput() *string
	AppLauncherVisible() interface{}
	SetAppLauncherVisible(val interface{})
	AppLauncherVisibleInput() interface{}
	Aud() *string
	AutoRedirectToIdentity() interface{}
	SetAutoRedirectToIdentity(val interface{})
	AutoRedirectToIdentityInput() interface{}
	BgColor() *string
	SetBgColor(val *string)
	BgColorInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsHeaders() ZeroTrustAccessApplicationCorsHeadersList
	CorsHeadersInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDenyMessage() *string
	SetCustomDenyMessage(val *string)
	CustomDenyMessageInput() *string
	CustomDenyUrl() *string
	SetCustomDenyUrl(val *string)
	CustomDenyUrlInput() *string
	CustomNonIdentityDenyUrl() *string
	SetCustomNonIdentityDenyUrl(val *string)
	CustomNonIdentityDenyUrlInput() *string
	CustomPages() *[]*string
	SetCustomPages(val *[]*string)
	CustomPagesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destinations() ZeroTrustAccessApplicationDestinationsList
	DestinationsInput() interface{}
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	DomainType() *string
	SetDomainType(val *string)
	DomainTypeInput() *string
	EnableBindingCookie() interface{}
	SetEnableBindingCookie(val interface{})
	EnableBindingCookieInput() interface{}
	FooterLinks() ZeroTrustAccessApplicationFooterLinksList
	FooterLinksInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeaderBgColor() *string
	SetHeaderBgColor(val *string)
	HeaderBgColorInput() *string
	HttpOnlyCookieAttribute() interface{}
	SetHttpOnlyCookieAttribute(val interface{})
	HttpOnlyCookieAttributeInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	LandingPageDesign() ZeroTrustAccessApplicationLandingPageDesignOutputReference
	LandingPageDesignInput() *ZeroTrustAccessApplicationLandingPageDesign
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoUrl() *string
	SetLogoUrl(val *string)
	LogoUrlInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OptionsPreflightBypass() interface{}
	SetOptionsPreflightBypass(val interface{})
	OptionsPreflightBypassInput() interface{}
	Policies() *[]*string
	SetPolicies(val *[]*string)
	PoliciesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SaasApp() ZeroTrustAccessApplicationSaasAppOutputReference
	SaasAppInput() *ZeroTrustAccessApplicationSaasApp
	SameSiteCookieAttribute() *string
	SetSameSiteCookieAttribute(val *string)
	SameSiteCookieAttributeInput() *string
	ScimConfig() ZeroTrustAccessApplicationScimConfigOutputReference
	ScimConfigInput() *ZeroTrustAccessApplicationScimConfig
	SelfHostedDomains() *[]*string
	SetSelfHostedDomains(val *[]*string)
	SelfHostedDomainsInput() *[]*string
	ServiceAuth401Redirect() interface{}
	SetServiceAuth401Redirect(val interface{})
	ServiceAuth401RedirectInput() interface{}
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	SkipAppLauncherLoginPage() interface{}
	SetSkipAppLauncherLoginPage(val interface{})
	SkipAppLauncherLoginPageInput() interface{}
	SkipInterstitial() interface{}
	SetSkipInterstitial(val interface{})
	SkipInterstitialInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TargetCriteria() ZeroTrustAccessApplicationTargetCriteriaList
	TargetCriteriaInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCorsHeaders(value interface{})
	PutDestinations(value interface{})
	PutFooterLinks(value interface{})
	PutLandingPageDesign(value *ZeroTrustAccessApplicationLandingPageDesign)
	PutSaasApp(value *ZeroTrustAccessApplicationSaasApp)
	PutScimConfig(value *ZeroTrustAccessApplicationScimConfig)
	PutTargetCriteria(value interface{})
	ResetAccountId()
	ResetAllowAuthenticateViaWarp()
	ResetAllowedIdps()
	ResetAppLauncherLogoUrl()
	ResetAppLauncherVisible()
	ResetAutoRedirectToIdentity()
	ResetBgColor()
	ResetCorsHeaders()
	ResetCustomDenyMessage()
	ResetCustomDenyUrl()
	ResetCustomNonIdentityDenyUrl()
	ResetCustomPages()
	ResetDestinations()
	ResetDomain()
	ResetDomainType()
	ResetEnableBindingCookie()
	ResetFooterLinks()
	ResetHeaderBgColor()
	ResetHttpOnlyCookieAttribute()
	ResetId()
	ResetLandingPageDesign()
	ResetLogoUrl()
	ResetName()
	ResetOptionsPreflightBypass()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicies()
	ResetSaasApp()
	ResetSameSiteCookieAttribute()
	ResetScimConfig()
	ResetSelfHostedDomains()
	ResetServiceAuth401Redirect()
	ResetSessionDuration()
	ResetSkipAppLauncherLoginPage()
	ResetSkipInterstitial()
	ResetTags()
	ResetTargetCriteria()
	ResetType()
	ResetZoneId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ZeroTrustAccessApplication
type jsiiProxy_ZeroTrustAccessApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowAuthenticateViaWarp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowAuthenticateViaWarpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowedIdps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowedIdpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherLogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherLogoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherLogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherLogoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLauncherVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLauncherVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Aud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AutoRedirectToIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AutoRedirectToIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) BgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) BgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CorsHeaders() ZeroTrustAccessApplicationCorsHeadersList {
	var returns ZeroTrustAccessApplicationCorsHeadersList
	_jsii_.Get(
		j,
		"corsHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CorsHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomNonIdentityDenyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customNonIdentityDenyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomNonIdentityDenyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customNonIdentityDenyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Destinations() ZeroTrustAccessApplicationDestinationsList {
	var returns ZeroTrustAccessApplicationDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DomainType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DomainTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) EnableBindingCookie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBindingCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) EnableBindingCookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBindingCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FooterLinks() ZeroTrustAccessApplicationFooterLinksList {
	var returns ZeroTrustAccessApplicationFooterLinksList
	_jsii_.Get(
		j,
		"footerLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FooterLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"footerLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HeaderBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HeaderBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HttpOnlyCookieAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpOnlyCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HttpOnlyCookieAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpOnlyCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LandingPageDesign() ZeroTrustAccessApplicationLandingPageDesignOutputReference {
	var returns ZeroTrustAccessApplicationLandingPageDesignOutputReference
	_jsii_.Get(
		j,
		"landingPageDesign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LandingPageDesignInput() *ZeroTrustAccessApplicationLandingPageDesign {
	var returns *ZeroTrustAccessApplicationLandingPageDesign
	_jsii_.Get(
		j,
		"landingPageDesignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) OptionsPreflightBypass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionsPreflightBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) OptionsPreflightBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionsPreflightBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) PoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SaasApp() ZeroTrustAccessApplicationSaasAppOutputReference {
	var returns ZeroTrustAccessApplicationSaasAppOutputReference
	_jsii_.Get(
		j,
		"saasApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SaasAppInput() *ZeroTrustAccessApplicationSaasApp {
	var returns *ZeroTrustAccessApplicationSaasApp
	_jsii_.Get(
		j,
		"saasAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SameSiteCookieAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SameSiteCookieAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ScimConfig() ZeroTrustAccessApplicationScimConfigOutputReference {
	var returns ZeroTrustAccessApplicationScimConfigOutputReference
	_jsii_.Get(
		j,
		"scimConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ScimConfigInput() *ZeroTrustAccessApplicationScimConfig {
	var returns *ZeroTrustAccessApplicationScimConfig
	_jsii_.Get(
		j,
		"scimConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SelfHostedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selfHostedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SelfHostedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selfHostedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ServiceAuth401Redirect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAuth401Redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ServiceAuth401RedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAuth401RedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipAppLauncherLoginPage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipAppLauncherLoginPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipAppLauncherLoginPageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipAppLauncherLoginPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipInterstitial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInterstitial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipInterstitialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInterstitialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TargetCriteria() ZeroTrustAccessApplicationTargetCriteriaList {
	var returns ZeroTrustAccessApplicationTargetCriteriaList
	_jsii_.Get(
		j,
		"targetCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TargetCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application} Resource.
func NewZeroTrustAccessApplication(scope constructs.Construct, id *string, config *ZeroTrustAccessApplicationConfig) ZeroTrustAccessApplication {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplication{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application} Resource.
func NewZeroTrustAccessApplication_Override(z ZeroTrustAccessApplication, scope constructs.Construct, id *string, config *ZeroTrustAccessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAllowAuthenticateViaWarp(val interface{}) {
	if err := j.validateSetAllowAuthenticateViaWarpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAuthenticateViaWarp",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAllowedIdps(val *[]*string) {
	if err := j.validateSetAllowedIdpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIdps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAppLauncherLogoUrl(val *string) {
	if err := j.validateSetAppLauncherLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherLogoUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAppLauncherVisible(val interface{}) {
	if err := j.validateSetAppLauncherVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherVisible",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetAutoRedirectToIdentity(val interface{}) {
	if err := j.validateSetAutoRedirectToIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRedirectToIdentity",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetBgColor(val *string) {
	if err := j.validateSetBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetCustomDenyMessage(val *string) {
	if err := j.validateSetCustomDenyMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyMessage",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetCustomDenyUrl(val *string) {
	if err := j.validateSetCustomDenyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetCustomNonIdentityDenyUrl(val *string) {
	if err := j.validateSetCustomNonIdentityDenyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customNonIdentityDenyUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetCustomPages(val *[]*string) {
	if err := j.validateSetCustomPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPages",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetDomainType(val *string) {
	if err := j.validateSetDomainTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetEnableBindingCookie(val interface{}) {
	if err := j.validateSetEnableBindingCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBindingCookie",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetHeaderBgColor(val *string) {
	if err := j.validateSetHeaderBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerBgColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetHttpOnlyCookieAttribute(val interface{}) {
	if err := j.validateSetHttpOnlyCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpOnlyCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetLogoUrl(val *string) {
	if err := j.validateSetLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetOptionsPreflightBypass(val interface{}) {
	if err := j.validateSetOptionsPreflightBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionsPreflightBypass",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetPolicies(val *[]*string) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetSameSiteCookieAttribute(val *string) {
	if err := j.validateSetSameSiteCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sameSiteCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetSelfHostedDomains(val *[]*string) {
	if err := j.validateSetSelfHostedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfHostedDomains",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetServiceAuth401Redirect(val interface{}) {
	if err := j.validateSetServiceAuth401RedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAuth401Redirect",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetSkipAppLauncherLoginPage(val interface{}) {
	if err := j.validateSetSkipAppLauncherLoginPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipAppLauncherLoginPage",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetSkipInterstitial(val interface{}) {
	if err := j.validateSetSkipInterstitialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipInterstitial",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustAccessApplication resource upon running "cdktf plan <stack-name>".
func ZeroTrustAccessApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ZeroTrustAccessApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustAccessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutCorsHeaders(value interface{}) {
	if err := z.validatePutCorsHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCorsHeaders",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutDestinations(value interface{}) {
	if err := z.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDestinations",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutFooterLinks(value interface{}) {
	if err := z.validatePutFooterLinksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putFooterLinks",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutLandingPageDesign(value *ZeroTrustAccessApplicationLandingPageDesign) {
	if err := z.validatePutLandingPageDesignParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLandingPageDesign",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutSaasApp(value *ZeroTrustAccessApplicationSaasApp) {
	if err := z.validatePutSaasAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaasApp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutScimConfig(value *ZeroTrustAccessApplicationScimConfig) {
	if err := z.validatePutScimConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putScimConfig",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutTargetCriteria(value interface{}) {
	if err := z.validatePutTargetCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putTargetCriteria",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAccountId() {
	_jsii_.InvokeVoid(
		z,
		"resetAccountId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAllowAuthenticateViaWarp() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowAuthenticateViaWarp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAllowedIdps() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedIdps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAppLauncherLogoUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetAppLauncherLogoUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAppLauncherVisible() {
	_jsii_.InvokeVoid(
		z,
		"resetAppLauncherVisible",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAutoRedirectToIdentity() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoRedirectToIdentity",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetBgColor() {
	_jsii_.InvokeVoid(
		z,
		"resetBgColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCorsHeaders() {
	_jsii_.InvokeVoid(
		z,
		"resetCorsHeaders",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomDenyMessage() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomDenyMessage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomDenyUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomDenyUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomNonIdentityDenyUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomNonIdentityDenyUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomPages() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomPages",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetDestinations() {
	_jsii_.InvokeVoid(
		z,
		"resetDestinations",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetDomainType() {
	_jsii_.InvokeVoid(
		z,
		"resetDomainType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetEnableBindingCookie() {
	_jsii_.InvokeVoid(
		z,
		"resetEnableBindingCookie",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetFooterLinks() {
	_jsii_.InvokeVoid(
		z,
		"resetFooterLinks",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetHeaderBgColor() {
	_jsii_.InvokeVoid(
		z,
		"resetHeaderBgColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetHttpOnlyCookieAttribute() {
	_jsii_.InvokeVoid(
		z,
		"resetHttpOnlyCookieAttribute",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetLandingPageDesign() {
	_jsii_.InvokeVoid(
		z,
		"resetLandingPageDesign",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetLogoUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetLogoUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetName() {
	_jsii_.InvokeVoid(
		z,
		"resetName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetOptionsPreflightBypass() {
	_jsii_.InvokeVoid(
		z,
		"resetOptionsPreflightBypass",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetPolicies() {
	_jsii_.InvokeVoid(
		z,
		"resetPolicies",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSaasApp() {
	_jsii_.InvokeVoid(
		z,
		"resetSaasApp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSameSiteCookieAttribute() {
	_jsii_.InvokeVoid(
		z,
		"resetSameSiteCookieAttribute",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetScimConfig() {
	_jsii_.InvokeVoid(
		z,
		"resetScimConfig",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSelfHostedDomains() {
	_jsii_.InvokeVoid(
		z,
		"resetSelfHostedDomains",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetServiceAuth401Redirect() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceAuth401Redirect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSkipAppLauncherLoginPage() {
	_jsii_.InvokeVoid(
		z,
		"resetSkipAppLauncherLoginPage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSkipInterstitial() {
	_jsii_.InvokeVoid(
		z,
		"resetSkipInterstitial",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetTags() {
	_jsii_.InvokeVoid(
		z,
		"resetTags",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetTargetCriteria() {
	_jsii_.InvokeVoid(
		z,
		"resetTargetCriteria",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetType() {
	_jsii_.InvokeVoid(
		z,
		"resetType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetZoneId() {
	_jsii_.InvokeVoid(
		z,
		"resetZoneId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustaccessidentityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessIdentityProviderConfigAOutputReference interface {
	cdktf.ComplexObject
	AppsDomain() *string
	SetAppsDomain(val *string)
	AppsDomainInput() *string
	Attributes() *[]*string
	SetAttributes(val *[]*string)
	AttributesInput() *[]*string
	AuthorizationServerId() *string
	SetAuthorizationServerId(val *string)
	AuthorizationServerIdInput() *string
	AuthUrl() *string
	SetAuthUrl(val *string)
	AuthUrlInput() *string
	CentrifyAccount() *string
	SetCentrifyAccount(val *string)
	CentrifyAccountInput() *string
	CentrifyAppId() *string
	SetCentrifyAppId(val *string)
	CentrifyAppIdInput() *string
	CertsUrl() *string
	SetCertsUrl(val *string)
	CertsUrlInput() *string
	Claims() *[]*string
	SetClaims(val *[]*string)
	ClaimsInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConditionalAccessEnabled() interface{}
	SetConditionalAccessEnabled(val interface{})
	ConditionalAccessEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	EmailAttributeName() *string
	SetEmailAttributeName(val *string)
	EmailAttributeNameInput() *string
	EmailClaimName() *string
	SetEmailClaimName(val *string)
	EmailClaimNameInput() *string
	// Experimental.
	Fqn() *string
	HeaderAttributes() ZeroTrustAccessIdentityProviderConfigHeaderAttributesList
	HeaderAttributesInput() interface{}
	IdpPublicCerts() *[]*string
	SetIdpPublicCerts(val *[]*string)
	IdpPublicCertsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssuerUrl() *string
	SetIssuerUrl(val *string)
	IssuerUrlInput() *string
	OktaAccount() *string
	SetOktaAccount(val *string)
	OktaAccountInput() *string
	OneloginAccount() *string
	SetOneloginAccount(val *string)
	OneloginAccountInput() *string
	PingEnvId() *string
	SetPingEnvId(val *string)
	PingEnvIdInput() *string
	PkceEnabled() interface{}
	SetPkceEnabled(val interface{})
	PkceEnabledInput() interface{}
	Prompt() *string
	SetPrompt(val *string)
	PromptInput() *string
	RedirectUrl() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SignRequest() interface{}
	SetSignRequest(val interface{})
	SignRequestInput() interface{}
	SsoTargetUrl() *string
	SetSsoTargetUrl(val *string)
	SsoTargetUrlInput() *string
	SupportGroups() interface{}
	SetSupportGroups(val interface{})
	SupportGroupsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutHeaderAttributes(value interface{})
	ResetAppsDomain()
	ResetAttributes()
	ResetAuthorizationServerId()
	ResetAuthUrl()
	ResetCentrifyAccount()
	ResetCentrifyAppId()
	ResetCertsUrl()
	ResetClaims()
	ResetClientId()
	ResetClientSecret()
	ResetConditionalAccessEnabled()
	ResetDirectoryId()
	ResetEmailAttributeName()
	ResetEmailClaimName()
	ResetHeaderAttributes()
	ResetIdpPublicCerts()
	ResetIssuerUrl()
	ResetOktaAccount()
	ResetOneloginAccount()
	ResetPingEnvId()
	ResetPkceEnabled()
	ResetPrompt()
	ResetScopes()
	ResetSignRequest()
	ResetSsoTargetUrl()
	ResetSupportGroups()
	ResetTokenUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessIdentityProviderConfigAOutputReference
type jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AppsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AppsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Attributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AuthorizationServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AuthorizationServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) AuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CentrifyAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CentrifyAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CentrifyAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CentrifyAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CertsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CertsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Claims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ClaimsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ConditionalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ConditionalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) EmailAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) EmailAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) EmailClaimName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailClaimName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) EmailClaimNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailClaimNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) HeaderAttributes() ZeroTrustAccessIdentityProviderConfigHeaderAttributesList {
	var returns ZeroTrustAccessIdentityProviderConfigHeaderAttributesList
	_jsii_.Get(
		j,
		"headerAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) HeaderAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) IdpPublicCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpPublicCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) IdpPublicCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpPublicCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) OktaAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) OktaAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) OneloginAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) OneloginAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PingEnvId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pingEnvId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PingEnvIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pingEnvIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PkceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PkceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SignRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SignRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SsoTargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SsoTargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SupportGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) SupportGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessIdentityProviderConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessIdentityProviderConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessIdentityProviderConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessIdentityProvider.ZeroTrustAccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessIdentityProviderConfigAOutputReference_Override(z ZeroTrustAccessIdentityProviderConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessIdentityProvider.ZeroTrustAccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetAppsDomain(val *string) {
	if err := j.validateSetAppsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appsDomain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetAttributes(val *[]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetAuthorizationServerId(val *string) {
	if err := j.validateSetAuthorizationServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationServerId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetAuthUrl(val *string) {
	if err := j.validateSetAuthUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetCentrifyAccount(val *string) {
	if err := j.validateSetCentrifyAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"centrifyAccount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetCentrifyAppId(val *string) {
	if err := j.validateSetCentrifyAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"centrifyAppId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetCertsUrl(val *string) {
	if err := j.validateSetCertsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certsUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetClaims(val *[]*string) {
	if err := j.validateSetClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claims",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetConditionalAccessEnabled(val interface{}) {
	if err := j.validateSetConditionalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetEmailAttributeName(val *string) {
	if err := j.validateSetEmailAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAttributeName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetEmailClaimName(val *string) {
	if err := j.validateSetEmailClaimNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailClaimName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetIdpPublicCerts(val *[]*string) {
	if err := j.validateSetIdpPublicCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpPublicCerts",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetOktaAccount(val *string) {
	if err := j.validateSetOktaAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaAccount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetOneloginAccount(val *string) {
	if err := j.validateSetOneloginAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oneloginAccount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetPingEnvId(val *string) {
	if err := j.validateSetPingEnvIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pingEnvId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetPkceEnabled(val interface{}) {
	if err := j.validateSetPkceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetPrompt(val *string) {
	if err := j.validateSetPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompt",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetSignRequest(val interface{}) {
	if err := j.validateSetSignRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signRequest",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetSsoTargetUrl(val *string) {
	if err := j.validateSetSsoTargetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoTargetUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetSupportGroups(val interface{}) {
	if err := j.validateSetSupportGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportGroups",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) PutHeaderAttributes(value interface{}) {
	if err := z.validatePutHeaderAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putHeaderAttributes",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetAppsDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetAppsDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		z,
		"resetAttributes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetAuthorizationServerId() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthorizationServerId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetAuthUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetCentrifyAccount() {
	_jsii_.InvokeVoid(
		z,
		"resetCentrifyAccount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetCentrifyAppId() {
	_jsii_.InvokeVoid(
		z,
		"resetCentrifyAppId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetCertsUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetCertsUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetClaims() {
	_jsii_.InvokeVoid(
		z,
		"resetClaims",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		z,
		"resetClientId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		z,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetConditionalAccessEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetConditionalAccessEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		z,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetEmailAttributeName() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailAttributeName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetEmailClaimName() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailClaimName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetHeaderAttributes() {
	_jsii_.InvokeVoid(
		z,
		"resetHeaderAttributes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetIdpPublicCerts() {
	_jsii_.InvokeVoid(
		z,
		"resetIdpPublicCerts",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetIssuerUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetIssuerUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetOktaAccount() {
	_jsii_.InvokeVoid(
		z,
		"resetOktaAccount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetOneloginAccount() {
	_jsii_.InvokeVoid(
		z,
		"resetOneloginAccount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetPingEnvId() {
	_jsii_.InvokeVoid(
		z,
		"resetPingEnvId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetPkceEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetPkceEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		z,
		"resetPrompt",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		z,
		"resetScopes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetSignRequest() {
	_jsii_.InvokeVoid(
		z,
		"resetSignRequest",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetSsoTargetUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetSsoTargetUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetSupportGroups() {
	_jsii_.InvokeVoid(
		z,
		"resetSupportGroups",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


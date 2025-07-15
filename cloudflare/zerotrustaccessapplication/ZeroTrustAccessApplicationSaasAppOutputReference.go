// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessApplicationSaasAppOutputReference interface {
	cdktf.ComplexObject
	AccessTokenLifetime() *string
	SetAccessTokenLifetime(val *string)
	AccessTokenLifetimeInput() *string
	AllowPkceWithoutClientSecret() interface{}
	SetAllowPkceWithoutClientSecret(val interface{})
	AllowPkceWithoutClientSecretInput() interface{}
	AppLauncherUrl() *string
	SetAppLauncherUrl(val *string)
	AppLauncherUrlInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	ClientId() *string
	ClientSecret() *string
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
	ConsumerServiceUrl() *string
	SetConsumerServiceUrl(val *string)
	ConsumerServiceUrlInput() *string
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAttributes() ZeroTrustAccessApplicationSaasAppCustomAttributesList
	CustomAttributesInput() interface{}
	CustomClaims() ZeroTrustAccessApplicationSaasAppCustomClaimsList
	CustomClaimsInput() interface{}
	DefaultRelayState() *string
	SetDefaultRelayState(val *string)
	DefaultRelayStateInput() *string
	// Experimental.
	Fqn() *string
	GrantTypes() *[]*string
	SetGrantTypes(val *[]*string)
	GrantTypesInput() *[]*string
	GroupFilterRegex() *string
	SetGroupFilterRegex(val *string)
	GroupFilterRegexInput() *string
	HybridAndImplicitOptions() ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	HybridAndImplicitOptionsInput() interface{}
	IdpEntityId() *string
	SetIdpEntityId(val *string)
	IdpEntityIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NameIdFormat() *string
	SetNameIdFormat(val *string)
	NameIdFormatInput() *string
	NameIdTransformJsonata() *string
	SetNameIdTransformJsonata(val *string)
	NameIdTransformJsonataInput() *string
	PublicKey() *string
	RedirectUris() *[]*string
	SetRedirectUris(val *[]*string)
	RedirectUrisInput() *[]*string
	RefreshTokenOptions() ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference
	RefreshTokenOptionsInput() interface{}
	SamlAttributeTransformJsonata() *string
	SetSamlAttributeTransformJsonata(val *string)
	SamlAttributeTransformJsonataInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SpEntityId() *string
	SetSpEntityId(val *string)
	SpEntityIdInput() *string
	SsoEndpoint() *string
	SetSsoEndpoint(val *string)
	SsoEndpointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
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
	PutCustomAttributes(value interface{})
	PutCustomClaims(value interface{})
	PutHybridAndImplicitOptions(value *ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions)
	PutRefreshTokenOptions(value *ZeroTrustAccessApplicationSaasAppRefreshTokenOptions)
	ResetAccessTokenLifetime()
	ResetAllowPkceWithoutClientSecret()
	ResetAppLauncherUrl()
	ResetAuthType()
	ResetConsumerServiceUrl()
	ResetCustomAttributes()
	ResetCustomClaims()
	ResetDefaultRelayState()
	ResetGrantTypes()
	ResetGroupFilterRegex()
	ResetHybridAndImplicitOptions()
	ResetIdpEntityId()
	ResetNameIdFormat()
	ResetNameIdTransformJsonata()
	ResetRedirectUris()
	ResetRefreshTokenOptions()
	ResetSamlAttributeTransformJsonata()
	ResetScopes()
	ResetSpEntityId()
	ResetSsoEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationSaasAppOutputReference
type jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AccessTokenLifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AccessTokenLifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AllowPkceWithoutClientSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPkceWithoutClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AllowPkceWithoutClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPkceWithoutClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AppLauncherUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AppLauncherUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ConsumerServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ConsumerServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CustomAttributes() ZeroTrustAccessApplicationSaasAppCustomAttributesList {
	var returns ZeroTrustAccessApplicationSaasAppCustomAttributesList
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CustomAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CustomClaims() ZeroTrustAccessApplicationSaasAppCustomClaimsList {
	var returns ZeroTrustAccessApplicationSaasAppCustomClaimsList
	_jsii_.Get(
		j,
		"customClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) CustomClaimsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) DefaultRelayStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GroupFilterRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GroupFilterRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) HybridAndImplicitOptions() ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference {
	var returns ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	_jsii_.Get(
		j,
		"hybridAndImplicitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) HybridAndImplicitOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hybridAndImplicitOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) IdpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) NameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) NameIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) NameIdTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) NameIdTransformJsonataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdTransformJsonataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) RefreshTokenOptions() ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference {
	var returns ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference
	_jsii_.Get(
		j,
		"refreshTokenOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) RefreshTokenOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshTokenOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SamlAttributeTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAttributeTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SamlAttributeTransformJsonataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAttributeTransformJsonataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SsoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) SsoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessApplicationSaasAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessApplicationSaasAppOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationSaasAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationSaasAppOutputReference_Override(z ZeroTrustAccessApplicationSaasAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetAccessTokenLifetime(val *string) {
	if err := j.validateSetAccessTokenLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenLifetime",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetAllowPkceWithoutClientSecret(val interface{}) {
	if err := j.validateSetAllowPkceWithoutClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPkceWithoutClientSecret",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetAppLauncherUrl(val *string) {
	if err := j.validateSetAppLauncherUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetConsumerServiceUrl(val *string) {
	if err := j.validateSetConsumerServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerServiceUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetDefaultRelayState(val *string) {
	if err := j.validateSetDefaultRelayStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRelayState",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetGroupFilterRegex(val *string) {
	if err := j.validateSetGroupFilterRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupFilterRegex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetIdpEntityId(val *string) {
	if err := j.validateSetIdpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpEntityId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetNameIdFormat(val *string) {
	if err := j.validateSetNameIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameIdFormat",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetNameIdTransformJsonata(val *string) {
	if err := j.validateSetNameIdTransformJsonataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameIdTransformJsonata",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetRedirectUris(val *[]*string) {
	if err := j.validateSetRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetSamlAttributeTransformJsonata(val *string) {
	if err := j.validateSetSamlAttributeTransformJsonataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlAttributeTransformJsonata",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetSpEntityId(val *string) {
	if err := j.validateSetSpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spEntityId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetSsoEndpoint(val *string) {
	if err := j.validateSetSsoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoEndpoint",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) PutCustomAttributes(value interface{}) {
	if err := z.validatePutCustomAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCustomAttributes",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) PutCustomClaims(value interface{}) {
	if err := z.validatePutCustomClaimsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCustomClaims",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) PutHybridAndImplicitOptions(value *ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions) {
	if err := z.validatePutHybridAndImplicitOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putHybridAndImplicitOptions",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) PutRefreshTokenOptions(value *ZeroTrustAccessApplicationSaasAppRefreshTokenOptions) {
	if err := z.validatePutRefreshTokenOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putRefreshTokenOptions",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetAccessTokenLifetime() {
	_jsii_.InvokeVoid(
		z,
		"resetAccessTokenLifetime",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetAllowPkceWithoutClientSecret() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowPkceWithoutClientSecret",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetAppLauncherUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetAppLauncherUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetConsumerServiceUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetConsumerServiceUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetCustomClaims() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomClaims",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetDefaultRelayState() {
	_jsii_.InvokeVoid(
		z,
		"resetDefaultRelayState",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetGrantTypes() {
	_jsii_.InvokeVoid(
		z,
		"resetGrantTypes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetGroupFilterRegex() {
	_jsii_.InvokeVoid(
		z,
		"resetGroupFilterRegex",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetHybridAndImplicitOptions() {
	_jsii_.InvokeVoid(
		z,
		"resetHybridAndImplicitOptions",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetIdpEntityId() {
	_jsii_.InvokeVoid(
		z,
		"resetIdpEntityId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetNameIdFormat() {
	_jsii_.InvokeVoid(
		z,
		"resetNameIdFormat",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetNameIdTransformJsonata() {
	_jsii_.InvokeVoid(
		z,
		"resetNameIdTransformJsonata",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetRedirectUris() {
	_jsii_.InvokeVoid(
		z,
		"resetRedirectUris",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetRefreshTokenOptions() {
	_jsii_.InvokeVoid(
		z,
		"resetRefreshTokenOptions",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetSamlAttributeTransformJsonata() {
	_jsii_.InvokeVoid(
		z,
		"resetSamlAttributeTransformJsonata",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		z,
		"resetScopes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetSpEntityId() {
	_jsii_.InvokeVoid(
		z,
		"resetSpEntityId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ResetSsoEndpoint() {
	_jsii_.InvokeVoid(
		z,
		"resetSsoEndpoint",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


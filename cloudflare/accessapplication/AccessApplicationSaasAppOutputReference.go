// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/accessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessApplicationSaasAppOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAttribute() AccessApplicationSaasAppCustomAttributeList
	CustomAttributeInput() interface{}
	CustomClaim() AccessApplicationSaasAppCustomClaimList
	CustomClaimInput() interface{}
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
	HybridAndImplicitOptions() AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	HybridAndImplicitOptionsInput() *AccessApplicationSaasAppHybridAndImplicitOptions
	IdpEntityId() *string
	InternalValue() *AccessApplicationSaasApp
	SetInternalValue(val *AccessApplicationSaasApp)
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
	RefreshTokenOptions() AccessApplicationSaasAppRefreshTokenOptionsList
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCustomAttribute(value interface{})
	PutCustomClaim(value interface{})
	PutHybridAndImplicitOptions(value *AccessApplicationSaasAppHybridAndImplicitOptions)
	PutRefreshTokenOptions(value interface{})
	ResetAllowPkceWithoutClientSecret()
	ResetAppLauncherUrl()
	ResetAuthType()
	ResetConsumerServiceUrl()
	ResetCustomAttribute()
	ResetCustomClaim()
	ResetDefaultRelayState()
	ResetGrantTypes()
	ResetGroupFilterRegex()
	ResetHybridAndImplicitOptions()
	ResetNameIdFormat()
	ResetNameIdTransformJsonata()
	ResetRedirectUris()
	ResetRefreshTokenOptions()
	ResetSamlAttributeTransformJsonata()
	ResetScopes()
	ResetSpEntityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessApplicationSaasAppOutputReference
type jsiiProxy_AccessApplicationSaasAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AllowPkceWithoutClientSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPkceWithoutClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AllowPkceWithoutClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPkceWithoutClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AppLauncherUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AppLauncherUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ConsumerServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ConsumerServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomAttribute() AccessApplicationSaasAppCustomAttributeList {
	var returns AccessApplicationSaasAppCustomAttributeList
	_jsii_.Get(
		j,
		"customAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomClaim() AccessApplicationSaasAppCustomClaimList {
	var returns AccessApplicationSaasAppCustomClaimList
	_jsii_.Get(
		j,
		"customClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) DefaultRelayStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) GroupFilterRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) GroupFilterRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) HybridAndImplicitOptions() AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference {
	var returns AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	_jsii_.Get(
		j,
		"hybridAndImplicitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) HybridAndImplicitOptionsInput() *AccessApplicationSaasAppHybridAndImplicitOptions {
	var returns *AccessApplicationSaasAppHybridAndImplicitOptions
	_jsii_.Get(
		j,
		"hybridAndImplicitOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) InternalValue() *AccessApplicationSaasApp {
	var returns *AccessApplicationSaasApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdTransformJsonataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdTransformJsonataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) RefreshTokenOptions() AccessApplicationSaasAppRefreshTokenOptionsList {
	var returns AccessApplicationSaasAppRefreshTokenOptionsList
	_jsii_.Get(
		j,
		"refreshTokenOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) RefreshTokenOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshTokenOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SamlAttributeTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAttributeTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SamlAttributeTransformJsonataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAttributeTransformJsonataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SsoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessApplicationSaasAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessApplicationSaasAppOutputReference {
	_init_.Initialize()

	if err := validateNewAccessApplicationSaasAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplicationSaasAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessApplicationSaasAppOutputReference_Override(a AccessApplicationSaasAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetAllowPkceWithoutClientSecret(val interface{}) {
	if err := j.validateSetAllowPkceWithoutClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPkceWithoutClientSecret",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetAppLauncherUrl(val *string) {
	if err := j.validateSetAppLauncherUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherUrl",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetConsumerServiceUrl(val *string) {
	if err := j.validateSetConsumerServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerServiceUrl",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetDefaultRelayState(val *string) {
	if err := j.validateSetDefaultRelayStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRelayState",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetGroupFilterRegex(val *string) {
	if err := j.validateSetGroupFilterRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupFilterRegex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetInternalValue(val *AccessApplicationSaasApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetNameIdFormat(val *string) {
	if err := j.validateSetNameIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameIdFormat",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetNameIdTransformJsonata(val *string) {
	if err := j.validateSetNameIdTransformJsonataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameIdTransformJsonata",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetRedirectUris(val *[]*string) {
	if err := j.validateSetRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetSamlAttributeTransformJsonata(val *string) {
	if err := j.validateSetSamlAttributeTransformJsonataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlAttributeTransformJsonata",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetSpEntityId(val *string) {
	if err := j.validateSetSpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spEntityId",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) PutCustomAttribute(value interface{}) {
	if err := a.validatePutCustomAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomAttribute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) PutCustomClaim(value interface{}) {
	if err := a.validatePutCustomClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomClaim",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) PutHybridAndImplicitOptions(value *AccessApplicationSaasAppHybridAndImplicitOptions) {
	if err := a.validatePutHybridAndImplicitOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHybridAndImplicitOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) PutRefreshTokenOptions(value interface{}) {
	if err := a.validatePutRefreshTokenOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRefreshTokenOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetAllowPkceWithoutClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowPkceWithoutClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetAppLauncherUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLauncherUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetConsumerServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetConsumerServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetCustomAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetCustomClaim() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomClaim",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetDefaultRelayState() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRelayState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetGrantTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetGrantTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetGroupFilterRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupFilterRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetHybridAndImplicitOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetHybridAndImplicitOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetNameIdFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetNameIdFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetNameIdTransformJsonata() {
	_jsii_.InvokeVoid(
		a,
		"resetNameIdTransformJsonata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetRedirectUris() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetRefreshTokenOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetSamlAttributeTransformJsonata() {
	_jsii_.InvokeVoid(
		a,
		"resetSamlAttributeTransformJsonata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetSpEntityId() {
	_jsii_.InvokeVoid(
		a,
		"resetSpEntityId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


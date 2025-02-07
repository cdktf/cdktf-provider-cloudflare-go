// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference interface {
	cdktf.ComplexObject
	AccessTokenLifetime() *string
	AllowPkceWithoutClientSecret() cdktf.IResolvable
	AppLauncherUrl() *string
	AuthType() *string
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAttributes() DataCloudflareZeroTrustAccessApplicationSaasAppCustomAttributesList
	CustomClaims() DataCloudflareZeroTrustAccessApplicationSaasAppCustomClaimsList
	DefaultRelayState() *string
	// Experimental.
	Fqn() *string
	GrantTypes() *[]*string
	GroupFilterRegex() *string
	HybridAndImplicitOptions() DataCloudflareZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	IdpEntityId() *string
	InternalValue() *DataCloudflareZeroTrustAccessApplicationSaasApp
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationSaasApp)
	NameIdFormat() *string
	NameIdTransformJsonata() *string
	PublicKey() *string
	RedirectUris() *[]*string
	RefreshTokenOptions() DataCloudflareZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference
	SamlAttributeTransformJsonata() *string
	Scopes() *[]*string
	SpEntityId() *string
	SsoEndpoint() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) AccessTokenLifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) AllowPkceWithoutClientSecret() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowPkceWithoutClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) AppLauncherUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ConsumerServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) CustomAttributes() DataCloudflareZeroTrustAccessApplicationSaasAppCustomAttributesList {
	var returns DataCloudflareZeroTrustAccessApplicationSaasAppCustomAttributesList
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) CustomClaims() DataCloudflareZeroTrustAccessApplicationSaasAppCustomClaimsList {
	var returns DataCloudflareZeroTrustAccessApplicationSaasAppCustomClaimsList
	_jsii_.Get(
		j,
		"customClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GroupFilterRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) HybridAndImplicitOptions() DataCloudflareZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
	_jsii_.Get(
		j,
		"hybridAndImplicitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationSaasApp {
	var returns *DataCloudflareZeroTrustAccessApplicationSaasApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) NameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) NameIdTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) RefreshTokenOptions() DataCloudflareZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference
	_jsii_.Get(
		j,
		"refreshTokenOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) SamlAttributeTransformJsonata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAttributeTransformJsonata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) SpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) SsoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationSaasAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationSaasAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationSaasAppOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationSaasApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationSaasAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


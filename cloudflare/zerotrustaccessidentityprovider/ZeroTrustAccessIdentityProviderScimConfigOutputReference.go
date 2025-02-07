// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustaccessidentityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessIdentityProviderScimConfigOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IdentityUpdateBehavior() *string
	SetIdentityUpdateBehavior(val *string)
	IdentityUpdateBehaviorInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScimBaseUrl() *string
	SeatDeprovision() interface{}
	SetSeatDeprovision(val interface{})
	SeatDeprovisionInput() interface{}
	Secret() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserDeprovision() interface{}
	SetUserDeprovision(val interface{})
	UserDeprovisionInput() interface{}
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
	ResetEnabled()
	ResetIdentityUpdateBehavior()
	ResetSeatDeprovision()
	ResetUserDeprovision()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessIdentityProviderScimConfigOutputReference
type jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) IdentityUpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityUpdateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) IdentityUpdateBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityUpdateBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ScimBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) SeatDeprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"seatDeprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) SeatDeprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"seatDeprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) UserDeprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDeprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) UserDeprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDeprovisionInput",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessIdentityProviderScimConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessIdentityProviderScimConfigOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessIdentityProviderScimConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessIdentityProvider.ZeroTrustAccessIdentityProviderScimConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessIdentityProviderScimConfigOutputReference_Override(z ZeroTrustAccessIdentityProviderScimConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessIdentityProvider.ZeroTrustAccessIdentityProviderScimConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetIdentityUpdateBehavior(val *string) {
	if err := j.validateSetIdentityUpdateBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityUpdateBehavior",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetSeatDeprovision(val interface{}) {
	if err := j.validateSetSeatDeprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seatDeprovision",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference)SetUserDeprovision(val interface{}) {
	if err := j.validateSetUserDeprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDeprovision",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ResetIdentityUpdateBehavior() {
	_jsii_.InvokeVoid(
		z,
		"resetIdentityUpdateBehavior",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ResetSeatDeprovision() {
	_jsii_.InvokeVoid(
		z,
		"resetSeatDeprovision",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ResetUserDeprovision() {
	_jsii_.InvokeVoid(
		z,
		"resetUserDeprovision",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessIdentityProviderScimConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


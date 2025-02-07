// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference interface {
	cdktf.ComplexObject
	AcId() *string
	SetAcId(val *string)
	AcIdInput() *string
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
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdentityProviderId() *string
	SetIdentityProviderId(val *string)
	IdentityProviderIdInput() *string
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference
type jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) AcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) AcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) IdentityProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) IdentityProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference_Override(z ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetAcId(val *string) {
	if err := j.validateSetAcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetIdentityProviderId(val *string) {
	if err := j.validateSetIdentityProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/accessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *AccessApplicationSaasAppHybridAndImplicitOptions
	SetInternalValue(val *AccessApplicationSaasAppHybridAndImplicitOptions)
	ReturnAccessTokenFromAuthorizationEndpoint() interface{}
	SetReturnAccessTokenFromAuthorizationEndpoint(val interface{})
	ReturnAccessTokenFromAuthorizationEndpointInput() interface{}
	ReturnIdTokenFromAuthorizationEndpoint() interface{}
	SetReturnIdTokenFromAuthorizationEndpoint(val interface{})
	ReturnIdTokenFromAuthorizationEndpointInput() interface{}
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
	ResetReturnAccessTokenFromAuthorizationEndpoint()
	ResetReturnIdTokenFromAuthorizationEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference
type jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) InternalValue() *AccessApplicationSaasAppHybridAndImplicitOptions {
	var returns *AccessApplicationSaasAppHybridAndImplicitOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ReturnAccessTokenFromAuthorizationEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAccessTokenFromAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ReturnAccessTokenFromAuthorizationEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAccessTokenFromAuthorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ReturnIdTokenFromAuthorizationEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnIdTokenFromAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ReturnIdTokenFromAuthorizationEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnIdTokenFromAuthorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessApplicationSaasAppHybridAndImplicitOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference_Override(a AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetInternalValue(val *AccessApplicationSaasAppHybridAndImplicitOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetReturnAccessTokenFromAuthorizationEndpoint(val interface{}) {
	if err := j.validateSetReturnAccessTokenFromAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnAccessTokenFromAuthorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetReturnIdTokenFromAuthorizationEndpoint(val interface{}) {
	if err := j.validateSetReturnIdTokenFromAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnIdTokenFromAuthorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ResetReturnAccessTokenFromAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetReturnAccessTokenFromAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ResetReturnIdTokenFromAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetReturnIdTokenFromAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessApplicationSaasAppHybridAndImplicitOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustdnslocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDnsLocationEndpointsOutputReference interface {
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
	Doh() ZeroTrustDnsLocationEndpointsDohOutputReference
	DohInput() interface{}
	Dot() ZeroTrustDnsLocationEndpointsDotOutputReference
	DotInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4() ZeroTrustDnsLocationEndpointsIpv4OutputReference
	Ipv4Input() interface{}
	Ipv6() ZeroTrustDnsLocationEndpointsIpv6OutputReference
	Ipv6Input() interface{}
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
	PutDoh(value *ZeroTrustDnsLocationEndpointsDoh)
	PutDot(value *ZeroTrustDnsLocationEndpointsDot)
	PutIpv4(value *ZeroTrustDnsLocationEndpointsIpv4)
	PutIpv6(value *ZeroTrustDnsLocationEndpointsIpv6)
	ResetDoh()
	ResetDot()
	ResetIpv4()
	ResetIpv6()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustDnsLocationEndpointsOutputReference
type jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Doh() ZeroTrustDnsLocationEndpointsDohOutputReference {
	var returns ZeroTrustDnsLocationEndpointsDohOutputReference
	_jsii_.Get(
		j,
		"doh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) DohInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dohInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Dot() ZeroTrustDnsLocationEndpointsDotOutputReference {
	var returns ZeroTrustDnsLocationEndpointsDotOutputReference
	_jsii_.Get(
		j,
		"dot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) DotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Ipv4() ZeroTrustDnsLocationEndpointsIpv4OutputReference {
	var returns ZeroTrustDnsLocationEndpointsIpv4OutputReference
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Ipv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Ipv6() ZeroTrustDnsLocationEndpointsIpv6OutputReference {
	var returns ZeroTrustDnsLocationEndpointsIpv6OutputReference
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustDnsLocationEndpointsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustDnsLocationEndpointsOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustDnsLocationEndpointsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocationEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustDnsLocationEndpointsOutputReference_Override(z ZeroTrustDnsLocationEndpointsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocationEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) PutDoh(value *ZeroTrustDnsLocationEndpointsDoh) {
	if err := z.validatePutDohParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDoh",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) PutDot(value *ZeroTrustDnsLocationEndpointsDot) {
	if err := z.validatePutDotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDot",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) PutIpv4(value *ZeroTrustDnsLocationEndpointsIpv4) {
	if err := z.validatePutIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpv4",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) PutIpv6(value *ZeroTrustDnsLocationEndpointsIpv6) {
	if err := z.validatePutIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpv6",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ResetDoh() {
	_jsii_.InvokeVoid(
		z,
		"resetDoh",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ResetDot() {
	_jsii_.InvokeVoid(
		z,
		"resetDot",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		z,
		"resetIpv4",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		z,
		"resetIpv6",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


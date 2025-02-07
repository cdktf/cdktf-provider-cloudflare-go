// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustdnslocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDnsLocationEndpointsIpv6NetworksList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ZeroTrustDnsLocationEndpointsIpv6NetworksOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustDnsLocationEndpointsIpv6NetworksList
type jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewZeroTrustDnsLocationEndpointsIpv6NetworksList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ZeroTrustDnsLocationEndpointsIpv6NetworksList {
	_init_.Initialize()

	if err := validateNewZeroTrustDnsLocationEndpointsIpv6NetworksListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocationEndpointsIpv6NetworksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewZeroTrustDnsLocationEndpointsIpv6NetworksList_Override(z ZeroTrustDnsLocationEndpointsIpv6NetworksList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocationEndpointsIpv6NetworksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := z.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		z,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) Get(index *float64) ZeroTrustDnsLocationEndpointsIpv6NetworksOutputReference {
	if err := z.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ZeroTrustDnsLocationEndpointsIpv6NetworksOutputReference

	_jsii_.Invoke(
		z,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustDnsLocationEndpointsIpv6NetworksList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


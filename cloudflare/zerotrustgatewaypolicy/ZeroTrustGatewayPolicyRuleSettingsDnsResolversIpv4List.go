// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustgatewaypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List interface {
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
	Get(index *float64) ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4OutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List
type jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4ListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List_Override(z ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) Get(index *float64) ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4OutputReference {
	if err := z.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4OutputReference

	_jsii_.Invoke(
		z,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4List) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


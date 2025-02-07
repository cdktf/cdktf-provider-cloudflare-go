// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustgatewaypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() *string
	SetIp(val *string)
	IpInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	RouteThroughPrivateNetwork() interface{}
	SetRouteThroughPrivateNetwork(val interface{})
	RouteThroughPrivateNetworkInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VnetId() *string
	SetVnetId(val *string)
	VnetIdInput() *string
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
	ResetPort()
	ResetRouteThroughPrivateNetwork()
	ResetVnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference
type jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) RouteThroughPrivateNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeThroughPrivateNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) RouteThroughPrivateNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeThroughPrivateNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) VnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) VnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetIdInput",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference_Override(z ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetIp(val *string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetRouteThroughPrivateNetwork(val interface{}) {
	if err := j.validateSetRouteThroughPrivateNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeThroughPrivateNetwork",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference)SetVnetId(val *string) {
	if err := j.validateSetVnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetId",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		z,
		"resetPort",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ResetRouteThroughPrivateNetwork() {
	_jsii_.InvokeVoid(
		z,
		"resetRouteThroughPrivateNetwork",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ResetVnetId() {
	_jsii_.InvokeVoid(
		z,
		"resetVnetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference interface {
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
	Dns() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDnsOutputReference
	DnsInput() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns
	// Experimental.
	Fqn() *string
	Http() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttpOutputReference
	HttpInput() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp
	InternalValue() *ZeroTrustGatewaySettingsLoggingSettingsByRuleType
	SetInternalValue(val *ZeroTrustGatewaySettingsLoggingSettingsByRuleType)
	L4() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4OutputReference
	L4Input() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4
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
	PutDns(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns)
	PutHttp(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp)
	PutL4(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference
type jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) Dns() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDnsOutputReference {
	var returns ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) DnsInput() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns {
	var returns *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) Http() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttpOutputReference {
	var returns ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) HttpInput() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp {
	var returns *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) InternalValue() *ZeroTrustGatewaySettingsLoggingSettingsByRuleType {
	var returns *ZeroTrustGatewaySettingsLoggingSettingsByRuleType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) L4() ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4OutputReference {
	var returns ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4OutputReference
	_jsii_.Get(
		j,
		"l4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) L4Input() *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4 {
	var returns *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4
	_jsii_.Get(
		j,
		"l4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference_Override(z ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference)SetInternalValue(val *ZeroTrustGatewaySettingsLoggingSettingsByRuleType) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) PutDns(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeDns) {
	if err := z.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDns",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) PutHttp(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeHttp) {
	if err := z.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putHttp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) PutL4(value *ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeL4) {
	if err := z.validatePutL4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putL4",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsLoggingSettingsByRuleTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


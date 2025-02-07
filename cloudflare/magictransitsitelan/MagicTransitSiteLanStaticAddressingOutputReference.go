// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/magictransitsitelan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteLanStaticAddressingOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
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
	DhcpRelay() MagicTransitSiteLanStaticAddressingDhcpRelayOutputReference
	DhcpRelayInput() interface{}
	DhcpServer() MagicTransitSiteLanStaticAddressingDhcpServerOutputReference
	DhcpServerInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecondaryAddress() *string
	SetSecondaryAddress(val *string)
	SecondaryAddressInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualAddress() *string
	SetVirtualAddress(val *string)
	VirtualAddressInput() *string
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
	PutDhcpRelay(value *MagicTransitSiteLanStaticAddressingDhcpRelay)
	PutDhcpServer(value *MagicTransitSiteLanStaticAddressingDhcpServer)
	ResetDhcpRelay()
	ResetDhcpServer()
	ResetSecondaryAddress()
	ResetVirtualAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MagicTransitSiteLanStaticAddressingOutputReference
type jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) DhcpRelay() MagicTransitSiteLanStaticAddressingDhcpRelayOutputReference {
	var returns MagicTransitSiteLanStaticAddressingDhcpRelayOutputReference
	_jsii_.Get(
		j,
		"dhcpRelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) DhcpRelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpRelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) DhcpServer() MagicTransitSiteLanStaticAddressingDhcpServerOutputReference {
	var returns MagicTransitSiteLanStaticAddressingDhcpServerOutputReference
	_jsii_.Get(
		j,
		"dhcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) DhcpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) SecondaryAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) SecondaryAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) VirtualAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) VirtualAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualAddressInput",
		&returns,
	)
	return returns
}


func NewMagicTransitSiteLanStaticAddressingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MagicTransitSiteLanStaticAddressingOutputReference {
	_init_.Initialize()

	if err := validateNewMagicTransitSiteLanStaticAddressingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanStaticAddressingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMagicTransitSiteLanStaticAddressingOutputReference_Override(m MagicTransitSiteLanStaticAddressingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanStaticAddressingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetSecondaryAddress(val *string) {
	if err := j.validateSetSecondaryAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryAddress",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference)SetVirtualAddress(val *string) {
	if err := j.validateSetVirtualAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualAddress",
		val,
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) PutDhcpRelay(value *MagicTransitSiteLanStaticAddressingDhcpRelay) {
	if err := m.validatePutDhcpRelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDhcpRelay",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) PutDhcpServer(value *MagicTransitSiteLanStaticAddressingDhcpServer) {
	if err := m.validatePutDhcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDhcpServer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ResetDhcpRelay() {
	_jsii_.InvokeVoid(
		m,
		"resetDhcpRelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ResetDhcpServer() {
	_jsii_.InvokeVoid(
		m,
		"resetDhcpServer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ResetSecondaryAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSecondaryAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ResetVirtualAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetVirtualAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


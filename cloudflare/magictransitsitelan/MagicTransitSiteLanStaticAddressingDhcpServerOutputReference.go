// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/magictransitsitelan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteLanStaticAddressingDhcpServerOutputReference interface {
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
	DhcpPoolEnd() *string
	SetDhcpPoolEnd(val *string)
	DhcpPoolEndInput() *string
	DhcpPoolStart() *string
	SetDhcpPoolStart(val *string)
	DhcpPoolStartInput() *string
	DnsServer() *string
	SetDnsServer(val *string)
	DnsServerInput() *string
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Reservations() *map[string]*string
	SetReservations(val *map[string]*string)
	ReservationsInput() *map[string]*string
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
	ResetDhcpPoolEnd()
	ResetDhcpPoolStart()
	ResetDnsServer()
	ResetDnsServers()
	ResetReservations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MagicTransitSiteLanStaticAddressingDhcpServerOutputReference
type jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DhcpPoolEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpPoolEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DhcpPoolEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpPoolEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DhcpPoolStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpPoolStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DhcpPoolStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpPoolStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DnsServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DnsServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) Reservations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"reservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ReservationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"reservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMagicTransitSiteLanStaticAddressingDhcpServerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MagicTransitSiteLanStaticAddressingDhcpServerOutputReference {
	_init_.Initialize()

	if err := validateNewMagicTransitSiteLanStaticAddressingDhcpServerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanStaticAddressingDhcpServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMagicTransitSiteLanStaticAddressingDhcpServerOutputReference_Override(m MagicTransitSiteLanStaticAddressingDhcpServerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanStaticAddressingDhcpServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetDhcpPoolEnd(val *string) {
	if err := j.validateSetDhcpPoolEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpPoolEnd",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetDhcpPoolStart(val *string) {
	if err := j.validateSetDhcpPoolStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpPoolStart",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetDnsServer(val *string) {
	if err := j.validateSetDnsServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServer",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetReservations(val *map[string]*string) {
	if err := j.validateSetReservationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservations",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ResetDhcpPoolEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetDhcpPoolEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ResetDhcpPoolStart() {
	_jsii_.InvokeVoid(
		m,
		"resetDhcpPoolStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ResetDnsServer() {
	_jsii_.InvokeVoid(
		m,
		"resetDnsServer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ResetDnsServers() {
	_jsii_.InvokeVoid(
		m,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ResetReservations() {
	_jsii_.InvokeVoid(
		m,
		"resetReservations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MagicTransitSiteLanStaticAddressingDhcpServerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


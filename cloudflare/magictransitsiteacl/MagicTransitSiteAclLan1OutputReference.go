// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsiteacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/magictransitsiteacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteAclLan1OutputReference interface {
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
	LanId() *string
	SetLanId(val *string)
	LanIdInput() *string
	LanName() *string
	SetLanName(val *string)
	LanNameInput() *string
	PortRanges() *[]*string
	SetPortRanges(val *[]*string)
	PortRangesInput() *[]*string
	Ports() *[]*float64
	SetPorts(val *[]*float64)
	PortsInput() *[]*float64
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
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
	ResetLanName()
	ResetPortRanges()
	ResetPorts()
	ResetSubnets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MagicTransitSiteAclLan1OutputReference
type jsiiProxy_MagicTransitSiteAclLan1OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) LanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) LanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) LanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) LanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) PortRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) PortRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) Ports() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) PortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMagicTransitSiteAclLan1OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MagicTransitSiteAclLan1OutputReference {
	_init_.Initialize()

	if err := validateNewMagicTransitSiteAclLan1OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicTransitSiteAclLan1OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteAcl.MagicTransitSiteAclLan1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMagicTransitSiteAclLan1OutputReference_Override(m MagicTransitSiteAclLan1OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteAcl.MagicTransitSiteAclLan1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetLanId(val *string) {
	if err := j.validateSetLanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanId",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetLanName(val *string) {
	if err := j.validateSetLanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanName",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetPortRanges(val *[]*string) {
	if err := j.validateSetPortRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRanges",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetPorts(val *[]*float64) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteAclLan1OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ResetLanName() {
	_jsii_.InvokeVoid(
		m,
		"resetLanName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ResetPortRanges() {
	_jsii_.InvokeVoid(
		m,
		"resetPortRanges",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		m,
		"resetPorts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ResetSubnets() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnets",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MagicTransitSiteAclLan1OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


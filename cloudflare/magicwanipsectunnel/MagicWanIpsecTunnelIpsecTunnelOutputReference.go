// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanipsectunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/magicwanipsectunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicWanIpsecTunnelIpsecTunnelOutputReference interface {
	cdktf.ComplexObject
	AllowNullCipher() cdktf.IResolvable
	CloudflareEndpoint() *string
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerEndpoint() *string
	Description() *string
	// Experimental.
	Fqn() *string
	HealthCheck() MagicWanIpsecTunnelIpsecTunnelHealthCheckOutputReference
	Id() *string
	InterfaceAddress() *string
	InternalValue() *MagicWanIpsecTunnelIpsecTunnel
	SetInternalValue(val *MagicWanIpsecTunnelIpsecTunnel)
	ModifiedOn() *string
	Name() *string
	PskMetadata() MagicWanIpsecTunnelIpsecTunnelPskMetadataOutputReference
	ReplayProtection() cdktf.IResolvable
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

// The jsii proxy struct for MagicWanIpsecTunnelIpsecTunnelOutputReference
type jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) AllowNullCipher() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowNullCipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) CloudflareEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) CustomerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) HealthCheck() MagicWanIpsecTunnelIpsecTunnelHealthCheckOutputReference {
	var returns MagicWanIpsecTunnelIpsecTunnelHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) InterfaceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) InternalValue() *MagicWanIpsecTunnelIpsecTunnel {
	var returns *MagicWanIpsecTunnelIpsecTunnel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) PskMetadata() MagicWanIpsecTunnelIpsecTunnelPskMetadataOutputReference {
	var returns MagicWanIpsecTunnelIpsecTunnelPskMetadataOutputReference
	_jsii_.Get(
		j,
		"pskMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ReplayProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"replayProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMagicWanIpsecTunnelIpsecTunnelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MagicWanIpsecTunnelIpsecTunnelOutputReference {
	_init_.Initialize()

	if err := validateNewMagicWanIpsecTunnelIpsecTunnelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnelIpsecTunnelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMagicWanIpsecTunnelIpsecTunnelOutputReference_Override(m MagicWanIpsecTunnelIpsecTunnelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnelIpsecTunnelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference)SetInternalValue(val *MagicWanIpsecTunnelIpsecTunnel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MagicWanIpsecTunnelIpsecTunnelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


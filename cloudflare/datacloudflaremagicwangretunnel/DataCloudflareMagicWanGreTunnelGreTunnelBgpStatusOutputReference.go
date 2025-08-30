// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremagicwangretunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflaremagicwangretunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference interface {
	cdktf.ComplexObject
	BgpState() *string
	CfSpeakerIp() *string
	CfSpeakerPort() *float64
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
	CustomerSpeakerIp() *string
	CustomerSpeakerPort() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareMagicWanGreTunnelGreTunnelBgpStatus
	SetInternalValue(val *DataCloudflareMagicWanGreTunnelGreTunnelBgpStatus)
	State() *string
	TcpEstablished() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
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

// The jsii proxy struct for DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference
type jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) BgpState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) CfSpeakerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfSpeakerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) CfSpeakerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cfSpeakerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) CustomerSpeakerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerSpeakerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) CustomerSpeakerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerSpeakerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) InternalValue() *DataCloudflareMagicWanGreTunnelGreTunnelBgpStatus {
	var returns *DataCloudflareMagicWanGreTunnelGreTunnelBgpStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) TcpEstablished() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tcpEstablished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareMagicWanGreTunnel.DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference_Override(d DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareMagicWanGreTunnel.DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference)SetInternalValue(val *DataCloudflareMagicWanGreTunnelGreTunnelBgpStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareMagicWanGreTunnelGreTunnelBgpStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


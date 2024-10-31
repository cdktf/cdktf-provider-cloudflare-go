// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustinfrastructureaccesstargets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/datacloudflarezerotrustinfrastructureaccesstargets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference interface {
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
	InternalValue() *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp
	SetInternalValue(val *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp)
	Ipv4() DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4OutputReference
	Ipv4Input() interface{}
	Ipv6() DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6OutputReference
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
	PutIpv4(value *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4)
	PutIpv6(value *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6)
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

// The jsii proxy struct for DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference
type jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) InternalValue() *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp {
	var returns *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Ipv4() DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4OutputReference {
	var returns DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4OutputReference
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Ipv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Ipv6() DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6OutputReference {
	var returns DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6OutputReference
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Ipv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustInfrastructureAccessTargets.DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference_Override(d DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustInfrastructureAccessTargets.DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference)SetInternalValue(val *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) PutIpv4(value *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4) {
	if err := d.validatePutIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIpv4",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) PutIpv6(value *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6) {
	if err := d.validatePutIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIpv6",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


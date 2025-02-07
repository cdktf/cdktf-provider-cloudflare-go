// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarednsrecords/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareDnsRecordsCommentOutputReference interface {
	cdktf.ComplexObject
	Absent() *string
	SetAbsent(val *string)
	AbsentInput() *string
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
	Contains() *string
	SetContains(val *string)
	ContainsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Endswith() *string
	SetEndswith(val *string)
	EndswithInput() *string
	Exact() *string
	SetExact(val *string)
	ExactInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Present() *string
	SetPresent(val *string)
	PresentInput() *string
	Startswith() *string
	SetStartswith(val *string)
	StartswithInput() *string
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
	ResetAbsent()
	ResetContains()
	ResetEndswith()
	ResetExact()
	ResetPresent()
	ResetStartswith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareDnsRecordsCommentOutputReference
type jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Absent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) AbsentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Contains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ContainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Endswith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endswith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) EndswithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endswithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Exact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Present() *string {
	var returns *string
	_jsii_.Get(
		j,
		"present",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) PresentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Startswith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startswith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) StartswithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startswithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareDnsRecordsCommentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareDnsRecordsCommentOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareDnsRecordsCommentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecords.DataCloudflareDnsRecordsCommentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareDnsRecordsCommentOutputReference_Override(d DataCloudflareDnsRecordsCommentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecords.DataCloudflareDnsRecordsCommentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetAbsent(val *string) {
	if err := j.validateSetAbsentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"absent",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetContains(val *string) {
	if err := j.validateSetContainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contains",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetEndswith(val *string) {
	if err := j.validateSetEndswithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endswith",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetExact(val *string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetPresent(val *string) {
	if err := j.validateSetPresentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"present",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetStartswith(val *string) {
	if err := j.validateSetStartswithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startswith",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetAbsent() {
	_jsii_.InvokeVoid(
		d,
		"resetAbsent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetContains() {
	_jsii_.InvokeVoid(
		d,
		"resetContains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetEndswith() {
	_jsii_.InvokeVoid(
		d,
		"resetEndswith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetExact() {
	_jsii_.InvokeVoid(
		d,
		"resetExact",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetPresent() {
	_jsii_.InvokeVoid(
		d,
		"resetPresent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ResetStartswith() {
	_jsii_.InvokeVoid(
		d,
		"resetStartswith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsCommentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


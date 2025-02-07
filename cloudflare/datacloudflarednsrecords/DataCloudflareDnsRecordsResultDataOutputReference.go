// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarednsrecords/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareDnsRecordsResultDataOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *float64
	Altitude() *float64
	Certificate() *string
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
	Digest() *string
	DigestType() *float64
	Fingerprint() *string
	Flags() cdktf.AnyMap
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareDnsRecordsResultData
	SetInternalValue(val *DataCloudflareDnsRecordsResultData)
	KeyTag() *float64
	LatDegrees() *float64
	LatDirection() *string
	LatMinutes() *float64
	LatSeconds() *float64
	LongDegrees() *float64
	LongDirection() *string
	LongMinutes() *float64
	LongSeconds() *float64
	MatchingType() *float64
	Order() *float64
	Port() *float64
	PrecisionHorz() *float64
	PrecisionVert() *float64
	Preference() *float64
	Priority() *float64
	Protocol() *float64
	PublicKey() *string
	Regex() *string
	Replacement() *string
	Selector() *float64
	Service() *string
	Size() *float64
	Tag() *string
	Target() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *float64
	Usage() *float64
	Value() *string
	Weight() *float64
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

// The jsii proxy struct for DataCloudflareDnsRecordsResultDataOutputReference
type jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Algorithm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Altitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"altitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Digest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) DigestType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Flags() cdktf.AnyMap {
	var returns cdktf.AnyMap
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) InternalValue() *DataCloudflareDnsRecordsResultData {
	var returns *DataCloudflareDnsRecordsResultData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) KeyTag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LatDegrees() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latDegrees",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LatDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LatMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LatSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LongDegrees() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longDegrees",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LongDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LongMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) LongSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) MatchingType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) PrecisionHorz() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionHorz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) PrecisionVert() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionVert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Preference() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Protocol() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Replacement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Selector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Usage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


func NewDataCloudflareDnsRecordsResultDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareDnsRecordsResultDataOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareDnsRecordsResultDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecords.DataCloudflareDnsRecordsResultDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareDnsRecordsResultDataOutputReference_Override(d DataCloudflareDnsRecordsResultDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecords.DataCloudflareDnsRecordsResultDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference)SetInternalValue(val *DataCloudflareDnsRecordsResultData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordsResultDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


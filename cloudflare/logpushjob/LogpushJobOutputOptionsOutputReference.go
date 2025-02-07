// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logpushjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/logpushjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogpushJobOutputOptionsOutputReference interface {
	cdktf.ComplexObject
	BatchPrefix() *string
	SetBatchPrefix(val *string)
	BatchPrefixInput() *string
	BatchSuffix() *string
	SetBatchSuffix(val *string)
	BatchSuffixInput() *string
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
	Cve20214428() interface{}
	SetCve20214428(val interface{})
	Cve20214428Input() interface{}
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	FieldNames() *[]*string
	SetFieldNames(val *[]*string)
	FieldNamesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
	RecordDelimiter() *string
	SetRecordDelimiter(val *string)
	RecordDelimiterInput() *string
	RecordPrefix() *string
	SetRecordPrefix(val *string)
	RecordPrefixInput() *string
	RecordSuffix() *string
	SetRecordSuffix(val *string)
	RecordSuffixInput() *string
	RecordTemplate() *string
	SetRecordTemplate(val *string)
	RecordTemplateInput() *string
	SampleRate() *float64
	SetSampleRate(val *float64)
	SampleRateInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampFormat() *string
	SetTimestampFormat(val *string)
	TimestampFormatInput() *string
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
	ResetBatchPrefix()
	ResetBatchSuffix()
	ResetCve20214428()
	ResetFieldDelimiter()
	ResetFieldNames()
	ResetOutputType()
	ResetRecordDelimiter()
	ResetRecordPrefix()
	ResetRecordSuffix()
	ResetRecordTemplate()
	ResetSampleRate()
	ResetTimestampFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogpushJobOutputOptionsOutputReference
type jsiiProxy_LogpushJobOutputOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) BatchPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) BatchPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) BatchSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) BatchSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) Cve20214428() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cve20214428",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) Cve20214428Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cve20214428Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) FieldNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fieldNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) FieldNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fieldNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) RecordTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) SampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) SampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}


func NewLogpushJobOutputOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogpushJobOutputOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewLogpushJobOutputOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogpushJobOutputOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJobOutputOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogpushJobOutputOptionsOutputReference_Override(l LogpushJobOutputOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJobOutputOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetBatchPrefix(val *string) {
	if err := j.validateSetBatchPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchPrefix",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetBatchSuffix(val *string) {
	if err := j.validateSetBatchSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSuffix",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetCve20214428(val interface{}) {
	if err := j.validateSetCve20214428Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cve20214428",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetFieldNames(val *[]*string) {
	if err := j.validateSetFieldNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldNames",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetRecordPrefix(val *string) {
	if err := j.validateSetRecordPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordPrefix",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetRecordSuffix(val *string) {
	if err := j.validateSetRecordSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSuffix",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetRecordTemplate(val *string) {
	if err := j.validateSetRecordTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordTemplate",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetSampleRate(val *float64) {
	if err := j.validateSetSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogpushJobOutputOptionsOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetBatchPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetBatchPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetBatchSuffix() {
	_jsii_.InvokeVoid(
		l,
		"resetBatchSuffix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetCve20214428() {
	_jsii_.InvokeVoid(
		l,
		"resetCve20214428",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		l,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetFieldNames() {
	_jsii_.InvokeVoid(
		l,
		"resetFieldNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		l,
		"resetOutputType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		l,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetRecordPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetRecordPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetRecordSuffix() {
	_jsii_.InvokeVoid(
		l,
		"resetRecordSuffix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetRecordTemplate() {
	_jsii_.InvokeVoid(
		l,
		"resetRecordTemplate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleRate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJobOutputOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


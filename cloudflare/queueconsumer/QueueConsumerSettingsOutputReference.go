// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package queueconsumer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/queueconsumer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QueueConsumerSettingsOutputReference interface {
	cdktf.ComplexObject
	BatchSize() *float64
	SetBatchSize(val *float64)
	BatchSizeInput() *float64
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
	MaxConcurrency() *float64
	SetMaxConcurrency(val *float64)
	MaxConcurrencyInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MaxWaitTimeMs() *float64
	SetMaxWaitTimeMs(val *float64)
	MaxWaitTimeMsInput() *float64
	RetryDelay() *float64
	SetRetryDelay(val *float64)
	RetryDelayInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VisibilityTimeoutMs() *float64
	SetVisibilityTimeoutMs(val *float64)
	VisibilityTimeoutMsInput() *float64
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
	ResetBatchSize()
	ResetMaxConcurrency()
	ResetMaxRetries()
	ResetMaxWaitTimeMs()
	ResetRetryDelay()
	ResetVisibilityTimeoutMs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueueConsumerSettingsOutputReference
type jsiiProxy_QueueConsumerSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxWaitTimeMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitTimeMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) MaxWaitTimeMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitTimeMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) RetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) RetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) VisibilityTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference) VisibilityTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutMsInput",
		&returns,
	)
	return returns
}


func NewQueueConsumerSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QueueConsumerSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewQueueConsumerSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueueConsumerSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.queueConsumer.QueueConsumerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQueueConsumerSettingsOutputReference_Override(q QueueConsumerSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.queueConsumer.QueueConsumerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetMaxConcurrency(val *float64) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetMaxWaitTimeMs(val *float64) {
	if err := j.validateSetMaxWaitTimeMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWaitTimeMs",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetRetryDelay(val *float64) {
	if err := j.validateSetRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDelay",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QueueConsumerSettingsOutputReference)SetVisibilityTimeoutMs(val *float64) {
	if err := j.validateSetVisibilityTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibilityTimeoutMs",
		val,
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		q,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		q,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		q,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetMaxWaitTimeMs() {
	_jsii_.InvokeVoid(
		q,
		"resetMaxWaitTimeMs",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetRetryDelay() {
	_jsii_.InvokeVoid(
		q,
		"resetRetryDelay",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ResetVisibilityTimeoutMs() {
	_jsii_.InvokeVoid(
		q,
		"resetVisibilityTimeoutMs",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueConsumerSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


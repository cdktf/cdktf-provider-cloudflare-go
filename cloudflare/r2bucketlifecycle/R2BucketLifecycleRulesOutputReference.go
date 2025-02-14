// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketlifecycle

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/r2bucketlifecycle/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type R2BucketLifecycleRulesOutputReference interface {
	cdktf.ComplexObject
	AbortMultipartUploadsTransition() R2BucketLifecycleRulesAbortMultipartUploadsTransitionOutputReference
	AbortMultipartUploadsTransitionInput() interface{}
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
	Conditions() R2BucketLifecycleRulesConditionsOutputReference
	ConditionsInput() *R2BucketLifecycleRulesConditions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteObjectsTransition() R2BucketLifecycleRulesDeleteObjectsTransitionOutputReference
	DeleteObjectsTransitionInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageClassTransitions() R2BucketLifecycleRulesStorageClassTransitionsList
	StorageClassTransitionsInput() interface{}
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
	PutAbortMultipartUploadsTransition(value *R2BucketLifecycleRulesAbortMultipartUploadsTransition)
	PutConditions(value *R2BucketLifecycleRulesConditions)
	PutDeleteObjectsTransition(value *R2BucketLifecycleRulesDeleteObjectsTransition)
	PutStorageClassTransitions(value interface{})
	ResetAbortMultipartUploadsTransition()
	ResetDeleteObjectsTransition()
	ResetStorageClassTransitions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for R2BucketLifecycleRulesOutputReference
type jsiiProxy_R2BucketLifecycleRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) AbortMultipartUploadsTransition() R2BucketLifecycleRulesAbortMultipartUploadsTransitionOutputReference {
	var returns R2BucketLifecycleRulesAbortMultipartUploadsTransitionOutputReference
	_jsii_.Get(
		j,
		"abortMultipartUploadsTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) AbortMultipartUploadsTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortMultipartUploadsTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) Conditions() R2BucketLifecycleRulesConditionsOutputReference {
	var returns R2BucketLifecycleRulesConditionsOutputReference
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) ConditionsInput() *R2BucketLifecycleRulesConditions {
	var returns *R2BucketLifecycleRulesConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) DeleteObjectsTransition() R2BucketLifecycleRulesDeleteObjectsTransitionOutputReference {
	var returns R2BucketLifecycleRulesDeleteObjectsTransitionOutputReference
	_jsii_.Get(
		j,
		"deleteObjectsTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) DeleteObjectsTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteObjectsTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) StorageClassTransitions() R2BucketLifecycleRulesStorageClassTransitionsList {
	var returns R2BucketLifecycleRulesStorageClassTransitionsList
	_jsii_.Get(
		j,
		"storageClassTransitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) StorageClassTransitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageClassTransitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewR2BucketLifecycleRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) R2BucketLifecycleRulesOutputReference {
	_init_.Initialize()

	if err := validateNewR2BucketLifecycleRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_R2BucketLifecycleRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.r2BucketLifecycle.R2BucketLifecycleRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewR2BucketLifecycleRulesOutputReference_Override(r R2BucketLifecycleRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.r2BucketLifecycle.R2BucketLifecycleRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_R2BucketLifecycleRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) PutAbortMultipartUploadsTransition(value *R2BucketLifecycleRulesAbortMultipartUploadsTransition) {
	if err := r.validatePutAbortMultipartUploadsTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAbortMultipartUploadsTransition",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) PutConditions(value *R2BucketLifecycleRulesConditions) {
	if err := r.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putConditions",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) PutDeleteObjectsTransition(value *R2BucketLifecycleRulesDeleteObjectsTransition) {
	if err := r.validatePutDeleteObjectsTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDeleteObjectsTransition",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) PutStorageClassTransitions(value interface{}) {
	if err := r.validatePutStorageClassTransitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putStorageClassTransitions",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) ResetAbortMultipartUploadsTransition() {
	_jsii_.InvokeVoid(
		r,
		"resetAbortMultipartUploadsTransition",
		nil, // no parameters
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) ResetDeleteObjectsTransition() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteObjectsTransition",
		nil, // no parameters
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) ResetStorageClassTransitions() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageClassTransitions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_R2BucketLifecycleRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


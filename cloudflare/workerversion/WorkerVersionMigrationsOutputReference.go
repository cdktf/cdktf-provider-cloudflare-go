// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/workerversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerVersionMigrationsOutputReference interface {
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
	DeletedClasses() *[]*string
	SetDeletedClasses(val *[]*string)
	DeletedClassesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NewClasses() *[]*string
	SetNewClasses(val *[]*string)
	NewClassesInput() *[]*string
	NewSqliteClasses() *[]*string
	SetNewSqliteClasses(val *[]*string)
	NewSqliteClassesInput() *[]*string
	NewTag() *string
	SetNewTag(val *string)
	NewTagInput() *string
	OldTag() *string
	SetOldTag(val *string)
	OldTagInput() *string
	RenamedClasses() WorkerVersionMigrationsRenamedClassesList
	RenamedClassesInput() interface{}
	Steps() WorkerVersionMigrationsStepsList
	StepsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferredClasses() WorkerVersionMigrationsTransferredClassesList
	TransferredClassesInput() interface{}
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
	PutRenamedClasses(value interface{})
	PutSteps(value interface{})
	PutTransferredClasses(value interface{})
	ResetDeletedClasses()
	ResetNewClasses()
	ResetNewSqliteClasses()
	ResetNewTag()
	ResetOldTag()
	ResetRenamedClasses()
	ResetSteps()
	ResetTransferredClasses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkerVersionMigrationsOutputReference
type jsiiProxy_WorkerVersionMigrationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) DeletedClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) DeletedClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewSqliteClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewSqliteClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) NewTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) OldTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) OldTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) RenamedClasses() WorkerVersionMigrationsRenamedClassesList {
	var returns WorkerVersionMigrationsRenamedClassesList
	_jsii_.Get(
		j,
		"renamedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) RenamedClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renamedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) Steps() WorkerVersionMigrationsStepsList {
	var returns WorkerVersionMigrationsStepsList
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) StepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) TransferredClasses() WorkerVersionMigrationsTransferredClassesList {
	var returns WorkerVersionMigrationsTransferredClassesList
	_jsii_.Get(
		j,
		"transferredClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference) TransferredClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferredClassesInput",
		&returns,
	)
	return returns
}


func NewWorkerVersionMigrationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkerVersionMigrationsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkerVersionMigrationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkerVersionMigrationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerVersion.WorkerVersionMigrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkerVersionMigrationsOutputReference_Override(w WorkerVersionMigrationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerVersion.WorkerVersionMigrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetDeletedClasses(val *[]*string) {
	if err := j.validateSetDeletedClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedClasses",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetNewClasses(val *[]*string) {
	if err := j.validateSetNewClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newClasses",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetNewSqliteClasses(val *[]*string) {
	if err := j.validateSetNewSqliteClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newSqliteClasses",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetNewTag(val *string) {
	if err := j.validateSetNewTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newTag",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetOldTag(val *string) {
	if err := j.validateSetOldTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oldTag",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionMigrationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) PutRenamedClasses(value interface{}) {
	if err := w.validatePutRenamedClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRenamedClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) PutSteps(value interface{}) {
	if err := w.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSteps",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) PutTransferredClasses(value interface{}) {
	if err := w.validatePutTransferredClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTransferredClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetDeletedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetDeletedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetNewClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetNewSqliteClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewSqliteClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetNewTag() {
	_jsii_.InvokeVoid(
		w,
		"resetNewTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetOldTag() {
	_jsii_.InvokeVoid(
		w,
		"resetOldTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetRenamedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetRenamedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetSteps() {
	_jsii_.InvokeVoid(
		w,
		"resetSteps",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ResetTransferredClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetTransferredClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionMigrationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


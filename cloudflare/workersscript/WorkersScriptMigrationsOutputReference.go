// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/workersscript/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersScriptMigrationsOutputReference interface {
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
	RenamedClasses() WorkersScriptMigrationsRenamedClassesList
	RenamedClassesInput() interface{}
	Steps() WorkersScriptMigrationsStepsList
	StepsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferredClasses() WorkersScriptMigrationsTransferredClassesList
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

// The jsii proxy struct for WorkersScriptMigrationsOutputReference
type jsiiProxy_WorkersScriptMigrationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) DeletedClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) DeletedClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewSqliteClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewSqliteClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) NewTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) OldTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) OldTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) RenamedClasses() WorkersScriptMigrationsRenamedClassesList {
	var returns WorkersScriptMigrationsRenamedClassesList
	_jsii_.Get(
		j,
		"renamedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) RenamedClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renamedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) Steps() WorkersScriptMigrationsStepsList {
	var returns WorkersScriptMigrationsStepsList
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) StepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) TransferredClasses() WorkersScriptMigrationsTransferredClassesList {
	var returns WorkersScriptMigrationsTransferredClassesList
	_jsii_.Get(
		j,
		"transferredClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference) TransferredClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferredClassesInput",
		&returns,
	)
	return returns
}


func NewWorkersScriptMigrationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkersScriptMigrationsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkersScriptMigrationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScriptMigrationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptMigrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkersScriptMigrationsOutputReference_Override(w WorkersScriptMigrationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptMigrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetDeletedClasses(val *[]*string) {
	if err := j.validateSetDeletedClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetNewClasses(val *[]*string) {
	if err := j.validateSetNewClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetNewSqliteClasses(val *[]*string) {
	if err := j.validateSetNewSqliteClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newSqliteClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetNewTag(val *string) {
	if err := j.validateSetNewTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newTag",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetOldTag(val *string) {
	if err := j.validateSetOldTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oldTag",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) PutRenamedClasses(value interface{}) {
	if err := w.validatePutRenamedClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRenamedClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) PutSteps(value interface{}) {
	if err := w.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSteps",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) PutTransferredClasses(value interface{}) {
	if err := w.validatePutTransferredClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTransferredClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetDeletedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetDeletedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetNewClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetNewSqliteClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewSqliteClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetNewTag() {
	_jsii_.InvokeVoid(
		w,
		"resetNewTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetOldTag() {
	_jsii_.InvokeVoid(
		w,
		"resetOldTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetRenamedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetRenamedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetSteps() {
	_jsii_.InvokeVoid(
		w,
		"resetSteps",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ResetTransferredClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetTransferredClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkersScriptMigrationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


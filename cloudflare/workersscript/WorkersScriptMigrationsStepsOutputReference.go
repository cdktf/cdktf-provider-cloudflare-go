// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/workersscript/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersScriptMigrationsStepsOutputReference interface {
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
	RenamedClasses() WorkersScriptMigrationsStepsRenamedClassesList
	RenamedClassesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferredClasses() WorkersScriptMigrationsStepsTransferredClassesList
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
	PutTransferredClasses(value interface{})
	ResetDeletedClasses()
	ResetNewClasses()
	ResetNewSqliteClasses()
	ResetRenamedClasses()
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

// The jsii proxy struct for WorkersScriptMigrationsStepsOutputReference
type jsiiProxy_WorkersScriptMigrationsStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) DeletedClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) DeletedClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) NewClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) NewClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) NewSqliteClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) NewSqliteClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newSqliteClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) RenamedClasses() WorkersScriptMigrationsStepsRenamedClassesList {
	var returns WorkersScriptMigrationsStepsRenamedClassesList
	_jsii_.Get(
		j,
		"renamedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) RenamedClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renamedClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) TransferredClasses() WorkersScriptMigrationsStepsTransferredClassesList {
	var returns WorkersScriptMigrationsStepsTransferredClassesList
	_jsii_.Get(
		j,
		"transferredClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) TransferredClassesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferredClassesInput",
		&returns,
	)
	return returns
}


func NewWorkersScriptMigrationsStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkersScriptMigrationsStepsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkersScriptMigrationsStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScriptMigrationsStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptMigrationsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkersScriptMigrationsStepsOutputReference_Override(w WorkersScriptMigrationsStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptMigrationsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetDeletedClasses(val *[]*string) {
	if err := j.validateSetDeletedClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetNewClasses(val *[]*string) {
	if err := j.validateSetNewClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetNewSqliteClasses(val *[]*string) {
	if err := j.validateSetNewSqliteClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newSqliteClasses",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) PutRenamedClasses(value interface{}) {
	if err := w.validatePutRenamedClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRenamedClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) PutTransferredClasses(value interface{}) {
	if err := w.validatePutTransferredClassesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTransferredClasses",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ResetDeletedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetDeletedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ResetNewClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ResetNewSqliteClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetNewSqliteClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ResetRenamedClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetRenamedClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ResetTransferredClasses() {
	_jsii_.InvokeVoid(
		w,
		"resetTransferredClasses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkersScriptMigrationsStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


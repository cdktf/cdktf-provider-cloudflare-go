// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecloudforceonerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarecloudforceonerequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestFilterOutputReference interface {
	cdktf.ComplexObject
	CompletedAfter() *string
	SetCompletedAfter(val *string)
	CompletedAfterInput() *string
	CompletedBefore() *string
	SetCompletedBefore(val *string)
	CompletedBeforeInput() *string
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
	CreatedAfter() *string
	SetCreatedAfter(val *string)
	CreatedAfterInput() *string
	CreatedBefore() *string
	SetCreatedBefore(val *string)
	CreatedBeforeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Page() *float64
	SetPage(val *float64)
	PageInput() *float64
	PerPage() *float64
	SetPerPage(val *float64)
	PerPageInput() *float64
	RequestType() *string
	SetRequestType(val *string)
	RequestTypeInput() *string
	SortBy() *string
	SetSortBy(val *string)
	SortByInput() *string
	SortOrder() *string
	SetSortOrder(val *string)
	SortOrderInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	ResetCompletedAfter()
	ResetCompletedBefore()
	ResetCreatedAfter()
	ResetCreatedBefore()
	ResetRequestType()
	ResetSortBy()
	ResetSortOrder()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareCloudforceOneRequestFilterOutputReference
type jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CompletedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CompletedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CompletedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CompletedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) Page() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"page",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) PageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) PerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) PerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) RequestType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) RequestTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) SortBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) SortByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) SortOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) SortOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareCloudforceOneRequestFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareCloudforceOneRequestFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareCloudforceOneRequestFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequest.DataCloudflareCloudforceOneRequestFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareCloudforceOneRequestFilterOutputReference_Override(d DataCloudflareCloudforceOneRequestFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequest.DataCloudflareCloudforceOneRequestFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetCompletedAfter(val *string) {
	if err := j.validateSetCompletedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completedAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetCompletedBefore(val *string) {
	if err := j.validateSetCompletedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completedBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetCreatedAfter(val *string) {
	if err := j.validateSetCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetPage(val *float64) {
	if err := j.validateSetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"page",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetPerPage(val *float64) {
	if err := j.validateSetPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perPage",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetRequestType(val *string) {
	if err := j.validateSetRequestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestType",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetSortBy(val *string) {
	if err := j.validateSetSortByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortBy",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetSortOrder(val *string) {
	if err := j.validateSetSortOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortOrder",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetCompletedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCompletedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetCompletedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCompletedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetCreatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetRequestType() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetSortBy() {
	_jsii_.InvokeVoid(
		d,
		"resetSortBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetSortOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


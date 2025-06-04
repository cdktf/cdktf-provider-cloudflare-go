// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezonelockdown

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezonelockdown/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZoneLockdownFilterOutputReference interface {
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
	CreatedOn() *string
	SetCreatedOn(val *string)
	CreatedOnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DescriptionSearch() *string
	SetDescriptionSearch(val *string)
	DescriptionSearchInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() *string
	SetIp(val *string)
	IpInput() *string
	IpRangeSearch() *string
	SetIpRangeSearch(val *string)
	IpRangeSearchInput() *string
	IpSearch() *string
	SetIpSearch(val *string)
	IpSearchInput() *string
	ModifiedOn() *string
	SetModifiedOn(val *string)
	ModifiedOnInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriSearch() *string
	SetUriSearch(val *string)
	UriSearchInput() *string
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
	ResetCreatedOn()
	ResetDescription()
	ResetDescriptionSearch()
	ResetIp()
	ResetIpRangeSearch()
	ResetIpSearch()
	ResetModifiedOn()
	ResetPriority()
	ResetUriSearch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZoneLockdownFilterOutputReference
type jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) CreatedOnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) DescriptionSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) DescriptionSearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) IpRangeSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangeSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) IpRangeSearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangeSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) IpSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) IpSearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ModifiedOnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) UriSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) UriSearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriSearchInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareZoneLockdownFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZoneLockdownFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZoneLockdownFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZoneLockdown.DataCloudflareZoneLockdownFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZoneLockdownFilterOutputReference_Override(d DataCloudflareZoneLockdownFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZoneLockdown.DataCloudflareZoneLockdownFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetCreatedOn(val *string) {
	if err := j.validateSetCreatedOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetDescriptionSearch(val *string) {
	if err := j.validateSetDescriptionSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptionSearch",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetIp(val *string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetIpRangeSearch(val *string) {
	if err := j.validateSetIpRangeSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRangeSearch",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetIpSearch(val *string) {
	if err := j.validateSetIpSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSearch",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetModifiedOn(val *string) {
	if err := j.validateSetModifiedOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference)SetUriSearch(val *string) {
	if err := j.validateSetUriSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriSearch",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetCreatedOn() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetDescriptionSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetDescriptionSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		d,
		"resetIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetIpRangeSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetIpRangeSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetIpSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetIpSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetModifiedOn() {
	_jsii_.InvokeVoid(
		d,
		"resetModifiedOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		d,
		"resetPriority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ResetUriSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetUriSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZoneLockdownFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecuritytrusteddomains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflareemailsecuritytrusteddomains/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference interface {
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
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsRecent() interface{}
	SetIsRecent(val interface{})
	IsRecentInput() interface{}
	IsSimilarity() interface{}
	SetIsSimilarity(val interface{})
	IsSimilarityInput() interface{}
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
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
	ResetDirection()
	ResetIsRecent()
	ResetIsSimilarity()
	ResetOrder()
	ResetSearch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference
type jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) IsRecent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) IsRecentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) IsSimilarity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSimilarity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) IsSimilarityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSimilarityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareEmailSecurityTrustedDomainsFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareEmailSecurityTrustedDomainsFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareEmailSecurityTrustedDomains.DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareEmailSecurityTrustedDomainsFilterOutputReference_Override(d DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareEmailSecurityTrustedDomains.DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetIsRecent(val interface{}) {
	if err := j.validateSetIsRecentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRecent",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetIsSimilarity(val interface{}) {
	if err := j.validateSetIsSimilarityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSimilarity",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ResetIsRecent() {
	_jsii_.InvokeVoid(
		d,
		"resetIsRecent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ResetIsSimilarity() {
	_jsii_.InvokeVoid(
		d,
		"resetIsSimilarity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareEmailSecurityTrustedDomainsFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarednsrecord/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareDnsRecordFilterOutputReference interface {
	cdktf.ComplexObject
	Comment() DataCloudflareDnsRecordFilterCommentOutputReference
	CommentInput() interface{}
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
	Content() DataCloudflareDnsRecordFilterContentOutputReference
	ContentInput() interface{}
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
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	Name() DataCloudflareDnsRecordFilterNameOutputReference
	NameInput() interface{}
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	Proxied() interface{}
	SetProxied(val interface{})
	ProxiedInput() interface{}
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	Tag() DataCloudflareDnsRecordFilterTagOutputReference
	TagInput() interface{}
	TagMatch() *string
	SetTagMatch(val *string)
	TagMatchInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutComment(value *DataCloudflareDnsRecordFilterComment)
	PutContent(value *DataCloudflareDnsRecordFilterContent)
	PutName(value *DataCloudflareDnsRecordFilterName)
	PutTag(value *DataCloudflareDnsRecordFilterTag)
	ResetComment()
	ResetContent()
	ResetDirection()
	ResetMatch()
	ResetName()
	ResetOrder()
	ResetProxied()
	ResetSearch()
	ResetTag()
	ResetTagMatch()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareDnsRecordFilterOutputReference
type jsiiProxy_DataCloudflareDnsRecordFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Comment() DataCloudflareDnsRecordFilterCommentOutputReference {
	var returns DataCloudflareDnsRecordFilterCommentOutputReference
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) CommentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Content() DataCloudflareDnsRecordFilterContentOutputReference {
	var returns DataCloudflareDnsRecordFilterContentOutputReference
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Name() DataCloudflareDnsRecordFilterNameOutputReference {
	var returns DataCloudflareDnsRecordFilterNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Proxied() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxied",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ProxiedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Tag() DataCloudflareDnsRecordFilterTagOutputReference {
	var returns DataCloudflareDnsRecordFilterTagOutputReference
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TagMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TagMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareDnsRecordFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareDnsRecordFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareDnsRecordFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareDnsRecordFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecord.DataCloudflareDnsRecordFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareDnsRecordFilterOutputReference_Override(d DataCloudflareDnsRecordFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareDnsRecord.DataCloudflareDnsRecordFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetProxied(val interface{}) {
	if err := j.validateSetProxiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxied",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetTagMatch(val *string) {
	if err := j.validateSetTagMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagMatch",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) PutComment(value *DataCloudflareDnsRecordFilterComment) {
	if err := d.validatePutCommentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) PutContent(value *DataCloudflareDnsRecordFilterContent) {
	if err := d.validatePutContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) PutName(value *DataCloudflareDnsRecordFilterName) {
	if err := d.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putName",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) PutTag(value *DataCloudflareDnsRecordFilterTag) {
	if err := d.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		d,
		"resetContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		d,
		"resetMatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetProxied() {
	_jsii_.InvokeVoid(
		d,
		"resetProxied",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		d,
		"resetTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetTagMatch() {
	_jsii_.InvokeVoid(
		d,
		"resetTagMatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareDnsRecordFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


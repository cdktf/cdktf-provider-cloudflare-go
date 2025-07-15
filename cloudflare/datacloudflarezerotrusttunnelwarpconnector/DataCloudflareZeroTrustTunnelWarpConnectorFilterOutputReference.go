// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelwarpconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrusttunnelwarpconnector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference interface {
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
	ExcludePrefix() *string
	SetExcludePrefix(val *string)
	ExcludePrefixInput() *string
	ExistedAt() *string
	SetExistedAt(val *string)
	ExistedAtInput() *string
	// Experimental.
	Fqn() *string
	IncludePrefix() *string
	SetIncludePrefix(val *string)
	IncludePrefixInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsDeleted() interface{}
	SetIsDeleted(val interface{})
	IsDeletedInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
	WasActiveAt() *string
	SetWasActiveAt(val *string)
	WasActiveAtInput() *string
	WasInactiveAt() *string
	SetWasInactiveAt(val *string)
	WasInactiveAtInput() *string
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
	ResetExcludePrefix()
	ResetExistedAt()
	ResetIncludePrefix()
	ResetIsDeleted()
	ResetName()
	ResetStatus()
	ResetUuid()
	ResetWasActiveAt()
	ResetWasInactiveAt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference
type jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ExcludePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ExcludePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ExistedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ExistedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) IncludePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) IncludePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) IsDeleted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) IsDeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) WasActiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasActiveAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) WasActiveAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasActiveAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) WasInactiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasInactiveAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) WasInactiveAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasInactiveAtInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelWarpConnector.DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference_Override(d DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelWarpConnector.DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetExcludePrefix(val *string) {
	if err := j.validateSetExcludePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefix",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetExistedAt(val *string) {
	if err := j.validateSetExistedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existedAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetIncludePrefix(val *string) {
	if err := j.validateSetIncludePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefix",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetIsDeleted(val interface{}) {
	if err := j.validateSetIsDeletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDeleted",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetWasActiveAt(val *string) {
	if err := j.validateSetWasActiveAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wasActiveAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference)SetWasInactiveAt(val *string) {
	if err := j.validateSetWasInactiveAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wasInactiveAt",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetExcludePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetExistedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetExistedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetIncludePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetIsDeleted() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDeleted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetWasActiveAt() {
	_jsii_.InvokeVoid(
		d,
		"resetWasActiveAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ResetWasInactiveAt() {
	_jsii_.InvokeVoid(
		d,
		"resetWasInactiveAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelWarpConnectorFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarewaitingroomevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarewaitingroomevents/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWaitingRoomEventsResultOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomPageHtml() *string
	Description() *string
	DisableSessionRenewal() cdktf.IResolvable
	EventEndTime() *string
	EventStartTime() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataCloudflareWaitingRoomEventsResult
	SetInternalValue(val *DataCloudflareWaitingRoomEventsResult)
	ModifiedOn() *string
	Name() *string
	NewUsersPerMinute() *float64
	PrequeueStartTime() *string
	QueueingMethod() *string
	SessionDuration() *float64
	ShuffleAtEventStart() cdktf.IResolvable
	Suspended() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalActiveUsers() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareWaitingRoomEventsResultOutputReference
type jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) CustomPageHtml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) DisableSessionRenewal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableSessionRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) EventEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) EventStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) InternalValue() *DataCloudflareWaitingRoomEventsResult {
	var returns *DataCloudflareWaitingRoomEventsResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) NewUsersPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) PrequeueStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prequeueStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) QueueingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) SessionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ShuffleAtEventStart() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"shuffleAtEventStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Suspended() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) TotalActiveUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsers",
		&returns,
	)
	return returns
}


func NewDataCloudflareWaitingRoomEventsResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareWaitingRoomEventsResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareWaitingRoomEventsResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWaitingRoomEvents.DataCloudflareWaitingRoomEventsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareWaitingRoomEventsResultOutputReference_Override(d DataCloudflareWaitingRoomEventsResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWaitingRoomEvents.DataCloudflareWaitingRoomEventsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference)SetInternalValue(val *DataCloudflareWaitingRoomEventsResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomEventsResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


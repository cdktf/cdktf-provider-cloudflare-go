// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarewaitingrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarewaitingrooms/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWaitingRoomsResultOutputReference interface {
	cdktf.ComplexObject
	AdditionalRoutes() DataCloudflareWaitingRoomsResultAdditionalRoutesList
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
	CookieAttributes() DataCloudflareWaitingRoomsResultCookieAttributesOutputReference
	CookieSuffix() *string
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomPageHtml() *string
	DefaultTemplateLanguage() *string
	Description() *string
	DisableSessionRenewal() cdktf.IResolvable
	EnabledOriginCommands() *[]*string
	// Experimental.
	Fqn() *string
	Host() *string
	Id() *string
	InternalValue() *DataCloudflareWaitingRoomsResult
	SetInternalValue(val *DataCloudflareWaitingRoomsResult)
	JsonResponseEnabled() cdktf.IResolvable
	ModifiedOn() *string
	Name() *string
	NewUsersPerMinute() *float64
	NextEventPrequeueStartTime() *string
	NextEventStartTime() *string
	Path() *string
	QueueAll() cdktf.IResolvable
	QueueingMethod() *string
	QueueingStatusCode() *float64
	SessionDuration() *float64
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
	TurnstileAction() *string
	TurnstileMode() *string
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

// The jsii proxy struct for DataCloudflareWaitingRoomsResultOutputReference
type jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) AdditionalRoutes() DataCloudflareWaitingRoomsResultAdditionalRoutesList {
	var returns DataCloudflareWaitingRoomsResultAdditionalRoutesList
	_jsii_.Get(
		j,
		"additionalRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) CookieAttributes() DataCloudflareWaitingRoomsResultCookieAttributesOutputReference {
	var returns DataCloudflareWaitingRoomsResultCookieAttributesOutputReference
	_jsii_.Get(
		j,
		"cookieAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) CookieSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) CustomPageHtml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) DefaultTemplateLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTemplateLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) DisableSessionRenewal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableSessionRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) EnabledOriginCommands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledOriginCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) InternalValue() *DataCloudflareWaitingRoomsResult {
	var returns *DataCloudflareWaitingRoomsResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) JsonResponseEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jsonResponseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) NewUsersPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) NextEventPrequeueStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextEventPrequeueStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) NextEventStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextEventStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) QueueAll() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"queueAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) QueueingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) QueueingStatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueingStatusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) SessionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Suspended() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) TotalActiveUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) TurnstileAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turnstileAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) TurnstileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turnstileMode",
		&returns,
	)
	return returns
}


func NewDataCloudflareWaitingRoomsResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareWaitingRoomsResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareWaitingRoomsResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWaitingRooms.DataCloudflareWaitingRoomsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareWaitingRoomsResultOutputReference_Override(d DataCloudflareWaitingRoomsResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWaitingRooms.DataCloudflareWaitingRoomsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference)SetInternalValue(val *DataCloudflareWaitingRoomsResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareWaitingRoomsResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


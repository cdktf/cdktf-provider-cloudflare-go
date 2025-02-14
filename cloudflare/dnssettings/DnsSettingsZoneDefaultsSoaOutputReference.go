// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/dnssettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsSettingsZoneDefaultsSoaOutputReference interface {
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
	Expire() *float64
	SetExpire(val *float64)
	ExpireInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinTtl() *float64
	SetMinTtl(val *float64)
	MinTtlInput() *float64
	Mname() *string
	SetMname(val *string)
	MnameInput() *string
	Refresh() *float64
	SetRefresh(val *float64)
	RefreshInput() *float64
	Retry() *float64
	SetRetry(val *float64)
	RetryInput() *float64
	Rname() *string
	SetRname(val *string)
	RnameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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

// The jsii proxy struct for DnsSettingsZoneDefaultsSoaOutputReference
type jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Expire() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expire",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) ExpireInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Mname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) MnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Refresh() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) RefreshInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Retry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) RetryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Rname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) RnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewDnsSettingsZoneDefaultsSoaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsSettingsZoneDefaultsSoaOutputReference {
	_init_.Initialize()

	if err := validateNewDnsSettingsZoneDefaultsSoaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsSettings.DnsSettingsZoneDefaultsSoaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsSettingsZoneDefaultsSoaOutputReference_Override(d DnsSettingsZoneDefaultsSoaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsSettings.DnsSettingsZoneDefaultsSoaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetExpire(val *float64) {
	if err := j.validateSetExpireParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expire",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetMinTtl(val *float64) {
	if err := j.validateSetMinTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetMname(val *string) {
	if err := j.validateSetMnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mname",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetRefresh(val *float64) {
	if err := j.validateSetRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refresh",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetRetry(val *float64) {
	if err := j.validateSetRetryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retry",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetRname(val *string) {
	if err := j.validateSetRnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rname",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsSoaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


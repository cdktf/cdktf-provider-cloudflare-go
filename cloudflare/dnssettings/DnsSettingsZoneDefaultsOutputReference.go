// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnssettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/dnssettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsSettingsZoneDefaultsOutputReference interface {
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
	FlattenAllCnames() interface{}
	SetFlattenAllCnames(val interface{})
	FlattenAllCnamesInput() interface{}
	FoundationDns() interface{}
	SetFoundationDns(val interface{})
	FoundationDnsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalDns() DnsSettingsZoneDefaultsInternalDnsOutputReference
	InternalDnsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiProvider() interface{}
	SetMultiProvider(val interface{})
	MultiProviderInput() interface{}
	Nameservers() DnsSettingsZoneDefaultsNameserversOutputReference
	NameserversInput() interface{}
	NsTtl() *float64
	SetNsTtl(val *float64)
	NsTtlInput() *float64
	SecondaryOverrides() interface{}
	SetSecondaryOverrides(val interface{})
	SecondaryOverridesInput() interface{}
	Soa() DnsSettingsZoneDefaultsSoaOutputReference
	SoaInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZoneMode() *string
	SetZoneMode(val *string)
	ZoneModeInput() *string
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
	PutInternalDns(value *DnsSettingsZoneDefaultsInternalDns)
	PutNameservers(value *DnsSettingsZoneDefaultsNameservers)
	PutSoa(value *DnsSettingsZoneDefaultsSoa)
	ResetFlattenAllCnames()
	ResetFoundationDns()
	ResetInternalDns()
	ResetMultiProvider()
	ResetNameservers()
	ResetNsTtl()
	ResetSecondaryOverrides()
	ResetSoa()
	ResetZoneMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsSettingsZoneDefaultsOutputReference
type jsiiProxy_DnsSettingsZoneDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) FlattenAllCnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) FlattenAllCnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) FoundationDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) FoundationDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) InternalDns() DnsSettingsZoneDefaultsInternalDnsOutputReference {
	var returns DnsSettingsZoneDefaultsInternalDnsOutputReference
	_jsii_.Get(
		j,
		"internalDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) InternalDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) MultiProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) MultiProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) Nameservers() DnsSettingsZoneDefaultsNameserversOutputReference {
	var returns DnsSettingsZoneDefaultsNameserversOutputReference
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) NameserversInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) NsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) NsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) SecondaryOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) SecondaryOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) Soa() DnsSettingsZoneDefaultsSoaOutputReference {
	var returns DnsSettingsZoneDefaultsSoaOutputReference
	_jsii_.Get(
		j,
		"soa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) SoaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"soaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneModeInput",
		&returns,
	)
	return returns
}


func NewDnsSettingsZoneDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsSettingsZoneDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewDnsSettingsZoneDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsSettingsZoneDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsSettings.DnsSettingsZoneDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsSettingsZoneDefaultsOutputReference_Override(d DnsSettingsZoneDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsSettings.DnsSettingsZoneDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetFlattenAllCnames(val interface{}) {
	if err := j.validateSetFlattenAllCnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenAllCnames",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetFoundationDns(val interface{}) {
	if err := j.validateSetFoundationDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationDns",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetMultiProvider(val interface{}) {
	if err := j.validateSetMultiProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiProvider",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetNsTtl(val *float64) {
	if err := j.validateSetNsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsTtl",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetSecondaryOverrides(val interface{}) {
	if err := j.validateSetSecondaryOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryOverrides",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DnsSettingsZoneDefaultsOutputReference)SetZoneMode(val *string) {
	if err := j.validateSetZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMode",
		val,
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) PutInternalDns(value *DnsSettingsZoneDefaultsInternalDns) {
	if err := d.validatePutInternalDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInternalDns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) PutNameservers(value *DnsSettingsZoneDefaultsNameservers) {
	if err := d.validatePutNameserversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNameservers",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) PutSoa(value *DnsSettingsZoneDefaultsSoa) {
	if err := d.validatePutSoaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSoa",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetFlattenAllCnames() {
	_jsii_.InvokeVoid(
		d,
		"resetFlattenAllCnames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetFoundationDns() {
	_jsii_.InvokeVoid(
		d,
		"resetFoundationDns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetInternalDns() {
	_jsii_.InvokeVoid(
		d,
		"resetInternalDns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetMultiProvider() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiProvider",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetNameservers() {
	_jsii_.InvokeVoid(
		d,
		"resetNameservers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetNsTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetNsTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetSecondaryOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetSecondaryOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetSoa() {
	_jsii_.InvokeVoid(
		d,
		"resetSoa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ResetZoneMode() {
	_jsii_.InvokeVoid(
		d,
		"resetZoneMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DnsSettingsZoneDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


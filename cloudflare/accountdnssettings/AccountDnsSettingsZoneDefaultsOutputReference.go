// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountdnssettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/accountdnssettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountDnsSettingsZoneDefaultsOutputReference interface {
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
	InternalDns() AccountDnsSettingsZoneDefaultsInternalDnsOutputReference
	InternalDnsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiProvider() interface{}
	SetMultiProvider(val interface{})
	MultiProviderInput() interface{}
	Nameservers() AccountDnsSettingsZoneDefaultsNameserversOutputReference
	NameserversInput() interface{}
	NsTtl() *float64
	SetNsTtl(val *float64)
	NsTtlInput() *float64
	SecondaryOverrides() interface{}
	SetSecondaryOverrides(val interface{})
	SecondaryOverridesInput() interface{}
	Soa() AccountDnsSettingsZoneDefaultsSoaOutputReference
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
	PutInternalDns(value *AccountDnsSettingsZoneDefaultsInternalDns)
	PutNameservers(value *AccountDnsSettingsZoneDefaultsNameservers)
	PutSoa(value *AccountDnsSettingsZoneDefaultsSoa)
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

// The jsii proxy struct for AccountDnsSettingsZoneDefaultsOutputReference
type jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FlattenAllCnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FlattenAllCnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FoundationDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FoundationDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalDns() AccountDnsSettingsZoneDefaultsInternalDnsOutputReference {
	var returns AccountDnsSettingsZoneDefaultsInternalDnsOutputReference
	_jsii_.Get(
		j,
		"internalDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) MultiProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) MultiProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Nameservers() AccountDnsSettingsZoneDefaultsNameserversOutputReference {
	var returns AccountDnsSettingsZoneDefaultsNameserversOutputReference
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NameserversInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SecondaryOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SecondaryOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Soa() AccountDnsSettingsZoneDefaultsSoaOutputReference {
	var returns AccountDnsSettingsZoneDefaultsSoaOutputReference
	_jsii_.Get(
		j,
		"soa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SoaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"soaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneModeInput",
		&returns,
	)
	return returns
}


func NewAccountDnsSettingsZoneDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountDnsSettingsZoneDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountDnsSettingsZoneDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountDnsSettings.AccountDnsSettingsZoneDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountDnsSettingsZoneDefaultsOutputReference_Override(a AccountDnsSettingsZoneDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountDnsSettings.AccountDnsSettingsZoneDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetFlattenAllCnames(val interface{}) {
	if err := j.validateSetFlattenAllCnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenAllCnames",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetFoundationDns(val interface{}) {
	if err := j.validateSetFoundationDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationDns",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetMultiProvider(val interface{}) {
	if err := j.validateSetMultiProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiProvider",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetNsTtl(val *float64) {
	if err := j.validateSetNsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsTtl",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetSecondaryOverrides(val interface{}) {
	if err := j.validateSetSecondaryOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryOverrides",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference)SetZoneMode(val *string) {
	if err := j.validateSetZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMode",
		val,
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutInternalDns(value *AccountDnsSettingsZoneDefaultsInternalDns) {
	if err := a.validatePutInternalDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInternalDns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutNameservers(value *AccountDnsSettingsZoneDefaultsNameservers) {
	if err := a.validatePutNameserversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNameservers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutSoa(value *AccountDnsSettingsZoneDefaultsSoa) {
	if err := a.validatePutSoaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoa",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetFlattenAllCnames() {
	_jsii_.InvokeVoid(
		a,
		"resetFlattenAllCnames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetFoundationDns() {
	_jsii_.InvokeVoid(
		a,
		"resetFoundationDns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetInternalDns() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalDns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetMultiProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetNameservers() {
	_jsii_.InvokeVoid(
		a,
		"resetNameservers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetNsTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetNsTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetSecondaryOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetSoa() {
	_jsii_.InvokeVoid(
		a,
		"resetSoa",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetZoneMode() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


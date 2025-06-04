// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonesubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zonesubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneSubscriptionRatePlanOutputReference interface {
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
	Currency() *string
	SetCurrency(val *string)
	CurrencyInput() *string
	ExternallyManaged() interface{}
	SetExternallyManaged(val interface{})
	ExternallyManagedInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsContract() interface{}
	SetIsContract(val interface{})
	IsContractInput() interface{}
	PublicName() *string
	SetPublicName(val *string)
	PublicNameInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Sets() *[]*string
	SetSets(val *[]*string)
	SetsInput() *[]*string
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
	ResetCurrency()
	ResetExternallyManaged()
	ResetId()
	ResetIsContract()
	ResetPublicName()
	ResetScope()
	ResetSets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSubscriptionRatePlanOutputReference
type jsiiProxy_ZoneSubscriptionRatePlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Currency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) CurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ExternallyManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externallyManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ExternallyManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externallyManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) IsContract() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) IsContractInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isContractInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) PublicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) PublicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Sets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) SetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"setsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZoneSubscriptionRatePlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZoneSubscriptionRatePlanOutputReference {
	_init_.Initialize()

	if err := validateNewZoneSubscriptionRatePlanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSubscriptionRatePlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneSubscription.ZoneSubscriptionRatePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZoneSubscriptionRatePlanOutputReference_Override(z ZoneSubscriptionRatePlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneSubscription.ZoneSubscriptionRatePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetCurrency(val *string) {
	if err := j.validateSetCurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currency",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetExternallyManaged(val interface{}) {
	if err := j.validateSetExternallyManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externallyManaged",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetIsContract(val interface{}) {
	if err := j.validateSetIsContractParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isContract",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetPublicName(val *string) {
	if err := j.validateSetPublicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicName",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetSets(val *[]*string) {
	if err := j.validateSetSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sets",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSubscriptionRatePlanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetCurrency() {
	_jsii_.InvokeVoid(
		z,
		"resetCurrency",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetExternallyManaged() {
	_jsii_.InvokeVoid(
		z,
		"resetExternallyManaged",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetIsContract() {
	_jsii_.InvokeVoid(
		z,
		"resetIsContract",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetPublicName() {
	_jsii_.InvokeVoid(
		z,
		"resetPublicName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		z,
		"resetScope",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ResetSets() {
	_jsii_.InvokeVoid(
		z,
		"resetSets",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSubscriptionRatePlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


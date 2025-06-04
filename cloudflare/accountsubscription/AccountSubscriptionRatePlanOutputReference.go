// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/accountsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountSubscriptionRatePlanOutputReference interface {
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

// The jsii proxy struct for AccountSubscriptionRatePlanOutputReference
type jsiiProxy_AccountSubscriptionRatePlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Currency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) CurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ExternallyManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externallyManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ExternallyManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externallyManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) IsContract() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isContract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) IsContractInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isContractInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) PublicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) PublicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Sets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) SetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"setsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSubscriptionRatePlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountSubscriptionRatePlanOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSubscriptionRatePlanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSubscriptionRatePlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountSubscription.AccountSubscriptionRatePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSubscriptionRatePlanOutputReference_Override(a AccountSubscriptionRatePlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountSubscription.AccountSubscriptionRatePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetCurrency(val *string) {
	if err := j.validateSetCurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currency",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetExternallyManaged(val interface{}) {
	if err := j.validateSetExternallyManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externallyManaged",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetIsContract(val interface{}) {
	if err := j.validateSetIsContractParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isContract",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetPublicName(val *string) {
	if err := j.validateSetPublicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicName",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetSets(val *[]*string) {
	if err := j.validateSetSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sets",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSubscriptionRatePlanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetCurrency() {
	_jsii_.InvokeVoid(
		a,
		"resetCurrency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetExternallyManaged() {
	_jsii_.InvokeVoid(
		a,
		"resetExternallyManaged",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetIsContract() {
	_jsii_.InvokeVoid(
		a,
		"resetIsContract",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetPublicName() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ResetSets() {
	_jsii_.InvokeVoid(
		a,
		"resetSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSubscriptionRatePlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


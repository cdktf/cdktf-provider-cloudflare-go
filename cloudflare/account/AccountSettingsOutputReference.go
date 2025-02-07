// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/account/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountSettingsOutputReference interface {
	cdktf.ComplexObject
	AbuseContactEmail() *string
	SetAbuseContactEmail(val *string)
	AbuseContactEmailInput() *string
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
	DefaultNameservers() *string
	SetDefaultNameservers(val *string)
	DefaultNameserversInput() *string
	EnforceTwofactor() interface{}
	SetEnforceTwofactor(val interface{})
	EnforceTwofactorInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseAccountCustomNsByDefault() interface{}
	SetUseAccountCustomNsByDefault(val interface{})
	UseAccountCustomNsByDefaultInput() interface{}
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
	ResetAbuseContactEmail()
	ResetDefaultNameservers()
	ResetEnforceTwofactor()
	ResetUseAccountCustomNsByDefault()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountSettingsOutputReference
type jsiiProxy_AccountSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountSettingsOutputReference) AbuseContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) AbuseContactEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) DefaultNameservers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) DefaultNameserversInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) EnforceTwofactor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceTwofactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) EnforceTwofactorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceTwofactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) UseAccountCustomNsByDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAccountCustomNsByDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingsOutputReference) UseAccountCustomNsByDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAccountCustomNsByDefaultInput",
		&returns,
	)
	return returns
}


func NewAccountSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.account.AccountSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingsOutputReference_Override(a AccountSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.account.AccountSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetAbuseContactEmail(val *string) {
	if err := j.validateSetAbuseContactEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abuseContactEmail",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetDefaultNameservers(val *string) {
	if err := j.validateSetDefaultNameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNameservers",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetEnforceTwofactor(val interface{}) {
	if err := j.validateSetEnforceTwofactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceTwofactor",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccountSettingsOutputReference)SetUseAccountCustomNsByDefault(val interface{}) {
	if err := j.validateSetUseAccountCustomNsByDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAccountCustomNsByDefault",
		val,
	)
}

func (a *jsiiProxy_AccountSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountSettingsOutputReference) ResetAbuseContactEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetAbuseContactEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingsOutputReference) ResetDefaultNameservers() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultNameservers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingsOutputReference) ResetEnforceTwofactor() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceTwofactor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingsOutputReference) ResetUseAccountCustomNsByDefault() {
	_jsii_.InvokeVoid(
		a,
		"resetUseAccountCustomNsByDefault",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


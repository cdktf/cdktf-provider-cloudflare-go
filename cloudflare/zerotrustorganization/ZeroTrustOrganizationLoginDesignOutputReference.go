// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustorganization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustOrganizationLoginDesignOutputReference interface {
	cdktf.ComplexObject
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
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
	FooterText() *string
	SetFooterText(val *string)
	FooterTextInput() *string
	// Experimental.
	Fqn() *string
	HeaderText() *string
	SetHeaderText(val *string)
	HeaderTextInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogoPath() *string
	SetLogoPath(val *string)
	LogoPathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextColor() *string
	SetTextColor(val *string)
	TextColorInput() *string
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
	ResetBackgroundColor()
	ResetFooterText()
	ResetHeaderText()
	ResetLogoPath()
	ResetTextColor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustOrganizationLoginDesignOutputReference
type jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) FooterText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) FooterTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) HeaderText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) HeaderTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) LogoPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) LogoPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) TextColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) TextColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColorInput",
		&returns,
	)
	return returns
}


func NewZeroTrustOrganizationLoginDesignOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustOrganizationLoginDesignOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustOrganizationLoginDesignOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganizationLoginDesignOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustOrganizationLoginDesignOutputReference_Override(z ZeroTrustOrganizationLoginDesignOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganizationLoginDesignOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetFooterText(val *string) {
	if err := j.validateSetFooterTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"footerText",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetHeaderText(val *string) {
	if err := j.validateSetHeaderTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerText",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetLogoPath(val *string) {
	if err := j.validateSetLogoPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoPath",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference)SetTextColor(val *string) {
	if err := j.validateSetTextColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textColor",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		z,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ResetFooterText() {
	_jsii_.InvokeVoid(
		z,
		"resetFooterText",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ResetHeaderText() {
	_jsii_.InvokeVoid(
		z,
		"resetHeaderText",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ResetLogoPath() {
	_jsii_.InvokeVoid(
		z,
		"resetLogoPath",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ResetTextColor() {
	_jsii_.InvokeVoid(
		z,
		"resetTextColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustOrganizationLoginDesignOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsSettingsBlockPageOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	MailtoAddress() *string
	SetMailtoAddress(val *string)
	MailtoAddressInput() *string
	MailtoSubject() *string
	SetMailtoSubject(val *string)
	MailtoSubjectInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	SuppressFooter() interface{}
	SetSuppressFooter(val interface{})
	SuppressFooterInput() interface{}
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
	ResetBackgroundColor()
	ResetEnabled()
	ResetFooterText()
	ResetHeaderText()
	ResetLogoPath()
	ResetMailtoAddress()
	ResetMailtoSubject()
	ResetName()
	ResetSuppressFooter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewaySettingsSettingsBlockPageOutputReference
type jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) FooterText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) FooterTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) HeaderText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) HeaderTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) LogoPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) LogoPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) MailtoAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailtoAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) MailtoAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailtoAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) MailtoSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailtoSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) MailtoSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailtoSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) SuppressFooter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressFooter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) SuppressFooterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressFooterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewaySettingsSettingsBlockPageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewaySettingsSettingsBlockPageOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsSettingsBlockPageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsBlockPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewaySettingsSettingsBlockPageOutputReference_Override(z ZeroTrustGatewaySettingsSettingsBlockPageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsBlockPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetFooterText(val *string) {
	if err := j.validateSetFooterTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"footerText",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetHeaderText(val *string) {
	if err := j.validateSetHeaderTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerText",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetLogoPath(val *string) {
	if err := j.validateSetLogoPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoPath",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetMailtoAddress(val *string) {
	if err := j.validateSetMailtoAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailtoAddress",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetMailtoSubject(val *string) {
	if err := j.validateSetMailtoSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailtoSubject",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetSuppressFooter(val interface{}) {
	if err := j.validateSetSuppressFooterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressFooter",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		z,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetFooterText() {
	_jsii_.InvokeVoid(
		z,
		"resetFooterText",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetHeaderText() {
	_jsii_.InvokeVoid(
		z,
		"resetHeaderText",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetLogoPath() {
	_jsii_.InvokeVoid(
		z,
		"resetLogoPath",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetMailtoAddress() {
	_jsii_.InvokeVoid(
		z,
		"resetMailtoAddress",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetMailtoSubject() {
	_jsii_.InvokeVoid(
		z,
		"resetMailtoSubject",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		z,
		"resetName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ResetSuppressFooter() {
	_jsii_.InvokeVoid(
		z,
		"resetSuppressFooter",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsBlockPageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


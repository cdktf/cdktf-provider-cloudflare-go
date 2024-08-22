// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustgatewaypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference interface {
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
	DisableClipboardRedirection() interface{}
	SetDisableClipboardRedirection(val interface{})
	DisableClipboardRedirectionInput() interface{}
	DisableCopyPaste() interface{}
	SetDisableCopyPaste(val interface{})
	DisableCopyPasteInput() interface{}
	DisableDownload() interface{}
	SetDisableDownload(val interface{})
	DisableDownloadInput() interface{}
	DisableKeyboard() interface{}
	SetDisableKeyboard(val interface{})
	DisableKeyboardInput() interface{}
	DisablePrinting() interface{}
	SetDisablePrinting(val interface{})
	DisablePrintingInput() interface{}
	DisableUpload() interface{}
	SetDisableUpload(val interface{})
	DisableUploadInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls
	SetInternalValue(val *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls)
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
	ResetDisableClipboardRedirection()
	ResetDisableCopyPaste()
	ResetDisableDownload()
	ResetDisableKeyboard()
	ResetDisablePrinting()
	ResetDisableUpload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference
type jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableClipboardRedirection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableClipboardRedirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableClipboardRedirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableClipboardRedirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableCopyPaste() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCopyPaste",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableCopyPasteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCopyPasteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableDownload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDownload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableDownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDownloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableKeyboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableKeyboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableKeyboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableKeyboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisablePrinting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrinting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisablePrintingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePrintingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableUpload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) DisableUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) InternalValue() *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls {
	var returns *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference_Override(z ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisableClipboardRedirection(val interface{}) {
	if err := j.validateSetDisableClipboardRedirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableClipboardRedirection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisableCopyPaste(val interface{}) {
	if err := j.validateSetDisableCopyPasteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCopyPaste",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisableDownload(val interface{}) {
	if err := j.validateSetDisableDownloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDownload",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisableKeyboard(val interface{}) {
	if err := j.validateSetDisableKeyboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableKeyboard",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisablePrinting(val interface{}) {
	if err := j.validateSetDisablePrintingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePrinting",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetDisableUpload(val interface{}) {
	if err := j.validateSetDisableUploadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUpload",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetInternalValue(val *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisableClipboardRedirection() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableClipboardRedirection",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisableCopyPaste() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableCopyPaste",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisableDownload() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableDownload",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisableKeyboard() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableKeyboard",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisablePrinting() {
	_jsii_.InvokeVoid(
		z,
		"resetDisablePrinting",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ResetDisableUpload() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableUpload",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


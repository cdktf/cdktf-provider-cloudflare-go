// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsSettingsAntivirusOutputReference interface {
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
	EnabledDownloadPhase() interface{}
	SetEnabledDownloadPhase(val interface{})
	EnabledDownloadPhaseInput() interface{}
	EnabledUploadPhase() interface{}
	SetEnabledUploadPhase(val interface{})
	EnabledUploadPhaseInput() interface{}
	FailClosed() interface{}
	SetFailClosed(val interface{})
	FailClosedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotificationSettings() ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsOutputReference
	NotificationSettingsInput() interface{}
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
	PutNotificationSettings(value *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings)
	ResetEnabledDownloadPhase()
	ResetEnabledUploadPhase()
	ResetFailClosed()
	ResetNotificationSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewaySettingsSettingsAntivirusOutputReference
type jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) EnabledDownloadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) EnabledDownloadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) EnabledUploadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) EnabledUploadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) FailClosed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) FailClosedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) NotificationSettings() ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) NotificationSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewaySettingsSettingsAntivirusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewaySettingsSettingsAntivirusOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsSettingsAntivirusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewaySettingsSettingsAntivirusOutputReference_Override(z ZeroTrustGatewaySettingsSettingsAntivirusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetEnabledDownloadPhase(val interface{}) {
	if err := j.validateSetEnabledDownloadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledDownloadPhase",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetEnabledUploadPhase(val interface{}) {
	if err := j.validateSetEnabledUploadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledUploadPhase",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetFailClosed(val interface{}) {
	if err := j.validateSetFailClosedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failClosed",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) PutNotificationSettings(value *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings) {
	if err := z.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ResetEnabledDownloadPhase() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabledDownloadPhase",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ResetEnabledUploadPhase() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabledUploadPhase",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ResetFailClosed() {
	_jsii_.InvokeVoid(
		z,
		"resetFailClosed",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		z,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsAntivirusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


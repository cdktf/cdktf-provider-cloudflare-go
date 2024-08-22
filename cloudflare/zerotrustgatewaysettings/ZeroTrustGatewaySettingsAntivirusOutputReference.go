// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsAntivirusOutputReference interface {
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
	InternalValue() *ZeroTrustGatewaySettingsAntivirus
	SetInternalValue(val *ZeroTrustGatewaySettingsAntivirus)
	NotificationSettings() ZeroTrustGatewaySettingsAntivirusNotificationSettingsOutputReference
	NotificationSettingsInput() *ZeroTrustGatewaySettingsAntivirusNotificationSettings
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
	PutNotificationSettings(value *ZeroTrustGatewaySettingsAntivirusNotificationSettings)
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

// The jsii proxy struct for ZeroTrustGatewaySettingsAntivirusOutputReference
type jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) EnabledDownloadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) EnabledDownloadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) EnabledUploadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) EnabledUploadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) FailClosed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) FailClosedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) InternalValue() *ZeroTrustGatewaySettingsAntivirus {
	var returns *ZeroTrustGatewaySettingsAntivirus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) NotificationSettings() ZeroTrustGatewaySettingsAntivirusNotificationSettingsOutputReference {
	var returns ZeroTrustGatewaySettingsAntivirusNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) NotificationSettingsInput() *ZeroTrustGatewaySettingsAntivirusNotificationSettings {
	var returns *ZeroTrustGatewaySettingsAntivirusNotificationSettings
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewaySettingsAntivirusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewaySettingsAntivirusOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsAntivirusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewaySettingsAntivirusOutputReference_Override(z ZeroTrustGatewaySettingsAntivirusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetEnabledDownloadPhase(val interface{}) {
	if err := j.validateSetEnabledDownloadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledDownloadPhase",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetEnabledUploadPhase(val interface{}) {
	if err := j.validateSetEnabledUploadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledUploadPhase",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetFailClosed(val interface{}) {
	if err := j.validateSetFailClosedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failClosed",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetInternalValue(val *ZeroTrustGatewaySettingsAntivirus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) PutNotificationSettings(value *ZeroTrustGatewaySettingsAntivirusNotificationSettings) {
	if err := z.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		z,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsAntivirusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


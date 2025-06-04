// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicepostureintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustdevicepostureintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDevicePostureIntegrationConfigAOutputReference interface {
	cdktf.ComplexObject
	AccessClientId() *string
	SetAccessClientId(val *string)
	AccessClientIdInput() *string
	AccessClientSecret() *string
	SetAccessClientSecret(val *string)
	AccessClientSecretInput() *string
	ApiUrl() *string
	SetApiUrl(val *string)
	ApiUrlInput() *string
	AuthUrl() *string
	SetAuthUrl(val *string)
	AuthUrlInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	CustomerId() *string
	SetCustomerId(val *string)
	CustomerIdInput() *string
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
	ResetAccessClientId()
	ResetAccessClientSecret()
	ResetApiUrl()
	ResetAuthUrl()
	ResetClientId()
	ResetClientKey()
	ResetClientSecret()
	ResetCustomerId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustDevicePostureIntegrationConfigAOutputReference
type jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AccessClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AccessClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AccessClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AccessClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) AuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) CustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) CustomerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustDevicePostureIntegrationConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustDevicePostureIntegrationConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustDevicePostureIntegrationConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureIntegration.ZeroTrustDevicePostureIntegrationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustDevicePostureIntegrationConfigAOutputReference_Override(z ZeroTrustDevicePostureIntegrationConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureIntegration.ZeroTrustDevicePostureIntegrationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetAccessClientId(val *string) {
	if err := j.validateSetAccessClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessClientId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetAccessClientSecret(val *string) {
	if err := j.validateSetAccessClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessClientSecret",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetApiUrl(val *string) {
	if err := j.validateSetApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetAuthUrl(val *string) {
	if err := j.validateSetAuthUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetCustomerId(val *string) {
	if err := j.validateSetCustomerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetAccessClientId() {
	_jsii_.InvokeVoid(
		z,
		"resetAccessClientId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetAccessClientSecret() {
	_jsii_.InvokeVoid(
		z,
		"resetAccessClientSecret",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetApiUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetAuthUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		z,
		"resetClientId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetClientKey() {
	_jsii_.InvokeVoid(
		z,
		"resetClientKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		z,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ResetCustomerId() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomerId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustDevicePostureIntegrationConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


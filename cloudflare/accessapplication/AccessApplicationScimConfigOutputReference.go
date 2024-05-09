// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/accessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessApplicationScimConfigOutputReference interface {
	cdktf.ComplexObject
	Authentication() AccessApplicationScimConfigAuthenticationOutputReference
	AuthenticationInput() *AccessApplicationScimConfigAuthentication
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
	DeactivateOnDelete() interface{}
	SetDeactivateOnDelete(val interface{})
	DeactivateOnDeleteInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IdpUid() *string
	SetIdpUid(val *string)
	IdpUidInput() *string
	InternalValue() *AccessApplicationScimConfig
	SetInternalValue(val *AccessApplicationScimConfig)
	Mappings() AccessApplicationScimConfigMappingsList
	MappingsInput() interface{}
	RemoteUri() *string
	SetRemoteUri(val *string)
	RemoteUriInput() *string
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
	PutAuthentication(value *AccessApplicationScimConfigAuthentication)
	PutMappings(value interface{})
	ResetAuthentication()
	ResetDeactivateOnDelete()
	ResetEnabled()
	ResetMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessApplicationScimConfigOutputReference
type jsiiProxy_AccessApplicationScimConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) Authentication() AccessApplicationScimConfigAuthenticationOutputReference {
	var returns AccessApplicationScimConfigAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) AuthenticationInput() *AccessApplicationScimConfigAuthentication {
	var returns *AccessApplicationScimConfigAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) DeactivateOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivateOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) DeactivateOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivateOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) IdpUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) IdpUidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) InternalValue() *AccessApplicationScimConfig {
	var returns *AccessApplicationScimConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) Mappings() AccessApplicationScimConfigMappingsList {
	var returns AccessApplicationScimConfigMappingsList
	_jsii_.Get(
		j,
		"mappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) MappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) RemoteUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) RemoteUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessApplicationScimConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessApplicationScimConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAccessApplicationScimConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplicationScimConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationScimConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessApplicationScimConfigOutputReference_Override(a AccessApplicationScimConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationScimConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetDeactivateOnDelete(val interface{}) {
	if err := j.validateSetDeactivateOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivateOnDelete",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetIdpUid(val *string) {
	if err := j.validateSetIdpUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpUid",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetInternalValue(val *AccessApplicationScimConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetRemoteUri(val *string) {
	if err := j.validateSetRemoteUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteUri",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationScimConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) PutAuthentication(value *AccessApplicationScimConfigAuthentication) {
	if err := a.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) PutMappings(value interface{}) {
	if err := a.validatePutMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMappings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ResetDeactivateOnDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDeactivateOnDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ResetMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessApplicationScimConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


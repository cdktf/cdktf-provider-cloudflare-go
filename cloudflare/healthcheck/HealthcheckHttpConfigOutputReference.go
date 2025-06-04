// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/healthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcheckHttpConfigOutputReference interface {
	cdktf.ComplexObject
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
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
	ExpectedBody() *string
	SetExpectedBody(val *string)
	ExpectedBodyInput() *string
	ExpectedCodes() *[]*string
	SetExpectedCodes(val *[]*string)
	ExpectedCodesInput() *[]*string
	FollowRedirects() interface{}
	SetFollowRedirects(val interface{})
	FollowRedirectsInput() interface{}
	// Experimental.
	Fqn() *string
	Header() interface{}
	SetHeader(val interface{})
	HeaderInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ResetAllowInsecure()
	ResetExpectedBody()
	ResetExpectedCodes()
	ResetFollowRedirects()
	ResetHeader()
	ResetMethod()
	ResetPath()
	ResetPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcheckHttpConfigOutputReference
type jsiiProxy_HealthcheckHttpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ExpectedBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ExpectedBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ExpectedCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) ExpectedCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) FollowRedirects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) FollowRedirectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) Header() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcheckHttpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcheckHttpConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcheckHttpConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcheckHttpConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.healthcheck.HealthcheckHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcheckHttpConfigOutputReference_Override(h HealthcheckHttpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.healthcheck.HealthcheckHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetExpectedBody(val *string) {
	if err := j.validateSetExpectedBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBody",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetExpectedCodes(val *[]*string) {
	if err := j.validateSetExpectedCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedCodes",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetFollowRedirects(val interface{}) {
	if err := j.validateSetFollowRedirectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"followRedirects",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetHeader(val interface{}) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcheckHttpConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetExpectedBody() {
	_jsii_.InvokeVoid(
		h,
		"resetExpectedBody",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetExpectedCodes() {
	_jsii_.InvokeVoid(
		h,
		"resetExpectedCodes",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetFollowRedirects() {
	_jsii_.InvokeVoid(
		h,
		"resetFollowRedirects",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		h,
		"resetHeader",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		h,
		"resetMethod",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		h,
		"resetPath",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		h,
		"resetPort",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcheckHttpConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apishieldoperation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/apishieldoperation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiShieldOperationFeaturesThresholdsOutputReference interface {
	cdktf.ComplexObject
	AuthIdTokens() *float64
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
	DataPoints() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ApiShieldOperationFeaturesThresholds
	SetInternalValue(val *ApiShieldOperationFeaturesThresholds)
	LastUpdated() *string
	P50() *float64
	P90() *float64
	P99() *float64
	PeriodSeconds() *float64
	Requests() *float64
	SuggestedThreshold() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiShieldOperationFeaturesThresholdsOutputReference
type jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) AuthIdTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authIdTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) DataPoints() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) InternalValue() *ApiShieldOperationFeaturesThresholds {
	var returns *ApiShieldOperationFeaturesThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) LastUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) P50() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"p50",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) P90() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"p90",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) P99() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"p99",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) Requests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) SuggestedThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suggestedThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiShieldOperationFeaturesThresholdsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiShieldOperationFeaturesThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewApiShieldOperationFeaturesThresholdsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.apiShieldOperation.ApiShieldOperationFeaturesThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiShieldOperationFeaturesThresholdsOutputReference_Override(a ApiShieldOperationFeaturesThresholdsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.apiShieldOperation.ApiShieldOperationFeaturesThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference)SetInternalValue(val *ApiShieldOperationFeaturesThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiShieldOperationFeaturesThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


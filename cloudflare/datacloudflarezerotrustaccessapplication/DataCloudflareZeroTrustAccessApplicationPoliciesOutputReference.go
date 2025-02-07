// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference interface {
	cdktf.ComplexObject
	ApprovalGroups() DataCloudflareZeroTrustAccessApplicationPoliciesApprovalGroupsList
	ApprovalRequired() cdktf.IResolvable
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
	ConnectionRules() DataCloudflareZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Decision() *string
	Exclude() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeList
	// Experimental.
	Fqn() *string
	Id() *string
	Include() DataCloudflareZeroTrustAccessApplicationPoliciesIncludeList
	InternalValue() *DataCloudflareZeroTrustAccessApplicationPolicies
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPolicies)
	IsolationRequired() cdktf.IResolvable
	Name() *string
	Precedence() *float64
	PurposeJustificationPrompt() *string
	PurposeJustificationRequired() cdktf.IResolvable
	Require() DataCloudflareZeroTrustAccessApplicationPoliciesRequireList
	SessionDuration() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ApprovalGroups() DataCloudflareZeroTrustAccessApplicationPoliciesApprovalGroupsList {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesApprovalGroupsList
	_jsii_.Get(
		j,
		"approvalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ApprovalRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"approvalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ConnectionRules() DataCloudflareZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference
	_jsii_.Get(
		j,
		"connectionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Decision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Exclude() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeList {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Include() DataCloudflareZeroTrustAccessApplicationPoliciesIncludeList {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationPolicies {
	var returns *DataCloudflareZeroTrustAccessApplicationPolicies
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) IsolationRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isolationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) PurposeJustificationPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) PurposeJustificationRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"purposeJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Require() DataCloudflareZeroTrustAccessApplicationPoliciesRequireList {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireList
	_jsii_.Get(
		j,
		"require",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationPoliciesOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPolicies) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


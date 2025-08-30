// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkersscripts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflareworkersscripts/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWorkersScriptsResultOutputReference interface {
	cdktf.ComplexObject
	CompatibilityDate() *string
	CompatibilityFlags() *[]*string
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Etag() *string
	// Experimental.
	Fqn() *string
	Handlers() *[]*string
	HasAssets() cdktf.IResolvable
	HasModules() cdktf.IResolvable
	Id() *string
	InternalValue() *DataCloudflareWorkersScriptsResult
	SetInternalValue(val *DataCloudflareWorkersScriptsResult)
	LastDeployedFrom() *string
	Logpush() cdktf.IResolvable
	MigrationTag() *string
	ModifiedOn() *string
	NamedHandlers() DataCloudflareWorkersScriptsResultNamedHandlersList
	Placement() DataCloudflareWorkersScriptsResultPlacementOutputReference
	PlacementMode() *string
	PlacementStatus() *string
	TailConsumers() DataCloudflareWorkersScriptsResultTailConsumersList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageModel() *string
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

// The jsii proxy struct for DataCloudflareWorkersScriptsResultOutputReference
type jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Handlers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"handlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) HasAssets() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) HasModules() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) InternalValue() *DataCloudflareWorkersScriptsResult {
	var returns *DataCloudflareWorkersScriptsResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) LastDeployedFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDeployedFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Logpush() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) MigrationTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) NamedHandlers() DataCloudflareWorkersScriptsResultNamedHandlersList {
	var returns DataCloudflareWorkersScriptsResultNamedHandlersList
	_jsii_.Get(
		j,
		"namedHandlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Placement() DataCloudflareWorkersScriptsResultPlacementOutputReference {
	var returns DataCloudflareWorkersScriptsResultPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) PlacementMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) PlacementStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) TailConsumers() DataCloudflareWorkersScriptsResultTailConsumersList {
	var returns DataCloudflareWorkersScriptsResultTailConsumersList
	_jsii_.Get(
		j,
		"tailConsumers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}


func NewDataCloudflareWorkersScriptsResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareWorkersScriptsResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareWorkersScriptsResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWorkersScripts.DataCloudflareWorkersScriptsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareWorkersScriptsResultOutputReference_Override(d DataCloudflareWorkersScriptsResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWorkersScripts.DataCloudflareWorkersScriptsResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference)SetInternalValue(val *DataCloudflareWorkersScriptsResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareWorkersScriptsResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


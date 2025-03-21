// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdevicecustomprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustdevicecustomprofiles/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference interface {
	cdktf.ComplexObject
	AllowedToLeave() cdktf.IResolvable
	AllowModeSwitch() cdktf.IResolvable
	AllowUpdates() cdktf.IResolvable
	AutoConnect() *float64
	CaptivePortal() *float64
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
	Default() cdktf.IResolvable
	Description() *string
	DisableAutoFallback() cdktf.IResolvable
	Enabled() cdktf.IResolvable
	Exclude() DataCloudflareZeroTrustDeviceCustomProfilesResultExcludeList
	ExcludeOfficeIps() cdktf.IResolvable
	FallbackDomains() DataCloudflareZeroTrustDeviceCustomProfilesResultFallbackDomainsList
	// Experimental.
	Fqn() *string
	GatewayUniqueId() *string
	Include() DataCloudflareZeroTrustDeviceCustomProfilesResultIncludeList
	InternalValue() *DataCloudflareZeroTrustDeviceCustomProfilesResult
	SetInternalValue(val *DataCloudflareZeroTrustDeviceCustomProfilesResult)
	LanAllowMinutes() *float64
	LanAllowSubnetSize() *float64
	Match() *string
	Name() *string
	PolicyId() *string
	Precedence() *float64
	RegisterInterfaceIpWithDns() cdktf.IResolvable
	ServiceModeV2() DataCloudflareZeroTrustDeviceCustomProfilesResultServiceModeV2OutputReference
	SupportUrl() *string
	SwitchLocked() cdktf.IResolvable
	TargetTests() DataCloudflareZeroTrustDeviceCustomProfilesResultTargetTestsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelProtocol() *string
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

// The jsii proxy struct for DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference
type jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) AllowedToLeave() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowedToLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) AllowModeSwitch() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowModeSwitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) AllowUpdates() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) AutoConnect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) CaptivePortal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Default() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) DisableAutoFallback() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableAutoFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Exclude() DataCloudflareZeroTrustDeviceCustomProfilesResultExcludeList {
	var returns DataCloudflareZeroTrustDeviceCustomProfilesResultExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ExcludeOfficeIps() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"excludeOfficeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) FallbackDomains() DataCloudflareZeroTrustDeviceCustomProfilesResultFallbackDomainsList {
	var returns DataCloudflareZeroTrustDeviceCustomProfilesResultFallbackDomainsList
	_jsii_.Get(
		j,
		"fallbackDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GatewayUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Include() DataCloudflareZeroTrustDeviceCustomProfilesResultIncludeList {
	var returns DataCloudflareZeroTrustDeviceCustomProfilesResultIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) InternalValue() *DataCloudflareZeroTrustDeviceCustomProfilesResult {
	var returns *DataCloudflareZeroTrustDeviceCustomProfilesResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) LanAllowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) LanAllowSubnetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowSubnetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) RegisterInterfaceIpWithDns() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"registerInterfaceIpWithDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ServiceModeV2() DataCloudflareZeroTrustDeviceCustomProfilesResultServiceModeV2OutputReference {
	var returns DataCloudflareZeroTrustDeviceCustomProfilesResultServiceModeV2OutputReference
	_jsii_.Get(
		j,
		"serviceModeV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) SwitchLocked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"switchLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) TargetTests() DataCloudflareZeroTrustDeviceCustomProfilesResultTargetTestsList {
	var returns DataCloudflareZeroTrustDeviceCustomProfilesResultTargetTestsList
	_jsii_.Get(
		j,
		"targetTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) TunnelProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocol",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustDeviceCustomProfilesResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDeviceCustomProfiles.DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference_Override(d DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDeviceCustomProfiles.DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference)SetInternalValue(val *DataCloudflareZeroTrustDeviceCustomProfilesResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDeviceCustomProfilesResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


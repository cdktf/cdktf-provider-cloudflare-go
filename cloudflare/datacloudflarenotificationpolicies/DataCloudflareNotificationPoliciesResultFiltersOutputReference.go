// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarenotificationpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarenotificationpolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareNotificationPoliciesResultFiltersOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	AffectedAsns() *[]*string
	AffectedComponents() *[]*string
	AffectedLocations() *[]*string
	AirportCode() *[]*string
	AlertTriggerPreferences() *[]*string
	AlertTriggerPreferencesValue() *[]*string
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
	Enabled() *[]*string
	Environment() *[]*string
	Event() *[]*string
	EventSource() *[]*string
	EventType() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	HealthCheckId() *[]*string
	IncidentImpact() *[]*string
	InputId() *[]*string
	InsightClass() *[]*string
	InternalValue() *DataCloudflareNotificationPoliciesResultFilters
	SetInternalValue(val *DataCloudflareNotificationPoliciesResultFilters)
	Limit() *[]*string
	LogoTag() *[]*string
	MegabitsPerSecond() *[]*string
	NewHealth() *[]*string
	NewStatus() *[]*string
	PacketsPerSecond() *[]*string
	PoolId() *[]*string
	PopNames() *[]*string
	Product() *[]*string
	ProjectId() *[]*string
	Protocol() *[]*string
	QueryTag() *[]*string
	RequestsPerSecond() *[]*string
	Selectors() *[]*string
	Services() *[]*string
	Slo() *[]*string
	Status() *[]*string
	TargetHostname() *[]*string
	TargetIp() *[]*string
	TargetZoneName() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficExclusions() *[]*string
	TunnelId() *[]*string
	TunnelName() *[]*string
	Where() *[]*string
	Zones() *[]*string
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

// The jsii proxy struct for DataCloudflareNotificationPoliciesResultFiltersOutputReference
type jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AffectedAsns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedAsns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AffectedComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AffectedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AirportCode() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"airportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AlertTriggerPreferences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) AlertTriggerPreferencesValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferencesValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Enabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Environment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Event() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) EventSource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) EventType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) HealthCheckId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) IncidentImpact() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"incidentImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) InputId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) InsightClass() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insightClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) InternalValue() *DataCloudflareNotificationPoliciesResultFilters {
	var returns *DataCloudflareNotificationPoliciesResultFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Limit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) LogoTag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) MegabitsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"megabitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) NewHealth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) NewStatus() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) PacketsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packetsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) PoolId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) PopNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"popNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Product() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) ProjectId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Protocol() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) QueryTag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) RequestsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Selectors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Slo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Status() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TargetHostname() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TargetIp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TargetZoneName() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TrafficExclusions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trafficExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TunnelId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) TunnelName() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Where() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"where",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}


func NewDataCloudflareNotificationPoliciesResultFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareNotificationPoliciesResultFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareNotificationPoliciesResultFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareNotificationPolicies.DataCloudflareNotificationPoliciesResultFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareNotificationPoliciesResultFiltersOutputReference_Override(d DataCloudflareNotificationPoliciesResultFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareNotificationPolicies.DataCloudflareNotificationPoliciesResultFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference)SetInternalValue(val *DataCloudflareNotificationPoliciesResultFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareNotificationPoliciesResultFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


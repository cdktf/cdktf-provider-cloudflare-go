// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/notificationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyFiltersOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
	AffectedAsns() *[]*string
	SetAffectedAsns(val *[]*string)
	AffectedAsnsInput() *[]*string
	AffectedComponents() *[]*string
	SetAffectedComponents(val *[]*string)
	AffectedComponentsInput() *[]*string
	AffectedLocations() *[]*string
	SetAffectedLocations(val *[]*string)
	AffectedLocationsInput() *[]*string
	AirportCode() *[]*string
	SetAirportCode(val *[]*string)
	AirportCodeInput() *[]*string
	AlertTriggerPreferences() *[]*string
	SetAlertTriggerPreferences(val *[]*string)
	AlertTriggerPreferencesInput() *[]*string
	AlertTriggerPreferencesValue() *[]*string
	SetAlertTriggerPreferencesValue(val *[]*string)
	AlertTriggerPreferencesValueInput() *[]*string
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
	SetEnabled(val *[]*string)
	EnabledInput() *[]*string
	Environment() *[]*string
	SetEnvironment(val *[]*string)
	EnvironmentInput() *[]*string
	Event() *[]*string
	SetEvent(val *[]*string)
	EventInput() *[]*string
	EventSource() *[]*string
	SetEventSource(val *[]*string)
	EventSourceInput() *[]*string
	EventType() *[]*string
	SetEventType(val *[]*string)
	EventTypeInput() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	HealthCheckId() *[]*string
	SetHealthCheckId(val *[]*string)
	HealthCheckIdInput() *[]*string
	IncidentImpact() *[]*string
	SetIncidentImpact(val *[]*string)
	IncidentImpactInput() *[]*string
	InputId() *[]*string
	SetInputId(val *[]*string)
	InputIdInput() *[]*string
	InsightClass() *[]*string
	SetInsightClass(val *[]*string)
	InsightClassInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *[]*string
	SetLimit(val *[]*string)
	LimitInput() *[]*string
	LogoTag() *[]*string
	SetLogoTag(val *[]*string)
	LogoTagInput() *[]*string
	MegabitsPerSecond() *[]*string
	SetMegabitsPerSecond(val *[]*string)
	MegabitsPerSecondInput() *[]*string
	NewHealth() *[]*string
	SetNewHealth(val *[]*string)
	NewHealthInput() *[]*string
	NewStatus() *[]*string
	SetNewStatus(val *[]*string)
	NewStatusInput() *[]*string
	PacketsPerSecond() *[]*string
	SetPacketsPerSecond(val *[]*string)
	PacketsPerSecondInput() *[]*string
	PoolId() *[]*string
	SetPoolId(val *[]*string)
	PoolIdInput() *[]*string
	PopNames() *[]*string
	SetPopNames(val *[]*string)
	PopNamesInput() *[]*string
	Product() *[]*string
	SetProduct(val *[]*string)
	ProductInput() *[]*string
	ProjectId() *[]*string
	SetProjectId(val *[]*string)
	ProjectIdInput() *[]*string
	Protocol() *[]*string
	SetProtocol(val *[]*string)
	ProtocolInput() *[]*string
	QueryTag() *[]*string
	SetQueryTag(val *[]*string)
	QueryTagInput() *[]*string
	RequestsPerSecond() *[]*string
	SetRequestsPerSecond(val *[]*string)
	RequestsPerSecondInput() *[]*string
	Selectors() *[]*string
	SetSelectors(val *[]*string)
	SelectorsInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
	Slo() *[]*string
	SetSlo(val *[]*string)
	SloInput() *[]*string
	Status() *[]*string
	SetStatus(val *[]*string)
	StatusInput() *[]*string
	TargetHostname() *[]*string
	SetTargetHostname(val *[]*string)
	TargetHostnameInput() *[]*string
	TargetIp() *[]*string
	SetTargetIp(val *[]*string)
	TargetIpInput() *[]*string
	TargetZoneName() *[]*string
	SetTargetZoneName(val *[]*string)
	TargetZoneNameInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficExclusions() *[]*string
	SetTrafficExclusions(val *[]*string)
	TrafficExclusionsInput() *[]*string
	TunnelId() *[]*string
	SetTunnelId(val *[]*string)
	TunnelIdInput() *[]*string
	TunnelName() *[]*string
	SetTunnelName(val *[]*string)
	TunnelNameInput() *[]*string
	Where() *[]*string
	SetWhere(val *[]*string)
	WhereInput() *[]*string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	ResetActions()
	ResetAffectedAsns()
	ResetAffectedComponents()
	ResetAffectedLocations()
	ResetAirportCode()
	ResetAlertTriggerPreferences()
	ResetAlertTriggerPreferencesValue()
	ResetEnabled()
	ResetEnvironment()
	ResetEvent()
	ResetEventSource()
	ResetEventType()
	ResetGroupBy()
	ResetHealthCheckId()
	ResetIncidentImpact()
	ResetInputId()
	ResetInsightClass()
	ResetLimit()
	ResetLogoTag()
	ResetMegabitsPerSecond()
	ResetNewHealth()
	ResetNewStatus()
	ResetPacketsPerSecond()
	ResetPoolId()
	ResetPopNames()
	ResetProduct()
	ResetProjectId()
	ResetProtocol()
	ResetQueryTag()
	ResetRequestsPerSecond()
	ResetSelectors()
	ResetServices()
	ResetSlo()
	ResetStatus()
	ResetTargetHostname()
	ResetTargetIp()
	ResetTargetZoneName()
	ResetTrafficExclusions()
	ResetTunnelId()
	ResetTunnelName()
	ResetWhere()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationPolicyFiltersOutputReference
type jsiiProxy_NotificationPolicyFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedAsns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedAsns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedAsnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedAsnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedComponentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedComponentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AffectedLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"affectedLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AirportCode() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"airportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AirportCodeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"airportCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AlertTriggerPreferences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AlertTriggerPreferencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AlertTriggerPreferencesValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferencesValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) AlertTriggerPreferencesValueInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertTriggerPreferencesValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Enabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Environment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EnvironmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Event() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventSource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventSourceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) IncidentImpact() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"incidentImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) IncidentImpactInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"incidentImpactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InputId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InputIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InsightClass() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insightClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InsightClassInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insightClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Limit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) LimitInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) LogoTag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) LogoTagInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) MegabitsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"megabitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) MegabitsPerSecondInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"megabitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewHealth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewHealthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewStatus() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewStatusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PacketsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packetsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PacketsPerSecondInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packetsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PopNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"popNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PopNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"popNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Product() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProductInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProjectId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProjectIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Protocol() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProtocolInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) QueryTag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) QueryTagInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) RequestsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) RequestsPerSecondInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Selectors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SelectorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Slo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SloInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sloInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Status() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) StatusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetHostname() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetHostnameInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetIp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetIpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetZoneName() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetZoneNameInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TrafficExclusions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trafficExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TrafficExclusionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trafficExclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TunnelId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TunnelIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TunnelName() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TunnelNameInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Where() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"where",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) WhereInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whereInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewNotificationPolicyFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationPolicyFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationPolicyFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationPolicyFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationPolicyFiltersOutputReference_Override(n NotificationPolicyFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAffectedAsns(val *[]*string) {
	if err := j.validateSetAffectedAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affectedAsns",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAffectedComponents(val *[]*string) {
	if err := j.validateSetAffectedComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affectedComponents",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAffectedLocations(val *[]*string) {
	if err := j.validateSetAffectedLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affectedLocations",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAirportCode(val *[]*string) {
	if err := j.validateSetAirportCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airportCode",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAlertTriggerPreferences(val *[]*string) {
	if err := j.validateSetAlertTriggerPreferencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertTriggerPreferences",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetAlertTriggerPreferencesValue(val *[]*string) {
	if err := j.validateSetAlertTriggerPreferencesValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertTriggerPreferencesValue",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEnabled(val *[]*string) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEnvironment(val *[]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEvent(val *[]*string) {
	if err := j.validateSetEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"event",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEventSource(val *[]*string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEventType(val *[]*string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetHealthCheckId(val *[]*string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetIncidentImpact(val *[]*string) {
	if err := j.validateSetIncidentImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incidentImpact",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetInputId(val *[]*string) {
	if err := j.validateSetInputIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetInsightClass(val *[]*string) {
	if err := j.validateSetInsightClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightClass",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetLimit(val *[]*string) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetLogoTag(val *[]*string) {
	if err := j.validateSetLogoTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoTag",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetMegabitsPerSecond(val *[]*string) {
	if err := j.validateSetMegabitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"megabitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetNewHealth(val *[]*string) {
	if err := j.validateSetNewHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newHealth",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetNewStatus(val *[]*string) {
	if err := j.validateSetNewStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newStatus",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetPacketsPerSecond(val *[]*string) {
	if err := j.validateSetPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packetsPerSecond",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetPoolId(val *[]*string) {
	if err := j.validateSetPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetPopNames(val *[]*string) {
	if err := j.validateSetPopNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"popNames",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetProduct(val *[]*string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetProjectId(val *[]*string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetProtocol(val *[]*string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetQueryTag(val *[]*string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetRequestsPerSecond(val *[]*string) {
	if err := j.validateSetRequestsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestsPerSecond",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetSelectors(val *[]*string) {
	if err := j.validateSetSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectors",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetServices(val *[]*string) {
	if err := j.validateSetServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetSlo(val *[]*string) {
	if err := j.validateSetSloParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slo",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetStatus(val *[]*string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTargetHostname(val *[]*string) {
	if err := j.validateSetTargetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetHostname",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTargetIp(val *[]*string) {
	if err := j.validateSetTargetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIp",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTargetZoneName(val *[]*string) {
	if err := j.validateSetTargetZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetZoneName",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTrafficExclusions(val *[]*string) {
	if err := j.validateSetTrafficExclusionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficExclusions",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTunnelId(val *[]*string) {
	if err := j.validateSetTunnelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTunnelName(val *[]*string) {
	if err := j.validateSetTunnelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelName",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetWhere(val *[]*string) {
	if err := j.validateSetWhereParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"where",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		n,
		"resetActions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAffectedAsns() {
	_jsii_.InvokeVoid(
		n,
		"resetAffectedAsns",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAffectedComponents() {
	_jsii_.InvokeVoid(
		n,
		"resetAffectedComponents",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAffectedLocations() {
	_jsii_.InvokeVoid(
		n,
		"resetAffectedLocations",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAirportCode() {
	_jsii_.InvokeVoid(
		n,
		"resetAirportCode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAlertTriggerPreferences() {
	_jsii_.InvokeVoid(
		n,
		"resetAlertTriggerPreferences",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetAlertTriggerPreferencesValue() {
	_jsii_.InvokeVoid(
		n,
		"resetAlertTriggerPreferencesValue",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		n,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		n,
		"resetEvent",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		n,
		"resetEventSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEventType() {
	_jsii_.InvokeVoid(
		n,
		"resetEventType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		n,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		n,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetIncidentImpact() {
	_jsii_.InvokeVoid(
		n,
		"resetIncidentImpact",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetInputId() {
	_jsii_.InvokeVoid(
		n,
		"resetInputId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetInsightClass() {
	_jsii_.InvokeVoid(
		n,
		"resetInsightClass",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		n,
		"resetLimit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetLogoTag() {
	_jsii_.InvokeVoid(
		n,
		"resetLogoTag",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetMegabitsPerSecond() {
	_jsii_.InvokeVoid(
		n,
		"resetMegabitsPerSecond",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetNewHealth() {
	_jsii_.InvokeVoid(
		n,
		"resetNewHealth",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetNewStatus() {
	_jsii_.InvokeVoid(
		n,
		"resetNewStatus",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPacketsPerSecond() {
	_jsii_.InvokeVoid(
		n,
		"resetPacketsPerSecond",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPoolId() {
	_jsii_.InvokeVoid(
		n,
		"resetPoolId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPopNames() {
	_jsii_.InvokeVoid(
		n,
		"resetPopNames",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProduct() {
	_jsii_.InvokeVoid(
		n,
		"resetProduct",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		n,
		"resetProjectId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetQueryTag() {
	_jsii_.InvokeVoid(
		n,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetRequestsPerSecond() {
	_jsii_.InvokeVoid(
		n,
		"resetRequestsPerSecond",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetSelectors() {
	_jsii_.InvokeVoid(
		n,
		"resetSelectors",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		n,
		"resetServices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetSlo() {
	_jsii_.InvokeVoid(
		n,
		"resetSlo",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		n,
		"resetStatus",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTargetHostname() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetHostname",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTargetIp() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTargetZoneName() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetZoneName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTrafficExclusions() {
	_jsii_.InvokeVoid(
		n,
		"resetTrafficExclusions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTunnelId() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTunnelName() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetWhere() {
	_jsii_.InvokeVoid(
		n,
		"resetWhere",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		n,
		"resetZones",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


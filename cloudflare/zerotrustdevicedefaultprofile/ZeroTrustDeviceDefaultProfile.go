// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustdevicedefaultprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_default_profile cloudflare_zero_trust_device_default_profile}.
type ZeroTrustDeviceDefaultProfile interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedToLeave() interface{}
	SetAllowedToLeave(val interface{})
	AllowedToLeaveInput() interface{}
	AllowModeSwitch() interface{}
	SetAllowModeSwitch(val interface{})
	AllowModeSwitchInput() interface{}
	AllowUpdates() interface{}
	SetAllowUpdates(val interface{})
	AllowUpdatesInput() interface{}
	AutoConnect() *float64
	SetAutoConnect(val *float64)
	AutoConnectInput() *float64
	CaptivePortal() *float64
	SetCaptivePortal(val *float64)
	CaptivePortalInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Default() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableAutoFallback() interface{}
	SetDisableAutoFallback(val interface{})
	DisableAutoFallbackInput() interface{}
	Enabled() cdktf.IResolvable
	Exclude() ZeroTrustDeviceDefaultProfileExcludeList
	ExcludeInput() interface{}
	ExcludeOfficeIps() interface{}
	SetExcludeOfficeIps(val interface{})
	ExcludeOfficeIpsInput() interface{}
	FallbackDomains() ZeroTrustDeviceDefaultProfileFallbackDomainsList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayUniqueId() *string
	Id() *string
	Include() ZeroTrustDeviceDefaultProfileIncludeList
	IncludeInput() interface{}
	LanAllowMinutes() *float64
	SetLanAllowMinutes(val *float64)
	LanAllowMinutesInput() *float64
	LanAllowSubnetSize() *float64
	SetLanAllowSubnetSize(val *float64)
	LanAllowSubnetSizeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RegisterInterfaceIpWithDns() interface{}
	SetRegisterInterfaceIpWithDns(val interface{})
	RegisterInterfaceIpWithDnsInput() interface{}
	SccmVpnBoundarySupport() interface{}
	SetSccmVpnBoundarySupport(val interface{})
	SccmVpnBoundarySupportInput() interface{}
	ServiceModeV2() ZeroTrustDeviceDefaultProfileServiceModeV2OutputReference
	ServiceModeV2Input() interface{}
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	SwitchLocked() interface{}
	SetSwitchLocked(val interface{})
	SwitchLockedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TunnelProtocol() *string
	SetTunnelProtocol(val *string)
	TunnelProtocolInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutExclude(value interface{})
	PutInclude(value interface{})
	PutServiceModeV2(value *ZeroTrustDeviceDefaultProfileServiceModeV2)
	ResetAllowedToLeave()
	ResetAllowModeSwitch()
	ResetAllowUpdates()
	ResetAutoConnect()
	ResetCaptivePortal()
	ResetDisableAutoFallback()
	ResetExclude()
	ResetExcludeOfficeIps()
	ResetInclude()
	ResetLanAllowMinutes()
	ResetLanAllowSubnetSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegisterInterfaceIpWithDns()
	ResetSccmVpnBoundarySupport()
	ResetServiceModeV2()
	ResetSupportUrl()
	ResetSwitchLocked()
	ResetTunnelProtocol()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ZeroTrustDeviceDefaultProfile
type jsiiProxy_ZeroTrustDeviceDefaultProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowedToLeave() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowedToLeaveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToLeaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowModeSwitch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowModeSwitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowModeSwitchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowModeSwitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AllowUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AutoConnect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) AutoConnectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) CaptivePortal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) CaptivePortalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Default() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) DisableAutoFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) DisableAutoFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Exclude() ZeroTrustDeviceDefaultProfileExcludeList {
	var returns ZeroTrustDeviceDefaultProfileExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ExcludeOfficeIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeOfficeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ExcludeOfficeIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeOfficeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) FallbackDomains() ZeroTrustDeviceDefaultProfileFallbackDomainsList {
	var returns ZeroTrustDeviceDefaultProfileFallbackDomainsList
	_jsii_.Get(
		j,
		"fallbackDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) GatewayUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Include() ZeroTrustDeviceDefaultProfileIncludeList {
	var returns ZeroTrustDeviceDefaultProfileIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) IncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) LanAllowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) LanAllowMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) LanAllowSubnetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowSubnetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) LanAllowSubnetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowSubnetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) RegisterInterfaceIpWithDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registerInterfaceIpWithDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) RegisterInterfaceIpWithDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registerInterfaceIpWithDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SccmVpnBoundarySupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sccmVpnBoundarySupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SccmVpnBoundarySupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sccmVpnBoundarySupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ServiceModeV2() ZeroTrustDeviceDefaultProfileServiceModeV2OutputReference {
	var returns ZeroTrustDeviceDefaultProfileServiceModeV2OutputReference
	_jsii_.Get(
		j,
		"serviceModeV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) ServiceModeV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceModeV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SwitchLocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) SwitchLockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) TunnelProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile) TunnelProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocolInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_default_profile cloudflare_zero_trust_device_default_profile} Resource.
func NewZeroTrustDeviceDefaultProfile(scope constructs.Construct, id *string, config *ZeroTrustDeviceDefaultProfileConfig) ZeroTrustDeviceDefaultProfile {
	_init_.Initialize()

	if err := validateNewZeroTrustDeviceDefaultProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDeviceDefaultProfile{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_device_default_profile cloudflare_zero_trust_device_default_profile} Resource.
func NewZeroTrustDeviceDefaultProfile_Override(z ZeroTrustDeviceDefaultProfile, scope constructs.Construct, id *string, config *ZeroTrustDeviceDefaultProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetAllowedToLeave(val interface{}) {
	if err := j.validateSetAllowedToLeaveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToLeave",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetAllowModeSwitch(val interface{}) {
	if err := j.validateSetAllowModeSwitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowModeSwitch",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetAllowUpdates(val interface{}) {
	if err := j.validateSetAllowUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUpdates",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetAutoConnect(val *float64) {
	if err := j.validateSetAutoConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConnect",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetCaptivePortal(val *float64) {
	if err := j.validateSetCaptivePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captivePortal",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetDisableAutoFallback(val interface{}) {
	if err := j.validateSetDisableAutoFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoFallback",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetExcludeOfficeIps(val interface{}) {
	if err := j.validateSetExcludeOfficeIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOfficeIps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetLanAllowMinutes(val *float64) {
	if err := j.validateSetLanAllowMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanAllowMinutes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetLanAllowSubnetSize(val *float64) {
	if err := j.validateSetLanAllowSubnetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanAllowSubnetSize",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetRegisterInterfaceIpWithDns(val interface{}) {
	if err := j.validateSetRegisterInterfaceIpWithDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registerInterfaceIpWithDns",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetSccmVpnBoundarySupport(val interface{}) {
	if err := j.validateSetSccmVpnBoundarySupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sccmVpnBoundarySupport",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetSwitchLocked(val interface{}) {
	if err := j.validateSetSwitchLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchLocked",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceDefaultProfile)SetTunnelProtocol(val *string) {
	if err := j.validateSetTunnelProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelProtocol",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDeviceDefaultProfile resource upon running "cdktf plan <stack-name>".
func ZeroTrustDeviceDefaultProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDeviceDefaultProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ZeroTrustDeviceDefaultProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceDefaultProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceDefaultProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceDefaultProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceDefaultProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceDefaultProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustDeviceDefaultProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDeviceDefaultProfile.ZeroTrustDeviceDefaultProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) PutExclude(value interface{}) {
	if err := z.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExclude",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) PutInclude(value interface{}) {
	if err := z.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInclude",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) PutServiceModeV2(value *ZeroTrustDeviceDefaultProfileServiceModeV2) {
	if err := z.validatePutServiceModeV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceModeV2",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetAllowedToLeave() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedToLeave",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetAllowModeSwitch() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowModeSwitch",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetAllowUpdates() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowUpdates",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetAutoConnect() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoConnect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetCaptivePortal() {
	_jsii_.InvokeVoid(
		z,
		"resetCaptivePortal",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetDisableAutoFallback() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableAutoFallback",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetExcludeOfficeIps() {
	_jsii_.InvokeVoid(
		z,
		"resetExcludeOfficeIps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetInclude() {
	_jsii_.InvokeVoid(
		z,
		"resetInclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetLanAllowMinutes() {
	_jsii_.InvokeVoid(
		z,
		"resetLanAllowMinutes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetLanAllowSubnetSize() {
	_jsii_.InvokeVoid(
		z,
		"resetLanAllowSubnetSize",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetRegisterInterfaceIpWithDns() {
	_jsii_.InvokeVoid(
		z,
		"resetRegisterInterfaceIpWithDns",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetSccmVpnBoundarySupport() {
	_jsii_.InvokeVoid(
		z,
		"resetSccmVpnBoundarySupport",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetServiceModeV2() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceModeV2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetSwitchLocked() {
	_jsii_.InvokeVoid(
		z,
		"resetSwitchLocked",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ResetTunnelProtocol() {
	_jsii_.InvokeVoid(
		z,
		"resetTunnelProtocol",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceDefaultProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


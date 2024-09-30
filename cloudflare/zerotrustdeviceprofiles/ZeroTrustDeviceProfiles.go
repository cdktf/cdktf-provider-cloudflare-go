// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustdeviceprofiles/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_device_profiles cloudflare_zero_trust_device_profiles}.
type ZeroTrustDeviceProfiles interface {
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
	Default() interface{}
	SetDefault(val interface{})
	DefaultInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutoFallback() interface{}
	SetDisableAutoFallback(val interface{})
	DisableAutoFallbackInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludeOfficeIps() interface{}
	SetExcludeOfficeIps(val interface{})
	ExcludeOfficeIpsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
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
	ServiceModeV2Mode() *string
	SetServiceModeV2Mode(val *string)
	ServiceModeV2ModeInput() *string
	ServiceModeV2Port() *float64
	SetServiceModeV2Port(val *float64)
	ServiceModeV2PortInput() *float64
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
	ResetAllowedToLeave()
	ResetAllowModeSwitch()
	ResetAllowUpdates()
	ResetAutoConnect()
	ResetCaptivePortal()
	ResetDefault()
	ResetDisableAutoFallback()
	ResetEnabled()
	ResetExcludeOfficeIps()
	ResetId()
	ResetMatch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrecedence()
	ResetServiceModeV2Mode()
	ResetServiceModeV2Port()
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

// The jsii proxy struct for ZeroTrustDeviceProfiles
type jsiiProxy_ZeroTrustDeviceProfiles struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowedToLeave() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowedToLeaveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToLeaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowModeSwitch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowModeSwitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowModeSwitchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowModeSwitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AllowUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AutoConnect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) AutoConnectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) CaptivePortal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) CaptivePortalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) DefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) DisableAutoFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) DisableAutoFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ExcludeOfficeIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeOfficeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ExcludeOfficeIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeOfficeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ServiceModeV2Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeV2Mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ServiceModeV2ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeV2ModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ServiceModeV2Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceModeV2Port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) ServiceModeV2PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceModeV2PortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) SwitchLocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) SwitchLockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) TunnelProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles) TunnelProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocolInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_device_profiles cloudflare_zero_trust_device_profiles} Resource.
func NewZeroTrustDeviceProfiles(scope constructs.Construct, id *string, config *ZeroTrustDeviceProfilesConfig) ZeroTrustDeviceProfiles {
	_init_.Initialize()

	if err := validateNewZeroTrustDeviceProfilesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDeviceProfiles{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_device_profiles cloudflare_zero_trust_device_profiles} Resource.
func NewZeroTrustDeviceProfiles_Override(z ZeroTrustDeviceProfiles, scope constructs.Construct, id *string, config *ZeroTrustDeviceProfilesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetAllowedToLeave(val interface{}) {
	if err := j.validateSetAllowedToLeaveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToLeave",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetAllowModeSwitch(val interface{}) {
	if err := j.validateSetAllowModeSwitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowModeSwitch",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetAllowUpdates(val interface{}) {
	if err := j.validateSetAllowUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUpdates",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetAutoConnect(val *float64) {
	if err := j.validateSetAutoConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConnect",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetCaptivePortal(val *float64) {
	if err := j.validateSetCaptivePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captivePortal",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetDefault(val interface{}) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetDisableAutoFallback(val interface{}) {
	if err := j.validateSetDisableAutoFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoFallback",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetExcludeOfficeIps(val interface{}) {
	if err := j.validateSetExcludeOfficeIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOfficeIps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetServiceModeV2Mode(val *string) {
	if err := j.validateSetServiceModeV2ModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceModeV2Mode",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetServiceModeV2Port(val *float64) {
	if err := j.validateSetServiceModeV2PortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceModeV2Port",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetSwitchLocked(val interface{}) {
	if err := j.validateSetSwitchLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchLocked",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceProfiles)SetTunnelProtocol(val *string) {
	if err := j.validateSetTunnelProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelProtocol",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDeviceProfiles resource upon running "cdktf plan <stack-name>".
func ZeroTrustDeviceProfiles_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDeviceProfiles_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
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
func ZeroTrustDeviceProfiles_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceProfiles_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceProfiles_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceProfiles_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceProfiles_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceProfiles_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustDeviceProfiles_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDeviceProfiles.ZeroTrustDeviceProfiles",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDeviceProfiles) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetAllowedToLeave() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedToLeave",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetAllowModeSwitch() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowModeSwitch",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetAllowUpdates() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowUpdates",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetAutoConnect() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoConnect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetCaptivePortal() {
	_jsii_.InvokeVoid(
		z,
		"resetCaptivePortal",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetDefault() {
	_jsii_.InvokeVoid(
		z,
		"resetDefault",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetDisableAutoFallback() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableAutoFallback",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetExcludeOfficeIps() {
	_jsii_.InvokeVoid(
		z,
		"resetExcludeOfficeIps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetMatch() {
	_jsii_.InvokeVoid(
		z,
		"resetMatch",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetPrecedence() {
	_jsii_.InvokeVoid(
		z,
		"resetPrecedence",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetServiceModeV2Mode() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceModeV2Mode",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetServiceModeV2Port() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceModeV2Port",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetSwitchLocked() {
	_jsii_.InvokeVoid(
		z,
		"resetSwitchLocked",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ResetTunnelProtocol() {
	_jsii_.InvokeVoid(
		z,
		"resetTunnelProtocol",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceProfiles) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


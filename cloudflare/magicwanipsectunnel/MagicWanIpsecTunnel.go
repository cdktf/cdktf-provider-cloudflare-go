// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanipsectunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/magicwanipsectunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/magic_wan_ipsec_tunnel cloudflare_magic_wan_ipsec_tunnel}.
type MagicWanIpsecTunnel interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowNullCipher() interface{}
	SetAllowNullCipher(val interface{})
	AllowNullCipherInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudflareEndpoint() *string
	SetCloudflareEndpoint(val *string)
	CloudflareEndpointInput() *string
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
	CustomerEndpoint() *string
	SetCustomerEndpoint(val *string)
	CustomerEndpointInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FqdnId() *string
	SetFqdnId(val *string)
	FqdnIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckDirection() *string
	SetHealthCheckDirection(val *string)
	HealthCheckDirectionInput() *string
	HealthCheckEnabled() interface{}
	SetHealthCheckEnabled(val interface{})
	HealthCheckEnabledInput() interface{}
	HealthCheckRate() *string
	SetHealthCheckRate(val *string)
	HealthCheckRateInput() *string
	HealthCheckTarget() *string
	SetHealthCheckTarget(val *string)
	HealthCheckTargetInput() *string
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	HexId() *string
	SetHexId(val *string)
	HexIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterfaceAddress() *string
	SetInterfaceAddress(val *string)
	InterfaceAddressInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Psk() *string
	SetPsk(val *string)
	PskInput() *string
	// Experimental.
	RawOverrides() interface{}
	RemoteId() *string
	SetRemoteId(val *string)
	RemoteIdInput() *string
	ReplayProtection() interface{}
	SetReplayProtection(val interface{})
	ReplayProtectionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
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
	ResetAccountId()
	ResetAllowNullCipher()
	ResetDescription()
	ResetFqdnId()
	ResetHealthCheckDirection()
	ResetHealthCheckEnabled()
	ResetHealthCheckRate()
	ResetHealthCheckTarget()
	ResetHealthCheckType()
	ResetHexId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPsk()
	ResetRemoteId()
	ResetReplayProtection()
	ResetUserId()
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

// The jsii proxy struct for MagicWanIpsecTunnel
type jsiiProxy_MagicWanIpsecTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MagicWanIpsecTunnel) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) AllowNullCipher() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNullCipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) AllowNullCipherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNullCipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) CloudflareEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) CloudflareEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) CustomerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) CustomerEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) FqdnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) FqdnIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) HexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) InterfaceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) InterfaceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) Psk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) PskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) RemoteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) RemoteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) ReplayProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replayProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) ReplayProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replayProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicWanIpsecTunnel) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/magic_wan_ipsec_tunnel cloudflare_magic_wan_ipsec_tunnel} Resource.
func NewMagicWanIpsecTunnel(scope constructs.Construct, id *string, config *MagicWanIpsecTunnelConfig) MagicWanIpsecTunnel {
	_init_.Initialize()

	if err := validateNewMagicWanIpsecTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicWanIpsecTunnel{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/magic_wan_ipsec_tunnel cloudflare_magic_wan_ipsec_tunnel} Resource.
func NewMagicWanIpsecTunnel_Override(m MagicWanIpsecTunnel, scope constructs.Construct, id *string, config *MagicWanIpsecTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetAllowNullCipher(val interface{}) {
	if err := j.validateSetAllowNullCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowNullCipher",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetCloudflareEndpoint(val *string) {
	if err := j.validateSetCloudflareEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudflareEndpoint",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetCustomerEndpoint(val *string) {
	if err := j.validateSetCustomerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerEndpoint",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetFqdnId(val *string) {
	if err := j.validateSetFqdnIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdnId",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHealthCheckDirection(val *string) {
	if err := j.validateSetHealthCheckDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckDirection",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHealthCheckEnabled(val interface{}) {
	if err := j.validateSetHealthCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHealthCheckRate(val *string) {
	if err := j.validateSetHealthCheckRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckRate",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHealthCheckTarget(val *string) {
	if err := j.validateSetHealthCheckTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckTarget",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetHexId(val *string) {
	if err := j.validateSetHexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hexId",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetInterfaceAddress(val *string) {
	if err := j.validateSetInterfaceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceAddress",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetPsk(val *string) {
	if err := j.validateSetPskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"psk",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetRemoteId(val *string) {
	if err := j.validateSetRemoteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteId",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetReplayProtection(val interface{}) {
	if err := j.validateSetReplayProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replayProtection",
		val,
	)
}

func (j *jsiiProxy_MagicWanIpsecTunnel)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Generates CDKTF code for importing a MagicWanIpsecTunnel resource upon running "cdktf plan <stack-name>".
func MagicWanIpsecTunnel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMagicWanIpsecTunnel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
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
func MagicWanIpsecTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMagicWanIpsecTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MagicWanIpsecTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMagicWanIpsecTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MagicWanIpsecTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMagicWanIpsecTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MagicWanIpsecTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.magicWanIpsecTunnel.MagicWanIpsecTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetAllowNullCipher() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowNullCipher",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetFqdnId() {
	_jsii_.InvokeVoid(
		m,
		"resetFqdnId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHealthCheckDirection() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckDirection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHealthCheckEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHealthCheckRate() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHealthCheckTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckTarget",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		m,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetHexId() {
	_jsii_.InvokeVoid(
		m,
		"resetHexId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetPsk() {
	_jsii_.InvokeVoid(
		m,
		"resetPsk",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetRemoteId() {
	_jsii_.InvokeVoid(
		m,
		"resetRemoteId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetReplayProtection() {
	_jsii_.InvokeVoid(
		m,
		"resetReplayProtection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ResetUserId() {
	_jsii_.InvokeVoid(
		m,
		"resetUserId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicWanIpsecTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicWanIpsecTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustdnslocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_dns_location cloudflare_zero_trust_dns_location}.
type ZeroTrustDnsLocation interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientDefault() interface{}
	SetClientDefault(val interface{})
	ClientDefaultInput() interface{}
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsDestinationIpsId() *string
	SetDnsDestinationIpsId(val *string)
	DnsDestinationIpsIdInput() *string
	DnsDestinationIpv6BlockId() *string
	DohSubdomain() *string
	EcsSupport() interface{}
	SetEcsSupport(val interface{})
	EcsSupportInput() interface{}
	Endpoints() ZeroTrustDnsLocationEndpointsOutputReference
	EndpointsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Ip() *string
	Ipv4Destination() *string
	Ipv4DestinationBackup() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Networks() ZeroTrustDnsLocationNetworksList
	NetworksInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
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
	PutEndpoints(value *ZeroTrustDnsLocationEndpoints)
	PutNetworks(value interface{})
	ResetClientDefault()
	ResetDnsDestinationIpsId()
	ResetEcsSupport()
	ResetEndpoints()
	ResetNetworks()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ZeroTrustDnsLocation
type jsiiProxy_ZeroTrustDnsLocation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDnsLocation) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) ClientDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) ClientDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) DnsDestinationIpsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDestinationIpsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) DnsDestinationIpsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDestinationIpsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) DnsDestinationIpv6BlockId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDestinationIpv6BlockId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) DohSubdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dohSubdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) EcsSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) EcsSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Endpoints() ZeroTrustDnsLocationEndpointsOutputReference {
	var returns ZeroTrustDnsLocationEndpointsOutputReference
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Ipv4Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Ipv4DestinationBackup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4DestinationBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Networks() ZeroTrustDnsLocationNetworksList {
	var returns ZeroTrustDnsLocationNetworksList
	_jsii_.Get(
		j,
		"networks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) NetworksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDnsLocation) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_dns_location cloudflare_zero_trust_dns_location} Resource.
func NewZeroTrustDnsLocation(scope constructs.Construct, id *string, config *ZeroTrustDnsLocationConfig) ZeroTrustDnsLocation {
	_init_.Initialize()

	if err := validateNewZeroTrustDnsLocationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDnsLocation{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_dns_location cloudflare_zero_trust_dns_location} Resource.
func NewZeroTrustDnsLocation_Override(z ZeroTrustDnsLocation, scope constructs.Construct, id *string, config *ZeroTrustDnsLocationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetClientDefault(val interface{}) {
	if err := j.validateSetClientDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientDefault",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetDnsDestinationIpsId(val *string) {
	if err := j.validateSetDnsDestinationIpsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsDestinationIpsId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetEcsSupport(val interface{}) {
	if err := j.validateSetEcsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsSupport",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDnsLocation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDnsLocation resource upon running "cdktf plan <stack-name>".
func ZeroTrustDnsLocation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDnsLocation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
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
func ZeroTrustDnsLocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDnsLocation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDnsLocation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDnsLocation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDnsLocation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDnsLocation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustDnsLocation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDnsLocation.ZeroTrustDnsLocation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDnsLocation) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) PutEndpoints(value *ZeroTrustDnsLocationEndpoints) {
	if err := z.validatePutEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEndpoints",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) PutNetworks(value interface{}) {
	if err := z.validatePutNetworksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putNetworks",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetClientDefault() {
	_jsii_.InvokeVoid(
		z,
		"resetClientDefault",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetDnsDestinationIpsId() {
	_jsii_.InvokeVoid(
		z,
		"resetDnsDestinationIpsId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetEcsSupport() {
	_jsii_.InvokeVoid(
		z,
		"resetEcsSupport",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetEndpoints() {
	_jsii_.InvokeVoid(
		z,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetNetworks() {
	_jsii_.InvokeVoid(
		z,
		"resetNetworks",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDnsLocation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDnsLocation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


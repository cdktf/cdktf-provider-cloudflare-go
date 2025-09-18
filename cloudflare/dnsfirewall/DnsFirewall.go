// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/dnsfirewall/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/dns_firewall cloudflare_dns_firewall}.
type DnsFirewall interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AttackMitigation() DnsFirewallAttackMitigationOutputReference
	AttackMitigationInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeprecateAnyRequests() interface{}
	SetDeprecateAnyRequests(val interface{})
	DeprecateAnyRequestsInput() interface{}
	DnsFirewallIps() *[]*string
	EcsFallback() interface{}
	SetEcsFallback(val interface{})
	EcsFallbackInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumCacheTtl() *float64
	SetMaximumCacheTtl(val *float64)
	MaximumCacheTtlInput() *float64
	MinimumCacheTtl() *float64
	SetMinimumCacheTtl(val *float64)
	MinimumCacheTtlInput() *float64
	ModifiedOn() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NegativeCacheTtl() *float64
	SetNegativeCacheTtl(val *float64)
	NegativeCacheTtlInput() *float64
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
	Ratelimit() *float64
	SetRatelimit(val *float64)
	RatelimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpstreamIps() *[]*string
	SetUpstreamIps(val *[]*string)
	UpstreamIpsInput() *[]*string
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
	PutAttackMitigation(value *DnsFirewallAttackMitigation)
	ResetAttackMitigation()
	ResetDeprecateAnyRequests()
	ResetEcsFallback()
	ResetMaximumCacheTtl()
	ResetMinimumCacheTtl()
	ResetNegativeCacheTtl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRatelimit()
	ResetRetries()
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

// The jsii proxy struct for DnsFirewall
type jsiiProxy_DnsFirewall struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DnsFirewall) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) AttackMitigation() DnsFirewallAttackMitigationOutputReference {
	var returns DnsFirewallAttackMitigationOutputReference
	_jsii_.Get(
		j,
		"attackMitigation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) AttackMitigationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attackMitigationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) DeprecateAnyRequests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecateAnyRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) DeprecateAnyRequestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecateAnyRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) DnsFirewallIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsFirewallIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) EcsFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) EcsFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) MaximumCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) MaximumCacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) MinimumCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) MinimumCacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) NegativeCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"negativeCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) NegativeCacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"negativeCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Ratelimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ratelimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) RatelimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ratelimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) UpstreamIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"upstreamIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsFirewall) UpstreamIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"upstreamIpsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/dns_firewall cloudflare_dns_firewall} Resource.
func NewDnsFirewall(scope constructs.Construct, id *string, config *DnsFirewallConfig) DnsFirewall {
	_init_.Initialize()

	if err := validateNewDnsFirewallParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsFirewall{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/dns_firewall cloudflare_dns_firewall} Resource.
func NewDnsFirewall_Override(d DnsFirewall, scope constructs.Construct, id *string, config *DnsFirewallConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DnsFirewall)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetDeprecateAnyRequests(val interface{}) {
	if err := j.validateSetDeprecateAnyRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprecateAnyRequests",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetEcsFallback(val interface{}) {
	if err := j.validateSetEcsFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsFallback",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetMaximumCacheTtl(val *float64) {
	if err := j.validateSetMaximumCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCacheTtl",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetMinimumCacheTtl(val *float64) {
	if err := j.validateSetMinimumCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumCacheTtl",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetNegativeCacheTtl(val *float64) {
	if err := j.validateSetNegativeCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCacheTtl",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetRatelimit(val *float64) {
	if err := j.validateSetRatelimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ratelimit",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_DnsFirewall)SetUpstreamIps(val *[]*string) {
	if err := j.validateSetUpstreamIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upstreamIps",
		val,
	)
}

// Generates CDKTF code for importing a DnsFirewall resource upon running "cdktf plan <stack-name>".
func DnsFirewall_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDnsFirewall_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
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
func DnsFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsFirewall_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsFirewall_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsFirewall_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsFirewall_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DnsFirewall_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dnsFirewall.DnsFirewall",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DnsFirewall) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DnsFirewall) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DnsFirewall) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsFirewall) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsFirewall) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsFirewall) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsFirewall) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsFirewall) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsFirewall) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsFirewall) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsFirewall) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsFirewall) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DnsFirewall) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DnsFirewall) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DnsFirewall) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DnsFirewall) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DnsFirewall) PutAttackMitigation(value *DnsFirewallAttackMitigation) {
	if err := d.validatePutAttackMitigationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttackMitigation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsFirewall) ResetAttackMitigation() {
	_jsii_.InvokeVoid(
		d,
		"resetAttackMitigation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetDeprecateAnyRequests() {
	_jsii_.InvokeVoid(
		d,
		"resetDeprecateAnyRequests",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetEcsFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetEcsFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetMaximumCacheTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetMaximumCacheTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetMinimumCacheTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetMinimumCacheTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetNegativeCacheTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetNegativeCacheTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetRatelimit() {
	_jsii_.InvokeVoid(
		d,
		"resetRatelimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) ResetRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsFirewall) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsFirewall) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


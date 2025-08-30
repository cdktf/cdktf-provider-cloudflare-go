// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonednssettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zonednssettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zone_dns_settings cloudflare_zone_dns_settings}.
type ZoneDnsSettings interface {
	cdktf.TerraformResource
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
	FlattenAllCnames() interface{}
	SetFlattenAllCnames(val interface{})
	FlattenAllCnamesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FoundationDns() interface{}
	SetFoundationDns(val interface{})
	FoundationDnsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	InternalDns() ZoneDnsSettingsInternalDnsOutputReference
	InternalDnsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiProvider() interface{}
	SetMultiProvider(val interface{})
	MultiProviderInput() interface{}
	Nameservers() ZoneDnsSettingsNameserversOutputReference
	NameserversInput() interface{}
	// The tree node.
	Node() constructs.Node
	NsTtl() *float64
	SetNsTtl(val *float64)
	NsTtlInput() *float64
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
	SecondaryOverrides() interface{}
	SetSecondaryOverrides(val interface{})
	SecondaryOverridesInput() interface{}
	Soa() ZoneDnsSettingsSoaOutputReference
	SoaInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
	ZoneMode() *string
	SetZoneMode(val *string)
	ZoneModeInput() *string
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
	PutInternalDns(value *ZoneDnsSettingsInternalDns)
	PutNameservers(value *ZoneDnsSettingsNameservers)
	PutSoa(value *ZoneDnsSettingsSoa)
	ResetFlattenAllCnames()
	ResetFoundationDns()
	ResetInternalDns()
	ResetMultiProvider()
	ResetNameservers()
	ResetNsTtl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecondaryOverrides()
	ResetSoa()
	ResetZoneMode()
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

// The jsii proxy struct for ZoneDnsSettings
type jsiiProxy_ZoneDnsSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZoneDnsSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) FlattenAllCnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) FlattenAllCnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenAllCnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) FoundationDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) FoundationDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foundationDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) InternalDns() ZoneDnsSettingsInternalDnsOutputReference {
	var returns ZoneDnsSettingsInternalDnsOutputReference
	_jsii_.Get(
		j,
		"internalDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) InternalDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) MultiProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) MultiProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Nameservers() ZoneDnsSettingsNameserversOutputReference {
	var returns ZoneDnsSettingsNameserversOutputReference
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) NameserversInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) NsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) NsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) SecondaryOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) SecondaryOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) Soa() ZoneDnsSettingsSoaOutputReference {
	var returns ZoneDnsSettingsSoaOutputReference
	_jsii_.Get(
		j,
		"soa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) SoaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"soaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneDnsSettings) ZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneModeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zone_dns_settings cloudflare_zone_dns_settings} Resource.
func NewZoneDnsSettings(scope constructs.Construct, id *string, config *ZoneDnsSettingsConfig) ZoneDnsSettings {
	_init_.Initialize()

	if err := validateNewZoneDnsSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneDnsSettings{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zone_dns_settings cloudflare_zone_dns_settings} Resource.
func NewZoneDnsSettings_Override(z ZoneDnsSettings, scope constructs.Construct, id *string, config *ZoneDnsSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetFlattenAllCnames(val interface{}) {
	if err := j.validateSetFlattenAllCnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenAllCnames",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetFoundationDns(val interface{}) {
	if err := j.validateSetFoundationDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationDns",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetMultiProvider(val interface{}) {
	if err := j.validateSetMultiProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiProvider",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetNsTtl(val *float64) {
	if err := j.validateSetNsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsTtl",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetSecondaryOverrides(val interface{}) {
	if err := j.validateSetSecondaryOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryOverrides",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (j *jsiiProxy_ZoneDnsSettings)SetZoneMode(val *string) {
	if err := j.validateSetZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMode",
		val,
	)
}

// Generates CDKTF code for importing a ZoneDnsSettings resource upon running "cdktf plan <stack-name>".
func ZoneDnsSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZoneDnsSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
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
func ZoneDnsSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneDnsSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZoneDnsSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneDnsSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZoneDnsSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneDnsSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZoneDnsSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zoneDnsSettings.ZoneDnsSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZoneDnsSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneDnsSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZoneDnsSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZoneDnsSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZoneDnsSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZoneDnsSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZoneDnsSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZoneDnsSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZoneDnsSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneDnsSettings) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) PutInternalDns(value *ZoneDnsSettingsInternalDns) {
	if err := z.validatePutInternalDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInternalDns",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) PutNameservers(value *ZoneDnsSettingsNameservers) {
	if err := z.validatePutNameserversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putNameservers",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) PutSoa(value *ZoneDnsSettingsSoa) {
	if err := z.validatePutSoaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSoa",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetFlattenAllCnames() {
	_jsii_.InvokeVoid(
		z,
		"resetFlattenAllCnames",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetFoundationDns() {
	_jsii_.InvokeVoid(
		z,
		"resetFoundationDns",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetInternalDns() {
	_jsii_.InvokeVoid(
		z,
		"resetInternalDns",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetMultiProvider() {
	_jsii_.InvokeVoid(
		z,
		"resetMultiProvider",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetNameservers() {
	_jsii_.InvokeVoid(
		z,
		"resetNameservers",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetNsTtl() {
	_jsii_.InvokeVoid(
		z,
		"resetNsTtl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetSecondaryOverrides() {
	_jsii_.InvokeVoid(
		z,
		"resetSecondaryOverrides",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetSoa() {
	_jsii_.InvokeVoid(
		z,
		"resetSoa",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) ResetZoneMode() {
	_jsii_.InvokeVoid(
		z,
		"resetZoneMode",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneDnsSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneDnsSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


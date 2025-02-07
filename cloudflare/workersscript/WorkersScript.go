// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/workersscript/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script cloudflare_workers_script}.
type WorkersScript interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Assets() WorkersScriptAssetsOutputReference
	AssetsInput() interface{}
	Bindings() WorkersScriptBindingsList
	BindingsInput() interface{}
	BodyPart() *string
	SetBodyPart(val *string)
	BodyPartInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedOn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasAssets() cdktf.IResolvable
	HasModules() cdktf.IResolvable
	Id() *string
	KeepAssets() interface{}
	SetKeepAssets(val interface{})
	KeepAssetsInput() interface{}
	KeepBindings() *[]*string
	SetKeepBindings(val *[]*string)
	KeepBindingsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logpush() cdktf.IResolvable
	MainModule() *string
	SetMainModule(val *string)
	MainModuleInput() *string
	Migrations() WorkersScriptMigrationsOutputReference
	MigrationsInput() interface{}
	ModifiedOn() *string
	// The tree node.
	Node() constructs.Node
	Observability() WorkersScriptObservabilityOutputReference
	ObservabilityInput() interface{}
	Placement() WorkersScriptPlacementOutputReference
	PlacementInput() interface{}
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
	ScriptName() *string
	SetScriptName(val *string)
	ScriptNameInput() *string
	StartupTimeMs() *float64
	TailConsumers() WorkersScriptTailConsumersList
	TailConsumersInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsageModel() *string
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
	PutAssets(value *WorkersScriptAssets)
	PutBindings(value interface{})
	PutMigrations(value *WorkersScriptMigrations)
	PutObservability(value *WorkersScriptObservability)
	PutPlacement(value *WorkersScriptPlacement)
	PutTailConsumers(value interface{})
	ResetAssets()
	ResetBindings()
	ResetBodyPart()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetKeepAssets()
	ResetKeepBindings()
	ResetMainModule()
	ResetMigrations()
	ResetObservability()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacement()
	ResetTailConsumers()
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

// The jsii proxy struct for WorkersScript
type jsiiProxy_WorkersScript struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkersScript) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Assets() WorkersScriptAssetsOutputReference {
	var returns WorkersScriptAssetsOutputReference
	_jsii_.Get(
		j,
		"assets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) AssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Bindings() WorkersScriptBindingsList {
	var returns WorkersScriptBindingsList
	_jsii_.Get(
		j,
		"bindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) BindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) BodyPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) BodyPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyPartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) HasAssets() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) HasModules() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KeepAssets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KeepAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepAssetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KeepBindings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keepBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KeepBindingsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keepBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Logpush() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) MainModule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainModule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) MainModuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainModuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Migrations() WorkersScriptMigrationsOutputReference {
	var returns WorkersScriptMigrationsOutputReference
	_jsii_.Get(
		j,
		"migrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) MigrationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Observability() WorkersScriptObservabilityOutputReference {
	var returns WorkersScriptObservabilityOutputReference
	_jsii_.Get(
		j,
		"observability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ObservabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"observabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Placement() WorkersScriptPlacementOutputReference {
	var returns WorkersScriptPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ScriptName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ScriptNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) StartupTimeMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startupTimeMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TailConsumers() WorkersScriptTailConsumersList {
	var returns WorkersScriptTailConsumersList
	_jsii_.Get(
		j,
		"tailConsumers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TailConsumersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tailConsumersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script cloudflare_workers_script} Resource.
func NewWorkersScript(scope constructs.Construct, id *string, config *WorkersScriptConfig) WorkersScript {
	_init_.Initialize()

	if err := validateNewWorkersScriptParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScript{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script cloudflare_workers_script} Resource.
func NewWorkersScript_Override(w WorkersScript, scope constructs.Construct, id *string, config *WorkersScriptConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkersScript)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetBodyPart(val *string) {
	if err := j.validateSetBodyPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyPart",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetKeepAssets(val interface{}) {
	if err := j.validateSetKeepAssetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAssets",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetKeepBindings(val *[]*string) {
	if err := j.validateSetKeepBindingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepBindings",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetMainModule(val *string) {
	if err := j.validateSetMainModuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainModule",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetScriptName(val *string) {
	if err := j.validateSetScriptNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptName",
		val,
	)
}

// Generates CDKTF code for importing a WorkersScript resource upon running "cdktf plan <stack-name>".
func WorkersScript_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkersScript_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
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
func WorkersScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkersScript_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkersScript_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkersScript_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkersScript) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkersScript) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkersScript) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkersScript) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkersScript) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkersScript) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkersScript) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkersScript) PutAssets(value *WorkersScriptAssets) {
	if err := w.validatePutAssetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAssets",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutBindings(value interface{}) {
	if err := w.validatePutBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBindings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutMigrations(value *WorkersScriptMigrations) {
	if err := w.validatePutMigrationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMigrations",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutObservability(value *WorkersScriptObservability) {
	if err := w.validatePutObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putObservability",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutPlacement(value *WorkersScriptPlacement) {
	if err := w.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlacement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutTailConsumers(value interface{}) {
	if err := w.validatePutTailConsumersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTailConsumers",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) ResetAssets() {
	_jsii_.InvokeVoid(
		w,
		"resetAssets",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetBindings() {
	_jsii_.InvokeVoid(
		w,
		"resetBindings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetBodyPart() {
	_jsii_.InvokeVoid(
		w,
		"resetBodyPart",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetKeepAssets() {
	_jsii_.InvokeVoid(
		w,
		"resetKeepAssets",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetKeepBindings() {
	_jsii_.InvokeVoid(
		w,
		"resetKeepBindings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetMainModule() {
	_jsii_.InvokeVoid(
		w,
		"resetMainModule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetMigrations() {
	_jsii_.InvokeVoid(
		w,
		"resetMigrations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetObservability() {
	_jsii_.InvokeVoid(
		w,
		"resetObservability",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetPlacement() {
	_jsii_.InvokeVoid(
		w,
		"resetPlacement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetTailConsumers() {
	_jsii_.InvokeVoid(
		w,
		"resetTailConsumers",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


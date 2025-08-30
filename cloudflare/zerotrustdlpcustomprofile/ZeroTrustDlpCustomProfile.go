// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpcustomprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustdlpcustomprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile}.
type ZeroTrustDlpCustomProfile interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AiContextEnabled() interface{}
	SetAiContextEnabled(val interface{})
	AiContextEnabledInput() interface{}
	AllowedMatchCount() *float64
	SetAllowedMatchCount(val *float64)
	AllowedMatchCountInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidenceThreshold() *string
	SetConfidenceThreshold(val *string)
	ConfidenceThresholdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContextAwareness() ZeroTrustDlpCustomProfileContextAwarenessOutputReference
	ContextAwarenessInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Entries() ZeroTrustDlpCustomProfileEntriesList
	EntriesInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OcrEnabled() interface{}
	SetOcrEnabled(val interface{})
	OcrEnabledInput() interface{}
	OpenAccess() cdktf.IResolvable
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
	SharedEntries() ZeroTrustDlpCustomProfileSharedEntriesList
	SharedEntriesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
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
	PutContextAwareness(value *ZeroTrustDlpCustomProfileContextAwareness)
	PutEntries(value interface{})
	PutSharedEntries(value interface{})
	ResetAiContextEnabled()
	ResetAllowedMatchCount()
	ResetConfidenceThreshold()
	ResetContextAwareness()
	ResetDescription()
	ResetEntries()
	ResetOcrEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSharedEntries()
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

// The jsii proxy struct for ZeroTrustDlpCustomProfile
type jsiiProxy_ZeroTrustDlpCustomProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AiContextEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aiContextEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AiContextEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aiContextEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AllowedMatchCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedMatchCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AllowedMatchCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedMatchCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConfidenceThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConfidenceThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ContextAwareness() ZeroTrustDlpCustomProfileContextAwarenessOutputReference {
	var returns ZeroTrustDlpCustomProfileContextAwarenessOutputReference
	_jsii_.Get(
		j,
		"contextAwareness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ContextAwarenessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextAwarenessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Entries() ZeroTrustDlpCustomProfileEntriesList {
	var returns ZeroTrustDlpCustomProfileEntriesList
	_jsii_.Get(
		j,
		"entries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) EntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OcrEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocrEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OcrEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocrEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OpenAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"openAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SharedEntries() ZeroTrustDlpCustomProfileSharedEntriesList {
	var returns ZeroTrustDlpCustomProfileSharedEntriesList
	_jsii_.Get(
		j,
		"sharedEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SharedEntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedEntriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile} Resource.
func NewZeroTrustDlpCustomProfile(scope constructs.Construct, id *string, config *ZeroTrustDlpCustomProfileConfig) ZeroTrustDlpCustomProfile {
	_init_.Initialize()

	if err := validateNewZeroTrustDlpCustomProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDlpCustomProfile{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile} Resource.
func NewZeroTrustDlpCustomProfile_Override(z ZeroTrustDlpCustomProfile, scope constructs.Construct, id *string, config *ZeroTrustDlpCustomProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetAiContextEnabled(val interface{}) {
	if err := j.validateSetAiContextEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiContextEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetAllowedMatchCount(val *float64) {
	if err := j.validateSetAllowedMatchCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMatchCount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetConfidenceThreshold(val *string) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetOcrEnabled(val interface{}) {
	if err := j.validateSetOcrEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocrEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDlpCustomProfile resource upon running "cdktf plan <stack-name>".
func ZeroTrustDlpCustomProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
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
func ZeroTrustDlpCustomProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDlpCustomProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDlpCustomProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustDlpCustomProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutContextAwareness(value *ZeroTrustDlpCustomProfileContextAwareness) {
	if err := z.validatePutContextAwarenessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putContextAwareness",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutEntries(value interface{}) {
	if err := z.validatePutEntriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEntries",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutSharedEntries(value interface{}) {
	if err := z.validatePutSharedEntriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSharedEntries",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetAiContextEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetAiContextEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetAllowedMatchCount() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedMatchCount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		z,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetContextAwareness() {
	_jsii_.InvokeVoid(
		z,
		"resetContextAwareness",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		z,
		"resetDescription",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetEntries() {
	_jsii_.InvokeVoid(
		z,
		"resetEntries",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetOcrEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetOcrEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetSharedEntries() {
	_jsii_.InvokeVoid(
		z,
		"resetSharedEntries",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


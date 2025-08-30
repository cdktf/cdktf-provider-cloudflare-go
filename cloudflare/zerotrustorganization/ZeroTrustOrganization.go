// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustorganization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_organization cloudflare_zero_trust_organization}.
type ZeroTrustOrganization interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowAuthenticateViaWarp() interface{}
	SetAllowAuthenticateViaWarp(val interface{})
	AllowAuthenticateViaWarpInput() interface{}
	AuthDomain() *string
	SetAuthDomain(val *string)
	AuthDomainInput() *string
	AutoRedirectToIdentity() interface{}
	SetAutoRedirectToIdentity(val interface{})
	AutoRedirectToIdentityInput() interface{}
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
	CustomPages() ZeroTrustOrganizationCustomPagesOutputReference
	CustomPagesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IsUiReadOnly() interface{}
	SetIsUiReadOnly(val interface{})
	IsUiReadOnlyInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoginDesign() ZeroTrustOrganizationLoginDesignOutputReference
	LoginDesignInput() interface{}
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
	// Experimental.
	RawOverrides() interface{}
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UiReadOnlyToggleReason() *string
	SetUiReadOnlyToggleReason(val *string)
	UiReadOnlyToggleReasonInput() *string
	UserSeatExpirationInactiveTime() *string
	SetUserSeatExpirationInactiveTime(val *string)
	UserSeatExpirationInactiveTimeInput() *string
	WarpAuthSessionDuration() *string
	SetWarpAuthSessionDuration(val *string)
	WarpAuthSessionDurationInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutCustomPages(value *ZeroTrustOrganizationCustomPages)
	PutLoginDesign(value *ZeroTrustOrganizationLoginDesign)
	ResetAccountId()
	ResetAllowAuthenticateViaWarp()
	ResetAuthDomain()
	ResetAutoRedirectToIdentity()
	ResetCustomPages()
	ResetIsUiReadOnly()
	ResetLoginDesign()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSessionDuration()
	ResetUiReadOnlyToggleReason()
	ResetUserSeatExpirationInactiveTime()
	ResetWarpAuthSessionDuration()
	ResetZoneId()
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

// The jsii proxy struct for ZeroTrustOrganization
type jsiiProxy_ZeroTrustOrganization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustOrganization) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AllowAuthenticateViaWarp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AllowAuthenticateViaWarpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AuthDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AuthDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AutoRedirectToIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) AutoRedirectToIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) CustomPages() ZeroTrustOrganizationCustomPagesOutputReference {
	var returns ZeroTrustOrganizationCustomPagesOutputReference
	_jsii_.Get(
		j,
		"customPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) CustomPagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) IsUiReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUiReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) IsUiReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUiReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) LoginDesign() ZeroTrustOrganizationLoginDesignOutputReference {
	var returns ZeroTrustOrganizationLoginDesignOutputReference
	_jsii_.Get(
		j,
		"loginDesign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) LoginDesignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginDesignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) UiReadOnlyToggleReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiReadOnlyToggleReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) UiReadOnlyToggleReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiReadOnlyToggleReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) UserSeatExpirationInactiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSeatExpirationInactiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) UserSeatExpirationInactiveTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSeatExpirationInactiveTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) WarpAuthSessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warpAuthSessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) WarpAuthSessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warpAuthSessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustOrganization) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_organization cloudflare_zero_trust_organization} Resource.
func NewZeroTrustOrganization(scope constructs.Construct, id *string, config *ZeroTrustOrganizationConfig) ZeroTrustOrganization {
	_init_.Initialize()

	if err := validateNewZeroTrustOrganizationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustOrganization{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_organization cloudflare_zero_trust_organization} Resource.
func NewZeroTrustOrganization_Override(z ZeroTrustOrganization, scope constructs.Construct, id *string, config *ZeroTrustOrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetAllowAuthenticateViaWarp(val interface{}) {
	if err := j.validateSetAllowAuthenticateViaWarpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAuthenticateViaWarp",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetAuthDomain(val *string) {
	if err := j.validateSetAuthDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authDomain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetAutoRedirectToIdentity(val interface{}) {
	if err := j.validateSetAutoRedirectToIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRedirectToIdentity",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetIsUiReadOnly(val interface{}) {
	if err := j.validateSetIsUiReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isUiReadOnly",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetUiReadOnlyToggleReason(val *string) {
	if err := j.validateSetUiReadOnlyToggleReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uiReadOnlyToggleReason",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetUserSeatExpirationInactiveTime(val *string) {
	if err := j.validateSetUserSeatExpirationInactiveTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSeatExpirationInactiveTime",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetWarpAuthSessionDuration(val *string) {
	if err := j.validateSetWarpAuthSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warpAuthSessionDuration",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustOrganization)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustOrganization resource upon running "cdktf plan <stack-name>".
func ZeroTrustOrganization_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustOrganization_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
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
func ZeroTrustOrganization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustOrganization_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustOrganization_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustOrganization_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustOrganization_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustOrganization_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustOrganization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustOrganization.ZeroTrustOrganization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustOrganization) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustOrganization) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustOrganization) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) PutCustomPages(value *ZeroTrustOrganizationCustomPages) {
	if err := z.validatePutCustomPagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCustomPages",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) PutLoginDesign(value *ZeroTrustOrganizationLoginDesign) {
	if err := z.validatePutLoginDesignParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginDesign",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetAccountId() {
	_jsii_.InvokeVoid(
		z,
		"resetAccountId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetAllowAuthenticateViaWarp() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowAuthenticateViaWarp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetAuthDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetAutoRedirectToIdentity() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoRedirectToIdentity",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetCustomPages() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomPages",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetIsUiReadOnly() {
	_jsii_.InvokeVoid(
		z,
		"resetIsUiReadOnly",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetLoginDesign() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginDesign",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetName() {
	_jsii_.InvokeVoid(
		z,
		"resetName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetUiReadOnlyToggleReason() {
	_jsii_.InvokeVoid(
		z,
		"resetUiReadOnlyToggleReason",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetUserSeatExpirationInactiveTime() {
	_jsii_.InvokeVoid(
		z,
		"resetUserSeatExpirationInactiveTime",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetWarpAuthSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetWarpAuthSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) ResetZoneId() {
	_jsii_.InvokeVoid(
		z,
		"resetZoneId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustOrganization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustOrganization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


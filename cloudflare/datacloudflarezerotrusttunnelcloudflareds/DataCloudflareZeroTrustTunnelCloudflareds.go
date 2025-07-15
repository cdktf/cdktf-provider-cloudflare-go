// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflareds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrusttunnelcloudflareds/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_tunnel_cloudflareds cloudflare_zero_trust_tunnel_cloudflareds}.
type DataCloudflareZeroTrustTunnelCloudflareds interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	ExcludePrefix() *string
	SetExcludePrefix(val *string)
	ExcludePrefixInput() *string
	ExistedAt() *string
	SetExistedAt(val *string)
	ExistedAtInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IncludePrefix() *string
	SetIncludePrefix(val *string)
	IncludePrefixInput() *string
	IsDeleted() interface{}
	SetIsDeleted(val interface{})
	IsDeletedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
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
	RawOverrides() interface{}
	Result() DataCloudflareZeroTrustTunnelCloudflaredsResultList
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
	WasActiveAt() *string
	SetWasActiveAt(val *string)
	WasActiveAtInput() *string
	WasInactiveAt() *string
	SetWasInactiveAt(val *string)
	WasInactiveAtInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetExcludePrefix()
	ResetExistedAt()
	ResetIncludePrefix()
	ResetIsDeleted()
	ResetMaxItems()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStatus()
	ResetUuid()
	ResetWasActiveAt()
	ResetWasInactiveAt()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataCloudflareZeroTrustTunnelCloudflareds
type jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ExcludePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ExcludePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ExistedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ExistedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) IncludePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) IncludePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) IsDeleted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) IsDeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Result() DataCloudflareZeroTrustTunnelCloudflaredsResultList {
	var returns DataCloudflareZeroTrustTunnelCloudflaredsResultList
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) WasActiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasActiveAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) WasActiveAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasActiveAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) WasInactiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasInactiveAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) WasInactiveAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wasInactiveAtInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_tunnel_cloudflareds cloudflare_zero_trust_tunnel_cloudflareds} Data Source.
func NewDataCloudflareZeroTrustTunnelCloudflareds(scope constructs.Construct, id *string, config *DataCloudflareZeroTrustTunnelCloudflaredsConfig) DataCloudflareZeroTrustTunnelCloudflareds {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustTunnelCloudflaredsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_tunnel_cloudflareds cloudflare_zero_trust_tunnel_cloudflareds} Data Source.
func NewDataCloudflareZeroTrustTunnelCloudflareds_Override(d DataCloudflareZeroTrustTunnelCloudflareds, scope constructs.Construct, id *string, config *DataCloudflareZeroTrustTunnelCloudflaredsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetExcludePrefix(val *string) {
	if err := j.validateSetExcludePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefix",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetExistedAt(val *string) {
	if err := j.validateSetExistedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existedAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetIncludePrefix(val *string) {
	if err := j.validateSetIncludePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefix",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetIsDeleted(val interface{}) {
	if err := j.validateSetIsDeletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDeleted",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetWasActiveAt(val *string) {
	if err := j.validateSetWasActiveAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wasActiveAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds)SetWasInactiveAt(val *string) {
	if err := j.validateSetWasInactiveAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wasInactiveAt",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareZeroTrustTunnelCloudflareds resource upon running "cdktf plan <stack-name>".
func DataCloudflareZeroTrustTunnelCloudflareds_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflareds_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
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
func DataCloudflareZeroTrustTunnelCloudflareds_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflareds_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustTunnelCloudflareds_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflareds_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustTunnelCloudflareds_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflareds_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareZeroTrustTunnelCloudflareds_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflareds.DataCloudflareZeroTrustTunnelCloudflareds",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetExcludePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetExistedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetExistedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetIncludePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetIsDeleted() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDeleted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetMaxItems() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetWasActiveAt() {
	_jsii_.InvokeVoid(
		d,
		"resetWasActiveAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ResetWasInactiveAt() {
	_jsii_.InvokeVoid(
		d,
		"resetWasInactiveAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflareds) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


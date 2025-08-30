// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustdlppredefinedentry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustdlppredefinedentry/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_dlp_predefined_entry cloudflare_zero_trust_dlp_predefined_entry}.
type DataCloudflareZeroTrustDlpPredefinedEntry interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	CaseSensitive() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Confidence() DataCloudflareZeroTrustDlpPredefinedEntryConfidenceOutputReference
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
	Enabled() cdktf.IResolvable
	EntryId() *string
	SetEntryId(val *string)
	EntryIdInput() *string
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
	// The tree node.
	Node() constructs.Node
	Pattern() DataCloudflareZeroTrustDlpPredefinedEntryPatternOutputReference
	ProfileId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Secret() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	UpdatedAt() *string
	Variant() DataCloudflareZeroTrustDlpPredefinedEntryVariantOutputReference
	WordList() *string
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
	ResetEntryId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataCloudflareZeroTrustDlpPredefinedEntry
type jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) CaseSensitive() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Confidence() DataCloudflareZeroTrustDlpPredefinedEntryConfidenceOutputReference {
	var returns DataCloudflareZeroTrustDlpPredefinedEntryConfidenceOutputReference
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) EntryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) EntryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Pattern() DataCloudflareZeroTrustDlpPredefinedEntryPatternOutputReference {
	var returns DataCloudflareZeroTrustDlpPredefinedEntryPatternOutputReference
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Secret() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) Variant() DataCloudflareZeroTrustDlpPredefinedEntryVariantOutputReference {
	var returns DataCloudflareZeroTrustDlpPredefinedEntryVariantOutputReference
	_jsii_.Get(
		j,
		"variant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) WordList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wordList",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_dlp_predefined_entry cloudflare_zero_trust_dlp_predefined_entry} Data Source.
func NewDataCloudflareZeroTrustDlpPredefinedEntry(scope constructs.Construct, id *string, config *DataCloudflareZeroTrustDlpPredefinedEntryConfig) DataCloudflareZeroTrustDlpPredefinedEntry {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustDlpPredefinedEntryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_dlp_predefined_entry cloudflare_zero_trust_dlp_predefined_entry} Data Source.
func NewDataCloudflareZeroTrustDlpPredefinedEntry_Override(d DataCloudflareZeroTrustDlpPredefinedEntry, scope constructs.Construct, id *string, config *DataCloudflareZeroTrustDlpPredefinedEntryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetEntryId(val *string) {
	if err := j.validateSetEntryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareZeroTrustDlpPredefinedEntry resource upon running "cdktf plan <stack-name>".
func DataCloudflareZeroTrustDlpPredefinedEntry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustDlpPredefinedEntry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
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
func DataCloudflareZeroTrustDlpPredefinedEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustDlpPredefinedEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustDlpPredefinedEntry_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustDlpPredefinedEntry_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustDlpPredefinedEntry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustDlpPredefinedEntry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareZeroTrustDlpPredefinedEntry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustDlpPredefinedEntry.DataCloudflareZeroTrustDlpPredefinedEntry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ResetEntryId() {
	_jsii_.InvokeVoid(
		d,
		"resetEntryId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustDlpPredefinedEntry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


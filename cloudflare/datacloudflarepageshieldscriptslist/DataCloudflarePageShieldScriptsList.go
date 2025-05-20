// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldscriptslist

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarepageshieldscriptslist/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/page_shield_scripts_list cloudflare_page_shield_scripts_list}.
type DataCloudflarePageShieldScriptsList interface {
	cdktf.TerraformDataSource
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
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	ExcludeCdnCgi() interface{}
	SetExcludeCdnCgi(val interface{})
	ExcludeCdnCgiInput() interface{}
	ExcludeDuplicates() interface{}
	SetExcludeDuplicates(val interface{})
	ExcludeDuplicatesInput() interface{}
	ExcludeUrls() *string
	SetExcludeUrls(val *string)
	ExcludeUrlsInput() *string
	Export() *string
	SetExport(val *string)
	ExportInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hosts() *string
	SetHosts(val *string)
	HostsInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
	// The tree node.
	Node() constructs.Node
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	Page() *string
	SetPage(val *string)
	PageInput() *string
	PageUrl() *string
	SetPageUrl(val *string)
	PageUrlInput() *string
	PerPage() *float64
	SetPerPage(val *float64)
	PerPageInput() *float64
	PrioritizeMalicious() interface{}
	SetPrioritizeMalicious(val interface{})
	PrioritizeMaliciousInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Result() DataCloudflarePageShieldScriptsListResultList
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Urls() *string
	SetUrls(val *string)
	UrlsInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	ResetDirection()
	ResetExcludeCdnCgi()
	ResetExcludeDuplicates()
	ResetExcludeUrls()
	ResetExport()
	ResetHosts()
	ResetMaxItems()
	ResetOrderBy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPage()
	ResetPageUrl()
	ResetPerPage()
	ResetPrioritizeMalicious()
	ResetStatus()
	ResetUrls()
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

// The jsii proxy struct for DataCloudflarePageShieldScriptsList
type jsiiProxy_DataCloudflarePageShieldScriptsList struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeCdnCgi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCdnCgi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeCdnCgiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCdnCgiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeDuplicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeDuplicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeDuplicatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeDuplicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeUrls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExcludeUrlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Export() *string {
	var returns *string
	_jsii_.Get(
		j,
		"export",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ExportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Hosts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) HostsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Page() *string {
	var returns *string
	_jsii_.Get(
		j,
		"page",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PrioritizeMalicious() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeMalicious",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) PrioritizeMaliciousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeMaliciousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Result() DataCloudflarePageShieldScriptsListResultList {
	var returns DataCloudflarePageShieldScriptsListResultList
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) Urls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) UrlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/page_shield_scripts_list cloudflare_page_shield_scripts_list} Data Source.
func NewDataCloudflarePageShieldScriptsList(scope constructs.Construct, id *string, config *DataCloudflarePageShieldScriptsListConfig) DataCloudflarePageShieldScriptsList {
	_init_.Initialize()

	if err := validateNewDataCloudflarePageShieldScriptsListParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePageShieldScriptsList{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/data-sources/page_shield_scripts_list cloudflare_page_shield_scripts_list} Data Source.
func NewDataCloudflarePageShieldScriptsList_Override(d DataCloudflarePageShieldScriptsList, scope constructs.Construct, id *string, config *DataCloudflarePageShieldScriptsListConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetExcludeCdnCgi(val interface{}) {
	if err := j.validateSetExcludeCdnCgiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCdnCgi",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetExcludeDuplicates(val interface{}) {
	if err := j.validateSetExcludeDuplicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeDuplicates",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetExcludeUrls(val *string) {
	if err := j.validateSetExcludeUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUrls",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetExport(val *string) {
	if err := j.validateSetExportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"export",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetHosts(val *string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetPage(val *string) {
	if err := j.validateSetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"page",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetPageUrl(val *string) {
	if err := j.validateSetPageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageUrl",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetPerPage(val *float64) {
	if err := j.validateSetPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perPage",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetPrioritizeMalicious(val interface{}) {
	if err := j.validateSetPrioritizeMaliciousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prioritizeMalicious",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetUrls(val *string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsList)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflarePageShieldScriptsList resource upon running "cdktf plan <stack-name>".
func DataCloudflarePageShieldScriptsList_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflarePageShieldScriptsList_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
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
func DataCloudflarePageShieldScriptsList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflarePageShieldScriptsList_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflarePageShieldScriptsList_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflarePageShieldScriptsList_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflarePageShieldScriptsList_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflarePageShieldScriptsList_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflarePageShieldScriptsList_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsList",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetExcludeCdnCgi() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeCdnCgi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetExcludeDuplicates() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeDuplicates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetExcludeUrls() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeUrls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetExport() {
	_jsii_.InvokeVoid(
		d,
		"resetExport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetHosts() {
	_jsii_.InvokeVoid(
		d,
		"resetHosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetMaxItems() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetPage() {
	_jsii_.InvokeVoid(
		d,
		"resetPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetPageUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetPageUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetPerPage() {
	_jsii_.InvokeVoid(
		d,
		"resetPerPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetPrioritizeMalicious() {
	_jsii_.InvokeVoid(
		d,
		"resetPrioritizeMalicious",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ResetUrls() {
	_jsii_.InvokeVoid(
		d,
		"resetUrls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsList) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


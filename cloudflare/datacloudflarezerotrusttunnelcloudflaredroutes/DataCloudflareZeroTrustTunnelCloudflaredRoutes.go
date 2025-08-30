// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflaredroutes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrusttunnelcloudflaredroutes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_tunnel_cloudflared_routes cloudflare_zero_trust_tunnel_cloudflared_routes}.
type DataCloudflareZeroTrustTunnelCloudflaredRoutes interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	NetworkSubset() *string
	SetNetworkSubset(val *string)
	NetworkSubsetInput() *string
	NetworkSuperset() *string
	SetNetworkSuperset(val *string)
	NetworkSupersetInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Result() DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList
	RouteId() *string
	SetRouteId(val *string)
	RouteIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TunnelId() *string
	SetTunnelId(val *string)
	TunnelIdInput() *string
	TunTypes() *[]*string
	SetTunTypes(val *[]*string)
	TunTypesInput() *[]*string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
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
	ResetComment()
	ResetExistedAt()
	ResetIsDeleted()
	ResetMaxItems()
	ResetNetworkSubset()
	ResetNetworkSuperset()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRouteId()
	ResetTunnelId()
	ResetTunTypes()
	ResetVirtualNetworkId()
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

// The jsii proxy struct for DataCloudflareZeroTrustTunnelCloudflaredRoutes
type jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ExistedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ExistedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) IsDeleted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) IsDeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) NetworkSubset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSubset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) NetworkSubsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSubsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) NetworkSuperset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSuperset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) NetworkSupersetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSupersetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) Result() DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList {
	var returns DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) RouteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TunnelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TunnelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TunTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) TunTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_tunnel_cloudflared_routes cloudflare_zero_trust_tunnel_cloudflared_routes} Data Source.
func NewDataCloudflareZeroTrustTunnelCloudflaredRoutes(scope constructs.Construct, id *string, config *DataCloudflareZeroTrustTunnelCloudflaredRoutesConfig) DataCloudflareZeroTrustTunnelCloudflaredRoutes {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustTunnelCloudflaredRoutesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zero_trust_tunnel_cloudflared_routes cloudflare_zero_trust_tunnel_cloudflared_routes} Data Source.
func NewDataCloudflareZeroTrustTunnelCloudflaredRoutes_Override(d DataCloudflareZeroTrustTunnelCloudflaredRoutes, scope constructs.Construct, id *string, config *DataCloudflareZeroTrustTunnelCloudflaredRoutesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetExistedAt(val *string) {
	if err := j.validateSetExistedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existedAt",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetIsDeleted(val interface{}) {
	if err := j.validateSetIsDeletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDeleted",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetNetworkSubset(val *string) {
	if err := j.validateSetNetworkSubsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSubset",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetNetworkSuperset(val *string) {
	if err := j.validateSetNetworkSupersetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSuperset",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetRouteId(val *string) {
	if err := j.validateSetRouteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetTunnelId(val *string) {
	if err := j.validateSetTunnelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetTunTypes(val *[]*string) {
	if err := j.validateSetTunTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunTypes",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareZeroTrustTunnelCloudflaredRoutes resource upon running "cdktf plan <stack-name>".
func DataCloudflareZeroTrustTunnelCloudflaredRoutes_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflaredRoutes_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
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
func DataCloudflareZeroTrustTunnelCloudflaredRoutes_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflaredRoutes_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustTunnelCloudflaredRoutes_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflaredRoutes_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustTunnelCloudflaredRoutes_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustTunnelCloudflaredRoutes_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareZeroTrustTunnelCloudflaredRoutes_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetExistedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetExistedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetIsDeleted() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDeleted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetMaxItems() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetNetworkSubset() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkSubset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetNetworkSuperset() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkSuperset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetRouteId() {
	_jsii_.InvokeVoid(
		d,
		"resetRouteId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetTunnelId() {
	_jsii_.InvokeVoid(
		d,
		"resetTunnelId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetTunTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetTunTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessinfrastructuretargets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessinfrastructuretargets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_access_infrastructure_targets cloudflare_zero_trust_access_infrastructure_targets}.
type DataCloudflareZeroTrustAccessInfrastructureTargets interface {
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
	CreatedAfter() *string
	SetCreatedAfter(val *string)
	CreatedAfterInput() *string
	CreatedBefore() *string
	SetCreatedBefore(val *string)
	CreatedBeforeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameContains() *string
	SetHostnameContains(val *string)
	HostnameContainsInput() *string
	HostnameInput() *string
	IpLike() *string
	SetIpLike(val *string)
	IpLikeInput() *string
	Ips() *[]*string
	SetIps(val *[]*string)
	IpsInput() *[]*string
	IpV4() *string
	SetIpV4(val *string)
	Ipv4End() *string
	SetIpv4End(val *string)
	Ipv4EndInput() *string
	IpV4Input() *string
	Ipv4Start() *string
	SetIpv4Start(val *string)
	Ipv4StartInput() *string
	IpV6() *string
	SetIpV6(val *string)
	Ipv6End() *string
	SetIpv6End(val *string)
	Ipv6EndInput() *string
	IpV6Input() *string
	Ipv6Start() *string
	SetIpv6Start(val *string)
	Ipv6StartInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
	ModifiedAfter() *string
	SetModifiedAfter(val *string)
	ModifiedAfterInput() *string
	ModifiedBefore() *string
	SetModifiedBefore(val *string)
	ModifiedBeforeInput() *string
	// The tree node.
	Node() constructs.Node
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Result() DataCloudflareZeroTrustAccessInfrastructureTargetsResultList
	TargetIds() *[]*string
	SetTargetIds(val *[]*string)
	TargetIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCreatedAfter()
	ResetCreatedBefore()
	ResetDirection()
	ResetHostname()
	ResetHostnameContains()
	ResetIpLike()
	ResetIps()
	ResetIpV4()
	ResetIpv4End()
	ResetIpv4Start()
	ResetIpV6()
	ResetIpv6End()
	ResetIpv6Start()
	ResetMaxItems()
	ResetModifiedAfter()
	ResetModifiedBefore()
	ResetOrder()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTargetIds()
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessInfrastructureTargets
type jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) CreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) CreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) HostnameContains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) HostnameContainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpLike() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpLikeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpV4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv4End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4End",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv4EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4EndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpV4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv4Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv4StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4StartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpV6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv6End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6End",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv6EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6EndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) IpV6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv6Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Ipv6StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6StartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ModifiedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ModifiedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) Result() DataCloudflareZeroTrustAccessInfrastructureTargetsResultList {
	var returns DataCloudflareZeroTrustAccessInfrastructureTargetsResultList
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) TargetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) TargetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_access_infrastructure_targets cloudflare_zero_trust_access_infrastructure_targets} Data Source.
func NewDataCloudflareZeroTrustAccessInfrastructureTargets(scope constructs.Construct, id *string, config *DataCloudflareZeroTrustAccessInfrastructureTargetsConfig) DataCloudflareZeroTrustAccessInfrastructureTargets {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessInfrastructureTargetsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/zero_trust_access_infrastructure_targets cloudflare_zero_trust_access_infrastructure_targets} Data Source.
func NewDataCloudflareZeroTrustAccessInfrastructureTargets_Override(d DataCloudflareZeroTrustAccessInfrastructureTargets, scope constructs.Construct, id *string, config *DataCloudflareZeroTrustAccessInfrastructureTargetsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetCreatedAfter(val *string) {
	if err := j.validateSetCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetHostnameContains(val *string) {
	if err := j.validateSetHostnameContainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameContains",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpLike(val *string) {
	if err := j.validateSetIpLikeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipLike",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIps(val *[]*string) {
	if err := j.validateSetIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ips",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpV4(val *string) {
	if err := j.validateSetIpV4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV4",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpv4End(val *string) {
	if err := j.validateSetIpv4EndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4End",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpv4Start(val *string) {
	if err := j.validateSetIpv4StartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Start",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpV6(val *string) {
	if err := j.validateSetIpV6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV6",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpv6End(val *string) {
	if err := j.validateSetIpv6EndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6End",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetIpv6Start(val *string) {
	if err := j.validateSetIpv6StartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Start",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetModifiedAfter(val *string) {
	if err := j.validateSetModifiedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetModifiedBefore(val *string) {
	if err := j.validateSetModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetTargetIds(val *[]*string) {
	if err := j.validateSetTargetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIds",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareZeroTrustAccessInfrastructureTargets resource upon running "cdktf plan <stack-name>".
func DataCloudflareZeroTrustAccessInfrastructureTargets_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustAccessInfrastructureTargets_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
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
func DataCloudflareZeroTrustAccessInfrastructureTargets_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustAccessInfrastructureTargets_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustAccessInfrastructureTargets_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustAccessInfrastructureTargets_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareZeroTrustAccessInfrastructureTargets_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareZeroTrustAccessInfrastructureTargets_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareZeroTrustAccessInfrastructureTargets_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTargets.DataCloudflareZeroTrustAccessInfrastructureTargets",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetCreatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetHostname() {
	_jsii_.InvokeVoid(
		d,
		"resetHostname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetHostnameContains() {
	_jsii_.InvokeVoid(
		d,
		"resetHostnameContains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpLike() {
	_jsii_.InvokeVoid(
		d,
		"resetIpLike",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIps() {
	_jsii_.InvokeVoid(
		d,
		"resetIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpV4() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV4",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpv4End() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4End",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpv4Start() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4Start",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpV6() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV6",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpv6End() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6End",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetIpv6Start() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Start",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetMaxItems() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetModifiedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetModifiedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetModifiedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetModifiedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetTargetIds() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargets) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


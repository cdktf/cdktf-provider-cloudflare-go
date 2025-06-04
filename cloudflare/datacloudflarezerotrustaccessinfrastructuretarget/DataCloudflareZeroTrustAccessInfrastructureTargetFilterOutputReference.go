// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessinfrastructuretarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessinfrastructuretarget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CreatedAfter() *string
	SetCreatedAfter(val *string)
	CreatedAfterInput() *string
	CreatedBefore() *string
	SetCreatedBefore(val *string)
	CreatedBeforeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameContains() *string
	SetHostnameContains(val *string)
	HostnameContainsInput() *string
	HostnameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ModifiedAfter() *string
	SetModifiedAfter(val *string)
	ModifiedAfterInput() *string
	ModifiedBefore() *string
	SetModifiedBefore(val *string)
	ModifiedBeforeInput() *string
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	TargetIds() *[]*string
	SetTargetIds(val *[]*string)
	TargetIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
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
	ResetModifiedAfter()
	ResetModifiedBefore()
	ResetOrder()
	ResetTargetIds()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) CreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) CreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) HostnameContains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) HostnameContainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpLike() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpLikeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ips() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpV4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv4End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4End",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv4EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4EndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpV4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv4Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv4StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4StartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpV6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv6End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6End",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv6EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6EndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) IpV6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipV6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv6Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Ipv6StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6StartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ModifiedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ModifiedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) TargetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) TargetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference_Override(d DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetCreatedAfter(val *string) {
	if err := j.validateSetCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetHostnameContains(val *string) {
	if err := j.validateSetHostnameContainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameContains",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpLike(val *string) {
	if err := j.validateSetIpLikeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipLike",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIps(val *[]*string) {
	if err := j.validateSetIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ips",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpV4(val *string) {
	if err := j.validateSetIpV4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV4",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpv4End(val *string) {
	if err := j.validateSetIpv4EndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4End",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpv4Start(val *string) {
	if err := j.validateSetIpv4StartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Start",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpV6(val *string) {
	if err := j.validateSetIpV6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipV6",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpv6End(val *string) {
	if err := j.validateSetIpv6EndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6End",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetIpv6Start(val *string) {
	if err := j.validateSetIpv6StartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Start",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetModifiedAfter(val *string) {
	if err := j.validateSetModifiedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedAfter",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetModifiedBefore(val *string) {
	if err := j.validateSetModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedBefore",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetTargetIds(val *[]*string) {
	if err := j.validateSetTargetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIds",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetCreatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		d,
		"resetHostname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetHostnameContains() {
	_jsii_.InvokeVoid(
		d,
		"resetHostnameContains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpLike() {
	_jsii_.InvokeVoid(
		d,
		"resetIpLike",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIps() {
	_jsii_.InvokeVoid(
		d,
		"resetIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpV4() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV4",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpv4End() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4End",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpv4Start() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4Start",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpV6() {
	_jsii_.InvokeVoid(
		d,
		"resetIpV6",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpv6End() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6End",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetIpv6Start() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Start",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetModifiedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetModifiedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetModifiedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetModifiedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetTargetIds() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


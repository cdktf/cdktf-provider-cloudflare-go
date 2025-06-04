// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountmember

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflareaccountmember/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccountMemberRolesPermissionsOutputReference interface {
	cdktf.ComplexObject
	Analytics() DataCloudflareAccountMemberRolesPermissionsAnalyticsOutputReference
	Billing() DataCloudflareAccountMemberRolesPermissionsBillingOutputReference
	CachePurge() DataCloudflareAccountMemberRolesPermissionsCachePurgeOutputReference
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dns() DataCloudflareAccountMemberRolesPermissionsDnsOutputReference
	DnsRecords() DataCloudflareAccountMemberRolesPermissionsDnsRecordsOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareAccountMemberRolesPermissions
	SetInternalValue(val *DataCloudflareAccountMemberRolesPermissions)
	Lb() DataCloudflareAccountMemberRolesPermissionsLbOutputReference
	Logs() DataCloudflareAccountMemberRolesPermissionsLogsOutputReference
	Organization() DataCloudflareAccountMemberRolesPermissionsOrganizationOutputReference
	Ssl() DataCloudflareAccountMemberRolesPermissionsSslOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Waf() DataCloudflareAccountMemberRolesPermissionsWafOutputReference
	Zones() DataCloudflareAccountMemberRolesPermissionsZonesOutputReference
	ZoneSettings() DataCloudflareAccountMemberRolesPermissionsZoneSettingsOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareAccountMemberRolesPermissionsOutputReference
type jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Analytics() DataCloudflareAccountMemberRolesPermissionsAnalyticsOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsAnalyticsOutputReference
	_jsii_.Get(
		j,
		"analytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Billing() DataCloudflareAccountMemberRolesPermissionsBillingOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsBillingOutputReference
	_jsii_.Get(
		j,
		"billing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) CachePurge() DataCloudflareAccountMemberRolesPermissionsCachePurgeOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsCachePurgeOutputReference
	_jsii_.Get(
		j,
		"cachePurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Dns() DataCloudflareAccountMemberRolesPermissionsDnsOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) DnsRecords() DataCloudflareAccountMemberRolesPermissionsDnsRecordsOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsDnsRecordsOutputReference
	_jsii_.Get(
		j,
		"dnsRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) InternalValue() *DataCloudflareAccountMemberRolesPermissions {
	var returns *DataCloudflareAccountMemberRolesPermissions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Lb() DataCloudflareAccountMemberRolesPermissionsLbOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsLbOutputReference
	_jsii_.Get(
		j,
		"lb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Logs() DataCloudflareAccountMemberRolesPermissionsLogsOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Organization() DataCloudflareAccountMemberRolesPermissionsOrganizationOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsOrganizationOutputReference
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Ssl() DataCloudflareAccountMemberRolesPermissionsSslOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsSslOutputReference
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Waf() DataCloudflareAccountMemberRolesPermissionsWafOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsWafOutputReference
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Zones() DataCloudflareAccountMemberRolesPermissionsZonesOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsZonesOutputReference
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) ZoneSettings() DataCloudflareAccountMemberRolesPermissionsZoneSettingsOutputReference {
	var returns DataCloudflareAccountMemberRolesPermissionsZoneSettingsOutputReference
	_jsii_.Get(
		j,
		"zoneSettings",
		&returns,
	)
	return returns
}


func NewDataCloudflareAccountMemberRolesPermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareAccountMemberRolesPermissionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAccountMemberRolesPermissionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAccountMember.DataCloudflareAccountMemberRolesPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareAccountMemberRolesPermissionsOutputReference_Override(d DataCloudflareAccountMemberRolesPermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAccountMember.DataCloudflareAccountMemberRolesPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference)SetInternalValue(val *DataCloudflareAccountMemberRolesPermissions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareAccountMemberRolesPermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


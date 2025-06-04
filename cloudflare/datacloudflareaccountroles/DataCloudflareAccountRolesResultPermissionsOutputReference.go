// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccountroles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflareaccountroles/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccountRolesResultPermissionsOutputReference interface {
	cdktf.ComplexObject
	Analytics() DataCloudflareAccountRolesResultPermissionsAnalyticsOutputReference
	Billing() DataCloudflareAccountRolesResultPermissionsBillingOutputReference
	CachePurge() DataCloudflareAccountRolesResultPermissionsCachePurgeOutputReference
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
	Dns() DataCloudflareAccountRolesResultPermissionsDnsOutputReference
	DnsRecords() DataCloudflareAccountRolesResultPermissionsDnsRecordsOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareAccountRolesResultPermissions
	SetInternalValue(val *DataCloudflareAccountRolesResultPermissions)
	Lb() DataCloudflareAccountRolesResultPermissionsLbOutputReference
	Logs() DataCloudflareAccountRolesResultPermissionsLogsOutputReference
	Organization() DataCloudflareAccountRolesResultPermissionsOrganizationOutputReference
	Ssl() DataCloudflareAccountRolesResultPermissionsSslOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Waf() DataCloudflareAccountRolesResultPermissionsWafOutputReference
	Zones() DataCloudflareAccountRolesResultPermissionsZonesOutputReference
	ZoneSettings() DataCloudflareAccountRolesResultPermissionsZoneSettingsOutputReference
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

// The jsii proxy struct for DataCloudflareAccountRolesResultPermissionsOutputReference
type jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Analytics() DataCloudflareAccountRolesResultPermissionsAnalyticsOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsAnalyticsOutputReference
	_jsii_.Get(
		j,
		"analytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Billing() DataCloudflareAccountRolesResultPermissionsBillingOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsBillingOutputReference
	_jsii_.Get(
		j,
		"billing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) CachePurge() DataCloudflareAccountRolesResultPermissionsCachePurgeOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsCachePurgeOutputReference
	_jsii_.Get(
		j,
		"cachePurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Dns() DataCloudflareAccountRolesResultPermissionsDnsOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) DnsRecords() DataCloudflareAccountRolesResultPermissionsDnsRecordsOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsDnsRecordsOutputReference
	_jsii_.Get(
		j,
		"dnsRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) InternalValue() *DataCloudflareAccountRolesResultPermissions {
	var returns *DataCloudflareAccountRolesResultPermissions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Lb() DataCloudflareAccountRolesResultPermissionsLbOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsLbOutputReference
	_jsii_.Get(
		j,
		"lb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Logs() DataCloudflareAccountRolesResultPermissionsLogsOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Organization() DataCloudflareAccountRolesResultPermissionsOrganizationOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsOrganizationOutputReference
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Ssl() DataCloudflareAccountRolesResultPermissionsSslOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsSslOutputReference
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Waf() DataCloudflareAccountRolesResultPermissionsWafOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsWafOutputReference
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Zones() DataCloudflareAccountRolesResultPermissionsZonesOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsZonesOutputReference
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) ZoneSettings() DataCloudflareAccountRolesResultPermissionsZoneSettingsOutputReference {
	var returns DataCloudflareAccountRolesResultPermissionsZoneSettingsOutputReference
	_jsii_.Get(
		j,
		"zoneSettings",
		&returns,
	)
	return returns
}


func NewDataCloudflareAccountRolesResultPermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareAccountRolesResultPermissionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAccountRolesResultPermissionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAccountRoles.DataCloudflareAccountRolesResultPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareAccountRolesResultPermissionsOutputReference_Override(d DataCloudflareAccountRolesResultPermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAccountRoles.DataCloudflareAccountRolesResultPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference)SetInternalValue(val *DataCloudflareAccountRolesResultPermissions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareAccountRolesResultPermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


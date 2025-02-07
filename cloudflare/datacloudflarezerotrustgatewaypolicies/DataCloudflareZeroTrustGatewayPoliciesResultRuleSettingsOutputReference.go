// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewaypolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustgatewaypolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHeaders() cdktf.StringMap
	AllowChildBypass() cdktf.IResolvable
	AuditSsh() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsAuditSshOutputReference
	BisoAdminControls() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsBisoAdminControlsOutputReference
	BlockPageEnabled() cdktf.IResolvable
	BlockReason() *string
	BypassParentRule() cdktf.IResolvable
	CheckSession() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsCheckSessionOutputReference
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
	DnsResolvers() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsDnsResolversOutputReference
	Egress() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsEgressOutputReference
	// Experimental.
	Fqn() *string
	IgnoreCnameCategoryMatches() cdktf.IResolvable
	InsecureDisableDnssecValidation() cdktf.IResolvable
	InternalValue() *DataCloudflareZeroTrustGatewayPoliciesResultRuleSettings
	SetInternalValue(val *DataCloudflareZeroTrustGatewayPoliciesResultRuleSettings)
	IpCategories() cdktf.IResolvable
	IpIndicatorFeeds() cdktf.IResolvable
	L4Override() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsL4OverrideOutputReference
	NotificationSettings() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsNotificationSettingsOutputReference
	OverrideHost() *string
	OverrideIps() *[]*string
	PayloadLog() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsPayloadLogOutputReference
	Quarantine() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsQuarantineOutputReference
	ResolveDnsInternally() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsResolveDnsInternallyOutputReference
	ResolveDnsThroughCloudflare() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedCert() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsUntrustedCertOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference
type jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) AddHeaders() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"addHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) AllowChildBypass() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowChildBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) AuditSsh() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsAuditSshOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsAuditSshOutputReference
	_jsii_.Get(
		j,
		"auditSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) BisoAdminControls() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsBisoAdminControlsOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsBisoAdminControlsOutputReference
	_jsii_.Get(
		j,
		"bisoAdminControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) BlockPageEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"blockPageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) BlockReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) BypassParentRule() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bypassParentRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) CheckSession() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsCheckSessionOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsCheckSessionOutputReference
	_jsii_.Get(
		j,
		"checkSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) DnsResolvers() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsDnsResolversOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsDnsResolversOutputReference
	_jsii_.Get(
		j,
		"dnsResolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) Egress() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsEgressOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsEgressOutputReference
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) IgnoreCnameCategoryMatches() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ignoreCnameCategoryMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) InsecureDisableDnssecValidation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) InternalValue() *DataCloudflareZeroTrustGatewayPoliciesResultRuleSettings {
	var returns *DataCloudflareZeroTrustGatewayPoliciesResultRuleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) IpCategories() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) IpIndicatorFeeds() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipIndicatorFeeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) L4Override() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsL4OverrideOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsL4OverrideOutputReference
	_jsii_.Get(
		j,
		"l4Override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) NotificationSettings() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsNotificationSettingsOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) OverrideHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) OverrideIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) PayloadLog() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsPayloadLogOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) Quarantine() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsQuarantineOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsQuarantineOutputReference
	_jsii_.Get(
		j,
		"quarantine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ResolveDnsInternally() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsResolveDnsInternallyOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsResolveDnsInternallyOutputReference
	_jsii_.Get(
		j,
		"resolveDnsInternally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ResolveDnsThroughCloudflare() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"resolveDnsThroughCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) UntrustedCert() DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsUntrustedCertOutputReference {
	var returns DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsUntrustedCertOutputReference
	_jsii_.Get(
		j,
		"untrustedCert",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewayPolicies.DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference_Override(d DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewayPolicies.DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference)SetInternalValue(val *DataCloudflareZeroTrustGatewayPoliciesResultRuleSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPoliciesResultRuleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


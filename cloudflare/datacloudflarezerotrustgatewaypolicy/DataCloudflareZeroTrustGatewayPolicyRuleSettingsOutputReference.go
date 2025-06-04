// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustgatewaypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHeaders() cdktf.StringMap
	AllowChildBypass() cdktf.IResolvable
	AuditSsh() DataCloudflareZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference
	BisoAdminControls() DataCloudflareZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference
	BlockPage() DataCloudflareZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference
	BlockPageEnabled() cdktf.IResolvable
	BlockReason() *string
	BypassParentRule() cdktf.IResolvable
	CheckSession() DataCloudflareZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference
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
	DnsResolvers() DataCloudflareZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference
	Egress() DataCloudflareZeroTrustGatewayPolicyRuleSettingsEgressOutputReference
	// Experimental.
	Fqn() *string
	IgnoreCnameCategoryMatches() cdktf.IResolvable
	InsecureDisableDnssecValidation() cdktf.IResolvable
	InternalValue() *DataCloudflareZeroTrustGatewayPolicyRuleSettings
	SetInternalValue(val *DataCloudflareZeroTrustGatewayPolicyRuleSettings)
	IpCategories() cdktf.IResolvable
	IpIndicatorFeeds() cdktf.IResolvable
	L4Override() DataCloudflareZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference
	NotificationSettings() DataCloudflareZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference
	OverrideHost() *string
	OverrideIps() *[]*string
	PayloadLog() DataCloudflareZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference
	Quarantine() DataCloudflareZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference
	Redirect() DataCloudflareZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference
	ResolveDnsInternally() DataCloudflareZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference
	ResolveDnsThroughCloudflare() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedCert() DataCloudflareZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference
type jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) AddHeaders() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"addHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) AllowChildBypass() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowChildBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) AuditSsh() DataCloudflareZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference
	_jsii_.Get(
		j,
		"auditSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) BisoAdminControls() DataCloudflareZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference
	_jsii_.Get(
		j,
		"bisoAdminControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPage() DataCloudflareZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPageEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"blockPageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) BypassParentRule() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bypassParentRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) CheckSession() DataCloudflareZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference
	_jsii_.Get(
		j,
		"checkSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) DnsResolvers() DataCloudflareZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference
	_jsii_.Get(
		j,
		"dnsResolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) Egress() DataCloudflareZeroTrustGatewayPolicyRuleSettingsEgressOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsEgressOutputReference
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) IgnoreCnameCategoryMatches() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ignoreCnameCategoryMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) InsecureDisableDnssecValidation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) InternalValue() *DataCloudflareZeroTrustGatewayPolicyRuleSettings {
	var returns *DataCloudflareZeroTrustGatewayPolicyRuleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) IpCategories() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) IpIndicatorFeeds() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipIndicatorFeeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) L4Override() DataCloudflareZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference
	_jsii_.Get(
		j,
		"l4Override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) NotificationSettings() DataCloudflareZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) PayloadLog() DataCloudflareZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) Quarantine() DataCloudflareZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference
	_jsii_.Get(
		j,
		"quarantine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) Redirect() DataCloudflareZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsInternally() DataCloudflareZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference
	_jsii_.Get(
		j,
		"resolveDnsInternally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsThroughCloudflare() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"resolveDnsThroughCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) UntrustedCert() DataCloudflareZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference {
	var returns DataCloudflareZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference
	_jsii_.Get(
		j,
		"untrustedCert",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewayPolicy.DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference_Override(d DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewayPolicy.DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference)SetInternalValue(val *DataCloudflareZeroTrustGatewayPolicyRuleSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewayPolicyRuleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


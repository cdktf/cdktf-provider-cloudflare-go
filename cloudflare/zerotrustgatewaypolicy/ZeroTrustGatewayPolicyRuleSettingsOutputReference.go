// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustgatewaypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPolicyRuleSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHeaders() *map[string]*string
	SetAddHeaders(val *map[string]*string)
	AddHeadersInput() *map[string]*string
	AllowChildBypass() interface{}
	SetAllowChildBypass(val interface{})
	AllowChildBypassInput() interface{}
	AuditSsh() ZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference
	AuditSshInput() interface{}
	BisoAdminControls() ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference
	BisoAdminControlsInput() interface{}
	BlockPage() ZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference
	BlockPageEnabled() interface{}
	SetBlockPageEnabled(val interface{})
	BlockPageEnabledInput() interface{}
	BlockPageInput() interface{}
	BlockReason() *string
	SetBlockReason(val *string)
	BlockReasonInput() *string
	BypassParentRule() interface{}
	SetBypassParentRule(val interface{})
	BypassParentRuleInput() interface{}
	CheckSession() ZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference
	CheckSessionInput() interface{}
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
	DnsResolvers() ZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference
	DnsResolversInput() interface{}
	Egress() ZeroTrustGatewayPolicyRuleSettingsEgressOutputReference
	EgressInput() interface{}
	// Experimental.
	Fqn() *string
	IgnoreCnameCategoryMatches() interface{}
	SetIgnoreCnameCategoryMatches(val interface{})
	IgnoreCnameCategoryMatchesInput() interface{}
	InsecureDisableDnssecValidation() interface{}
	SetInsecureDisableDnssecValidation(val interface{})
	InsecureDisableDnssecValidationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpCategories() interface{}
	SetIpCategories(val interface{})
	IpCategoriesInput() interface{}
	IpIndicatorFeeds() interface{}
	SetIpIndicatorFeeds(val interface{})
	IpIndicatorFeedsInput() interface{}
	L4Override() ZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference
	L4OverrideInput() interface{}
	NotificationSettings() ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference
	NotificationSettingsInput() interface{}
	OverrideHost() *string
	SetOverrideHost(val *string)
	OverrideHostInput() *string
	OverrideIps() *[]*string
	SetOverrideIps(val *[]*string)
	OverrideIpsInput() *[]*string
	PayloadLog() ZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference
	PayloadLogInput() interface{}
	Quarantine() ZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference
	QuarantineInput() interface{}
	Redirect() ZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference
	RedirectInput() interface{}
	ResolveDnsInternally() ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference
	ResolveDnsInternallyInput() interface{}
	ResolveDnsThroughCloudflare() interface{}
	SetResolveDnsThroughCloudflare(val interface{})
	ResolveDnsThroughCloudflareInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedCert() ZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference
	UntrustedCertInput() interface{}
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
	PutAuditSsh(value *ZeroTrustGatewayPolicyRuleSettingsAuditSsh)
	PutBisoAdminControls(value *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls)
	PutBlockPage(value *ZeroTrustGatewayPolicyRuleSettingsBlockPage)
	PutCheckSession(value *ZeroTrustGatewayPolicyRuleSettingsCheckSession)
	PutDnsResolvers(value *ZeroTrustGatewayPolicyRuleSettingsDnsResolvers)
	PutEgress(value *ZeroTrustGatewayPolicyRuleSettingsEgress)
	PutL4Override(value *ZeroTrustGatewayPolicyRuleSettingsL4Override)
	PutNotificationSettings(value *ZeroTrustGatewayPolicyRuleSettingsNotificationSettings)
	PutPayloadLog(value *ZeroTrustGatewayPolicyRuleSettingsPayloadLog)
	PutQuarantine(value *ZeroTrustGatewayPolicyRuleSettingsQuarantine)
	PutRedirect(value *ZeroTrustGatewayPolicyRuleSettingsRedirect)
	PutResolveDnsInternally(value *ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally)
	PutUntrustedCert(value *ZeroTrustGatewayPolicyRuleSettingsUntrustedCert)
	ResetAddHeaders()
	ResetAllowChildBypass()
	ResetAuditSsh()
	ResetBisoAdminControls()
	ResetBlockPage()
	ResetBlockPageEnabled()
	ResetBlockReason()
	ResetBypassParentRule()
	ResetCheckSession()
	ResetDnsResolvers()
	ResetEgress()
	ResetIgnoreCnameCategoryMatches()
	ResetInsecureDisableDnssecValidation()
	ResetIpCategories()
	ResetIpIndicatorFeeds()
	ResetL4Override()
	ResetNotificationSettings()
	ResetOverrideHost()
	ResetOverrideIps()
	ResetPayloadLog()
	ResetQuarantine()
	ResetRedirect()
	ResetResolveDnsInternally()
	ResetResolveDnsThroughCloudflare()
	ResetUntrustedCert()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewayPolicyRuleSettingsOutputReference
type jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AddHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"addHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AddHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"addHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AllowChildBypass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowChildBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AllowChildBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowChildBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AuditSsh() ZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsAuditSshOutputReference
	_jsii_.Get(
		j,
		"auditSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) AuditSshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BisoAdminControls() ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsBisoAdminControlsOutputReference
	_jsii_.Get(
		j,
		"bisoAdminControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BisoAdminControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisoAdminControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPage() ZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockPageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BlockReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BypassParentRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassParentRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) BypassParentRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassParentRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) CheckSession() ZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsCheckSessionOutputReference
	_jsii_.Get(
		j,
		"checkSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) CheckSessionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) DnsResolvers() ZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsDnsResolversOutputReference
	_jsii_.Get(
		j,
		"dnsResolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) DnsResolversInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsResolversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) Egress() ZeroTrustGatewayPolicyRuleSettingsEgressOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsEgressOutputReference
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) EgressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IgnoreCnameCategoryMatches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCnameCategoryMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IgnoreCnameCategoryMatchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCnameCategoryMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) InsecureDisableDnssecValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) InsecureDisableDnssecValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IpCategories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IpCategoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IpIndicatorFeeds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipIndicatorFeeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) IpIndicatorFeedsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipIndicatorFeedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) L4Override() ZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsL4OverrideOutputReference
	_jsii_.Get(
		j,
		"l4Override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) L4OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"l4OverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) NotificationSettings() ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) NotificationSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) OverrideIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PayloadLog() ZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PayloadLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payloadLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) Quarantine() ZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsQuarantineOutputReference
	_jsii_.Get(
		j,
		"quarantine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) QuarantineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quarantineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) Redirect() ZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) RedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsInternally() ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternallyOutputReference
	_jsii_.Get(
		j,
		"resolveDnsInternally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsInternallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveDnsInternallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsThroughCloudflare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveDnsThroughCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResolveDnsThroughCloudflareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveDnsThroughCloudflareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) UntrustedCert() ZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference {
	var returns ZeroTrustGatewayPolicyRuleSettingsUntrustedCertOutputReference
	_jsii_.Get(
		j,
		"untrustedCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) UntrustedCertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"untrustedCertInput",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewayPolicyRuleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewayPolicyRuleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewayPolicyRuleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewayPolicyRuleSettingsOutputReference_Override(z ZeroTrustGatewayPolicyRuleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetAddHeaders(val *map[string]*string) {
	if err := j.validateSetAddHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addHeaders",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetAllowChildBypass(val interface{}) {
	if err := j.validateSetAllowChildBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowChildBypass",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetBlockPageEnabled(val interface{}) {
	if err := j.validateSetBlockPageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPageEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetBlockReason(val *string) {
	if err := j.validateSetBlockReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockReason",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetBypassParentRule(val interface{}) {
	if err := j.validateSetBypassParentRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassParentRule",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetIgnoreCnameCategoryMatches(val interface{}) {
	if err := j.validateSetIgnoreCnameCategoryMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCnameCategoryMatches",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetInsecureDisableDnssecValidation(val interface{}) {
	if err := j.validateSetInsecureDisableDnssecValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureDisableDnssecValidation",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetIpCategories(val interface{}) {
	if err := j.validateSetIpCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipCategories",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetIpIndicatorFeeds(val interface{}) {
	if err := j.validateSetIpIndicatorFeedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipIndicatorFeeds",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetOverrideHost(val *string) {
	if err := j.validateSetOverrideHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideHost",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetOverrideIps(val *[]*string) {
	if err := j.validateSetOverrideIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideIps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetResolveDnsThroughCloudflare(val interface{}) {
	if err := j.validateSetResolveDnsThroughCloudflareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveDnsThroughCloudflare",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutAuditSsh(value *ZeroTrustGatewayPolicyRuleSettingsAuditSsh) {
	if err := z.validatePutAuditSshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuditSsh",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutBisoAdminControls(value *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls) {
	if err := z.validatePutBisoAdminControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBisoAdminControls",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutBlockPage(value *ZeroTrustGatewayPolicyRuleSettingsBlockPage) {
	if err := z.validatePutBlockPageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBlockPage",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutCheckSession(value *ZeroTrustGatewayPolicyRuleSettingsCheckSession) {
	if err := z.validatePutCheckSessionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCheckSession",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutDnsResolvers(value *ZeroTrustGatewayPolicyRuleSettingsDnsResolvers) {
	if err := z.validatePutDnsResolversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDnsResolvers",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutEgress(value *ZeroTrustGatewayPolicyRuleSettingsEgress) {
	if err := z.validatePutEgressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEgress",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutL4Override(value *ZeroTrustGatewayPolicyRuleSettingsL4Override) {
	if err := z.validatePutL4OverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putL4Override",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutNotificationSettings(value *ZeroTrustGatewayPolicyRuleSettingsNotificationSettings) {
	if err := z.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutPayloadLog(value *ZeroTrustGatewayPolicyRuleSettingsPayloadLog) {
	if err := z.validatePutPayloadLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putPayloadLog",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutQuarantine(value *ZeroTrustGatewayPolicyRuleSettingsQuarantine) {
	if err := z.validatePutQuarantineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putQuarantine",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutRedirect(value *ZeroTrustGatewayPolicyRuleSettingsRedirect) {
	if err := z.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putRedirect",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutResolveDnsInternally(value *ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally) {
	if err := z.validatePutResolveDnsInternallyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putResolveDnsInternally",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) PutUntrustedCert(value *ZeroTrustGatewayPolicyRuleSettingsUntrustedCert) {
	if err := z.validatePutUntrustedCertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putUntrustedCert",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetAddHeaders() {
	_jsii_.InvokeVoid(
		z,
		"resetAddHeaders",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetAllowChildBypass() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowChildBypass",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetAuditSsh() {
	_jsii_.InvokeVoid(
		z,
		"resetAuditSsh",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetBisoAdminControls() {
	_jsii_.InvokeVoid(
		z,
		"resetBisoAdminControls",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetBlockPage() {
	_jsii_.InvokeVoid(
		z,
		"resetBlockPage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetBlockPageEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetBlockPageEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetBlockReason() {
	_jsii_.InvokeVoid(
		z,
		"resetBlockReason",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetBypassParentRule() {
	_jsii_.InvokeVoid(
		z,
		"resetBypassParentRule",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetCheckSession() {
	_jsii_.InvokeVoid(
		z,
		"resetCheckSession",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetDnsResolvers() {
	_jsii_.InvokeVoid(
		z,
		"resetDnsResolvers",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetEgress() {
	_jsii_.InvokeVoid(
		z,
		"resetEgress",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetIgnoreCnameCategoryMatches() {
	_jsii_.InvokeVoid(
		z,
		"resetIgnoreCnameCategoryMatches",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetInsecureDisableDnssecValidation() {
	_jsii_.InvokeVoid(
		z,
		"resetInsecureDisableDnssecValidation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetIpCategories() {
	_jsii_.InvokeVoid(
		z,
		"resetIpCategories",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetIpIndicatorFeeds() {
	_jsii_.InvokeVoid(
		z,
		"resetIpIndicatorFeeds",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetL4Override() {
	_jsii_.InvokeVoid(
		z,
		"resetL4Override",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		z,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetOverrideHost() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideHost",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetOverrideIps() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideIps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetPayloadLog() {
	_jsii_.InvokeVoid(
		z,
		"resetPayloadLog",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetQuarantine() {
	_jsii_.InvokeVoid(
		z,
		"resetQuarantine",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		z,
		"resetRedirect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetResolveDnsInternally() {
	_jsii_.InvokeVoid(
		z,
		"resetResolveDnsInternally",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetResolveDnsThroughCloudflare() {
	_jsii_.InvokeVoid(
		z,
		"resetResolveDnsThroughCloudflare",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ResetUntrustedCert() {
	_jsii_.InvokeVoid(
		z,
		"resetUntrustedCert",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyRuleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_gateway_settings cloudflare_zero_trust_gateway_settings}.
type ZeroTrustGatewaySettings interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ActivityLogEnabled() interface{}
	SetActivityLogEnabled(val interface{})
	ActivityLogEnabledInput() interface{}
	Antivirus() ZeroTrustGatewaySettingsAntivirusOutputReference
	AntivirusInput() *ZeroTrustGatewaySettingsAntivirus
	BlockPage() ZeroTrustGatewaySettingsBlockPageOutputReference
	BlockPageInput() *ZeroTrustGatewaySettingsBlockPage
	BodyScanning() ZeroTrustGatewaySettingsBodyScanningOutputReference
	BodyScanningInput() *ZeroTrustGatewaySettingsBodyScanning
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() ZeroTrustGatewaySettingsCertificateOutputReference
	CertificateInput() *ZeroTrustGatewaySettingsCertificate
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomCertificate() ZeroTrustGatewaySettingsCustomCertificateOutputReference
	CustomCertificateInput() *ZeroTrustGatewaySettingsCustomCertificate
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExtendedEmailMatching() ZeroTrustGatewaySettingsExtendedEmailMatchingOutputReference
	ExtendedEmailMatchingInput() *ZeroTrustGatewaySettingsExtendedEmailMatching
	Fips() ZeroTrustGatewaySettingsFipsOutputReference
	FipsInput() *ZeroTrustGatewaySettingsFips
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() ZeroTrustGatewaySettingsLoggingOutputReference
	LoggingInput() *ZeroTrustGatewaySettingsLogging
	// The tree node.
	Node() constructs.Node
	NonIdentityBrowserIsolationEnabled() interface{}
	SetNonIdentityBrowserIsolationEnabled(val interface{})
	NonIdentityBrowserIsolationEnabledInput() interface{}
	PayloadLog() ZeroTrustGatewaySettingsPayloadLogOutputReference
	PayloadLogInput() *ZeroTrustGatewaySettingsPayloadLog
	ProtocolDetectionEnabled() interface{}
	SetProtocolDetectionEnabled(val interface{})
	ProtocolDetectionEnabledInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxy() ZeroTrustGatewaySettingsProxyOutputReference
	ProxyInput() *ZeroTrustGatewaySettingsProxy
	// Experimental.
	RawOverrides() interface{}
	SshSessionLog() ZeroTrustGatewaySettingsSshSessionLogOutputReference
	SshSessionLogInput() *ZeroTrustGatewaySettingsSshSessionLog
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsDecryptEnabled() interface{}
	SetTlsDecryptEnabled(val interface{})
	TlsDecryptEnabledInput() interface{}
	UrlBrowserIsolationEnabled() interface{}
	SetUrlBrowserIsolationEnabled(val interface{})
	UrlBrowserIsolationEnabledInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAntivirus(value *ZeroTrustGatewaySettingsAntivirus)
	PutBlockPage(value *ZeroTrustGatewaySettingsBlockPage)
	PutBodyScanning(value *ZeroTrustGatewaySettingsBodyScanning)
	PutCertificate(value *ZeroTrustGatewaySettingsCertificate)
	PutCustomCertificate(value *ZeroTrustGatewaySettingsCustomCertificate)
	PutExtendedEmailMatching(value *ZeroTrustGatewaySettingsExtendedEmailMatching)
	PutFips(value *ZeroTrustGatewaySettingsFips)
	PutLogging(value *ZeroTrustGatewaySettingsLogging)
	PutPayloadLog(value *ZeroTrustGatewaySettingsPayloadLog)
	PutProxy(value *ZeroTrustGatewaySettingsProxy)
	PutSshSessionLog(value *ZeroTrustGatewaySettingsSshSessionLog)
	ResetActivityLogEnabled()
	ResetAntivirus()
	ResetBlockPage()
	ResetBodyScanning()
	ResetCertificate()
	ResetCustomCertificate()
	ResetExtendedEmailMatching()
	ResetFips()
	ResetId()
	ResetLogging()
	ResetNonIdentityBrowserIsolationEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPayloadLog()
	ResetProtocolDetectionEnabled()
	ResetProxy()
	ResetSshSessionLog()
	ResetTlsDecryptEnabled()
	ResetUrlBrowserIsolationEnabled()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for ZeroTrustGatewaySettings
type jsiiProxy_ZeroTrustGatewaySettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ActivityLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ActivityLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Antivirus() ZeroTrustGatewaySettingsAntivirusOutputReference {
	var returns ZeroTrustGatewaySettingsAntivirusOutputReference
	_jsii_.Get(
		j,
		"antivirus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) AntivirusInput() *ZeroTrustGatewaySettingsAntivirus {
	var returns *ZeroTrustGatewaySettingsAntivirus
	_jsii_.Get(
		j,
		"antivirusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) BlockPage() ZeroTrustGatewaySettingsBlockPageOutputReference {
	var returns ZeroTrustGatewaySettingsBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) BlockPageInput() *ZeroTrustGatewaySettingsBlockPage {
	var returns *ZeroTrustGatewaySettingsBlockPage
	_jsii_.Get(
		j,
		"blockPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) BodyScanning() ZeroTrustGatewaySettingsBodyScanningOutputReference {
	var returns ZeroTrustGatewaySettingsBodyScanningOutputReference
	_jsii_.Get(
		j,
		"bodyScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) BodyScanningInput() *ZeroTrustGatewaySettingsBodyScanning {
	var returns *ZeroTrustGatewaySettingsBodyScanning
	_jsii_.Get(
		j,
		"bodyScanningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Certificate() ZeroTrustGatewaySettingsCertificateOutputReference {
	var returns ZeroTrustGatewaySettingsCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) CertificateInput() *ZeroTrustGatewaySettingsCertificate {
	var returns *ZeroTrustGatewaySettingsCertificate
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) CustomCertificate() ZeroTrustGatewaySettingsCustomCertificateOutputReference {
	var returns ZeroTrustGatewaySettingsCustomCertificateOutputReference
	_jsii_.Get(
		j,
		"customCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) CustomCertificateInput() *ZeroTrustGatewaySettingsCustomCertificate {
	var returns *ZeroTrustGatewaySettingsCustomCertificate
	_jsii_.Get(
		j,
		"customCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ExtendedEmailMatching() ZeroTrustGatewaySettingsExtendedEmailMatchingOutputReference {
	var returns ZeroTrustGatewaySettingsExtendedEmailMatchingOutputReference
	_jsii_.Get(
		j,
		"extendedEmailMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ExtendedEmailMatchingInput() *ZeroTrustGatewaySettingsExtendedEmailMatching {
	var returns *ZeroTrustGatewaySettingsExtendedEmailMatching
	_jsii_.Get(
		j,
		"extendedEmailMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Fips() ZeroTrustGatewaySettingsFipsOutputReference {
	var returns ZeroTrustGatewaySettingsFipsOutputReference
	_jsii_.Get(
		j,
		"fips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) FipsInput() *ZeroTrustGatewaySettingsFips {
	var returns *ZeroTrustGatewaySettingsFips
	_jsii_.Get(
		j,
		"fipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Logging() ZeroTrustGatewaySettingsLoggingOutputReference {
	var returns ZeroTrustGatewaySettingsLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) LoggingInput() *ZeroTrustGatewaySettingsLogging {
	var returns *ZeroTrustGatewaySettingsLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) NonIdentityBrowserIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonIdentityBrowserIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) NonIdentityBrowserIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonIdentityBrowserIsolationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) PayloadLog() ZeroTrustGatewaySettingsPayloadLogOutputReference {
	var returns ZeroTrustGatewaySettingsPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) PayloadLogInput() *ZeroTrustGatewaySettingsPayloadLog {
	var returns *ZeroTrustGatewaySettingsPayloadLog
	_jsii_.Get(
		j,
		"payloadLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ProtocolDetectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolDetectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ProtocolDetectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolDetectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) Proxy() ZeroTrustGatewaySettingsProxyOutputReference {
	var returns ZeroTrustGatewaySettingsProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) ProxyInput() *ZeroTrustGatewaySettingsProxy {
	var returns *ZeroTrustGatewaySettingsProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) SshSessionLog() ZeroTrustGatewaySettingsSshSessionLogOutputReference {
	var returns ZeroTrustGatewaySettingsSshSessionLogOutputReference
	_jsii_.Get(
		j,
		"sshSessionLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) SshSessionLogInput() *ZeroTrustGatewaySettingsSshSessionLog {
	var returns *ZeroTrustGatewaySettingsSshSessionLog
	_jsii_.Get(
		j,
		"sshSessionLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) TlsDecryptEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsDecryptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) TlsDecryptEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsDecryptEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) UrlBrowserIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlBrowserIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettings) UrlBrowserIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlBrowserIsolationEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_gateway_settings cloudflare_zero_trust_gateway_settings} Resource.
func NewZeroTrustGatewaySettings(scope constructs.Construct, id *string, config *ZeroTrustGatewaySettingsConfig) ZeroTrustGatewaySettings {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettings{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/zero_trust_gateway_settings cloudflare_zero_trust_gateway_settings} Resource.
func NewZeroTrustGatewaySettings_Override(z ZeroTrustGatewaySettings, scope constructs.Construct, id *string, config *ZeroTrustGatewaySettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetActivityLogEnabled(val interface{}) {
	if err := j.validateSetActivityLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activityLogEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetNonIdentityBrowserIsolationEnabled(val interface{}) {
	if err := j.validateSetNonIdentityBrowserIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonIdentityBrowserIsolationEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetProtocolDetectionEnabled(val interface{}) {
	if err := j.validateSetProtocolDetectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolDetectionEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetTlsDecryptEnabled(val interface{}) {
	if err := j.validateSetTlsDecryptEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsDecryptEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettings)SetUrlBrowserIsolationEnabled(val interface{}) {
	if err := j.validateSetUrlBrowserIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlBrowserIsolationEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustGatewaySettings resource upon running "cdktf plan <stack-name>".
func ZeroTrustGatewaySettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustGatewaySettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
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
func ZeroTrustGatewaySettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustGatewaySettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustGatewaySettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustGatewaySettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustGatewaySettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZeroTrustGatewaySettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZeroTrustGatewaySettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutAntivirus(value *ZeroTrustGatewaySettingsAntivirus) {
	if err := z.validatePutAntivirusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAntivirus",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutBlockPage(value *ZeroTrustGatewaySettingsBlockPage) {
	if err := z.validatePutBlockPageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBlockPage",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutBodyScanning(value *ZeroTrustGatewaySettingsBodyScanning) {
	if err := z.validatePutBodyScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBodyScanning",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutCertificate(value *ZeroTrustGatewaySettingsCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutCustomCertificate(value *ZeroTrustGatewaySettingsCustomCertificate) {
	if err := z.validatePutCustomCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCustomCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutExtendedEmailMatching(value *ZeroTrustGatewaySettingsExtendedEmailMatching) {
	if err := z.validatePutExtendedEmailMatchingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExtendedEmailMatching",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutFips(value *ZeroTrustGatewaySettingsFips) {
	if err := z.validatePutFipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putFips",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutLogging(value *ZeroTrustGatewaySettingsLogging) {
	if err := z.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLogging",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutPayloadLog(value *ZeroTrustGatewaySettingsPayloadLog) {
	if err := z.validatePutPayloadLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putPayloadLog",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutProxy(value *ZeroTrustGatewaySettingsProxy) {
	if err := z.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putProxy",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) PutSshSessionLog(value *ZeroTrustGatewaySettingsSshSessionLog) {
	if err := z.validatePutSshSessionLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSshSessionLog",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetActivityLogEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetActivityLogEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetAntivirus() {
	_jsii_.InvokeVoid(
		z,
		"resetAntivirus",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetBlockPage() {
	_jsii_.InvokeVoid(
		z,
		"resetBlockPage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetBodyScanning() {
	_jsii_.InvokeVoid(
		z,
		"resetBodyScanning",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetCustomCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetExtendedEmailMatching() {
	_jsii_.InvokeVoid(
		z,
		"resetExtendedEmailMatching",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetFips() {
	_jsii_.InvokeVoid(
		z,
		"resetFips",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetLogging() {
	_jsii_.InvokeVoid(
		z,
		"resetLogging",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetNonIdentityBrowserIsolationEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetNonIdentityBrowserIsolationEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetPayloadLog() {
	_jsii_.InvokeVoid(
		z,
		"resetPayloadLog",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetProtocolDetectionEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetProtocolDetectionEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetProxy() {
	_jsii_.InvokeVoid(
		z,
		"resetProxy",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetSshSessionLog() {
	_jsii_.InvokeVoid(
		z,
		"resetSshSessionLog",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetTlsDecryptEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetTlsDecryptEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ResetUrlBrowserIsolationEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetUrlBrowserIsolationEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


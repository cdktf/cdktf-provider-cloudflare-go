// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustgatewaysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivityLog() ZeroTrustGatewaySettingsSettingsActivityLogOutputReference
	ActivityLogInput() interface{}
	Antivirus() ZeroTrustGatewaySettingsSettingsAntivirusOutputReference
	AntivirusInput() interface{}
	BlockPage() ZeroTrustGatewaySettingsSettingsBlockPageOutputReference
	BlockPageInput() interface{}
	BodyScanning() ZeroTrustGatewaySettingsSettingsBodyScanningOutputReference
	BodyScanningInput() interface{}
	BrowserIsolation() ZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference
	BrowserIsolationInput() interface{}
	Certificate() ZeroTrustGatewaySettingsSettingsCertificateOutputReference
	CertificateInput() interface{}
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
	CustomCertificate() ZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference
	CustomCertificateInput() interface{}
	ExtendedEmailMatching() ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference
	ExtendedEmailMatchingInput() interface{}
	Fips() ZeroTrustGatewaySettingsSettingsFipsOutputReference
	FipsInput() interface{}
	// Experimental.
	Fqn() *string
	HostSelector() ZeroTrustGatewaySettingsSettingsHostSelectorOutputReference
	HostSelectorInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProtocolDetection() ZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference
	ProtocolDetectionInput() interface{}
	Sandbox() ZeroTrustGatewaySettingsSettingsSandboxOutputReference
	SandboxInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsDecrypt() ZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference
	TlsDecryptInput() interface{}
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
	PutActivityLog(value *ZeroTrustGatewaySettingsSettingsActivityLog)
	PutAntivirus(value *ZeroTrustGatewaySettingsSettingsAntivirus)
	PutBlockPage(value *ZeroTrustGatewaySettingsSettingsBlockPage)
	PutBodyScanning(value *ZeroTrustGatewaySettingsSettingsBodyScanning)
	PutBrowserIsolation(value *ZeroTrustGatewaySettingsSettingsBrowserIsolation)
	PutCertificate(value *ZeroTrustGatewaySettingsSettingsCertificate)
	PutCustomCertificate(value *ZeroTrustGatewaySettingsSettingsCustomCertificate)
	PutExtendedEmailMatching(value *ZeroTrustGatewaySettingsSettingsExtendedEmailMatching)
	PutFips(value *ZeroTrustGatewaySettingsSettingsFips)
	PutHostSelector(value *ZeroTrustGatewaySettingsSettingsHostSelector)
	PutProtocolDetection(value *ZeroTrustGatewaySettingsSettingsProtocolDetection)
	PutSandbox(value *ZeroTrustGatewaySettingsSettingsSandbox)
	PutTlsDecrypt(value *ZeroTrustGatewaySettingsSettingsTlsDecrypt)
	ResetActivityLog()
	ResetAntivirus()
	ResetBlockPage()
	ResetBodyScanning()
	ResetBrowserIsolation()
	ResetCertificate()
	ResetCustomCertificate()
	ResetExtendedEmailMatching()
	ResetFips()
	ResetHostSelector()
	ResetProtocolDetection()
	ResetSandbox()
	ResetTlsDecrypt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewaySettingsSettingsOutputReference
type jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ActivityLog() ZeroTrustGatewaySettingsSettingsActivityLogOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsActivityLogOutputReference
	_jsii_.Get(
		j,
		"activityLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ActivityLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Antivirus() ZeroTrustGatewaySettingsSettingsAntivirusOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsAntivirusOutputReference
	_jsii_.Get(
		j,
		"antivirus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) AntivirusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antivirusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BlockPage() ZeroTrustGatewaySettingsSettingsBlockPageOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BlockPageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BodyScanning() ZeroTrustGatewaySettingsSettingsBodyScanningOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsBodyScanningOutputReference
	_jsii_.Get(
		j,
		"bodyScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BodyScanningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyScanningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BrowserIsolation() ZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference
	_jsii_.Get(
		j,
		"browserIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) BrowserIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Certificate() ZeroTrustGatewaySettingsSettingsCertificateOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) CustomCertificate() ZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference
	_jsii_.Get(
		j,
		"customCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) CustomCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ExtendedEmailMatching() ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference
	_jsii_.Get(
		j,
		"extendedEmailMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ExtendedEmailMatchingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedEmailMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Fips() ZeroTrustGatewaySettingsSettingsFipsOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsFipsOutputReference
	_jsii_.Get(
		j,
		"fips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) FipsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) HostSelector() ZeroTrustGatewaySettingsSettingsHostSelectorOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsHostSelectorOutputReference
	_jsii_.Get(
		j,
		"hostSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) HostSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ProtocolDetection() ZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference
	_jsii_.Get(
		j,
		"protocolDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ProtocolDetectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Sandbox() ZeroTrustGatewaySettingsSettingsSandboxOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsSandboxOutputReference
	_jsii_.Get(
		j,
		"sandbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) SandboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sandboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) TlsDecrypt() ZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference {
	var returns ZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference
	_jsii_.Get(
		j,
		"tlsDecrypt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) TlsDecryptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsDecryptInput",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewaySettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewaySettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewaySettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewaySettingsSettingsOutputReference_Override(z ZeroTrustGatewaySettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewaySettings.ZeroTrustGatewaySettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutActivityLog(value *ZeroTrustGatewaySettingsSettingsActivityLog) {
	if err := z.validatePutActivityLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putActivityLog",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutAntivirus(value *ZeroTrustGatewaySettingsSettingsAntivirus) {
	if err := z.validatePutAntivirusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAntivirus",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutBlockPage(value *ZeroTrustGatewaySettingsSettingsBlockPage) {
	if err := z.validatePutBlockPageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBlockPage",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutBodyScanning(value *ZeroTrustGatewaySettingsSettingsBodyScanning) {
	if err := z.validatePutBodyScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBodyScanning",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutBrowserIsolation(value *ZeroTrustGatewaySettingsSettingsBrowserIsolation) {
	if err := z.validatePutBrowserIsolationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putBrowserIsolation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutCertificate(value *ZeroTrustGatewaySettingsSettingsCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutCustomCertificate(value *ZeroTrustGatewaySettingsSettingsCustomCertificate) {
	if err := z.validatePutCustomCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCustomCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutExtendedEmailMatching(value *ZeroTrustGatewaySettingsSettingsExtendedEmailMatching) {
	if err := z.validatePutExtendedEmailMatchingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExtendedEmailMatching",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutFips(value *ZeroTrustGatewaySettingsSettingsFips) {
	if err := z.validatePutFipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putFips",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutHostSelector(value *ZeroTrustGatewaySettingsSettingsHostSelector) {
	if err := z.validatePutHostSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putHostSelector",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutProtocolDetection(value *ZeroTrustGatewaySettingsSettingsProtocolDetection) {
	if err := z.validatePutProtocolDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putProtocolDetection",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutSandbox(value *ZeroTrustGatewaySettingsSettingsSandbox) {
	if err := z.validatePutSandboxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSandbox",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) PutTlsDecrypt(value *ZeroTrustGatewaySettingsSettingsTlsDecrypt) {
	if err := z.validatePutTlsDecryptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putTlsDecrypt",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetActivityLog() {
	_jsii_.InvokeVoid(
		z,
		"resetActivityLog",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetAntivirus() {
	_jsii_.InvokeVoid(
		z,
		"resetAntivirus",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetBlockPage() {
	_jsii_.InvokeVoid(
		z,
		"resetBlockPage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetBodyScanning() {
	_jsii_.InvokeVoid(
		z,
		"resetBodyScanning",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetBrowserIsolation() {
	_jsii_.InvokeVoid(
		z,
		"resetBrowserIsolation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetCustomCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetExtendedEmailMatching() {
	_jsii_.InvokeVoid(
		z,
		"resetExtendedEmailMatching",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetFips() {
	_jsii_.InvokeVoid(
		z,
		"resetFips",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetHostSelector() {
	_jsii_.InvokeVoid(
		z,
		"resetHostSelector",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetProtocolDetection() {
	_jsii_.InvokeVoid(
		z,
		"resetProtocolDetection",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetSandbox() {
	_jsii_.InvokeVoid(
		z,
		"resetSandbox",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ResetTlsDecrypt() {
	_jsii_.InvokeVoid(
		z,
		"resetTlsDecrypt",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustGatewaySettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


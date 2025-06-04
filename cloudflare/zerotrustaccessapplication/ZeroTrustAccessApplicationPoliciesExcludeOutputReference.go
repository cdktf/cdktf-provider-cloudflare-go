// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessApplicationPoliciesExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference
	AuthContextInput() interface{}
	AuthMethod() ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference
	AuthMethodInput() interface{}
	AzureAd() ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference
	AzureAdInput() interface{}
	Certificate() ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference
	CertificateInput() interface{}
	CommonName() ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference
	CommonNameInput() interface{}
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
	DevicePosture() ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference
	DevicePostureInput() interface{}
	Email() ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference
	EmailDomain() ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference
	EmailDomainInput() interface{}
	EmailInput() interface{}
	EmailList() ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference
	EmailListInput() interface{}
	Everyone() ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference
	GeoInput() interface{}
	GithubOrganization() ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference
	GithubOrganizationInput() interface{}
	Group() ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference
	GroupInput() interface{}
	Gsuite() ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference
	IpInput() interface{}
	IpList() ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference
	IpListInput() interface{}
	LoginMethod() ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference
	LoginMethodInput() interface{}
	Okta() ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference
	OktaInput() interface{}
	Saml() ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference
	SamlInput() interface{}
	ServiceToken() ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference
	ServiceTokenInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessApplicationPoliciesExcludeAuthContext)
	PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesExcludeAuthMethod)
	PutAzureAd(value *ZeroTrustAccessApplicationPoliciesExcludeAzureAd)
	PutCertificate(value *ZeroTrustAccessApplicationPoliciesExcludeCertificate)
	PutCommonName(value *ZeroTrustAccessApplicationPoliciesExcludeCommonName)
	PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesExcludeDevicePosture)
	PutEmail(value *ZeroTrustAccessApplicationPoliciesExcludeEmail)
	PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesExcludeEmailDomain)
	PutEmailList(value *ZeroTrustAccessApplicationPoliciesExcludeEmailListStruct)
	PutEveryone(value *ZeroTrustAccessApplicationPoliciesExcludeEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluation)
	PutGeo(value *ZeroTrustAccessApplicationPoliciesExcludeGeo)
	PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesExcludeGithubOrganization)
	PutGroup(value *ZeroTrustAccessApplicationPoliciesExcludeGroup)
	PutGsuite(value *ZeroTrustAccessApplicationPoliciesExcludeGsuite)
	PutIp(value *ZeroTrustAccessApplicationPoliciesExcludeIp)
	PutIpList(value *ZeroTrustAccessApplicationPoliciesExcludeIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessApplicationPoliciesExcludeLoginMethod)
	PutOkta(value *ZeroTrustAccessApplicationPoliciesExcludeOkta)
	PutSaml(value *ZeroTrustAccessApplicationPoliciesExcludeSaml)
	PutServiceToken(value *ZeroTrustAccessApplicationPoliciesExcludeServiceToken)
	ResetAnyValidServiceToken()
	ResetAuthContext()
	ResetAuthMethod()
	ResetAzureAd()
	ResetCertificate()
	ResetCommonName()
	ResetDevicePosture()
	ResetEmail()
	ResetEmailDomain()
	ResetEmailList()
	ResetEveryone()
	ResetExternalEvaluation()
	ResetGeo()
	ResetGithubOrganization()
	ResetGroup()
	ResetGsuite()
	ResetIp()
	ResetIpList()
	ResetLoginMethod()
	ResetOkta()
	ResetSaml()
	ResetServiceToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationPoliciesExcludeOutputReference
type jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthContext() ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthMethod() ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AzureAd() ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) AzureAdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Certificate() ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) CommonName() ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) CommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) DevicePosture() ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) DevicePostureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Email() ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailDomain() ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailList() ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Everyone() ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ExternalEvaluation() ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Geo() ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GithubOrganization() ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GithubOrganizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Group() ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Gsuite() ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Ip() ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) IpList() ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) IpListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) LoginMethod() ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) LoginMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Okta() ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Saml() ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ServiceToken() ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessApplicationPoliciesExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessApplicationPoliciesExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationPoliciesExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationPoliciesExcludeOutputReference_Override(z ZeroTrustAccessApplicationPoliciesExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutAuthContext(value *ZeroTrustAccessApplicationPoliciesExcludeAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesExcludeAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutAzureAd(value *ZeroTrustAccessApplicationPoliciesExcludeAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutCertificate(value *ZeroTrustAccessApplicationPoliciesExcludeCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutCommonName(value *ZeroTrustAccessApplicationPoliciesExcludeCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesExcludeDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutEmail(value *ZeroTrustAccessApplicationPoliciesExcludeEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesExcludeEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutEmailList(value *ZeroTrustAccessApplicationPoliciesExcludeEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutEveryone(value *ZeroTrustAccessApplicationPoliciesExcludeEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutGeo(value *ZeroTrustAccessApplicationPoliciesExcludeGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesExcludeGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutGroup(value *ZeroTrustAccessApplicationPoliciesExcludeGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutGsuite(value *ZeroTrustAccessApplicationPoliciesExcludeGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutIp(value *ZeroTrustAccessApplicationPoliciesExcludeIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutIpList(value *ZeroTrustAccessApplicationPoliciesExcludeIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutLoginMethod(value *ZeroTrustAccessApplicationPoliciesExcludeLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutOkta(value *ZeroTrustAccessApplicationPoliciesExcludeOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutSaml(value *ZeroTrustAccessApplicationPoliciesExcludeSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) PutServiceToken(value *ZeroTrustAccessApplicationPoliciesExcludeServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


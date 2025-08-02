// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessPolicyExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessPolicyExcludeAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessPolicyExcludeAuthContextOutputReference
	AuthContextInput() interface{}
	AuthMethod() ZeroTrustAccessPolicyExcludeAuthMethodOutputReference
	AuthMethodInput() interface{}
	AzureAd() ZeroTrustAccessPolicyExcludeAzureAdOutputReference
	AzureAdInput() interface{}
	Certificate() ZeroTrustAccessPolicyExcludeCertificateOutputReference
	CertificateInput() interface{}
	CommonName() ZeroTrustAccessPolicyExcludeCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessPolicyExcludeDevicePostureOutputReference
	DevicePostureInput() interface{}
	Email() ZeroTrustAccessPolicyExcludeEmailOutputReference
	EmailDomain() ZeroTrustAccessPolicyExcludeEmailDomainOutputReference
	EmailDomainInput() interface{}
	EmailInput() interface{}
	EmailList() ZeroTrustAccessPolicyExcludeEmailListStructOutputReference
	EmailListInput() interface{}
	Everyone() ZeroTrustAccessPolicyExcludeEveryoneOutputReference
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessPolicyExcludeExternalEvaluationOutputReference
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessPolicyExcludeGeoOutputReference
	GeoInput() interface{}
	GithubOrganization() ZeroTrustAccessPolicyExcludeGithubOrganizationOutputReference
	GithubOrganizationInput() interface{}
	Group() ZeroTrustAccessPolicyExcludeGroupOutputReference
	GroupInput() interface{}
	Gsuite() ZeroTrustAccessPolicyExcludeGsuiteOutputReference
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() ZeroTrustAccessPolicyExcludeIpOutputReference
	IpInput() interface{}
	IpList() ZeroTrustAccessPolicyExcludeIpListStructOutputReference
	IpListInput() interface{}
	LinkedAppToken() ZeroTrustAccessPolicyExcludeLinkedAppTokenOutputReference
	LinkedAppTokenInput() interface{}
	LoginMethod() ZeroTrustAccessPolicyExcludeLoginMethodOutputReference
	LoginMethodInput() interface{}
	Oidc() ZeroTrustAccessPolicyExcludeOidcOutputReference
	OidcInput() interface{}
	Okta() ZeroTrustAccessPolicyExcludeOktaOutputReference
	OktaInput() interface{}
	Saml() ZeroTrustAccessPolicyExcludeSamlOutputReference
	SamlInput() interface{}
	ServiceToken() ZeroTrustAccessPolicyExcludeServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessPolicyExcludeAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessPolicyExcludeAuthContext)
	PutAuthMethod(value *ZeroTrustAccessPolicyExcludeAuthMethod)
	PutAzureAd(value *ZeroTrustAccessPolicyExcludeAzureAd)
	PutCertificate(value *ZeroTrustAccessPolicyExcludeCertificate)
	PutCommonName(value *ZeroTrustAccessPolicyExcludeCommonName)
	PutDevicePosture(value *ZeroTrustAccessPolicyExcludeDevicePosture)
	PutEmail(value *ZeroTrustAccessPolicyExcludeEmail)
	PutEmailDomain(value *ZeroTrustAccessPolicyExcludeEmailDomain)
	PutEmailList(value *ZeroTrustAccessPolicyExcludeEmailListStruct)
	PutEveryone(value *ZeroTrustAccessPolicyExcludeEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessPolicyExcludeExternalEvaluation)
	PutGeo(value *ZeroTrustAccessPolicyExcludeGeo)
	PutGithubOrganization(value *ZeroTrustAccessPolicyExcludeGithubOrganization)
	PutGroup(value *ZeroTrustAccessPolicyExcludeGroup)
	PutGsuite(value *ZeroTrustAccessPolicyExcludeGsuite)
	PutIp(value *ZeroTrustAccessPolicyExcludeIp)
	PutIpList(value *ZeroTrustAccessPolicyExcludeIpListStruct)
	PutLinkedAppToken(value *ZeroTrustAccessPolicyExcludeLinkedAppToken)
	PutLoginMethod(value *ZeroTrustAccessPolicyExcludeLoginMethod)
	PutOidc(value *ZeroTrustAccessPolicyExcludeOidc)
	PutOkta(value *ZeroTrustAccessPolicyExcludeOkta)
	PutSaml(value *ZeroTrustAccessPolicyExcludeSaml)
	PutServiceToken(value *ZeroTrustAccessPolicyExcludeServiceToken)
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
	ResetLinkedAppToken()
	ResetLoginMethod()
	ResetOidc()
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

// The jsii proxy struct for ZeroTrustAccessPolicyExcludeOutputReference
type jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AnyValidServiceToken() ZeroTrustAccessPolicyExcludeAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthContext() ZeroTrustAccessPolicyExcludeAuthContextOutputReference {
	var returns ZeroTrustAccessPolicyExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthMethod() ZeroTrustAccessPolicyExcludeAuthMethodOutputReference {
	var returns ZeroTrustAccessPolicyExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AzureAd() ZeroTrustAccessPolicyExcludeAzureAdOutputReference {
	var returns ZeroTrustAccessPolicyExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AzureAdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Certificate() ZeroTrustAccessPolicyExcludeCertificateOutputReference {
	var returns ZeroTrustAccessPolicyExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonName() ZeroTrustAccessPolicyExcludeCommonNameOutputReference {
	var returns ZeroTrustAccessPolicyExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) DevicePosture() ZeroTrustAccessPolicyExcludeDevicePostureOutputReference {
	var returns ZeroTrustAccessPolicyExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) DevicePostureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Email() ZeroTrustAccessPolicyExcludeEmailOutputReference {
	var returns ZeroTrustAccessPolicyExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailDomain() ZeroTrustAccessPolicyExcludeEmailDomainOutputReference {
	var returns ZeroTrustAccessPolicyExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailList() ZeroTrustAccessPolicyExcludeEmailListStructOutputReference {
	var returns ZeroTrustAccessPolicyExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Everyone() ZeroTrustAccessPolicyExcludeEveryoneOutputReference {
	var returns ZeroTrustAccessPolicyExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ExternalEvaluation() ZeroTrustAccessPolicyExcludeExternalEvaluationOutputReference {
	var returns ZeroTrustAccessPolicyExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Geo() ZeroTrustAccessPolicyExcludeGeoOutputReference {
	var returns ZeroTrustAccessPolicyExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GithubOrganization() ZeroTrustAccessPolicyExcludeGithubOrganizationOutputReference {
	var returns ZeroTrustAccessPolicyExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GithubOrganizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Group() ZeroTrustAccessPolicyExcludeGroupOutputReference {
	var returns ZeroTrustAccessPolicyExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Gsuite() ZeroTrustAccessPolicyExcludeGsuiteOutputReference {
	var returns ZeroTrustAccessPolicyExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Ip() ZeroTrustAccessPolicyExcludeIpOutputReference {
	var returns ZeroTrustAccessPolicyExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpList() ZeroTrustAccessPolicyExcludeIpListStructOutputReference {
	var returns ZeroTrustAccessPolicyExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LinkedAppToken() ZeroTrustAccessPolicyExcludeLinkedAppTokenOutputReference {
	var returns ZeroTrustAccessPolicyExcludeLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LinkedAppTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedAppTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LoginMethod() ZeroTrustAccessPolicyExcludeLoginMethodOutputReference {
	var returns ZeroTrustAccessPolicyExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LoginMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Oidc() ZeroTrustAccessPolicyExcludeOidcOutputReference {
	var returns ZeroTrustAccessPolicyExcludeOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) OidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Okta() ZeroTrustAccessPolicyExcludeOktaOutputReference {
	var returns ZeroTrustAccessPolicyExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Saml() ZeroTrustAccessPolicyExcludeSamlOutputReference {
	var returns ZeroTrustAccessPolicyExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ServiceToken() ZeroTrustAccessPolicyExcludeServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessPolicyExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessPolicyExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessPolicyExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessPolicyExcludeOutputReference_Override(z ZeroTrustAccessPolicyExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessPolicyExcludeAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAuthContext(value *ZeroTrustAccessPolicyExcludeAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAuthMethod(value *ZeroTrustAccessPolicyExcludeAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAzureAd(value *ZeroTrustAccessPolicyExcludeAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutCertificate(value *ZeroTrustAccessPolicyExcludeCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutCommonName(value *ZeroTrustAccessPolicyExcludeCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutDevicePosture(value *ZeroTrustAccessPolicyExcludeDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutEmail(value *ZeroTrustAccessPolicyExcludeEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutEmailDomain(value *ZeroTrustAccessPolicyExcludeEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutEmailList(value *ZeroTrustAccessPolicyExcludeEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutEveryone(value *ZeroTrustAccessPolicyExcludeEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutExternalEvaluation(value *ZeroTrustAccessPolicyExcludeExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGeo(value *ZeroTrustAccessPolicyExcludeGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGithubOrganization(value *ZeroTrustAccessPolicyExcludeGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGroup(value *ZeroTrustAccessPolicyExcludeGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGsuite(value *ZeroTrustAccessPolicyExcludeGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutIp(value *ZeroTrustAccessPolicyExcludeIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutIpList(value *ZeroTrustAccessPolicyExcludeIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutLinkedAppToken(value *ZeroTrustAccessPolicyExcludeLinkedAppToken) {
	if err := z.validatePutLinkedAppTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLinkedAppToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutLoginMethod(value *ZeroTrustAccessPolicyExcludeLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutOidc(value *ZeroTrustAccessPolicyExcludeOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutOkta(value *ZeroTrustAccessPolicyExcludeOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutSaml(value *ZeroTrustAccessPolicyExcludeSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutServiceToken(value *ZeroTrustAccessPolicyExcludeServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetLinkedAppToken() {
	_jsii_.InvokeVoid(
		z,
		"resetLinkedAppToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


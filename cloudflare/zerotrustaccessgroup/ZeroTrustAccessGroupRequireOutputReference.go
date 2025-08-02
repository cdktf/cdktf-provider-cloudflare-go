// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zerotrustaccessgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessGroupRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessGroupRequireAuthContextOutputReference
	AuthContextInput() interface{}
	AuthMethod() ZeroTrustAccessGroupRequireAuthMethodOutputReference
	AuthMethodInput() interface{}
	AzureAd() ZeroTrustAccessGroupRequireAzureAdOutputReference
	AzureAdInput() interface{}
	Certificate() ZeroTrustAccessGroupRequireCertificateOutputReference
	CertificateInput() interface{}
	CommonName() ZeroTrustAccessGroupRequireCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessGroupRequireDevicePostureOutputReference
	DevicePostureInput() interface{}
	Email() ZeroTrustAccessGroupRequireEmailOutputReference
	EmailDomain() ZeroTrustAccessGroupRequireEmailDomainOutputReference
	EmailDomainInput() interface{}
	EmailInput() interface{}
	EmailList() ZeroTrustAccessGroupRequireEmailListStructOutputReference
	EmailListInput() interface{}
	Everyone() ZeroTrustAccessGroupRequireEveryoneOutputReference
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessGroupRequireExternalEvaluationOutputReference
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessGroupRequireGeoOutputReference
	GeoInput() interface{}
	GithubOrganization() ZeroTrustAccessGroupRequireGithubOrganizationOutputReference
	GithubOrganizationInput() interface{}
	Group() ZeroTrustAccessGroupRequireGroupOutputReference
	GroupInput() interface{}
	Gsuite() ZeroTrustAccessGroupRequireGsuiteOutputReference
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() ZeroTrustAccessGroupRequireIpOutputReference
	IpInput() interface{}
	IpList() ZeroTrustAccessGroupRequireIpListStructOutputReference
	IpListInput() interface{}
	LinkedAppToken() ZeroTrustAccessGroupRequireLinkedAppTokenOutputReference
	LinkedAppTokenInput() interface{}
	LoginMethod() ZeroTrustAccessGroupRequireLoginMethodOutputReference
	LoginMethodInput() interface{}
	Oidc() ZeroTrustAccessGroupRequireOidcOutputReference
	OidcInput() interface{}
	Okta() ZeroTrustAccessGroupRequireOktaOutputReference
	OktaInput() interface{}
	Saml() ZeroTrustAccessGroupRequireSamlOutputReference
	SamlInput() interface{}
	ServiceToken() ZeroTrustAccessGroupRequireServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessGroupRequireAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessGroupRequireAuthContext)
	PutAuthMethod(value *ZeroTrustAccessGroupRequireAuthMethod)
	PutAzureAd(value *ZeroTrustAccessGroupRequireAzureAd)
	PutCertificate(value *ZeroTrustAccessGroupRequireCertificate)
	PutCommonName(value *ZeroTrustAccessGroupRequireCommonName)
	PutDevicePosture(value *ZeroTrustAccessGroupRequireDevicePosture)
	PutEmail(value *ZeroTrustAccessGroupRequireEmail)
	PutEmailDomain(value *ZeroTrustAccessGroupRequireEmailDomain)
	PutEmailList(value *ZeroTrustAccessGroupRequireEmailListStruct)
	PutEveryone(value *ZeroTrustAccessGroupRequireEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessGroupRequireExternalEvaluation)
	PutGeo(value *ZeroTrustAccessGroupRequireGeo)
	PutGithubOrganization(value *ZeroTrustAccessGroupRequireGithubOrganization)
	PutGroup(value *ZeroTrustAccessGroupRequireGroup)
	PutGsuite(value *ZeroTrustAccessGroupRequireGsuite)
	PutIp(value *ZeroTrustAccessGroupRequireIp)
	PutIpList(value *ZeroTrustAccessGroupRequireIpListStruct)
	PutLinkedAppToken(value *ZeroTrustAccessGroupRequireLinkedAppToken)
	PutLoginMethod(value *ZeroTrustAccessGroupRequireLoginMethod)
	PutOidc(value *ZeroTrustAccessGroupRequireOidc)
	PutOkta(value *ZeroTrustAccessGroupRequireOkta)
	PutSaml(value *ZeroTrustAccessGroupRequireSaml)
	PutServiceToken(value *ZeroTrustAccessGroupRequireServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessGroupRequireOutputReference
type jsiiProxy_ZeroTrustAccessGroupRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AnyValidServiceToken() ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthContext() ZeroTrustAccessGroupRequireAuthContextOutputReference {
	var returns ZeroTrustAccessGroupRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthMethod() ZeroTrustAccessGroupRequireAuthMethodOutputReference {
	var returns ZeroTrustAccessGroupRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AzureAd() ZeroTrustAccessGroupRequireAzureAdOutputReference {
	var returns ZeroTrustAccessGroupRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AzureAdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Certificate() ZeroTrustAccessGroupRequireCertificateOutputReference {
	var returns ZeroTrustAccessGroupRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CommonName() ZeroTrustAccessGroupRequireCommonNameOutputReference {
	var returns ZeroTrustAccessGroupRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) DevicePosture() ZeroTrustAccessGroupRequireDevicePostureOutputReference {
	var returns ZeroTrustAccessGroupRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) DevicePostureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Email() ZeroTrustAccessGroupRequireEmailOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailDomain() ZeroTrustAccessGroupRequireEmailDomainOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailList() ZeroTrustAccessGroupRequireEmailListStructOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Everyone() ZeroTrustAccessGroupRequireEveryoneOutputReference {
	var returns ZeroTrustAccessGroupRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ExternalEvaluation() ZeroTrustAccessGroupRequireExternalEvaluationOutputReference {
	var returns ZeroTrustAccessGroupRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Geo() ZeroTrustAccessGroupRequireGeoOutputReference {
	var returns ZeroTrustAccessGroupRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GithubOrganization() ZeroTrustAccessGroupRequireGithubOrganizationOutputReference {
	var returns ZeroTrustAccessGroupRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GithubOrganizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Group() ZeroTrustAccessGroupRequireGroupOutputReference {
	var returns ZeroTrustAccessGroupRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Gsuite() ZeroTrustAccessGroupRequireGsuiteOutputReference {
	var returns ZeroTrustAccessGroupRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Ip() ZeroTrustAccessGroupRequireIpOutputReference {
	var returns ZeroTrustAccessGroupRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpList() ZeroTrustAccessGroupRequireIpListStructOutputReference {
	var returns ZeroTrustAccessGroupRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LinkedAppToken() ZeroTrustAccessGroupRequireLinkedAppTokenOutputReference {
	var returns ZeroTrustAccessGroupRequireLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LinkedAppTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedAppTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LoginMethod() ZeroTrustAccessGroupRequireLoginMethodOutputReference {
	var returns ZeroTrustAccessGroupRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LoginMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Oidc() ZeroTrustAccessGroupRequireOidcOutputReference {
	var returns ZeroTrustAccessGroupRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) OidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Okta() ZeroTrustAccessGroupRequireOktaOutputReference {
	var returns ZeroTrustAccessGroupRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Saml() ZeroTrustAccessGroupRequireSamlOutputReference {
	var returns ZeroTrustAccessGroupRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ServiceToken() ZeroTrustAccessGroupRequireServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessGroupRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessGroupRequireOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessGroupRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessGroupRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessGroupRequireOutputReference_Override(z ZeroTrustAccessGroupRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessGroupRequireAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAuthContext(value *ZeroTrustAccessGroupRequireAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAuthMethod(value *ZeroTrustAccessGroupRequireAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAzureAd(value *ZeroTrustAccessGroupRequireAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutCertificate(value *ZeroTrustAccessGroupRequireCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutCommonName(value *ZeroTrustAccessGroupRequireCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutDevicePosture(value *ZeroTrustAccessGroupRequireDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmail(value *ZeroTrustAccessGroupRequireEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmailDomain(value *ZeroTrustAccessGroupRequireEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmailList(value *ZeroTrustAccessGroupRequireEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEveryone(value *ZeroTrustAccessGroupRequireEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutExternalEvaluation(value *ZeroTrustAccessGroupRequireExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGeo(value *ZeroTrustAccessGroupRequireGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGithubOrganization(value *ZeroTrustAccessGroupRequireGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGroup(value *ZeroTrustAccessGroupRequireGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGsuite(value *ZeroTrustAccessGroupRequireGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutIp(value *ZeroTrustAccessGroupRequireIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutIpList(value *ZeroTrustAccessGroupRequireIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutLinkedAppToken(value *ZeroTrustAccessGroupRequireLinkedAppToken) {
	if err := z.validatePutLinkedAppTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLinkedAppToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutLoginMethod(value *ZeroTrustAccessGroupRequireLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutOidc(value *ZeroTrustAccessGroupRequireOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutOkta(value *ZeroTrustAccessGroupRequireOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutSaml(value *ZeroTrustAccessGroupRequireSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutServiceToken(value *ZeroTrustAccessGroupRequireServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetLinkedAppToken() {
	_jsii_.InvokeVoid(
		z,
		"resetLinkedAppToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


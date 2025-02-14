// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessPolicyIncludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessPolicyIncludeAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessPolicyIncludeAuthContextOutputReference
	AuthContextInput() interface{}
	AuthMethod() ZeroTrustAccessPolicyIncludeAuthMethodOutputReference
	AuthMethodInput() interface{}
	AzureAd() ZeroTrustAccessPolicyIncludeAzureAdOutputReference
	AzureAdInput() interface{}
	Certificate() ZeroTrustAccessPolicyIncludeCertificateOutputReference
	CertificateInput() interface{}
	CommonName() ZeroTrustAccessPolicyIncludeCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessPolicyIncludeDevicePostureOutputReference
	DevicePostureInput() interface{}
	Email() ZeroTrustAccessPolicyIncludeEmailOutputReference
	EmailDomain() ZeroTrustAccessPolicyIncludeEmailDomainOutputReference
	EmailDomainInput() interface{}
	EmailInput() interface{}
	EmailList() ZeroTrustAccessPolicyIncludeEmailListStructOutputReference
	EmailListInput() interface{}
	Everyone() ZeroTrustAccessPolicyIncludeEveryoneOutputReference
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessPolicyIncludeExternalEvaluationOutputReference
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessPolicyIncludeGeoOutputReference
	GeoInput() interface{}
	GithubOrganization() ZeroTrustAccessPolicyIncludeGithubOrganizationOutputReference
	GithubOrganizationInput() interface{}
	Group() ZeroTrustAccessPolicyIncludeGroupOutputReference
	GroupInput() interface{}
	Gsuite() ZeroTrustAccessPolicyIncludeGsuiteOutputReference
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() ZeroTrustAccessPolicyIncludeIpOutputReference
	IpInput() interface{}
	IpList() ZeroTrustAccessPolicyIncludeIpListStructOutputReference
	IpListInput() interface{}
	LoginMethod() ZeroTrustAccessPolicyIncludeLoginMethodOutputReference
	LoginMethodInput() interface{}
	Okta() ZeroTrustAccessPolicyIncludeOktaOutputReference
	OktaInput() interface{}
	Saml() ZeroTrustAccessPolicyIncludeSamlOutputReference
	SamlInput() interface{}
	ServiceToken() ZeroTrustAccessPolicyIncludeServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessPolicyIncludeAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessPolicyIncludeAuthContext)
	PutAuthMethod(value *ZeroTrustAccessPolicyIncludeAuthMethod)
	PutAzureAd(value *ZeroTrustAccessPolicyIncludeAzureAd)
	PutCertificate(value *ZeroTrustAccessPolicyIncludeCertificate)
	PutCommonName(value *ZeroTrustAccessPolicyIncludeCommonName)
	PutDevicePosture(value *ZeroTrustAccessPolicyIncludeDevicePosture)
	PutEmail(value *ZeroTrustAccessPolicyIncludeEmail)
	PutEmailDomain(value *ZeroTrustAccessPolicyIncludeEmailDomain)
	PutEmailList(value *ZeroTrustAccessPolicyIncludeEmailListStruct)
	PutEveryone(value *ZeroTrustAccessPolicyIncludeEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessPolicyIncludeExternalEvaluation)
	PutGeo(value *ZeroTrustAccessPolicyIncludeGeo)
	PutGithubOrganization(value *ZeroTrustAccessPolicyIncludeGithubOrganization)
	PutGroup(value *ZeroTrustAccessPolicyIncludeGroup)
	PutGsuite(value *ZeroTrustAccessPolicyIncludeGsuite)
	PutIp(value *ZeroTrustAccessPolicyIncludeIp)
	PutIpList(value *ZeroTrustAccessPolicyIncludeIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessPolicyIncludeLoginMethod)
	PutOkta(value *ZeroTrustAccessPolicyIncludeOkta)
	PutSaml(value *ZeroTrustAccessPolicyIncludeSaml)
	PutServiceToken(value *ZeroTrustAccessPolicyIncludeServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessPolicyIncludeOutputReference
type jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AnyValidServiceToken() ZeroTrustAccessPolicyIncludeAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyIncludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AuthContext() ZeroTrustAccessPolicyIncludeAuthContextOutputReference {
	var returns ZeroTrustAccessPolicyIncludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AuthMethod() ZeroTrustAccessPolicyIncludeAuthMethodOutputReference {
	var returns ZeroTrustAccessPolicyIncludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AuthMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AzureAd() ZeroTrustAccessPolicyIncludeAzureAdOutputReference {
	var returns ZeroTrustAccessPolicyIncludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) AzureAdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Certificate() ZeroTrustAccessPolicyIncludeCertificateOutputReference {
	var returns ZeroTrustAccessPolicyIncludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) CommonName() ZeroTrustAccessPolicyIncludeCommonNameOutputReference {
	var returns ZeroTrustAccessPolicyIncludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) CommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) DevicePosture() ZeroTrustAccessPolicyIncludeDevicePostureOutputReference {
	var returns ZeroTrustAccessPolicyIncludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) DevicePostureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Email() ZeroTrustAccessPolicyIncludeEmailOutputReference {
	var returns ZeroTrustAccessPolicyIncludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EmailDomain() ZeroTrustAccessPolicyIncludeEmailDomainOutputReference {
	var returns ZeroTrustAccessPolicyIncludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EmailDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EmailList() ZeroTrustAccessPolicyIncludeEmailListStructOutputReference {
	var returns ZeroTrustAccessPolicyIncludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EmailListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Everyone() ZeroTrustAccessPolicyIncludeEveryoneOutputReference {
	var returns ZeroTrustAccessPolicyIncludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ExternalEvaluation() ZeroTrustAccessPolicyIncludeExternalEvaluationOutputReference {
	var returns ZeroTrustAccessPolicyIncludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Geo() ZeroTrustAccessPolicyIncludeGeoOutputReference {
	var returns ZeroTrustAccessPolicyIncludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GithubOrganization() ZeroTrustAccessPolicyIncludeGithubOrganizationOutputReference {
	var returns ZeroTrustAccessPolicyIncludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GithubOrganizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Group() ZeroTrustAccessPolicyIncludeGroupOutputReference {
	var returns ZeroTrustAccessPolicyIncludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Gsuite() ZeroTrustAccessPolicyIncludeGsuiteOutputReference {
	var returns ZeroTrustAccessPolicyIncludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Ip() ZeroTrustAccessPolicyIncludeIpOutputReference {
	var returns ZeroTrustAccessPolicyIncludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) IpList() ZeroTrustAccessPolicyIncludeIpListStructOutputReference {
	var returns ZeroTrustAccessPolicyIncludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) IpListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) LoginMethod() ZeroTrustAccessPolicyIncludeLoginMethodOutputReference {
	var returns ZeroTrustAccessPolicyIncludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) LoginMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Okta() ZeroTrustAccessPolicyIncludeOktaOutputReference {
	var returns ZeroTrustAccessPolicyIncludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Saml() ZeroTrustAccessPolicyIncludeSamlOutputReference {
	var returns ZeroTrustAccessPolicyIncludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ServiceToken() ZeroTrustAccessPolicyIncludeServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyIncludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessPolicyIncludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessPolicyIncludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessPolicyIncludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessPolicyIncludeOutputReference_Override(z ZeroTrustAccessPolicyIncludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessPolicyIncludeAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutAuthContext(value *ZeroTrustAccessPolicyIncludeAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutAuthMethod(value *ZeroTrustAccessPolicyIncludeAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutAzureAd(value *ZeroTrustAccessPolicyIncludeAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutCertificate(value *ZeroTrustAccessPolicyIncludeCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutCommonName(value *ZeroTrustAccessPolicyIncludeCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutDevicePosture(value *ZeroTrustAccessPolicyIncludeDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutEmail(value *ZeroTrustAccessPolicyIncludeEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutEmailDomain(value *ZeroTrustAccessPolicyIncludeEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutEmailList(value *ZeroTrustAccessPolicyIncludeEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutEveryone(value *ZeroTrustAccessPolicyIncludeEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutExternalEvaluation(value *ZeroTrustAccessPolicyIncludeExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutGeo(value *ZeroTrustAccessPolicyIncludeGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutGithubOrganization(value *ZeroTrustAccessPolicyIncludeGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutGroup(value *ZeroTrustAccessPolicyIncludeGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutGsuite(value *ZeroTrustAccessPolicyIncludeGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutIp(value *ZeroTrustAccessPolicyIncludeIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutIpList(value *ZeroTrustAccessPolicyIncludeIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutLoginMethod(value *ZeroTrustAccessPolicyIncludeLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutOkta(value *ZeroTrustAccessPolicyIncludeOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutSaml(value *ZeroTrustAccessPolicyIncludeSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) PutServiceToken(value *ZeroTrustAccessPolicyIncludeServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyIncludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


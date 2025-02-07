// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessApplicationPoliciesIncludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference
	AuthContextInput() interface{}
	AuthMethod() ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference
	AuthMethodInput() interface{}
	AzureAd() ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference
	AzureAdInput() interface{}
	Certificate() ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference
	CertificateInput() interface{}
	CommonName() ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference
	DevicePostureInput() interface{}
	Email() ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference
	EmailDomain() ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference
	EmailDomainInput() interface{}
	EmailInput() interface{}
	EmailList() ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference
	EmailListInput() interface{}
	Everyone() ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference
	GeoInput() interface{}
	GithubOrganization() ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference
	GithubOrganizationInput() interface{}
	Group() ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference
	GroupInput() interface{}
	Gsuite() ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference
	IpInput() interface{}
	IpList() ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference
	IpListInput() interface{}
	Okta() ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference
	OktaInput() interface{}
	Saml() ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference
	SamlInput() interface{}
	ServiceToken() ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessApplicationPoliciesIncludeAuthContext)
	PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesIncludeAuthMethod)
	PutAzureAd(value *ZeroTrustAccessApplicationPoliciesIncludeAzureAd)
	PutCertificate(value *ZeroTrustAccessApplicationPoliciesIncludeCertificate)
	PutCommonName(value *ZeroTrustAccessApplicationPoliciesIncludeCommonName)
	PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesIncludeDevicePosture)
	PutEmail(value *ZeroTrustAccessApplicationPoliciesIncludeEmail)
	PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesIncludeEmailDomain)
	PutEmailList(value *ZeroTrustAccessApplicationPoliciesIncludeEmailListStruct)
	PutEveryone(value *ZeroTrustAccessApplicationPoliciesIncludeEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation)
	PutGeo(value *ZeroTrustAccessApplicationPoliciesIncludeGeo)
	PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesIncludeGithubOrganization)
	PutGroup(value *ZeroTrustAccessApplicationPoliciesIncludeGroup)
	PutGsuite(value *ZeroTrustAccessApplicationPoliciesIncludeGsuite)
	PutIp(value *ZeroTrustAccessApplicationPoliciesIncludeIp)
	PutIpList(value *ZeroTrustAccessApplicationPoliciesIncludeIpListStruct)
	PutOkta(value *ZeroTrustAccessApplicationPoliciesIncludeOkta)
	PutSaml(value *ZeroTrustAccessApplicationPoliciesIncludeSaml)
	PutServiceToken(value *ZeroTrustAccessApplicationPoliciesIncludeServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessApplicationPoliciesIncludeOutputReference
type jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AuthContext() ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AuthMethod() ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AuthMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AzureAd() ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) AzureAdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Certificate() ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) CommonName() ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) CommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) DevicePosture() ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) DevicePostureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Email() ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EmailDomain() ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EmailDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EmailList() ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EmailListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Everyone() ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ExternalEvaluation() ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Geo() ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GithubOrganization() ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GithubOrganizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Group() ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Gsuite() ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Ip() ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) IpList() ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) IpListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Okta() ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Saml() ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ServiceToken() ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessApplicationPoliciesIncludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessApplicationPoliciesIncludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationPoliciesIncludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationPoliciesIncludeOutputReference_Override(z ZeroTrustAccessApplicationPoliciesIncludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutAuthContext(value *ZeroTrustAccessApplicationPoliciesIncludeAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesIncludeAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutAzureAd(value *ZeroTrustAccessApplicationPoliciesIncludeAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutCertificate(value *ZeroTrustAccessApplicationPoliciesIncludeCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutCommonName(value *ZeroTrustAccessApplicationPoliciesIncludeCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesIncludeDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutEmail(value *ZeroTrustAccessApplicationPoliciesIncludeEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesIncludeEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutEmailList(value *ZeroTrustAccessApplicationPoliciesIncludeEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutEveryone(value *ZeroTrustAccessApplicationPoliciesIncludeEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutGeo(value *ZeroTrustAccessApplicationPoliciesIncludeGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesIncludeGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutGroup(value *ZeroTrustAccessApplicationPoliciesIncludeGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutGsuite(value *ZeroTrustAccessApplicationPoliciesIncludeGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutIp(value *ZeroTrustAccessApplicationPoliciesIncludeIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutIpList(value *ZeroTrustAccessApplicationPoliciesIncludeIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutOkta(value *ZeroTrustAccessApplicationPoliciesIncludeOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutSaml(value *ZeroTrustAccessApplicationPoliciesIncludeSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) PutServiceToken(value *ZeroTrustAccessApplicationPoliciesIncludeServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


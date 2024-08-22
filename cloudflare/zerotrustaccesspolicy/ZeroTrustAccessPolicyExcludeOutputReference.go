// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessPolicyExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() interface{}
	SetAnyValidServiceToken(val interface{})
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessPolicyExcludeAuthContextList
	AuthContextInput() interface{}
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
	Azure() ZeroTrustAccessPolicyExcludeAzureList
	AzureInput() interface{}
	Certificate() interface{}
	SetCertificate(val interface{})
	CertificateInput() interface{}
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
	CommonNames() *[]*string
	SetCommonNames(val *[]*string)
	CommonNamesInput() *[]*string
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
	DevicePosture() *[]*string
	SetDevicePosture(val *[]*string)
	DevicePostureInput() *[]*string
	Email() *[]*string
	SetEmail(val *[]*string)
	EmailDomain() *[]*string
	SetEmailDomain(val *[]*string)
	EmailDomainInput() *[]*string
	EmailInput() *[]*string
	EmailList() *[]*string
	SetEmailList(val *[]*string)
	EmailListInput() *[]*string
	Everyone() interface{}
	SetEveryone(val interface{})
	EveryoneInput() interface{}
	ExternalEvaluation() ZeroTrustAccessPolicyExcludeExternalEvaluationList
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() *[]*string
	SetGeo(val *[]*string)
	GeoInput() *[]*string
	Github() ZeroTrustAccessPolicyExcludeGithubList
	GithubInput() interface{}
	Group() *[]*string
	SetGroup(val *[]*string)
	GroupInput() *[]*string
	Gsuite() ZeroTrustAccessPolicyExcludeGsuiteList
	GsuiteInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() *[]*string
	SetIp(val *[]*string)
	IpInput() *[]*string
	IpList() *[]*string
	SetIpList(val *[]*string)
	IpListInput() *[]*string
	LoginMethod() *[]*string
	SetLoginMethod(val *[]*string)
	LoginMethodInput() *[]*string
	Okta() ZeroTrustAccessPolicyExcludeOktaList
	OktaInput() interface{}
	Saml() ZeroTrustAccessPolicyExcludeSamlList
	SamlInput() interface{}
	ServiceToken() *[]*string
	SetServiceToken(val *[]*string)
	ServiceTokenInput() *[]*string
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
	PutAuthContext(value interface{})
	PutAzure(value interface{})
	PutExternalEvaluation(value interface{})
	PutGithub(value interface{})
	PutGsuite(value interface{})
	PutOkta(value interface{})
	PutSaml(value interface{})
	ResetAnyValidServiceToken()
	ResetAuthContext()
	ResetAuthMethod()
	ResetAzure()
	ResetCertificate()
	ResetCommonName()
	ResetCommonNames()
	ResetDevicePosture()
	ResetEmail()
	ResetEmailDomain()
	ResetEmailList()
	ResetEveryone()
	ResetExternalEvaluation()
	ResetGeo()
	ResetGithub()
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

// The jsii proxy struct for ZeroTrustAccessPolicyExcludeOutputReference
type jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AnyValidServiceToken() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthContext() ZeroTrustAccessPolicyExcludeAuthContextList {
	var returns ZeroTrustAccessPolicyExcludeAuthContextList
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Azure() ZeroTrustAccessPolicyExcludeAzureList {
	var returns ZeroTrustAccessPolicyExcludeAzureList
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Certificate() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) CommonNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commonNamesInput",
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) DevicePosture() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) DevicePostureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Email() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailDomain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailDomainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) EmailListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Everyone() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ExternalEvaluation() ZeroTrustAccessPolicyExcludeExternalEvaluationList {
	var returns ZeroTrustAccessPolicyExcludeExternalEvaluationList
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Geo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GeoInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Github() ZeroTrustAccessPolicyExcludeGithubList {
	var returns ZeroTrustAccessPolicyExcludeGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GithubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Group() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) GroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Gsuite() ZeroTrustAccessPolicyExcludeGsuiteList {
	var returns ZeroTrustAccessPolicyExcludeGsuiteList
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Ip() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) IpListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LoginMethod() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) LoginMethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Okta() ZeroTrustAccessPolicyExcludeOktaList {
	var returns ZeroTrustAccessPolicyExcludeOktaList
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) Saml() ZeroTrustAccessPolicyExcludeSamlList {
	var returns ZeroTrustAccessPolicyExcludeSamlList
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ServiceToken() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ServiceTokenInput() *[]*string {
	var returns *[]*string
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetAnyValidServiceToken(val interface{}) {
	if err := j.validateSetAnyValidServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyValidServiceToken",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetCertificate(val interface{}) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetCommonNames(val *[]*string) {
	if err := j.validateSetCommonNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonNames",
		val,
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetDevicePosture(val *[]*string) {
	if err := j.validateSetDevicePostureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devicePosture",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetEmail(val *[]*string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetEmailDomain(val *[]*string) {
	if err := j.validateSetEmailDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailDomain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetEmailList(val *[]*string) {
	if err := j.validateSetEmailListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailList",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetEveryone(val interface{}) {
	if err := j.validateSetEveryoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"everyone",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetGeo(val *[]*string) {
	if err := j.validateSetGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geo",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetGroup(val *[]*string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
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

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetIp(val *[]*string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetIpList(val *[]*string) {
	if err := j.validateSetIpListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipList",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetLoginMethod(val *[]*string) {
	if err := j.validateSetLoginMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMethod",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference)SetServiceToken(val *[]*string) {
	if err := j.validateSetServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceToken",
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAuthContext(value interface{}) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutAzure(value interface{}) {
	if err := z.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzure",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutExternalEvaluation(value interface{}) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGithub(value interface{}) {
	if err := z.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithub",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutGsuite(value interface{}) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutOkta(value interface{}) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) PutSaml(value interface{}) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		z,
		"resetAzure",
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetCommonNames() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonNames",
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		z,
		"resetGithub",
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

func (z *jsiiProxy_ZeroTrustAccessPolicyExcludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
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


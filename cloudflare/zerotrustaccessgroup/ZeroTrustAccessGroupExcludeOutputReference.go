// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustaccessgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessGroupExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() interface{}
	SetAnyValidServiceToken(val interface{})
	AnyValidServiceTokenInput() interface{}
	AuthContext() ZeroTrustAccessGroupExcludeAuthContextList
	AuthContextInput() interface{}
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
	Azure() ZeroTrustAccessGroupExcludeAzureList
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
	ExternalEvaluation() ZeroTrustAccessGroupExcludeExternalEvaluationList
	ExternalEvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() *[]*string
	SetGeo(val *[]*string)
	GeoInput() *[]*string
	Github() ZeroTrustAccessGroupExcludeGithubList
	GithubInput() interface{}
	Group() *[]*string
	SetGroup(val *[]*string)
	GroupInput() *[]*string
	Gsuite() ZeroTrustAccessGroupExcludeGsuiteList
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
	Okta() ZeroTrustAccessGroupExcludeOktaList
	OktaInput() interface{}
	Saml() ZeroTrustAccessGroupExcludeSamlList
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

// The jsii proxy struct for ZeroTrustAccessGroupExcludeOutputReference
type jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AnyValidServiceToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AnyValidServiceTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AuthContext() ZeroTrustAccessGroupExcludeAuthContextList {
	var returns ZeroTrustAccessGroupExcludeAuthContextList
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AuthContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Azure() ZeroTrustAccessGroupExcludeAzureList {
	var returns ZeroTrustAccessGroupExcludeAzureList
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Certificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CommonNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CommonNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) DevicePosture() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) DevicePostureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Email() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EmailDomain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EmailDomainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EmailInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EmailList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EmailListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Everyone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) EveryoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ExternalEvaluation() ZeroTrustAccessGroupExcludeExternalEvaluationList {
	var returns ZeroTrustAccessGroupExcludeExternalEvaluationList
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ExternalEvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Geo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GeoInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Github() ZeroTrustAccessGroupExcludeGithubList {
	var returns ZeroTrustAccessGroupExcludeGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GithubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Group() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Gsuite() ZeroTrustAccessGroupExcludeGsuiteList {
	var returns ZeroTrustAccessGroupExcludeGsuiteList
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GsuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Ip() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) IpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) IpList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) IpListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) LoginMethod() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) LoginMethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Okta() ZeroTrustAccessGroupExcludeOktaList {
	var returns ZeroTrustAccessGroupExcludeOktaList
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Saml() ZeroTrustAccessGroupExcludeSamlList {
	var returns ZeroTrustAccessGroupExcludeSamlList
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ServiceToken() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ServiceTokenInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustAccessGroupExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessGroupExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessGroupExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessGroupExcludeOutputReference_Override(z ZeroTrustAccessGroupExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetAnyValidServiceToken(val interface{}) {
	if err := j.validateSetAnyValidServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyValidServiceToken",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetCertificate(val interface{}) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetCommonNames(val *[]*string) {
	if err := j.validateSetCommonNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonNames",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetDevicePosture(val *[]*string) {
	if err := j.validateSetDevicePostureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devicePosture",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetEmail(val *[]*string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetEmailDomain(val *[]*string) {
	if err := j.validateSetEmailDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailDomain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetEmailList(val *[]*string) {
	if err := j.validateSetEmailListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailList",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetEveryone(val interface{}) {
	if err := j.validateSetEveryoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"everyone",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetGeo(val *[]*string) {
	if err := j.validateSetGeoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geo",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetGroup(val *[]*string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetIp(val *[]*string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetIpList(val *[]*string) {
	if err := j.validateSetIpListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipList",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetLoginMethod(val *[]*string) {
	if err := j.validateSetLoginMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMethod",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetServiceToken(val *[]*string) {
	if err := j.validateSetServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceToken",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutAuthContext(value interface{}) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutAzure(value interface{}) {
	if err := z.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzure",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutExternalEvaluation(value interface{}) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutGithub(value interface{}) {
	if err := z.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithub",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutGsuite(value interface{}) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutOkta(value interface{}) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) PutSaml(value interface{}) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		z,
		"resetAzure",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetCommonNames() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonNames",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		z,
		"resetGithub",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustAccessGroupExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccesspolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccesspolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessPoliciesResultRequireAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessPoliciesResultRequireAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessPoliciesResultRequireAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessPoliciesResultRequireAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessPoliciesResultRequireCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessPoliciesResultRequireCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessPoliciesResultRequireDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessPoliciesResultRequireEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessPoliciesResultRequireExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessPoliciesResultRequireGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessPoliciesResultRequireGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessPoliciesResultRequireGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessPoliciesResultRequireGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessPoliciesResultRequire
	SetInternalValue(val *DataCloudflareZeroTrustAccessPoliciesResultRequire)
	Ip() DataCloudflareZeroTrustAccessPoliciesResultRequireIpOutputReference
	IpList() DataCloudflareZeroTrustAccessPoliciesResultRequireIpListStructOutputReference
	LinkedAppToken() DataCloudflareZeroTrustAccessPoliciesResultRequireLinkedAppTokenOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessPoliciesResultRequireLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessPoliciesResultRequireOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessPoliciesResultRequireOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessPoliciesResultRequireSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessPoliciesResultRequireServiceTokenOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessPoliciesResultRequireAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) AuthContext() DataCloudflareZeroTrustAccessPoliciesResultRequireAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) AuthMethod() DataCloudflareZeroTrustAccessPoliciesResultRequireAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) AzureAd() DataCloudflareZeroTrustAccessPoliciesResultRequireAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Certificate() DataCloudflareZeroTrustAccessPoliciesResultRequireCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) CommonName() DataCloudflareZeroTrustAccessPoliciesResultRequireCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) DevicePosture() DataCloudflareZeroTrustAccessPoliciesResultRequireDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Email() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) EmailDomain() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) EmailList() DataCloudflareZeroTrustAccessPoliciesResultRequireEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Everyone() DataCloudflareZeroTrustAccessPoliciesResultRequireEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessPoliciesResultRequireExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Geo() DataCloudflareZeroTrustAccessPoliciesResultRequireGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessPoliciesResultRequireGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Group() DataCloudflareZeroTrustAccessPoliciesResultRequireGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Gsuite() DataCloudflareZeroTrustAccessPoliciesResultRequireGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) InternalValue() *DataCloudflareZeroTrustAccessPoliciesResultRequire {
	var returns *DataCloudflareZeroTrustAccessPoliciesResultRequire
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Ip() DataCloudflareZeroTrustAccessPoliciesResultRequireIpOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) IpList() DataCloudflareZeroTrustAccessPoliciesResultRequireIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) LinkedAppToken() DataCloudflareZeroTrustAccessPoliciesResultRequireLinkedAppTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) LoginMethod() DataCloudflareZeroTrustAccessPoliciesResultRequireLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Oidc() DataCloudflareZeroTrustAccessPoliciesResultRequireOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Okta() DataCloudflareZeroTrustAccessPoliciesResultRequireOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Saml() DataCloudflareZeroTrustAccessPoliciesResultRequireSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ServiceToken() DataCloudflareZeroTrustAccessPoliciesResultRequireServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessPoliciesResultRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicies.DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference_Override(d DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicies.DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessPoliciesResultRequire) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


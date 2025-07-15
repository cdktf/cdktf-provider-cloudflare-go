// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessApplicationPoliciesRequire
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPoliciesRequire)
	Ip() DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpOutputReference
	IpList() DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessApplicationPoliciesRequireOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessApplicationPoliciesRequireOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessApplicationPoliciesRequireSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthContext() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthMethod() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) AzureAd() DataCloudflareZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Certificate() DataCloudflareZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) CommonName() DataCloudflareZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) DevicePosture() DataCloudflareZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Email() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailDomain() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailList() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Everyone() DataCloudflareZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Geo() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Group() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Gsuite() DataCloudflareZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationPoliciesRequire {
	var returns *DataCloudflareZeroTrustAccessApplicationPoliciesRequire
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Ip() DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) IpList() DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) LoginMethod() DataCloudflareZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Oidc() DataCloudflareZeroTrustAccessApplicationPoliciesRequireOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Okta() DataCloudflareZeroTrustAccessApplicationPoliciesRequireOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Saml() DataCloudflareZeroTrustAccessApplicationPoliciesRequireSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPoliciesRequire) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


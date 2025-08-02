// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesInclude
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesInclude)
	Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpOutputReference
	IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpListStructOutputReference
	LinkedAppToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLinkedAppTokenOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesInclude {
	var returns *DataCloudflareZeroTrustAccessApplicationsResultPoliciesInclude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) LinkedAppToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLinkedAppTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) LoginMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Oidc() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesInclude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesIncludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


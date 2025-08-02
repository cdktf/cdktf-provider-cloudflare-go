// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessApplicationPoliciesExclude
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPoliciesExclude)
	Ip() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpOutputReference
	IpList() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference
	LinkedAppToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLinkedAppTokenOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthContext() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) AuthMethod() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) AzureAd() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Certificate() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) CommonName() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) DevicePosture() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Email() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailDomain() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) EmailList() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Everyone() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Geo() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Group() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Gsuite() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationPoliciesExclude {
	var returns *DataCloudflareZeroTrustAccessApplicationPoliciesExclude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Ip() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) IpList() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) LinkedAppToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLinkedAppTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLinkedAppTokenOutputReference
	_jsii_.Get(
		j,
		"linkedAppToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) LoginMethod() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Oidc() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Okta() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Saml() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ServiceToken() DataCloudflareZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplication.DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationPoliciesExclude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationPoliciesExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


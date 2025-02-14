// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccesspolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccesspolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessPoliciesResultExcludeAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessPoliciesResultExcludeAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessPoliciesResultExcludeCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessPoliciesResultExcludeCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessPoliciesResultExcludeDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessPoliciesResultExcludeEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessPoliciesResultExcludeExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessPoliciesResultExcludeGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessPoliciesResultExcludeGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessPoliciesResultExcludeGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessPoliciesResultExcludeGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessPoliciesResultExclude
	SetInternalValue(val *DataCloudflareZeroTrustAccessPoliciesResultExclude)
	Ip() DataCloudflareZeroTrustAccessPoliciesResultExcludeIpOutputReference
	IpList() DataCloudflareZeroTrustAccessPoliciesResultExcludeIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessPoliciesResultExcludeLoginMethodOutputReference
	Okta() DataCloudflareZeroTrustAccessPoliciesResultExcludeOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessPoliciesResultExcludeSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessPoliciesResultExcludeServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessPoliciesResultExcludeAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) AuthContext() DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) AuthMethod() DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) AzureAd() DataCloudflareZeroTrustAccessPoliciesResultExcludeAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Certificate() DataCloudflareZeroTrustAccessPoliciesResultExcludeCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) CommonName() DataCloudflareZeroTrustAccessPoliciesResultExcludeCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) DevicePosture() DataCloudflareZeroTrustAccessPoliciesResultExcludeDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Email() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) EmailDomain() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) EmailList() DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Everyone() DataCloudflareZeroTrustAccessPoliciesResultExcludeEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessPoliciesResultExcludeExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Geo() DataCloudflareZeroTrustAccessPoliciesResultExcludeGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessPoliciesResultExcludeGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Group() DataCloudflareZeroTrustAccessPoliciesResultExcludeGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Gsuite() DataCloudflareZeroTrustAccessPoliciesResultExcludeGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) InternalValue() *DataCloudflareZeroTrustAccessPoliciesResultExclude {
	var returns *DataCloudflareZeroTrustAccessPoliciesResultExclude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Ip() DataCloudflareZeroTrustAccessPoliciesResultExcludeIpOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) IpList() DataCloudflareZeroTrustAccessPoliciesResultExcludeIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) LoginMethod() DataCloudflareZeroTrustAccessPoliciesResultExcludeLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Okta() DataCloudflareZeroTrustAccessPoliciesResultExcludeOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Saml() DataCloudflareZeroTrustAccessPoliciesResultExcludeSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ServiceToken() DataCloudflareZeroTrustAccessPoliciesResultExcludeServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPoliciesResultExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicies.DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference_Override(d DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicies.DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessPoliciesResultExclude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPoliciesResultExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


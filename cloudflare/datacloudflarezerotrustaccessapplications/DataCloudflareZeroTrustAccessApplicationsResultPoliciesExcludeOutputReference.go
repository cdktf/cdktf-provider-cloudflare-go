// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesExclude
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesExclude)
	Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpOutputReference
	IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeLoginMethodOutputReference
	Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesExclude {
	var returns *DataCloudflareZeroTrustAccessApplicationsResultPoliciesExclude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) LoginMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesExclude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


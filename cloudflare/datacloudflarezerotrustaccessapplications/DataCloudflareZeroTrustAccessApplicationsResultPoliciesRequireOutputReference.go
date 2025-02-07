// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccessapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequire
	SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequire)
	Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpOutputReference
	IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpListStructOutputReference
	Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) AuthContext() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) AuthMethod() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) AzureAd() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Certificate() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) CommonName() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) DevicePosture() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Email() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) EmailDomain() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) EmailList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Everyone() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Geo() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Group() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Gsuite() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) InternalValue() *DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequire {
	var returns *DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequire
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Ip() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) IpList() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Okta() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Saml() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ServiceToken() DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference_Override(d DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessApplications.DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequire) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessApplicationsResultPoliciesRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


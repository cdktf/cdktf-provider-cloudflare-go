// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessPolicyRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessPolicyRequireAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessPolicyRequireAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessPolicyRequireAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessPolicyRequireCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessPolicyRequireCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessPolicyRequireDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessPolicyRequireEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessPolicyRequireEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessPolicyRequireEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessPolicyRequireEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessPolicyRequireExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessPolicyRequireGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessPolicyRequireGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessPolicyRequireGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessPolicyRequireGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessPolicyRequire
	SetInternalValue(val *DataCloudflareZeroTrustAccessPolicyRequire)
	Ip() DataCloudflareZeroTrustAccessPolicyRequireIpOutputReference
	IpList() DataCloudflareZeroTrustAccessPolicyRequireIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessPolicyRequireLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessPolicyRequireOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessPolicyRequireOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessPolicyRequireSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessPolicyRequireServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessPolicyRequireOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) AuthContext() DataCloudflareZeroTrustAccessPolicyRequireAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) AuthMethod() DataCloudflareZeroTrustAccessPolicyRequireAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) AzureAd() DataCloudflareZeroTrustAccessPolicyRequireAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Certificate() DataCloudflareZeroTrustAccessPolicyRequireCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) CommonName() DataCloudflareZeroTrustAccessPolicyRequireCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) DevicePosture() DataCloudflareZeroTrustAccessPolicyRequireDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Email() DataCloudflareZeroTrustAccessPolicyRequireEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) EmailDomain() DataCloudflareZeroTrustAccessPolicyRequireEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) EmailList() DataCloudflareZeroTrustAccessPolicyRequireEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Everyone() DataCloudflareZeroTrustAccessPolicyRequireEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessPolicyRequireExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Geo() DataCloudflareZeroTrustAccessPolicyRequireGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessPolicyRequireGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Group() DataCloudflareZeroTrustAccessPolicyRequireGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Gsuite() DataCloudflareZeroTrustAccessPolicyRequireGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) InternalValue() *DataCloudflareZeroTrustAccessPolicyRequire {
	var returns *DataCloudflareZeroTrustAccessPolicyRequire
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Ip() DataCloudflareZeroTrustAccessPolicyRequireIpOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) IpList() DataCloudflareZeroTrustAccessPolicyRequireIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) LoginMethod() DataCloudflareZeroTrustAccessPolicyRequireLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Oidc() DataCloudflareZeroTrustAccessPolicyRequireOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Okta() DataCloudflareZeroTrustAccessPolicyRequireOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Saml() DataCloudflareZeroTrustAccessPolicyRequireSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ServiceToken() DataCloudflareZeroTrustAccessPolicyRequireServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessPolicyRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessPolicyRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessPolicyRequireOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessPolicyRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicy.DataCloudflareZeroTrustAccessPolicyRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessPolicyRequireOutputReference_Override(d DataCloudflareZeroTrustAccessPolicyRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessPolicy.DataCloudflareZeroTrustAccessPolicyRequireOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessPolicyRequire) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessPolicyRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


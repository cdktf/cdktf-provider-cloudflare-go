// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessgroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflarezerotrustaccessgroups/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessGroupsResultIsDefaultCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessGroupsResultIsDefaultCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessGroupsResultIsDefaultDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultIsDefaultExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessGroupsResultIsDefault
	SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultIsDefault)
	Ip() DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpOutputReference
	IpList() DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessGroupsResultIsDefaultLoginMethodOutputReference
	Oidc() DataCloudflareZeroTrustAccessGroupsResultIsDefaultOidcOutputReference
	Okta() DataCloudflareZeroTrustAccessGroupsResultIsDefaultOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessGroupsResultIsDefaultSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessGroupsResultIsDefaultServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) AuthContext() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) AuthMethod() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) AzureAd() DataCloudflareZeroTrustAccessGroupsResultIsDefaultAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Certificate() DataCloudflareZeroTrustAccessGroupsResultIsDefaultCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) CommonName() DataCloudflareZeroTrustAccessGroupsResultIsDefaultCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) DevicePosture() DataCloudflareZeroTrustAccessGroupsResultIsDefaultDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Email() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) EmailDomain() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) EmailList() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Everyone() DataCloudflareZeroTrustAccessGroupsResultIsDefaultEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultIsDefaultExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Geo() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Group() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Gsuite() DataCloudflareZeroTrustAccessGroupsResultIsDefaultGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) InternalValue() *DataCloudflareZeroTrustAccessGroupsResultIsDefault {
	var returns *DataCloudflareZeroTrustAccessGroupsResultIsDefault
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Ip() DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) IpList() DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) LoginMethod() DataCloudflareZeroTrustAccessGroupsResultIsDefaultLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Oidc() DataCloudflareZeroTrustAccessGroupsResultIsDefaultOidcOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Okta() DataCloudflareZeroTrustAccessGroupsResultIsDefaultOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Saml() DataCloudflareZeroTrustAccessGroupsResultIsDefaultSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ServiceToken() DataCloudflareZeroTrustAccessGroupsResultIsDefaultServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultIsDefaultServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference_Override(d DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultIsDefault) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultIsDefaultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


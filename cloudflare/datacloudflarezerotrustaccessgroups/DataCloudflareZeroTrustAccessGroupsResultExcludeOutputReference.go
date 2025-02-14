// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessgroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccessgroups/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultExcludeAnyValidServiceTokenOutputReference
	AuthContext() DataCloudflareZeroTrustAccessGroupsResultExcludeAuthContextOutputReference
	AuthMethod() DataCloudflareZeroTrustAccessGroupsResultExcludeAuthMethodOutputReference
	AzureAd() DataCloudflareZeroTrustAccessGroupsResultExcludeAzureAdOutputReference
	Certificate() DataCloudflareZeroTrustAccessGroupsResultExcludeCertificateOutputReference
	CommonName() DataCloudflareZeroTrustAccessGroupsResultExcludeCommonNameOutputReference
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
	DevicePosture() DataCloudflareZeroTrustAccessGroupsResultExcludeDevicePostureOutputReference
	Email() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailOutputReference
	EmailDomain() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailDomainOutputReference
	EmailList() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailListStructOutputReference
	Everyone() DataCloudflareZeroTrustAccessGroupsResultExcludeEveryoneOutputReference
	ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultExcludeExternalEvaluationOutputReference
	// Experimental.
	Fqn() *string
	Geo() DataCloudflareZeroTrustAccessGroupsResultExcludeGeoOutputReference
	GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultExcludeGithubOrganizationOutputReference
	Group() DataCloudflareZeroTrustAccessGroupsResultExcludeGroupOutputReference
	Gsuite() DataCloudflareZeroTrustAccessGroupsResultExcludeGsuiteOutputReference
	InternalValue() *DataCloudflareZeroTrustAccessGroupsResultExclude
	SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultExclude)
	Ip() DataCloudflareZeroTrustAccessGroupsResultExcludeIpOutputReference
	IpList() DataCloudflareZeroTrustAccessGroupsResultExcludeIpListStructOutputReference
	LoginMethod() DataCloudflareZeroTrustAccessGroupsResultExcludeLoginMethodOutputReference
	Okta() DataCloudflareZeroTrustAccessGroupsResultExcludeOktaOutputReference
	Saml() DataCloudflareZeroTrustAccessGroupsResultExcludeSamlOutputReference
	ServiceToken() DataCloudflareZeroTrustAccessGroupsResultExcludeServiceTokenOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) AnyValidServiceToken() DataCloudflareZeroTrustAccessGroupsResultExcludeAnyValidServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) AuthContext() DataCloudflareZeroTrustAccessGroupsResultExcludeAuthContextOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) AuthMethod() DataCloudflareZeroTrustAccessGroupsResultExcludeAuthMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) AzureAd() DataCloudflareZeroTrustAccessGroupsResultExcludeAzureAdOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Certificate() DataCloudflareZeroTrustAccessGroupsResultExcludeCertificateOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) CommonName() DataCloudflareZeroTrustAccessGroupsResultExcludeCommonNameOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) DevicePosture() DataCloudflareZeroTrustAccessGroupsResultExcludeDevicePostureOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Email() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) EmailDomain() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailDomainOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) EmailList() DataCloudflareZeroTrustAccessGroupsResultExcludeEmailListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Everyone() DataCloudflareZeroTrustAccessGroupsResultExcludeEveryoneOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ExternalEvaluation() DataCloudflareZeroTrustAccessGroupsResultExcludeExternalEvaluationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Geo() DataCloudflareZeroTrustAccessGroupsResultExcludeGeoOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GithubOrganization() DataCloudflareZeroTrustAccessGroupsResultExcludeGithubOrganizationOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Group() DataCloudflareZeroTrustAccessGroupsResultExcludeGroupOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Gsuite() DataCloudflareZeroTrustAccessGroupsResultExcludeGsuiteOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) InternalValue() *DataCloudflareZeroTrustAccessGroupsResultExclude {
	var returns *DataCloudflareZeroTrustAccessGroupsResultExclude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Ip() DataCloudflareZeroTrustAccessGroupsResultExcludeIpOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) IpList() DataCloudflareZeroTrustAccessGroupsResultExcludeIpListStructOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) LoginMethod() DataCloudflareZeroTrustAccessGroupsResultExcludeLoginMethodOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Okta() DataCloudflareZeroTrustAccessGroupsResultExcludeOktaOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Saml() DataCloudflareZeroTrustAccessGroupsResultExcludeSamlOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ServiceToken() DataCloudflareZeroTrustAccessGroupsResultExcludeServiceTokenOutputReference {
	var returns DataCloudflareZeroTrustAccessGroupsResultExcludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessGroupsResultExcludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference_Override(d DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessGroups.DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessGroupsResultExclude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessGroupsResultExcludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


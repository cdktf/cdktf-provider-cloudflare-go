// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessidentityproviders

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrustaccessidentityproviders/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference interface {
	cdktf.ComplexObject
	AppsDomain() *string
	Attributes() *[]*string
	AuthorizationServerId() *string
	AuthUrl() *string
	CentrifyAccount() *string
	CentrifyAppId() *string
	CertsUrl() *string
	Claims() *[]*string
	ClientId() *string
	ClientSecret() *string
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
	ConditionalAccessEnabled() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DirectoryId() *string
	EmailAttributeName() *string
	EmailClaimName() *string
	// Experimental.
	Fqn() *string
	HeaderAttributes() DataCloudflareZeroTrustAccessIdentityProvidersResultConfigHeaderAttributesList
	IdpPublicCerts() *[]*string
	InternalValue() *DataCloudflareZeroTrustAccessIdentityProvidersResultConfig
	SetInternalValue(val *DataCloudflareZeroTrustAccessIdentityProvidersResultConfig)
	IssuerUrl() *string
	OktaAccount() *string
	OneloginAccount() *string
	PingEnvId() *string
	PkceEnabled() cdktf.IResolvable
	Prompt() *string
	Scopes() *[]*string
	SignRequest() cdktf.IResolvable
	SsoTargetUrl() *string
	SupportGroups() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenUrl() *string
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

// The jsii proxy struct for DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference
type jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) AppsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Attributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) AuthorizationServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) CentrifyAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) CentrifyAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) CertsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Claims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ConditionalAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"conditionalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) EmailAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) EmailClaimName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailClaimName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) HeaderAttributes() DataCloudflareZeroTrustAccessIdentityProvidersResultConfigHeaderAttributesList {
	var returns DataCloudflareZeroTrustAccessIdentityProvidersResultConfigHeaderAttributesList
	_jsii_.Get(
		j,
		"headerAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) IdpPublicCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpPublicCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) InternalValue() *DataCloudflareZeroTrustAccessIdentityProvidersResultConfig {
	var returns *DataCloudflareZeroTrustAccessIdentityProvidersResultConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) OktaAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) OneloginAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) PingEnvId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pingEnvId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) PkceEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pkceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) SignRequest() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"signRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) SsoTargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) SupportGroups() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessIdentityProviders.DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference_Override(d DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessIdentityProviders.DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference)SetInternalValue(val *DataCloudflareZeroTrustAccessIdentityProvidersResultConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustAccessIdentityProvidersResultConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


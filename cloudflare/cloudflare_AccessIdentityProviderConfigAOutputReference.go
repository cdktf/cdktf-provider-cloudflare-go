// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessIdentityProviderConfigAOutputReference interface {
	cdktf.ComplexObject
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	AppsDomain() *string
	SetAppsDomain(val *string)
	AppsDomainInput() *string
	Attributes() *[]*string
	SetAttributes(val *[]*string)
	AttributesInput() *[]*string
	AuthUrl() *string
	SetAuthUrl(val *string)
	AuthUrlInput() *string
	CentrifyAccount() *string
	SetCentrifyAccount(val *string)
	CentrifyAccountInput() *string
	CentrifyAppId() *string
	SetCentrifyAppId(val *string)
	CentrifyAppIdInput() *string
	CertsUrl() *string
	SetCertsUrl(val *string)
	CertsUrlInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	EmailAttributeName() *string
	SetEmailAttributeName(val *string)
	EmailAttributeNameInput() *string
	// Experimental.
	Fqn() *string
	IdpPublicCert() *string
	SetIdpPublicCert(val *string)
	IdpPublicCertInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssuerUrl() *string
	SetIssuerUrl(val *string)
	IssuerUrlInput() *string
	OktaAccount() *string
	SetOktaAccount(val *string)
	OktaAccountInput() *string
	OneloginAccount() *string
	SetOneloginAccount(val *string)
	OneloginAccountInput() *string
	PkceEnabled() interface{}
	SetPkceEnabled(val interface{})
	PkceEnabledInput() interface{}
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlInput() *string
	SignRequest() interface{}
	SetSignRequest(val interface{})
	SignRequestInput() interface{}
	SsoTargetUrl() *string
	SetSsoTargetUrl(val *string)
	SsoTargetUrlInput() *string
	SupportGroups() interface{}
	SetSupportGroups(val interface{})
	SupportGroupsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
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
	ResetApiToken()
	ResetAppsDomain()
	ResetAttributes()
	ResetAuthUrl()
	ResetCentrifyAccount()
	ResetCentrifyAppId()
	ResetCertsUrl()
	ResetClientId()
	ResetClientSecret()
	ResetDirectoryId()
	ResetEmailAttributeName()
	ResetIdpPublicCert()
	ResetIssuerUrl()
	ResetOktaAccount()
	ResetOneloginAccount()
	ResetPkceEnabled()
	ResetRedirectUrl()
	ResetSignRequest()
	ResetSsoTargetUrl()
	ResetSupportGroups()
	ResetTokenUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessIdentityProviderConfigAOutputReference
type jsiiProxy_AccessIdentityProviderConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AppsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AppsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Attributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) AuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CentrifyAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centrifyAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CertsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CertsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) EmailAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IdpPublicCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpPublicCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IdpPublicCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpPublicCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OktaAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OktaAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OneloginAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) OneloginAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneloginAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PkceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) PkceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SignRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SignRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SsoTargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SsoTargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoTargetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SupportGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SupportGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewAccessIdentityProviderConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessIdentityProviderConfigAOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AccessIdentityProviderConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.AccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessIdentityProviderConfigAOutputReference_Override(a AccessIdentityProviderConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.AccessIdentityProviderConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetAppsDomain(val *string) {
	_jsii_.Set(
		j,
		"appsDomain",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetAuthUrl(val *string) {
	_jsii_.Set(
		j,
		"authUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetCentrifyAccount(val *string) {
	_jsii_.Set(
		j,
		"centrifyAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetCentrifyAppId(val *string) {
	_jsii_.Set(
		j,
		"centrifyAppId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetCertsUrl(val *string) {
	_jsii_.Set(
		j,
		"certsUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetEmailAttributeName(val *string) {
	_jsii_.Set(
		j,
		"emailAttributeName",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetIdpPublicCert(val *string) {
	_jsii_.Set(
		j,
		"idpPublicCert",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetIssuerUrl(val *string) {
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetOktaAccount(val *string) {
	_jsii_.Set(
		j,
		"oktaAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetOneloginAccount(val *string) {
	_jsii_.Set(
		j,
		"oneloginAccount",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetPkceEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"pkceEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetSignRequest(val interface{}) {
	_jsii_.Set(
		j,
		"signRequest",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetSsoTargetUrl(val *string) {
	_jsii_.Set(
		j,
		"ssoTargetUrl",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetSupportGroups(val interface{}) {
	_jsii_.Set(
		j,
		"supportGroups",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccessIdentityProviderConfigAOutputReference) SetTokenUrl(val *string) {
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetApiToken() {
	_jsii_.InvokeVoid(
		a,
		"resetApiToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAppsDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetAuthUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCentrifyAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetCentrifyAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCentrifyAppId() {
	_jsii_.InvokeVoid(
		a,
		"resetCentrifyAppId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetCertsUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetCertsUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetEmailAttributeName() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAttributeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetIdpPublicCert() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpPublicCert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetIssuerUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuerUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetOktaAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetOktaAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetOneloginAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetOneloginAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetPkceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPkceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSignRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetSignRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSsoTargetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSsoTargetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetSupportGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessIdentityProviderConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

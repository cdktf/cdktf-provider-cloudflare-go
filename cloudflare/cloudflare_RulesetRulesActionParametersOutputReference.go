// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRulesActionParametersOutputReference interface {
	cdktf.ComplexObject
	BrowserTtl() RulesetRulesActionParametersBrowserTtlOutputReference
	BrowserTtlInput() *RulesetRulesActionParametersBrowserTtl
	Cache() interface{}
	SetCache(val interface{})
	CacheInput() interface{}
	CacheKey() RulesetRulesActionParametersCacheKeyOutputReference
	CacheKeyInput() *RulesetRulesActionParametersCacheKey
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	CookieFields() *[]*string
	SetCookieFields(val *[]*string)
	CookieFieldsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EdgeTtl() RulesetRulesActionParametersEdgeTtlOutputReference
	EdgeTtlInput() *RulesetRulesActionParametersEdgeTtl
	// Experimental.
	Fqn() *string
	FromList() RulesetRulesActionParametersFromListOutputReference
	FromListInput() *RulesetRulesActionParametersFromList
	FromValue() RulesetRulesActionParametersFromValueOutputReference
	FromValueInput() *RulesetRulesActionParametersFromValue
	Headers() RulesetRulesActionParametersHeadersList
	HeadersInput() interface{}
	HostHeader() *string
	SetHostHeader(val *string)
	HostHeaderInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Increment() *float64
	SetIncrement(val *float64)
	IncrementInput() *float64
	InternalValue() *RulesetRulesActionParameters
	SetInternalValue(val *RulesetRulesActionParameters)
	MatchedData() RulesetRulesActionParametersMatchedDataOutputReference
	MatchedDataInput() *RulesetRulesActionParametersMatchedData
	Origin() RulesetRulesActionParametersOriginOutputReference
	OriginErrorPagePassthru() interface{}
	SetOriginErrorPagePassthru(val interface{})
	OriginErrorPagePassthruInput() interface{}
	OriginInput() *RulesetRulesActionParametersOrigin
	Overrides() RulesetRulesActionParametersOverridesOutputReference
	OverridesInput() *RulesetRulesActionParametersOverrides
	Phases() *[]*string
	SetPhases(val *[]*string)
	PhasesInput() *[]*string
	Products() *[]*string
	SetProducts(val *[]*string)
	ProductsInput() *[]*string
	RequestFields() *[]*string
	SetRequestFields(val *[]*string)
	RequestFieldsInput() *[]*string
	RespectStrongEtags() interface{}
	SetRespectStrongEtags(val interface{})
	RespectStrongEtagsInput() interface{}
	Response() RulesetRulesActionParametersResponseList
	ResponseFields() *[]*string
	SetResponseFields(val *[]*string)
	ResponseFieldsInput() *[]*string
	ResponseInput() interface{}
	Rules() *map[string]*string
	SetRules(val *map[string]*string)
	Ruleset() *string
	SetRuleset(val *string)
	RulesetInput() *string
	Rulesets() *[]*string
	SetRulesets(val *[]*string)
	RulesetsInput() *[]*string
	RulesInput() *map[string]*string
	ServeStale() RulesetRulesActionParametersServeStaleOutputReference
	ServeStaleInput() *RulesetRulesActionParametersServeStale
	Sni() RulesetRulesActionParametersSniOutputReference
	SniInput() *RulesetRulesActionParametersSni
	StatusCode() *float64
	SetStatusCode(val *float64)
	StatusCodeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() RulesetRulesActionParametersUriOutputReference
	UriInput() *RulesetRulesActionParametersUri
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutBrowserTtl(value *RulesetRulesActionParametersBrowserTtl)
	PutCacheKey(value *RulesetRulesActionParametersCacheKey)
	PutEdgeTtl(value *RulesetRulesActionParametersEdgeTtl)
	PutFromList(value *RulesetRulesActionParametersFromList)
	PutFromValue(value *RulesetRulesActionParametersFromValue)
	PutHeaders(value interface{})
	PutMatchedData(value *RulesetRulesActionParametersMatchedData)
	PutOrigin(value *RulesetRulesActionParametersOrigin)
	PutOverrides(value *RulesetRulesActionParametersOverrides)
	PutResponse(value interface{})
	PutServeStale(value *RulesetRulesActionParametersServeStale)
	PutSni(value *RulesetRulesActionParametersSni)
	PutUri(value *RulesetRulesActionParametersUri)
	ResetBrowserTtl()
	ResetCache()
	ResetCacheKey()
	ResetContent()
	ResetContentType()
	ResetCookieFields()
	ResetEdgeTtl()
	ResetFromList()
	ResetFromValue()
	ResetHeaders()
	ResetHostHeader()
	ResetId()
	ResetIncrement()
	ResetMatchedData()
	ResetOrigin()
	ResetOriginErrorPagePassthru()
	ResetOverrides()
	ResetPhases()
	ResetProducts()
	ResetRequestFields()
	ResetRespectStrongEtags()
	ResetResponse()
	ResetResponseFields()
	ResetRules()
	ResetRuleset()
	ResetRulesets()
	ResetServeStale()
	ResetSni()
	ResetStatusCode()
	ResetUri()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesActionParametersOutputReference
type jsiiProxy_RulesetRulesActionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) BrowserTtl() RulesetRulesActionParametersBrowserTtlOutputReference {
	var returns RulesetRulesActionParametersBrowserTtlOutputReference
	_jsii_.Get(
		j,
		"browserTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) BrowserTtlInput() *RulesetRulesActionParametersBrowserTtl {
	var returns *RulesetRulesActionParametersBrowserTtl
	_jsii_.Get(
		j,
		"browserTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Cache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheKey() RulesetRulesActionParametersCacheKeyOutputReference {
	var returns RulesetRulesActionParametersCacheKeyOutputReference
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheKeyInput() *RulesetRulesActionParametersCacheKey {
	var returns *RulesetRulesActionParametersCacheKey
	_jsii_.Get(
		j,
		"cacheKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CookieFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookieFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CookieFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookieFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EdgeTtl() RulesetRulesActionParametersEdgeTtlOutputReference {
	var returns RulesetRulesActionParametersEdgeTtlOutputReference
	_jsii_.Get(
		j,
		"edgeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EdgeTtlInput() *RulesetRulesActionParametersEdgeTtl {
	var returns *RulesetRulesActionParametersEdgeTtl
	_jsii_.Get(
		j,
		"edgeTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromList() RulesetRulesActionParametersFromListOutputReference {
	var returns RulesetRulesActionParametersFromListOutputReference
	_jsii_.Get(
		j,
		"fromList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromListInput() *RulesetRulesActionParametersFromList {
	var returns *RulesetRulesActionParametersFromList
	_jsii_.Get(
		j,
		"fromListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromValue() RulesetRulesActionParametersFromValueOutputReference {
	var returns RulesetRulesActionParametersFromValueOutputReference
	_jsii_.Get(
		j,
		"fromValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromValueInput() *RulesetRulesActionParametersFromValue {
	var returns *RulesetRulesActionParametersFromValue
	_jsii_.Get(
		j,
		"fromValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Headers() RulesetRulesActionParametersHeadersList {
	var returns RulesetRulesActionParametersHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Increment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"increment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) IncrementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) InternalValue() *RulesetRulesActionParameters {
	var returns *RulesetRulesActionParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) MatchedData() RulesetRulesActionParametersMatchedDataOutputReference {
	var returns RulesetRulesActionParametersMatchedDataOutputReference
	_jsii_.Get(
		j,
		"matchedData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) MatchedDataInput() *RulesetRulesActionParametersMatchedData {
	var returns *RulesetRulesActionParametersMatchedData
	_jsii_.Get(
		j,
		"matchedDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Origin() RulesetRulesActionParametersOriginOutputReference {
	var returns RulesetRulesActionParametersOriginOutputReference
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginErrorPagePassthru() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originErrorPagePassthru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginErrorPagePassthruInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originErrorPagePassthruInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginInput() *RulesetRulesActionParametersOrigin {
	var returns *RulesetRulesActionParametersOrigin
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Overrides() RulesetRulesActionParametersOverridesOutputReference {
	var returns RulesetRulesActionParametersOverridesOutputReference
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OverridesInput() *RulesetRulesActionParametersOverrides {
	var returns *RulesetRulesActionParametersOverrides
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Phases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) PhasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Products() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"products",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ProductsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RequestFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RequestFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RespectStrongEtags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectStrongEtags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RespectStrongEtagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectStrongEtagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Response() RulesetRulesActionParametersResponseList {
	var returns RulesetRulesActionParametersResponseList
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Rules() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Ruleset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Rulesets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServeStale() RulesetRulesActionParametersServeStaleOutputReference {
	var returns RulesetRulesActionParametersServeStaleOutputReference
	_jsii_.Get(
		j,
		"serveStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServeStaleInput() *RulesetRulesActionParametersServeStale {
	var returns *RulesetRulesActionParametersServeStale
	_jsii_.Get(
		j,
		"serveStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Sni() RulesetRulesActionParametersSniOutputReference {
	var returns RulesetRulesActionParametersSniOutputReference
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SniInput() *RulesetRulesActionParametersSni {
	var returns *RulesetRulesActionParametersSni
	_jsii_.Get(
		j,
		"sniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) StatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) StatusCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Uri() RulesetRulesActionParametersUriOutputReference {
	var returns RulesetRulesActionParametersUriOutputReference
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) UriInput() *RulesetRulesActionParametersUri {
	var returns *RulesetRulesActionParametersUri
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewRulesetRulesActionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RulesetRulesActionParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RulesetRulesActionParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRulesetRulesActionParametersOutputReference_Override(r RulesetRulesActionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetCache(val interface{}) {
	_jsii_.Set(
		j,
		"cache",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetCookieFields(val *[]*string) {
	_jsii_.Set(
		j,
		"cookieFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetHostHeader(val *string) {
	_jsii_.Set(
		j,
		"hostHeader",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetIncrement(val *float64) {
	_jsii_.Set(
		j,
		"increment",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetInternalValue(val *RulesetRulesActionParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetOriginErrorPagePassthru(val interface{}) {
	_jsii_.Set(
		j,
		"originErrorPagePassthru",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetPhases(val *[]*string) {
	_jsii_.Set(
		j,
		"phases",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetProducts(val *[]*string) {
	_jsii_.Set(
		j,
		"products",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetRequestFields(val *[]*string) {
	_jsii_.Set(
		j,
		"requestFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetRespectStrongEtags(val interface{}) {
	_jsii_.Set(
		j,
		"respectStrongEtags",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetResponseFields(val *[]*string) {
	_jsii_.Set(
		j,
		"responseFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetRules(val *map[string]*string) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetRuleset(val *string) {
	_jsii_.Set(
		j,
		"ruleset",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetRulesets(val *[]*string) {
	_jsii_.Set(
		j,
		"rulesets",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetStatusCode(val *float64) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutBrowserTtl(value *RulesetRulesActionParametersBrowserTtl) {
	_jsii_.InvokeVoid(
		r,
		"putBrowserTtl",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutCacheKey(value *RulesetRulesActionParametersCacheKey) {
	_jsii_.InvokeVoid(
		r,
		"putCacheKey",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutEdgeTtl(value *RulesetRulesActionParametersEdgeTtl) {
	_jsii_.InvokeVoid(
		r,
		"putEdgeTtl",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutFromList(value *RulesetRulesActionParametersFromList) {
	_jsii_.InvokeVoid(
		r,
		"putFromList",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutFromValue(value *RulesetRulesActionParametersFromValue) {
	_jsii_.InvokeVoid(
		r,
		"putFromValue",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putHeaders",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutMatchedData(value *RulesetRulesActionParametersMatchedData) {
	_jsii_.InvokeVoid(
		r,
		"putMatchedData",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutOrigin(value *RulesetRulesActionParametersOrigin) {
	_jsii_.InvokeVoid(
		r,
		"putOrigin",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutOverrides(value *RulesetRulesActionParametersOverrides) {
	_jsii_.InvokeVoid(
		r,
		"putOverrides",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutResponse(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putResponse",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutServeStale(value *RulesetRulesActionParametersServeStale) {
	_jsii_.InvokeVoid(
		r,
		"putServeStale",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutSni(value *RulesetRulesActionParametersSni) {
	_jsii_.InvokeVoid(
		r,
		"putSni",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutUri(value *RulesetRulesActionParametersUri) {
	_jsii_.InvokeVoid(
		r,
		"putUri",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetBrowserTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetBrowserTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCache() {
	_jsii_.InvokeVoid(
		r,
		"resetCache",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCacheKey() {
	_jsii_.InvokeVoid(
		r,
		"resetCacheKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		r,
		"resetContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		r,
		"resetContentType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCookieFields() {
	_jsii_.InvokeVoid(
		r,
		"resetCookieFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetEdgeTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetEdgeTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetFromList() {
	_jsii_.InvokeVoid(
		r,
		"resetFromList",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetFromValue() {
	_jsii_.InvokeVoid(
		r,
		"resetFromValue",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		r,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetIncrement() {
	_jsii_.InvokeVoid(
		r,
		"resetIncrement",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetMatchedData() {
	_jsii_.InvokeVoid(
		r,
		"resetMatchedData",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		r,
		"resetOrigin",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOriginErrorPagePassthru() {
	_jsii_.InvokeVoid(
		r,
		"resetOriginErrorPagePassthru",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOverrides() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrides",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetPhases() {
	_jsii_.InvokeVoid(
		r,
		"resetPhases",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetProducts() {
	_jsii_.InvokeVoid(
		r,
		"resetProducts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRequestFields() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRespectStrongEtags() {
	_jsii_.InvokeVoid(
		r,
		"resetRespectStrongEtags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		r,
		"resetResponse",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetResponseFields() {
	_jsii_.InvokeVoid(
		r,
		"resetResponseFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		r,
		"resetRules",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRuleset() {
	_jsii_.InvokeVoid(
		r,
		"resetRuleset",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRulesets() {
	_jsii_.InvokeVoid(
		r,
		"resetRulesets",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetServeStale() {
	_jsii_.InvokeVoid(
		r,
		"resetServeStale",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetSni() {
	_jsii_.InvokeVoid(
		r,
		"resetSni",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		r,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetUri() {
	_jsii_.InvokeVoid(
		r,
		"resetUri",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/datacloudflareruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareRulesetRulesActionParametersOutputReference interface {
	cdktf.ComplexObject
	AdditionalCacheablePorts() *[]*float64
	Algorithms() DataCloudflareRulesetRulesActionParametersAlgorithmsList
	AssetName() *string
	AutomaticHttpsRewrites() cdktf.IResolvable
	Autominify() DataCloudflareRulesetRulesActionParametersAutominifyOutputReference
	Bic() cdktf.IResolvable
	BrowserTtl() DataCloudflareRulesetRulesActionParametersBrowserTtlOutputReference
	Cache() cdktf.IResolvable
	CacheKey() DataCloudflareRulesetRulesActionParametersCacheKeyOutputReference
	CacheReserve() DataCloudflareRulesetRulesActionParametersCacheReserveOutputReference
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
	ContentType() *string
	CookieFields() DataCloudflareRulesetRulesActionParametersCookieFieldsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableApps() cdktf.IResolvable
	DisableRum() cdktf.IResolvable
	DisableZaraz() cdktf.IResolvable
	EdgeTtl() DataCloudflareRulesetRulesActionParametersEdgeTtlOutputReference
	EmailObfuscation() cdktf.IResolvable
	Fonts() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	FromList() DataCloudflareRulesetRulesActionParametersFromListStructOutputReference
	FromValue() DataCloudflareRulesetRulesActionParametersFromValueOutputReference
	Headers() DataCloudflareRulesetRulesActionParametersHeadersMap
	HostHeader() *string
	HotlinkProtection() cdktf.IResolvable
	Id() *string
	Increment() *float64
	InternalValue() *DataCloudflareRulesetRulesActionParameters
	SetInternalValue(val *DataCloudflareRulesetRulesActionParameters)
	MatchedData() DataCloudflareRulesetRulesActionParametersMatchedDataOutputReference
	Mirage() cdktf.IResolvable
	OpportunisticEncryption() cdktf.IResolvable
	Origin() DataCloudflareRulesetRulesActionParametersOriginOutputReference
	OriginCacheControl() cdktf.IResolvable
	OriginErrorPagePassthru() cdktf.IResolvable
	Overrides() DataCloudflareRulesetRulesActionParametersOverridesOutputReference
	Phases() *[]*string
	Polish() *string
	Products() *[]*string
	RawResponseFields() DataCloudflareRulesetRulesActionParametersRawResponseFieldsList
	ReadTimeout() *float64
	RequestFields() DataCloudflareRulesetRulesActionParametersRequestFieldsList
	RespectStrongEtags() cdktf.IResolvable
	Response() DataCloudflareRulesetRulesActionParametersResponseOutputReference
	ResponseFields() DataCloudflareRulesetRulesActionParametersResponseFieldsList
	RocketLoader() cdktf.IResolvable
	Rules() cdktf.StringListMap
	Ruleset() *string
	Rulesets() *[]*string
	SecurityLevel() *string
	ServerSideExcludes() cdktf.IResolvable
	ServeStale() DataCloudflareRulesetRulesActionParametersServeStaleOutputReference
	Sni() DataCloudflareRulesetRulesActionParametersSniOutputReference
	Ssl() *string
	StatusCode() *float64
	Sxg() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformedRequestFields() DataCloudflareRulesetRulesActionParametersTransformedRequestFieldsList
	Uri() DataCloudflareRulesetRulesActionParametersUriOutputReference
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

// The jsii proxy struct for DataCloudflareRulesetRulesActionParametersOutputReference
type jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) AdditionalCacheablePorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"additionalCacheablePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Algorithms() DataCloudflareRulesetRulesActionParametersAlgorithmsList {
	var returns DataCloudflareRulesetRulesActionParametersAlgorithmsList
	_jsii_.Get(
		j,
		"algorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) AssetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) AutomaticHttpsRewrites() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Autominify() DataCloudflareRulesetRulesActionParametersAutominifyOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersAutominifyOutputReference
	_jsii_.Get(
		j,
		"autominify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Bic() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) BrowserTtl() DataCloudflareRulesetRulesActionParametersBrowserTtlOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersBrowserTtlOutputReference
	_jsii_.Get(
		j,
		"browserTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Cache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) CacheKey() DataCloudflareRulesetRulesActionParametersCacheKeyOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersCacheKeyOutputReference
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) CacheReserve() DataCloudflareRulesetRulesActionParametersCacheReserveOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersCacheReserveOutputReference
	_jsii_.Get(
		j,
		"cacheReserve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) CookieFields() DataCloudflareRulesetRulesActionParametersCookieFieldsList {
	var returns DataCloudflareRulesetRulesActionParametersCookieFieldsList
	_jsii_.Get(
		j,
		"cookieFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) DisableApps() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) DisableRum() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableRum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) DisableZaraz() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableZaraz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) EdgeTtl() DataCloudflareRulesetRulesActionParametersEdgeTtlOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersEdgeTtlOutputReference
	_jsii_.Get(
		j,
		"edgeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) EmailObfuscation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Fonts() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fonts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) FromList() DataCloudflareRulesetRulesActionParametersFromListStructOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersFromListStructOutputReference
	_jsii_.Get(
		j,
		"fromList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) FromValue() DataCloudflareRulesetRulesActionParametersFromValueOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersFromValueOutputReference
	_jsii_.Get(
		j,
		"fromValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Headers() DataCloudflareRulesetRulesActionParametersHeadersMap {
	var returns DataCloudflareRulesetRulesActionParametersHeadersMap
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) HotlinkProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hotlinkProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Increment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"increment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) InternalValue() *DataCloudflareRulesetRulesActionParameters {
	var returns *DataCloudflareRulesetRulesActionParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) MatchedData() DataCloudflareRulesetRulesActionParametersMatchedDataOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersMatchedDataOutputReference
	_jsii_.Get(
		j,
		"matchedData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Mirage() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) OpportunisticEncryption() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Origin() DataCloudflareRulesetRulesActionParametersOriginOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersOriginOutputReference
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) OriginCacheControl() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"originCacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) OriginErrorPagePassthru() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"originErrorPagePassthru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Overrides() DataCloudflareRulesetRulesActionParametersOverridesOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersOverridesOutputReference
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Phases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Products() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"products",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) RawResponseFields() DataCloudflareRulesetRulesActionParametersRawResponseFieldsList {
	var returns DataCloudflareRulesetRulesActionParametersRawResponseFieldsList
	_jsii_.Get(
		j,
		"rawResponseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) RequestFields() DataCloudflareRulesetRulesActionParametersRequestFieldsList {
	var returns DataCloudflareRulesetRulesActionParametersRequestFieldsList
	_jsii_.Get(
		j,
		"requestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) RespectStrongEtags() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"respectStrongEtags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Response() DataCloudflareRulesetRulesActionParametersResponseOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersResponseOutputReference
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ResponseFields() DataCloudflareRulesetRulesActionParametersResponseFieldsList {
	var returns DataCloudflareRulesetRulesActionParametersResponseFieldsList
	_jsii_.Get(
		j,
		"responseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) RocketLoader() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Rules() cdktf.StringListMap {
	var returns cdktf.StringListMap
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Ruleset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Rulesets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ServerSideExcludes() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"serverSideExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ServeStale() DataCloudflareRulesetRulesActionParametersServeStaleOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersServeStaleOutputReference
	_jsii_.Get(
		j,
		"serveStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Sni() DataCloudflareRulesetRulesActionParametersSniOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersSniOutputReference
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) StatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Sxg() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sxg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) TransformedRequestFields() DataCloudflareRulesetRulesActionParametersTransformedRequestFieldsList {
	var returns DataCloudflareRulesetRulesActionParametersTransformedRequestFieldsList
	_jsii_.Get(
		j,
		"transformedRequestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Uri() DataCloudflareRulesetRulesActionParametersUriOutputReference {
	var returns DataCloudflareRulesetRulesActionParametersUriOutputReference
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


func NewDataCloudflareRulesetRulesActionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareRulesetRulesActionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareRulesetRulesActionParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareRuleset.DataCloudflareRulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareRulesetRulesActionParametersOutputReference_Override(d DataCloudflareRulesetRulesActionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareRuleset.DataCloudflareRulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference)SetInternalValue(val *DataCloudflareRulesetRulesActionParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareRulesetRulesActionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


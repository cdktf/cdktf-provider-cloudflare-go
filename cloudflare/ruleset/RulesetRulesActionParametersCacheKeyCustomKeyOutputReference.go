package ruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/ruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRulesActionParametersCacheKeyCustomKeyOutputReference interface {
	cdktf.ComplexObject
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
	Cookie() RulesetRulesActionParametersCacheKeyCustomKeyCookieOutputReference
	CookieInput() *RulesetRulesActionParametersCacheKeyCustomKeyCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Header() RulesetRulesActionParametersCacheKeyCustomKeyHeaderOutputReference
	HeaderInput() *RulesetRulesActionParametersCacheKeyCustomKeyHeader
	Host() RulesetRulesActionParametersCacheKeyCustomKeyHostOutputReference
	HostInput() *RulesetRulesActionParametersCacheKeyCustomKeyHost
	InternalValue() *RulesetRulesActionParametersCacheKeyCustomKey
	SetInternalValue(val *RulesetRulesActionParametersCacheKeyCustomKey)
	QueryString() RulesetRulesActionParametersCacheKeyCustomKeyQueryStringOutputReference
	QueryStringInput() *RulesetRulesActionParametersCacheKeyCustomKeyQueryString
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() RulesetRulesActionParametersCacheKeyCustomKeyUserOutputReference
	UserInput() *RulesetRulesActionParametersCacheKeyCustomKeyUser
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
	PutCookie(value *RulesetRulesActionParametersCacheKeyCustomKeyCookie)
	PutHeader(value *RulesetRulesActionParametersCacheKeyCustomKeyHeader)
	PutHost(value *RulesetRulesActionParametersCacheKeyCustomKeyHost)
	PutQueryString(value *RulesetRulesActionParametersCacheKeyCustomKeyQueryString)
	PutUser(value *RulesetRulesActionParametersCacheKeyCustomKeyUser)
	ResetCookie()
	ResetHeader()
	ResetHost()
	ResetQueryString()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesActionParametersCacheKeyCustomKeyOutputReference
type jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) Cookie() RulesetRulesActionParametersCacheKeyCustomKeyCookieOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) CookieInput() *RulesetRulesActionParametersCacheKeyCustomKeyCookie {
	var returns *RulesetRulesActionParametersCacheKeyCustomKeyCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) Header() RulesetRulesActionParametersCacheKeyCustomKeyHeaderOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) HeaderInput() *RulesetRulesActionParametersCacheKeyCustomKeyHeader {
	var returns *RulesetRulesActionParametersCacheKeyCustomKeyHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) Host() RulesetRulesActionParametersCacheKeyCustomKeyHostOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) HostInput() *RulesetRulesActionParametersCacheKeyCustomKeyHost {
	var returns *RulesetRulesActionParametersCacheKeyCustomKeyHost
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) InternalValue() *RulesetRulesActionParametersCacheKeyCustomKey {
	var returns *RulesetRulesActionParametersCacheKeyCustomKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) QueryString() RulesetRulesActionParametersCacheKeyCustomKeyQueryStringOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) QueryStringInput() *RulesetRulesActionParametersCacheKeyCustomKeyQueryString {
	var returns *RulesetRulesActionParametersCacheKeyCustomKeyQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) User() RulesetRulesActionParametersCacheKeyCustomKeyUserOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyUserOutputReference
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) UserInput() *RulesetRulesActionParametersCacheKeyCustomKeyUser {
	var returns *RulesetRulesActionParametersCacheKeyCustomKeyUser
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewRulesetRulesActionParametersCacheKeyCustomKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RulesetRulesActionParametersCacheKeyCustomKeyOutputReference {
	_init_.Initialize()

	if err := validateNewRulesetRulesActionParametersCacheKeyCustomKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.ruleset.RulesetRulesActionParametersCacheKeyCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRulesetRulesActionParametersCacheKeyCustomKeyOutputReference_Override(r RulesetRulesActionParametersCacheKeyCustomKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.ruleset.RulesetRulesActionParametersCacheKeyCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference)SetInternalValue(val *RulesetRulesActionParametersCacheKeyCustomKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) PutCookie(value *RulesetRulesActionParametersCacheKeyCustomKeyCookie) {
	if err := r.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCookie",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) PutHeader(value *RulesetRulesActionParametersCacheKeyCustomKeyHeader) {
	if err := r.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHeader",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) PutHost(value *RulesetRulesActionParametersCacheKeyCustomKeyHost) {
	if err := r.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHost",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) PutQueryString(value *RulesetRulesActionParametersCacheKeyCustomKeyQueryString) {
	if err := r.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putQueryString",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) PutUser(value *RulesetRulesActionParametersCacheKeyCustomKeyUser) {
	if err := r.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putUser",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		r,
		"resetCookie",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		r,
		"resetHeader",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		r,
		"resetHost",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		r,
		"resetQueryString",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		r,
		"resetUser",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyCustomKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


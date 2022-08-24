// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRulesActionParametersAutominifyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Css() interface{}
	SetCss(val interface{})
	CssInput() interface{}
	// Experimental.
	Fqn() *string
	Html() interface{}
	SetHtml(val interface{})
	HtmlInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Js() interface{}
	SetJs(val interface{})
	JsInput() interface{}
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
	ResetCss()
	ResetHtml()
	ResetJs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesActionParametersAutominifyOutputReference
type jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) Css() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"css",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) CssInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) Html() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"html",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) HtmlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"htmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) Js() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"js",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) JsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRulesetRulesActionParametersAutominifyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RulesetRulesActionParametersAutominifyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersAutominifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRulesetRulesActionParametersAutominifyOutputReference_Override(r RulesetRulesActionParametersAutominifyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersAutominifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetCss(val interface{}) {
	_jsii_.Set(
		j,
		"css",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetHtml(val interface{}) {
	_jsii_.Set(
		j,
		"html",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetJs(val interface{}) {
	_jsii_.Set(
		j,
		"js",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ResetCss() {
	_jsii_.InvokeVoid(
		r,
		"resetCss",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ResetHtml() {
	_jsii_.InvokeVoid(
		r,
		"resetHtml",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ResetJs() {
	_jsii_.InvokeVoid(
		r,
		"resetJs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersAutominifyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


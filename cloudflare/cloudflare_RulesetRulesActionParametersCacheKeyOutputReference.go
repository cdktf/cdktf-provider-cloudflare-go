// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRulesActionParametersCacheKeyOutputReference interface {
	cdktf.ComplexObject
	CacheByDeviceType() interface{}
	SetCacheByDeviceType(val interface{})
	CacheByDeviceTypeInput() interface{}
	CacheDeceptionArmor() interface{}
	SetCacheDeceptionArmor(val interface{})
	CacheDeceptionArmorInput() interface{}
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
	CustomKey() RulesetRulesActionParametersCacheKeyCustomKeyOutputReference
	CustomKeyInput() *RulesetRulesActionParametersCacheKeyCustomKey
	// Experimental.
	Fqn() *string
	IgnoreQueryStringsOrder() interface{}
	SetIgnoreQueryStringsOrder(val interface{})
	IgnoreQueryStringsOrderInput() interface{}
	InternalValue() *RulesetRulesActionParametersCacheKey
	SetInternalValue(val *RulesetRulesActionParametersCacheKey)
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
	PutCustomKey(value *RulesetRulesActionParametersCacheKeyCustomKey)
	ResetCacheByDeviceType()
	ResetCacheDeceptionArmor()
	ResetCustomKey()
	ResetIgnoreQueryStringsOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesActionParametersCacheKeyOutputReference
type jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CacheByDeviceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheByDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CacheByDeviceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheByDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CacheDeceptionArmor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDeceptionArmor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CacheDeceptionArmorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDeceptionArmorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CustomKey() RulesetRulesActionParametersCacheKeyCustomKeyOutputReference {
	var returns RulesetRulesActionParametersCacheKeyCustomKeyOutputReference
	_jsii_.Get(
		j,
		"customKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) CustomKeyInput() *RulesetRulesActionParametersCacheKeyCustomKey {
	var returns *RulesetRulesActionParametersCacheKeyCustomKey
	_jsii_.Get(
		j,
		"customKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) IgnoreQueryStringsOrder() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryStringsOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) IgnoreQueryStringsOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreQueryStringsOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) InternalValue() *RulesetRulesActionParametersCacheKey {
	var returns *RulesetRulesActionParametersCacheKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRulesetRulesActionParametersCacheKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RulesetRulesActionParametersCacheKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersCacheKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRulesetRulesActionParametersCacheKeyOutputReference_Override(r RulesetRulesActionParametersCacheKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesActionParametersCacheKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetCacheByDeviceType(val interface{}) {
	_jsii_.Set(
		j,
		"cacheByDeviceType",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetCacheDeceptionArmor(val interface{}) {
	_jsii_.Set(
		j,
		"cacheDeceptionArmor",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetIgnoreQueryStringsOrder(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreQueryStringsOrder",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetInternalValue(val *RulesetRulesActionParametersCacheKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) PutCustomKey(value *RulesetRulesActionParametersCacheKeyCustomKey) {
	_jsii_.InvokeVoid(
		r,
		"putCustomKey",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ResetCacheByDeviceType() {
	_jsii_.InvokeVoid(
		r,
		"resetCacheByDeviceType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ResetCacheDeceptionArmor() {
	_jsii_.InvokeVoid(
		r,
		"resetCacheDeceptionArmor",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ResetCustomKey() {
	_jsii_.InvokeVoid(
		r,
		"resetCustomKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ResetIgnoreQueryStringsOrder() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreQueryStringsOrder",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersCacheKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


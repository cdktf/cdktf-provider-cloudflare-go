// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRulesRatelimitOutputReference interface {
	cdktf.ComplexObject
	Characteristics() *[]*string
	SetCharacteristics(val *[]*string)
	CharacteristicsInput() *[]*string
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
	CountingExpression() *string
	SetCountingExpression(val *string)
	CountingExpressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *RulesetRulesRatelimit
	SetInternalValue(val *RulesetRulesRatelimit)
	MitigationTimeout() *float64
	SetMitigationTimeout(val *float64)
	MitigationTimeoutInput() *float64
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	RequestsPerPeriod() *float64
	SetRequestsPerPeriod(val *float64)
	RequestsPerPeriodInput() *float64
	RequestsToOrigin() interface{}
	SetRequestsToOrigin(val interface{})
	RequestsToOriginInput() interface{}
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
	ResetCharacteristics()
	ResetCountingExpression()
	ResetMitigationTimeout()
	ResetPeriod()
	ResetRequestsPerPeriod()
	ResetRequestsToOrigin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesRatelimitOutputReference
type jsiiProxy_RulesetRulesRatelimitOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) Characteristics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"characteristics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) CharacteristicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"characteristicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) CountingExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countingExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) CountingExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countingExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) InternalValue() *RulesetRulesRatelimit {
	var returns *RulesetRulesRatelimit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) MitigationTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mitigationTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) MitigationTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mitigationTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) RequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) RequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) RequestsToOrigin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestsToOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) RequestsToOriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestsToOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRulesetRulesRatelimitOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RulesetRulesRatelimitOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RulesetRulesRatelimitOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesRatelimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRulesetRulesRatelimitOutputReference_Override(r RulesetRulesRatelimitOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.RulesetRulesRatelimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetCharacteristics(val *[]*string) {
	_jsii_.Set(
		j,
		"characteristics",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetCountingExpression(val *string) {
	_jsii_.Set(
		j,
		"countingExpression",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetInternalValue(val *RulesetRulesRatelimit) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetMitigationTimeout(val *float64) {
	_jsii_.Set(
		j,
		"mitigationTimeout",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetPeriod(val *float64) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetRequestsPerPeriod(val *float64) {
	_jsii_.Set(
		j,
		"requestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetRequestsToOrigin(val interface{}) {
	_jsii_.Set(
		j,
		"requestsToOrigin",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesRatelimitOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetCharacteristics() {
	_jsii_.InvokeVoid(
		r,
		"resetCharacteristics",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetCountingExpression() {
	_jsii_.InvokeVoid(
		r,
		"resetCountingExpression",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetMitigationTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetMitigationTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestsPerPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ResetRequestsToOrigin() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestsToOrigin",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesRatelimitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

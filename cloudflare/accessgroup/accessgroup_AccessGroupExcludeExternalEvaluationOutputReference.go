package accessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/accessgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessGroupExcludeExternalEvaluationOutputReference interface {
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
	EvaluateUrl() *string
	SetEvaluateUrl(val *string)
	EvaluateUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AccessGroupExcludeExternalEvaluation
	SetInternalValue(val *AccessGroupExcludeExternalEvaluation)
	KeysUrl() *string
	SetKeysUrl(val *string)
	KeysUrlInput() *string
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
	ResetEvaluateUrl()
	ResetKeysUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessGroupExcludeExternalEvaluationOutputReference
type jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) EvaluateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) EvaluateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) InternalValue() *AccessGroupExcludeExternalEvaluation {
	var returns *AccessGroupExcludeExternalEvaluation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) KeysUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keysUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) KeysUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keysUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessGroupExcludeExternalEvaluationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessGroupExcludeExternalEvaluationOutputReference {
	_init_.Initialize()

	if err := validateNewAccessGroupExcludeExternalEvaluationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessGroup.AccessGroupExcludeExternalEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessGroupExcludeExternalEvaluationOutputReference_Override(a AccessGroupExcludeExternalEvaluationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessGroup.AccessGroupExcludeExternalEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetEvaluateUrl(val *string) {
	if err := j.validateSetEvaluateUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateUrl",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetInternalValue(val *AccessGroupExcludeExternalEvaluation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetKeysUrl(val *string) {
	if err := j.validateSetKeysUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keysUrl",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ResetEvaluateUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluateUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ResetKeysUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetKeysUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessGroupExcludeExternalEvaluationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


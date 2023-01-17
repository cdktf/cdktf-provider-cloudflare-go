package tunnelconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v5/tunnelconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TunnelConfigConfigAOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IngressRule() TunnelConfigConfigIngressRuleList
	IngressRuleInput() interface{}
	InternalValue() *TunnelConfigConfigA
	SetInternalValue(val *TunnelConfigConfigA)
	OriginRequest() TunnelConfigConfigOriginRequestOutputReference
	OriginRequestInput() *TunnelConfigConfigOriginRequest
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarpRouting() TunnelConfigConfigWarpRoutingOutputReference
	WarpRoutingInput() *TunnelConfigConfigWarpRouting
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
	PutIngressRule(value interface{})
	PutOriginRequest(value *TunnelConfigConfigOriginRequest)
	PutWarpRouting(value *TunnelConfigConfigWarpRouting)
	ResetOriginRequest()
	ResetWarpRouting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TunnelConfigConfigAOutputReference
type jsiiProxy_TunnelConfigConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) IngressRule() TunnelConfigConfigIngressRuleList {
	var returns TunnelConfigConfigIngressRuleList
	_jsii_.Get(
		j,
		"ingressRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) IngressRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) InternalValue() *TunnelConfigConfigA {
	var returns *TunnelConfigConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) OriginRequest() TunnelConfigConfigOriginRequestOutputReference {
	var returns TunnelConfigConfigOriginRequestOutputReference
	_jsii_.Get(
		j,
		"originRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) OriginRequestInput() *TunnelConfigConfigOriginRequest {
	var returns *TunnelConfigConfigOriginRequest
	_jsii_.Get(
		j,
		"originRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) WarpRouting() TunnelConfigConfigWarpRoutingOutputReference {
	var returns TunnelConfigConfigWarpRoutingOutputReference
	_jsii_.Get(
		j,
		"warpRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference) WarpRoutingInput() *TunnelConfigConfigWarpRouting {
	var returns *TunnelConfigConfigWarpRouting
	_jsii_.Get(
		j,
		"warpRoutingInput",
		&returns,
	)
	return returns
}


func NewTunnelConfigConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TunnelConfigConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewTunnelConfigConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TunnelConfigConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTunnelConfigConfigAOutputReference_Override(t TunnelConfigConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference)SetInternalValue(val *TunnelConfigConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) PutIngressRule(value interface{}) {
	if err := t.validatePutIngressRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIngressRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) PutOriginRequest(value *TunnelConfigConfigOriginRequest) {
	if err := t.validatePutOriginRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOriginRequest",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) PutWarpRouting(value *TunnelConfigConfigWarpRouting) {
	if err := t.validatePutWarpRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWarpRouting",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) ResetOriginRequest() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginRequest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) ResetWarpRouting() {
	_jsii_.InvokeVoid(
		t,
		"resetWarpRouting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


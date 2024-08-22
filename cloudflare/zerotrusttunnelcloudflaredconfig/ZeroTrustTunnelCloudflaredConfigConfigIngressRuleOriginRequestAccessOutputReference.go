// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrusttunnelcloudflaredconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference interface {
	cdktf.ComplexObject
	AudTag() *[]*string
	SetAudTag(val *[]*string)
	AudTagInput() *[]*string
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
	InternalValue() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess
	SetInternalValue(val *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess)
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	TeamName() *string
	SetTeamName(val *string)
	TeamNameInput() *string
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
	ResetAudTag()
	ResetRequired()
	ResetTeamName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference
type jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) AudTag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) AudTagInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) InternalValue() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess {
	var returns *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) TeamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) TeamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference_Override(z ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetAudTag(val *[]*string) {
	if err := j.validateSetAudTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audTag",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetInternalValue(val *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetTeamName(val *string) {
	if err := j.validateSetTeamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ResetAudTag() {
	_jsii_.InvokeVoid(
		z,
		"resetAudTag",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ResetTeamName() {
	_jsii_.InvokeVoid(
		z,
		"resetTeamName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


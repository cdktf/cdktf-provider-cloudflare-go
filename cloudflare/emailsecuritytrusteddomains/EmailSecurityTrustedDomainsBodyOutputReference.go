// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailsecuritytrusteddomains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/emailsecuritytrusteddomains/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailSecurityTrustedDomainsBodyOutputReference interface {
	cdktf.ComplexObject
	Comments() *string
	SetComments(val *string)
	CommentsInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsRecent() interface{}
	SetIsRecent(val interface{})
	IsRecentInput() interface{}
	IsRegex() interface{}
	SetIsRegex(val interface{})
	IsRegexInput() interface{}
	IsSimilarity() interface{}
	SetIsSimilarity(val interface{})
	IsSimilarityInput() interface{}
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
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
	ResetComments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmailSecurityTrustedDomainsBodyOutputReference
type jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) Comments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) CommentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsRecent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsRecentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRecentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsRegexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsSimilarity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSimilarity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) IsSimilarityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSimilarityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmailSecurityTrustedDomainsBodyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmailSecurityTrustedDomainsBodyOutputReference {
	_init_.Initialize()

	if err := validateNewEmailSecurityTrustedDomainsBodyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomainsBodyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmailSecurityTrustedDomainsBodyOutputReference_Override(e EmailSecurityTrustedDomainsBodyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomainsBodyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetComments(val *string) {
	if err := j.validateSetCommentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comments",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetIsRecent(val interface{}) {
	if err := j.validateSetIsRecentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRecent",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetIsRegex(val interface{}) {
	if err := j.validateSetIsRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRegex",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetIsSimilarity(val interface{}) {
	if err := j.validateSetIsSimilarityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSimilarity",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) ResetComments() {
	_jsii_.InvokeVoid(
		e,
		"resetComments",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomainsBodyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


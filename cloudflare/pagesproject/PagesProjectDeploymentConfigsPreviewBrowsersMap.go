// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/pagesproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PagesProjectDeploymentConfigsPreviewBrowsersMap interface {
	cdktf.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) PagesProjectDeploymentConfigsPreviewBrowsersOutputReference
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

// The jsii proxy struct for PagesProjectDeploymentConfigsPreviewBrowsersMap
type jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPagesProjectDeploymentConfigsPreviewBrowsersMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectDeploymentConfigsPreviewBrowsersMap {
	_init_.Initialize()

	if err := validateNewPagesProjectDeploymentConfigsPreviewBrowsersMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewBrowsersMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectDeploymentConfigsPreviewBrowsersMap_Override(p PagesProjectDeploymentConfigsPreviewBrowsersMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewBrowsersMap",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) Get(key *string) PagesProjectDeploymentConfigsPreviewBrowsersOutputReference {
	if err := p.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns PagesProjectDeploymentConfigsPreviewBrowsersOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewBrowsersMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


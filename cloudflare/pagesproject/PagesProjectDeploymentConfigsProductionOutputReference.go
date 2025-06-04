// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/pagesproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PagesProjectDeploymentConfigsProductionOutputReference interface {
	cdktf.ComplexObject
	AiBindings() PagesProjectDeploymentConfigsProductionAiBindingsMap
	AiBindingsInput() interface{}
	AnalyticsEngineDatasets() PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsMap
	AnalyticsEngineDatasetsInput() interface{}
	Browsers() PagesProjectDeploymentConfigsProductionBrowsersMap
	BrowsersInput() interface{}
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
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
	D1Databases() PagesProjectDeploymentConfigsProductionD1DatabasesMap
	D1DatabasesInput() interface{}
	DurableObjectNamespaces() PagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap
	DurableObjectNamespacesInput() interface{}
	EnvVars() PagesProjectDeploymentConfigsProductionEnvVarsMap
	EnvVarsInput() interface{}
	// Experimental.
	Fqn() *string
	HyperdriveBindings() PagesProjectDeploymentConfigsProductionHyperdriveBindingsMap
	HyperdriveBindingsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KvNamespaces() PagesProjectDeploymentConfigsProductionKvNamespacesMap
	KvNamespacesInput() interface{}
	MtlsCertificates() PagesProjectDeploymentConfigsProductionMtlsCertificatesMap
	MtlsCertificatesInput() interface{}
	Placement() PagesProjectDeploymentConfigsProductionPlacementOutputReference
	PlacementInput() interface{}
	QueueProducers() PagesProjectDeploymentConfigsProductionQueueProducersMap
	QueueProducersInput() interface{}
	R2Buckets() PagesProjectDeploymentConfigsProductionR2BucketsMap
	R2BucketsInput() interface{}
	Services() PagesProjectDeploymentConfigsProductionServicesMap
	ServicesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VectorizeBindings() PagesProjectDeploymentConfigsProductionVectorizeBindingsMap
	VectorizeBindingsInput() interface{}
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
	PutAiBindings(value interface{})
	PutAnalyticsEngineDatasets(value interface{})
	PutBrowsers(value interface{})
	PutD1Databases(value interface{})
	PutDurableObjectNamespaces(value interface{})
	PutEnvVars(value interface{})
	PutHyperdriveBindings(value interface{})
	PutKvNamespaces(value interface{})
	PutMtlsCertificates(value interface{})
	PutPlacement(value *PagesProjectDeploymentConfigsProductionPlacement)
	PutQueueProducers(value interface{})
	PutR2Buckets(value interface{})
	PutServices(value interface{})
	PutVectorizeBindings(value interface{})
	ResetAiBindings()
	ResetAnalyticsEngineDatasets()
	ResetBrowsers()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetD1Databases()
	ResetDurableObjectNamespaces()
	ResetEnvVars()
	ResetHyperdriveBindings()
	ResetKvNamespaces()
	ResetMtlsCertificates()
	ResetPlacement()
	ResetQueueProducers()
	ResetR2Buckets()
	ResetServices()
	ResetVectorizeBindings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PagesProjectDeploymentConfigsProductionOutputReference
type jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) AiBindings() PagesProjectDeploymentConfigsProductionAiBindingsMap {
	var returns PagesProjectDeploymentConfigsProductionAiBindingsMap
	_jsii_.Get(
		j,
		"aiBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) AiBindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aiBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) AnalyticsEngineDatasets() PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsMap {
	var returns PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsMap
	_jsii_.Get(
		j,
		"analyticsEngineDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) AnalyticsEngineDatasetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsEngineDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) Browsers() PagesProjectDeploymentConfigsProductionBrowsersMap {
	var returns PagesProjectDeploymentConfigsProductionBrowsersMap
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) BrowsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) D1Databases() PagesProjectDeploymentConfigsProductionD1DatabasesMap {
	var returns PagesProjectDeploymentConfigsProductionD1DatabasesMap
	_jsii_.Get(
		j,
		"d1Databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) D1DatabasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"d1DatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) DurableObjectNamespaces() PagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap {
	var returns PagesProjectDeploymentConfigsProductionDurableObjectNamespacesMap
	_jsii_.Get(
		j,
		"durableObjectNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) DurableObjectNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"durableObjectNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) EnvVars() PagesProjectDeploymentConfigsProductionEnvVarsMap {
	var returns PagesProjectDeploymentConfigsProductionEnvVarsMap
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) EnvVarsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) HyperdriveBindings() PagesProjectDeploymentConfigsProductionHyperdriveBindingsMap {
	var returns PagesProjectDeploymentConfigsProductionHyperdriveBindingsMap
	_jsii_.Get(
		j,
		"hyperdriveBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) HyperdriveBindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperdriveBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) KvNamespaces() PagesProjectDeploymentConfigsProductionKvNamespacesMap {
	var returns PagesProjectDeploymentConfigsProductionKvNamespacesMap
	_jsii_.Get(
		j,
		"kvNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) KvNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kvNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) MtlsCertificates() PagesProjectDeploymentConfigsProductionMtlsCertificatesMap {
	var returns PagesProjectDeploymentConfigsProductionMtlsCertificatesMap
	_jsii_.Get(
		j,
		"mtlsCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) MtlsCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) Placement() PagesProjectDeploymentConfigsProductionPlacementOutputReference {
	var returns PagesProjectDeploymentConfigsProductionPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) QueueProducers() PagesProjectDeploymentConfigsProductionQueueProducersMap {
	var returns PagesProjectDeploymentConfigsProductionQueueProducersMap
	_jsii_.Get(
		j,
		"queueProducers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) QueueProducersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueProducersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) R2Buckets() PagesProjectDeploymentConfigsProductionR2BucketsMap {
	var returns PagesProjectDeploymentConfigsProductionR2BucketsMap
	_jsii_.Get(
		j,
		"r2Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) R2BucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"r2BucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) Services() PagesProjectDeploymentConfigsProductionServicesMap {
	var returns PagesProjectDeploymentConfigsProductionServicesMap
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) VectorizeBindings() PagesProjectDeploymentConfigsProductionVectorizeBindingsMap {
	var returns PagesProjectDeploymentConfigsProductionVectorizeBindingsMap
	_jsii_.Get(
		j,
		"vectorizeBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) VectorizeBindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorizeBindingsInput",
		&returns,
	)
	return returns
}


func NewPagesProjectDeploymentConfigsProductionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectDeploymentConfigsProductionOutputReference {
	_init_.Initialize()

	if err := validateNewPagesProjectDeploymentConfigsProductionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsProductionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectDeploymentConfigsProductionOutputReference_Override(p PagesProjectDeploymentConfigsProductionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsProductionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutAiBindings(value interface{}) {
	if err := p.validatePutAiBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAiBindings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutAnalyticsEngineDatasets(value interface{}) {
	if err := p.validatePutAnalyticsEngineDatasetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAnalyticsEngineDatasets",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutBrowsers(value interface{}) {
	if err := p.validatePutBrowsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBrowsers",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutD1Databases(value interface{}) {
	if err := p.validatePutD1DatabasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putD1Databases",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutDurableObjectNamespaces(value interface{}) {
	if err := p.validatePutDurableObjectNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDurableObjectNamespaces",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutEnvVars(value interface{}) {
	if err := p.validatePutEnvVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvVars",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutHyperdriveBindings(value interface{}) {
	if err := p.validatePutHyperdriveBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHyperdriveBindings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutKvNamespaces(value interface{}) {
	if err := p.validatePutKvNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKvNamespaces",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutMtlsCertificates(value interface{}) {
	if err := p.validatePutMtlsCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMtlsCertificates",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutPlacement(value *PagesProjectDeploymentConfigsProductionPlacement) {
	if err := p.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPlacement",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutQueueProducers(value interface{}) {
	if err := p.validatePutQueueProducersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueueProducers",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutR2Buckets(value interface{}) {
	if err := p.validatePutR2BucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putR2Buckets",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutServices(value interface{}) {
	if err := p.validatePutServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServices",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) PutVectorizeBindings(value interface{}) {
	if err := p.validatePutVectorizeBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVectorizeBindings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetAiBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetAiBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetAnalyticsEngineDatasets() {
	_jsii_.InvokeVoid(
		p,
		"resetAnalyticsEngineDatasets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetBrowsers() {
	_jsii_.InvokeVoid(
		p,
		"resetBrowsers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetD1Databases() {
	_jsii_.InvokeVoid(
		p,
		"resetD1Databases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetDurableObjectNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetDurableObjectNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetEnvVars() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvVars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetHyperdriveBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetHyperdriveBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetKvNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetKvNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetMtlsCertificates() {
	_jsii_.InvokeVoid(
		p,
		"resetMtlsCertificates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		p,
		"resetPlacement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetQueueProducers() {
	_jsii_.InvokeVoid(
		p,
		"resetQueueProducers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetR2Buckets() {
	_jsii_.InvokeVoid(
		p,
		"resetR2Buckets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		p,
		"resetServices",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ResetVectorizeBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetVectorizeBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PagesProjectDeploymentConfigsProductionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


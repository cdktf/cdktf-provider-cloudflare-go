// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/workersscript/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersScriptBindingsOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	ClassName() *string
	SetClassName(val *string)
	ClassNameInput() *string
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
	Dataset() *string
	SetDataset(val *string)
	DatasetInput() *string
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Json() *string
	SetJson(val *string)
	JsonInput() *string
	KeyBase64() *string
	SetKeyBase64(val *string)
	KeyBase64Input() *string
	KeyJwk() *string
	SetKeyJwk(val *string)
	KeyJwkInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceId() *string
	SetNamespaceId(val *string)
	NamespaceIdInput() *string
	NamespaceInput() *string
	Outbound() WorkersScriptBindingsOutboundOutputReference
	OutboundInput() interface{}
	Pipeline() *string
	SetPipeline(val *string)
	PipelineInput() *string
	QueueName() *string
	SetQueueName(val *string)
	QueueNameInput() *string
	ScriptName() *string
	SetScriptName(val *string)
	ScriptNameInput() *string
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	StoreId() *string
	SetStoreId(val *string)
	StoreIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Usages() *[]*string
	SetUsages(val *[]*string)
	UsagesInput() *[]*string
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
	PutOutbound(value *WorkersScriptBindingsOutbound)
	ResetAlgorithm()
	ResetBucketName()
	ResetCertificateId()
	ResetClassName()
	ResetDataset()
	ResetEnvironment()
	ResetFormat()
	ResetId()
	ResetIndexName()
	ResetJson()
	ResetKeyBase64()
	ResetKeyJwk()
	ResetNamespace()
	ResetNamespaceId()
	ResetOutbound()
	ResetPipeline()
	ResetQueueName()
	ResetScriptName()
	ResetSecretName()
	ResetService()
	ResetStoreId()
	ResetText()
	ResetUsages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkersScriptBindingsOutputReference
type jsiiProxy_WorkersScriptBindingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Json() *string {
	var returns *string
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) JsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) KeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) KeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) KeyJwk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyJwk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) KeyJwkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyJwkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) NamespaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Outbound() WorkersScriptBindingsOutboundOutputReference {
	var returns WorkersScriptBindingsOutboundOutputReference
	_jsii_.Get(
		j,
		"outbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) OutboundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Pipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) PipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) QueueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ScriptName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ScriptNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) StoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) Usages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference) UsagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usagesInput",
		&returns,
	)
	return returns
}


func NewWorkersScriptBindingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkersScriptBindingsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkersScriptBindingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScriptBindingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptBindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkersScriptBindingsOutputReference_Override(w WorkersScriptBindingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScriptBindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetClassName(val *string) {
	if err := j.validateSetClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"className",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetJson(val *string) {
	if err := j.validateSetJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetKeyBase64(val *string) {
	if err := j.validateSetKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBase64",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetKeyJwk(val *string) {
	if err := j.validateSetKeyJwkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyJwk",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetNamespaceId(val *string) {
	if err := j.validateSetNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetPipeline(val *string) {
	if err := j.validateSetPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipeline",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetQueueName(val *string) {
	if err := j.validateSetQueueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetScriptName(val *string) {
	if err := j.validateSetScriptNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptName",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetSecretName(val *string) {
	if err := j.validateSetSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetStoreId(val *string) {
	if err := j.validateSetStoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeId",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_WorkersScriptBindingsOutputReference)SetUsages(val *[]*string) {
	if err := j.validateSetUsagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usages",
		val,
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) PutOutbound(value *WorkersScriptBindingsOutbound) {
	if err := w.validatePutOutboundParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOutbound",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		w,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		w,
		"resetBucketName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetCertificateId() {
	_jsii_.InvokeVoid(
		w,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetClassName() {
	_jsii_.InvokeVoid(
		w,
		"resetClassName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		w,
		"resetDataset",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		w,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		w,
		"resetFormat",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		w,
		"resetIndexName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		w,
		"resetJson",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetKeyBase64() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyBase64",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetKeyJwk() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyJwk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetNamespaceId() {
	_jsii_.InvokeVoid(
		w,
		"resetNamespaceId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetOutbound() {
	_jsii_.InvokeVoid(
		w,
		"resetOutbound",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetPipeline() {
	_jsii_.InvokeVoid(
		w,
		"resetPipeline",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetQueueName() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetScriptName() {
	_jsii_.InvokeVoid(
		w,
		"resetScriptName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetSecretName() {
	_jsii_.InvokeVoid(
		w,
		"resetSecretName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		w,
		"resetService",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetStoreId() {
	_jsii_.InvokeVoid(
		w,
		"resetStoreId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		w,
		"resetText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ResetUsages() {
	_jsii_.InvokeVoid(
		w,
		"resetUsages",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScriptBindingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


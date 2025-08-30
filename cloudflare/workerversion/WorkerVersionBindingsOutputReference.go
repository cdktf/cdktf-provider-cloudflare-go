// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/workerversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerVersionBindingsOutputReference interface {
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
	Outbound() WorkerVersionBindingsOutboundOutputReference
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
	WorkflowName() *string
	SetWorkflowName(val *string)
	WorkflowNameInput() *string
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
	PutOutbound(value *WorkerVersionBindingsOutbound)
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
	ResetWorkflowName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkerVersionBindingsOutputReference
type jsiiProxy_WorkerVersionBindingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Json() *string {
	var returns *string
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) JsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) KeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) KeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) KeyJwk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyJwk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) KeyJwkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyJwkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) NamespaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Outbound() WorkerVersionBindingsOutboundOutputReference {
	var returns WorkerVersionBindingsOutboundOutputReference
	_jsii_.Get(
		j,
		"outbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) OutboundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Pipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) PipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) QueueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ScriptName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ScriptNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) StoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) Usages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) UsagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference) WorkflowNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowNameInput",
		&returns,
	)
	return returns
}


func NewWorkerVersionBindingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkerVersionBindingsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkerVersionBindingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkerVersionBindingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerVersion.WorkerVersionBindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkerVersionBindingsOutputReference_Override(w WorkerVersionBindingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerVersion.WorkerVersionBindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetClassName(val *string) {
	if err := j.validateSetClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"className",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetJson(val *string) {
	if err := j.validateSetJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetKeyBase64(val *string) {
	if err := j.validateSetKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBase64",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetKeyJwk(val *string) {
	if err := j.validateSetKeyJwkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyJwk",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetNamespaceId(val *string) {
	if err := j.validateSetNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetPipeline(val *string) {
	if err := j.validateSetPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipeline",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetQueueName(val *string) {
	if err := j.validateSetQueueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetScriptName(val *string) {
	if err := j.validateSetScriptNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptName",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetSecretName(val *string) {
	if err := j.validateSetSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetStoreId(val *string) {
	if err := j.validateSetStoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeId",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetUsages(val *[]*string) {
	if err := j.validateSetUsagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usages",
		val,
	)
}

func (j *jsiiProxy_WorkerVersionBindingsOutputReference)SetWorkflowName(val *string) {
	if err := j.validateSetWorkflowNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowName",
		val,
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) PutOutbound(value *WorkerVersionBindingsOutbound) {
	if err := w.validatePutOutboundParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOutbound",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		w,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		w,
		"resetBucketName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetCertificateId() {
	_jsii_.InvokeVoid(
		w,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetClassName() {
	_jsii_.InvokeVoid(
		w,
		"resetClassName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetDataset() {
	_jsii_.InvokeVoid(
		w,
		"resetDataset",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		w,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		w,
		"resetFormat",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		w,
		"resetIndexName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		w,
		"resetJson",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetKeyBase64() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyBase64",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetKeyJwk() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyJwk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetNamespaceId() {
	_jsii_.InvokeVoid(
		w,
		"resetNamespaceId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetOutbound() {
	_jsii_.InvokeVoid(
		w,
		"resetOutbound",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetPipeline() {
	_jsii_.InvokeVoid(
		w,
		"resetPipeline",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetQueueName() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetScriptName() {
	_jsii_.InvokeVoid(
		w,
		"resetScriptName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetSecretName() {
	_jsii_.InvokeVoid(
		w,
		"resetSecretName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		w,
		"resetService",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetStoreId() {
	_jsii_.InvokeVoid(
		w,
		"resetStoreId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		w,
		"resetText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetUsages() {
	_jsii_.InvokeVoid(
		w,
		"resetUsages",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ResetWorkflowName() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkflowName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkerVersionBindingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


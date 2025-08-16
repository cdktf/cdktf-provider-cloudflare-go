// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamliveinput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/streamliveinput/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/stream_live_input cloudflare_stream_live_input}.
type StreamLiveInput interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Created() *string
	DefaultCreator() *string
	SetDefaultCreator(val *string)
	DefaultCreatorInput() *string
	DeleteRecordingAfterDays() *float64
	SetDeleteRecordingAfterDays(val *float64)
	DeleteRecordingAfterDaysInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LiveInputIdentifier() *string
	SetLiveInputIdentifier(val *string)
	LiveInputIdentifierInput() *string
	Meta() *string
	SetMeta(val *string)
	MetaInput() *string
	Modified() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Recording() StreamLiveInputRecordingOutputReference
	RecordingInput() interface{}
	Rtmps() StreamLiveInputRtmpsOutputReference
	RtmpsPlayback() StreamLiveInputRtmpsPlaybackOutputReference
	Srt() StreamLiveInputSrtOutputReference
	SrtPlayback() StreamLiveInputSrtPlaybackOutputReference
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uid() *string
	WebRtc() StreamLiveInputWebRtcOutputReference
	WebRtcPlayback() StreamLiveInputWebRtcPlaybackOutputReference
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutRecording(value *StreamLiveInputRecording)
	ResetDefaultCreator()
	ResetDeleteRecordingAfterDays()
	ResetLiveInputIdentifier()
	ResetMeta()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecording()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StreamLiveInput
type jsiiProxy_StreamLiveInput struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StreamLiveInput) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) DefaultCreator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCreator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) DefaultCreatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCreatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) DeleteRecordingAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteRecordingAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) DeleteRecordingAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteRecordingAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) LiveInputIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveInputIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) LiveInputIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveInputIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Meta() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) MetaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Modified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Recording() StreamLiveInputRecordingOutputReference {
	var returns StreamLiveInputRecordingOutputReference
	_jsii_.Get(
		j,
		"recording",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) RecordingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Rtmps() StreamLiveInputRtmpsOutputReference {
	var returns StreamLiveInputRtmpsOutputReference
	_jsii_.Get(
		j,
		"rtmps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) RtmpsPlayback() StreamLiveInputRtmpsPlaybackOutputReference {
	var returns StreamLiveInputRtmpsPlaybackOutputReference
	_jsii_.Get(
		j,
		"rtmpsPlayback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Srt() StreamLiveInputSrtOutputReference {
	var returns StreamLiveInputSrtOutputReference
	_jsii_.Get(
		j,
		"srt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) SrtPlayback() StreamLiveInputSrtPlaybackOutputReference {
	var returns StreamLiveInputSrtPlaybackOutputReference
	_jsii_.Get(
		j,
		"srtPlayback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) WebRtc() StreamLiveInputWebRtcOutputReference {
	var returns StreamLiveInputWebRtcOutputReference
	_jsii_.Get(
		j,
		"webRtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamLiveInput) WebRtcPlayback() StreamLiveInputWebRtcPlaybackOutputReference {
	var returns StreamLiveInputWebRtcPlaybackOutputReference
	_jsii_.Get(
		j,
		"webRtcPlayback",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/stream_live_input cloudflare_stream_live_input} Resource.
func NewStreamLiveInput(scope constructs.Construct, id *string, config *StreamLiveInputConfig) StreamLiveInput {
	_init_.Initialize()

	if err := validateNewStreamLiveInputParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamLiveInput{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/stream_live_input cloudflare_stream_live_input} Resource.
func NewStreamLiveInput_Override(s StreamLiveInput, scope constructs.Construct, id *string, config *StreamLiveInputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetDefaultCreator(val *string) {
	if err := j.validateSetDefaultCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCreator",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetDeleteRecordingAfterDays(val *float64) {
	if err := j.validateSetDeleteRecordingAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRecordingAfterDays",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetLiveInputIdentifier(val *string) {
	if err := j.validateSetLiveInputIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveInputIdentifier",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetMeta(val *string) {
	if err := j.validateSetMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meta",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StreamLiveInput)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a StreamLiveInput resource upon running "cdktf plan <stack-name>".
func StreamLiveInput_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStreamLiveInput_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StreamLiveInput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamLiveInput_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamLiveInput_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamLiveInput_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamLiveInput_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamLiveInput_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StreamLiveInput_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.streamLiveInput.StreamLiveInput",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StreamLiveInput) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StreamLiveInput) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StreamLiveInput) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StreamLiveInput) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamLiveInput) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StreamLiveInput) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamLiveInput) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StreamLiveInput) PutRecording(value *StreamLiveInputRecording) {
	if err := s.validatePutRecordingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRecording",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetDefaultCreator() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultCreator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetDeleteRecordingAfterDays() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteRecordingAfterDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetLiveInputIdentifier() {
	_jsii_.InvokeVoid(
		s,
		"resetLiveInputIdentifier",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetMeta() {
	_jsii_.InvokeVoid(
		s,
		"resetMeta",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) ResetRecording() {
	_jsii_.InvokeVoid(
		s,
		"resetRecording",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamLiveInput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamLiveInput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/stream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/stream cloudflare_stream}.
type Stream interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	Creator() *string
	SetCreator(val *string)
	CreatorInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Duration() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	Input() StreamInputOutputReference
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LiveInput() *string
	MaxDurationSeconds() *float64
	SetMaxDurationSeconds(val *float64)
	MaxDurationSecondsInput() *float64
	Meta() *string
	SetMeta(val *string)
	MetaInput() *string
	Modified() *string
	// The tree node.
	Node() constructs.Node
	Playback() StreamPlaybackOutputReference
	Preview() *string
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
	ReadyToStream() cdktf.IResolvable
	ReadyToStreamAt() *string
	RequireSignedUrls() interface{}
	SetRequireSignedUrls(val interface{})
	RequireSignedUrlsInput() interface{}
	ScheduledDeletion() *string
	SetScheduledDeletion(val *string)
	ScheduledDeletionInput() *string
	Size() *float64
	Status() StreamStatusOutputReference
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Thumbnail() *string
	ThumbnailTimestampPct() *float64
	SetThumbnailTimestampPct(val *float64)
	ThumbnailTimestampPctInput() *float64
	Uid() *string
	Uploaded() *string
	UploadExpiry() *string
	SetUploadExpiry(val *string)
	UploadExpiryInput() *string
	Watermark() StreamWatermarkOutputReference
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
	ResetAllowedOrigins()
	ResetCreator()
	ResetIdentifier()
	ResetMaxDurationSeconds()
	ResetMeta()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequireSignedUrls()
	ResetScheduledDeletion()
	ResetThumbnailTimestampPct()
	ResetUploadExpiry()
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

// The jsii proxy struct for Stream
type jsiiProxy_Stream struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Stream) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) CreatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Input() StreamInputOutputReference {
	var returns StreamInputOutputReference
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) LiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) MaxDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) MaxDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Meta() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) MetaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Modified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Playback() StreamPlaybackOutputReference {
	var returns StreamPlaybackOutputReference
	_jsii_.Get(
		j,
		"playback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Preview() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ReadyToStream() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readyToStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ReadyToStreamAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readyToStreamAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) RequireSignedUrls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) RequireSignedUrlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ScheduledDeletion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ScheduledDeletionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Status() StreamStatusOutputReference {
	var returns StreamStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Thumbnail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ThumbnailTimestampPct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thumbnailTimestampPct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) ThumbnailTimestampPctInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thumbnailTimestampPctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Uploaded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploaded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) UploadExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) UploadExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stream) Watermark() StreamWatermarkOutputReference {
	var returns StreamWatermarkOutputReference
	_jsii_.Get(
		j,
		"watermark",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/stream cloudflare_stream} Resource.
func NewStream(scope constructs.Construct, id *string, config *StreamConfig) Stream {
	_init_.Initialize()

	if err := validateNewStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stream{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.stream.Stream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/stream cloudflare_stream} Resource.
func NewStream_Override(s Stream, scope constructs.Construct, id *string, config *StreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.stream.Stream",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Stream)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_Stream)SetAllowedOrigins(val *[]*string) {
	if err := j.validateSetAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_Stream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Stream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Stream)SetCreator(val *string) {
	if err := j.validateSetCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creator",
		val,
	)
}

func (j *jsiiProxy_Stream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Stream)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Stream)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_Stream)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Stream)SetMaxDurationSeconds(val *float64) {
	if err := j.validateSetMaxDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_Stream)SetMeta(val *string) {
	if err := j.validateSetMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meta",
		val,
	)
}

func (j *jsiiProxy_Stream)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Stream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Stream)SetRequireSignedUrls(val interface{}) {
	if err := j.validateSetRequireSignedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSignedUrls",
		val,
	)
}

func (j *jsiiProxy_Stream)SetScheduledDeletion(val *string) {
	if err := j.validateSetScheduledDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledDeletion",
		val,
	)
}

func (j *jsiiProxy_Stream)SetThumbnailTimestampPct(val *float64) {
	if err := j.validateSetThumbnailTimestampPctParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thumbnailTimestampPct",
		val,
	)
}

func (j *jsiiProxy_Stream)SetUploadExpiry(val *string) {
	if err := j.validateSetUploadExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadExpiry",
		val,
	)
}

// Generates CDKTF code for importing a Stream resource upon running "cdktf plan <stack-name>".
func Stream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.stream.Stream",
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
func Stream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.stream.Stream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Stream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.stream.Stream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Stream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.stream.Stream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Stream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.stream.Stream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Stream) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Stream) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Stream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_Stream) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_Stream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_Stream) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_Stream) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_Stream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_Stream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_Stream) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_Stream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_Stream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Stream) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_Stream) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Stream) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Stream) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Stream) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Stream) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetCreator() {
	_jsii_.InvokeVoid(
		s,
		"resetCreator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetIdentifier() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentifier",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetMaxDurationSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDurationSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetMeta() {
	_jsii_.InvokeVoid(
		s,
		"resetMeta",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetRequireSignedUrls() {
	_jsii_.InvokeVoid(
		s,
		"resetRequireSignedUrls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetScheduledDeletion() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduledDeletion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetThumbnailTimestampPct() {
	_jsii_.InvokeVoid(
		s,
		"resetThumbnailTimestampPct",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) ResetUploadExpiry() {
	_jsii_.InvokeVoid(
		s,
		"resetUploadExpiry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/dnsrecord/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordDataOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *float64
	SetAlgorithm(val *float64)
	AlgorithmInput() *float64
	Altitude() *float64
	SetAltitude(val *float64)
	AltitudeInput() *float64
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
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
	Digest() *string
	SetDigest(val *string)
	DigestInput() *string
	DigestType() *float64
	SetDigestType(val *float64)
	DigestTypeInput() *float64
	Fingerprint() *string
	SetFingerprint(val *string)
	FingerprintInput() *string
	Flags() *map[string]interface{}
	SetFlags(val *map[string]interface{})
	FlagsInput() *map[string]interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyTag() *float64
	SetKeyTag(val *float64)
	KeyTagInput() *float64
	LatDegrees() *float64
	SetLatDegrees(val *float64)
	LatDegreesInput() *float64
	LatDirection() *string
	SetLatDirection(val *string)
	LatDirectionInput() *string
	LatMinutes() *float64
	SetLatMinutes(val *float64)
	LatMinutesInput() *float64
	LatSeconds() *float64
	SetLatSeconds(val *float64)
	LatSecondsInput() *float64
	LongDegrees() *float64
	SetLongDegrees(val *float64)
	LongDegreesInput() *float64
	LongDirection() *string
	SetLongDirection(val *string)
	LongDirectionInput() *string
	LongMinutes() *float64
	SetLongMinutes(val *float64)
	LongMinutesInput() *float64
	LongSeconds() *float64
	SetLongSeconds(val *float64)
	LongSecondsInput() *float64
	MatchingType() *float64
	SetMatchingType(val *float64)
	MatchingTypeInput() *float64
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrecisionHorz() *float64
	SetPrecisionHorz(val *float64)
	PrecisionHorzInput() *float64
	PrecisionVert() *float64
	SetPrecisionVert(val *float64)
	PrecisionVertInput() *float64
	Preference() *float64
	SetPreference(val *float64)
	PreferenceInput() *float64
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Protocol() *float64
	SetProtocol(val *float64)
	ProtocolInput() *float64
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
	Replacement() *string
	SetReplacement(val *string)
	ReplacementInput() *string
	Selector() *float64
	SetSelector(val *float64)
	SelectorInput() *float64
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *float64
	SetType(val *float64)
	TypeInput() *float64
	Usage() *float64
	SetUsage(val *float64)
	UsageInput() *float64
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetAlgorithm()
	ResetAltitude()
	ResetCertificate()
	ResetDigest()
	ResetDigestType()
	ResetFingerprint()
	ResetFlags()
	ResetKeyTag()
	ResetLatDegrees()
	ResetLatDirection()
	ResetLatMinutes()
	ResetLatSeconds()
	ResetLongDegrees()
	ResetLongDirection()
	ResetLongMinutes()
	ResetLongSeconds()
	ResetMatchingType()
	ResetOrder()
	ResetPort()
	ResetPrecisionHorz()
	ResetPrecisionVert()
	ResetPreference()
	ResetPriority()
	ResetProtocol()
	ResetPublicKey()
	ResetRegex()
	ResetReplacement()
	ResetSelector()
	ResetService()
	ResetSize()
	ResetTag()
	ResetTarget()
	ResetType()
	ResetUsage()
	ResetValue()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsRecordDataOutputReference
type jsiiProxy_DnsRecordDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Algorithm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) AlgorithmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Altitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"altitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) AltitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"altitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Digest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) DigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) DigestType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) DigestTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) FingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Flags() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) FlagsInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) KeyTag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) KeyTagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatDegrees() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latDegrees",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatDegreesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latDegreesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LatSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongDegrees() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longDegrees",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongDegreesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longDegreesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) LongSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) MatchingType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) MatchingTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"matchingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PrecisionHorz() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionHorz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PrecisionHorzInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionHorzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PrecisionVert() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionVert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PrecisionVertInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionVertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Preference() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PreferenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Protocol() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ProtocolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Replacement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ReplacementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Selector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) SelectorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) TypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Usage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) UsageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordDataOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewDnsRecordDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsRecordDataOutputReference {
	_init_.Initialize()

	if err := validateNewDnsRecordDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsRecordDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsRecord.DnsRecordDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsRecordDataOutputReference_Override(d DnsRecordDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dnsRecord.DnsRecordDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetAlgorithm(val *float64) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetAltitude(val *float64) {
	if err := j.validateSetAltitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"altitude",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetDigest(val *string) {
	if err := j.validateSetDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digest",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetDigestType(val *float64) {
	if err := j.validateSetDigestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digestType",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetFingerprint(val *string) {
	if err := j.validateSetFingerprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fingerprint",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetFlags(val *map[string]interface{}) {
	if err := j.validateSetFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flags",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetKeyTag(val *float64) {
	if err := j.validateSetKeyTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyTag",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLatDegrees(val *float64) {
	if err := j.validateSetLatDegreesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latDegrees",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLatDirection(val *string) {
	if err := j.validateSetLatDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latDirection",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLatMinutes(val *float64) {
	if err := j.validateSetLatMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latMinutes",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLatSeconds(val *float64) {
	if err := j.validateSetLatSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latSeconds",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLongDegrees(val *float64) {
	if err := j.validateSetLongDegreesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longDegrees",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLongDirection(val *string) {
	if err := j.validateSetLongDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longDirection",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLongMinutes(val *float64) {
	if err := j.validateSetLongMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longMinutes",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetLongSeconds(val *float64) {
	if err := j.validateSetLongSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longSeconds",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetMatchingType(val *float64) {
	if err := j.validateSetMatchingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingType",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPrecisionHorz(val *float64) {
	if err := j.validateSetPrecisionHorzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precisionHorz",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPrecisionVert(val *float64) {
	if err := j.validateSetPrecisionVertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precisionVert",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPreference(val *float64) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetProtocol(val *float64) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetReplacement(val *string) {
	if err := j.validateSetReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacement",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetSelector(val *float64) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetType(val *float64) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetUsage(val *float64) {
	if err := j.validateSetUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usage",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (j *jsiiProxy_DnsRecordDataOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetAltitude() {
	_jsii_.InvokeVoid(
		d,
		"resetAltitude",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetDigest() {
	_jsii_.InvokeVoid(
		d,
		"resetDigest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetDigestType() {
	_jsii_.InvokeVoid(
		d,
		"resetDigestType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetFingerprint() {
	_jsii_.InvokeVoid(
		d,
		"resetFingerprint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetFlags() {
	_jsii_.InvokeVoid(
		d,
		"resetFlags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetKeyTag() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLatDegrees() {
	_jsii_.InvokeVoid(
		d,
		"resetLatDegrees",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLatDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetLatDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLatMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetLatMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLatSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetLatSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLongDegrees() {
	_jsii_.InvokeVoid(
		d,
		"resetLongDegrees",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLongDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetLongDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLongMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetLongMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetLongSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetLongSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetMatchingType() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchingType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPrecisionHorz() {
	_jsii_.InvokeVoid(
		d,
		"resetPrecisionHorz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPrecisionVert() {
	_jsii_.InvokeVoid(
		d,
		"resetPrecisionVert",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPreference() {
	_jsii_.InvokeVoid(
		d,
		"resetPreference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		d,
		"resetPriority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetPublicKey() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetReplacement() {
	_jsii_.InvokeVoid(
		d,
		"resetReplacement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		d,
		"resetService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		d,
		"resetSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		d,
		"resetTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetUsage() {
	_jsii_.InvokeVoid(
		d,
		"resetUsage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		d,
		"resetValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


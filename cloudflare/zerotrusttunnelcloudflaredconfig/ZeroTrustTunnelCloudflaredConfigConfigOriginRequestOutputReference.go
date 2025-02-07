// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/zerotrusttunnelcloudflaredconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference interface {
	cdktf.ComplexObject
	Access() ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference
	AccessInput() interface{}
	CaPool() *string
	SetCaPool(val *string)
	CaPoolInput() *string
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
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableChunkedEncoding() interface{}
	SetDisableChunkedEncoding(val interface{})
	DisableChunkedEncodingInput() interface{}
	// Experimental.
	Fqn() *string
	Http2Origin() interface{}
	SetHttp2Origin(val interface{})
	Http2OriginInput() interface{}
	HttpHostHeader() *string
	SetHttpHostHeader(val *string)
	HttpHostHeaderInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeepAliveConnections() *float64
	SetKeepAliveConnections(val *float64)
	KeepAliveConnectionsInput() *float64
	KeepAliveTimeout() *float64
	SetKeepAliveTimeout(val *float64)
	KeepAliveTimeoutInput() *float64
	NoHappyEyeballs() interface{}
	SetNoHappyEyeballs(val interface{})
	NoHappyEyeballsInput() interface{}
	NoTlsVerify() interface{}
	SetNoTlsVerify(val interface{})
	NoTlsVerifyInput() interface{}
	OriginServerName() *string
	SetOriginServerName(val *string)
	OriginServerNameInput() *string
	ProxyType() *string
	SetProxyType(val *string)
	ProxyTypeInput() *string
	TcpKeepAlive() *float64
	SetTcpKeepAlive(val *float64)
	TcpKeepAliveInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsTimeout() *float64
	SetTlsTimeout(val *float64)
	TlsTimeoutInput() *float64
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
	PutAccess(value *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccess)
	ResetAccess()
	ResetCaPool()
	ResetConnectTimeout()
	ResetDisableChunkedEncoding()
	ResetHttp2Origin()
	ResetHttpHostHeader()
	ResetKeepAliveConnections()
	ResetKeepAliveTimeout()
	ResetNoHappyEyeballs()
	ResetNoTlsVerify()
	ResetOriginServerName()
	ResetProxyType()
	ResetTcpKeepAlive()
	ResetTlsTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference
type jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Access() ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference {
	var returns ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) AccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) CaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) CaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) DisableChunkedEncoding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) DisableChunkedEncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Http2Origin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Http2OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2OriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) HttpHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) HttpHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoHappyEyeballs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoHappyEyeballsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) OriginServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) OriginServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ProxyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TcpKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TcpKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TlsTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TlsTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsTimeoutInput",
		&returns,
	)
	return returns
}


func NewZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference_Override(z ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetCaPool(val *string) {
	if err := j.validateSetCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPool",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetDisableChunkedEncoding(val interface{}) {
	if err := j.validateSetDisableChunkedEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableChunkedEncoding",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetHttp2Origin(val interface{}) {
	if err := j.validateSetHttp2OriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Origin",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetHttpHostHeader(val *string) {
	if err := j.validateSetHttpHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHostHeader",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetKeepAliveConnections(val *float64) {
	if err := j.validateSetKeepAliveConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveConnections",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetKeepAliveTimeout(val *float64) {
	if err := j.validateSetKeepAliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetNoHappyEyeballs(val interface{}) {
	if err := j.validateSetNoHappyEyeballsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHappyEyeballs",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetNoTlsVerify(val interface{}) {
	if err := j.validateSetNoTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTlsVerify",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetOriginServerName(val *string) {
	if err := j.validateSetOriginServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originServerName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetProxyType(val *string) {
	if err := j.validateSetProxyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTcpKeepAlive(val *float64) {
	if err := j.validateSetTcpKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpKeepAlive",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTlsTimeout(val *float64) {
	if err := j.validateSetTlsTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsTimeout",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) PutAccess(value *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccess) {
	if err := z.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAccess",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		z,
		"resetAccess",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetCaPool() {
	_jsii_.InvokeVoid(
		z,
		"resetCaPool",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetDisableChunkedEncoding() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableChunkedEncoding",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetHttp2Origin() {
	_jsii_.InvokeVoid(
		z,
		"resetHttp2Origin",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetHttpHostHeader() {
	_jsii_.InvokeVoid(
		z,
		"resetHttpHostHeader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetKeepAliveConnections() {
	_jsii_.InvokeVoid(
		z,
		"resetKeepAliveConnections",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetNoHappyEyeballs() {
	_jsii_.InvokeVoid(
		z,
		"resetNoHappyEyeballs",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetNoTlsVerify() {
	_jsii_.InvokeVoid(
		z,
		"resetNoTlsVerify",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetOriginServerName() {
	_jsii_.InvokeVoid(
		z,
		"resetOriginServerName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetProxyType() {
	_jsii_.InvokeVoid(
		z,
		"resetProxyType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetTcpKeepAlive() {
	_jsii_.InvokeVoid(
		z,
		"resetTcpKeepAlive",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ResetTlsTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetTlsTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


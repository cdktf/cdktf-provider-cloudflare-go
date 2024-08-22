// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrusttunnelcloudflaredconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference interface {
	cdktf.ComplexObject
	Access() ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference
	AccessInput() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess
	BastionMode() interface{}
	SetBastionMode(val interface{})
	BastionModeInput() interface{}
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
	ConnectTimeout() *string
	SetConnectTimeout(val *string)
	ConnectTimeoutInput() *string
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
	InternalValue() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest
	SetInternalValue(val *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest)
	IpRules() ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestIpRulesList
	IpRulesInput() interface{}
	KeepAliveConnections() *float64
	SetKeepAliveConnections(val *float64)
	KeepAliveConnectionsInput() *float64
	KeepAliveTimeout() *string
	SetKeepAliveTimeout(val *string)
	KeepAliveTimeoutInput() *string
	NoHappyEyeballs() interface{}
	SetNoHappyEyeballs(val interface{})
	NoHappyEyeballsInput() interface{}
	NoTlsVerify() interface{}
	SetNoTlsVerify(val interface{})
	NoTlsVerifyInput() interface{}
	OriginServerName() *string
	SetOriginServerName(val *string)
	OriginServerNameInput() *string
	ProxyAddress() *string
	SetProxyAddress(val *string)
	ProxyAddressInput() *string
	ProxyPort() *float64
	SetProxyPort(val *float64)
	ProxyPortInput() *float64
	ProxyType() *string
	SetProxyType(val *string)
	ProxyTypeInput() *string
	TcpKeepAlive() *string
	SetTcpKeepAlive(val *string)
	TcpKeepAliveInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsTimeout() *string
	SetTlsTimeout(val *string)
	TlsTimeoutInput() *string
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
	PutAccess(value *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess)
	PutIpRules(value interface{})
	ResetAccess()
	ResetBastionMode()
	ResetCaPool()
	ResetConnectTimeout()
	ResetDisableChunkedEncoding()
	ResetHttp2Origin()
	ResetHttpHostHeader()
	ResetIpRules()
	ResetKeepAliveConnections()
	ResetKeepAliveTimeout()
	ResetNoHappyEyeballs()
	ResetNoTlsVerify()
	ResetOriginServerName()
	ResetProxyAddress()
	ResetProxyPort()
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

// The jsii proxy struct for ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference
type jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) Access() ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference {
	var returns ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) AccessInput() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess {
	var returns *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) BastionMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) BastionModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) CaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) CaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ConnectTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ConnectTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) DisableChunkedEncoding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) DisableChunkedEncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) Http2Origin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) Http2OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2OriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) HttpHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) HttpHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) InternalValue() *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest {
	var returns *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) IpRules() ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestIpRulesList {
	var returns ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestIpRulesList
	_jsii_.Get(
		j,
		"ipRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) IpRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) NoHappyEyeballs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) NoHappyEyeballsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) NoTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) NoTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) OriginServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) OriginServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ProxyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TcpKeepAlive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TcpKeepAliveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TlsTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) TlsTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeoutInput",
		&returns,
	)
	return returns
}


func NewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference_Override(z ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustTunnelCloudflaredConfig.ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetBastionMode(val interface{}) {
	if err := j.validateSetBastionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bastionMode",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetCaPool(val *string) {
	if err := j.validateSetCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPool",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetConnectTimeout(val *string) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetDisableChunkedEncoding(val interface{}) {
	if err := j.validateSetDisableChunkedEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableChunkedEncoding",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetHttp2Origin(val interface{}) {
	if err := j.validateSetHttp2OriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Origin",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetHttpHostHeader(val *string) {
	if err := j.validateSetHttpHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHostHeader",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetInternalValue(val *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetKeepAliveConnections(val *float64) {
	if err := j.validateSetKeepAliveConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveConnections",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetKeepAliveTimeout(val *string) {
	if err := j.validateSetKeepAliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetNoHappyEyeballs(val interface{}) {
	if err := j.validateSetNoHappyEyeballsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHappyEyeballs",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetNoTlsVerify(val interface{}) {
	if err := j.validateSetNoTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTlsVerify",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetOriginServerName(val *string) {
	if err := j.validateSetOriginServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originServerName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetProxyAddress(val *string) {
	if err := j.validateSetProxyAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyAddress",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetProxyPort(val *float64) {
	if err := j.validateSetProxyPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyPort",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetProxyType(val *string) {
	if err := j.validateSetProxyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyType",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetTcpKeepAlive(val *string) {
	if err := j.validateSetTcpKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpKeepAlive",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference)SetTlsTimeout(val *string) {
	if err := j.validateSetTlsTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsTimeout",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) PutAccess(value *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess) {
	if err := z.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAccess",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) PutIpRules(value interface{}) {
	if err := z.validatePutIpRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpRules",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		z,
		"resetAccess",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetBastionMode() {
	_jsii_.InvokeVoid(
		z,
		"resetBastionMode",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetCaPool() {
	_jsii_.InvokeVoid(
		z,
		"resetCaPool",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetDisableChunkedEncoding() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableChunkedEncoding",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetHttp2Origin() {
	_jsii_.InvokeVoid(
		z,
		"resetHttp2Origin",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetHttpHostHeader() {
	_jsii_.InvokeVoid(
		z,
		"resetHttpHostHeader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetIpRules() {
	_jsii_.InvokeVoid(
		z,
		"resetIpRules",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetKeepAliveConnections() {
	_jsii_.InvokeVoid(
		z,
		"resetKeepAliveConnections",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetNoHappyEyeballs() {
	_jsii_.InvokeVoid(
		z,
		"resetNoHappyEyeballs",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetNoTlsVerify() {
	_jsii_.InvokeVoid(
		z,
		"resetNoTlsVerify",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetOriginServerName() {
	_jsii_.InvokeVoid(
		z,
		"resetOriginServerName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyAddress() {
	_jsii_.InvokeVoid(
		z,
		"resetProxyAddress",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyPort() {
	_jsii_.InvokeVoid(
		z,
		"resetProxyPort",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyType() {
	_jsii_.InvokeVoid(
		z,
		"resetProxyType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetTcpKeepAlive() {
	_jsii_.InvokeVoid(
		z,
		"resetTcpKeepAlive",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ResetTlsTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetTlsTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


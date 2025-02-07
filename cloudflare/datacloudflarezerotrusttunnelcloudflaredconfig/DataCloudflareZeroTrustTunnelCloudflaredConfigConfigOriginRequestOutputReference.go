// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflaredconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarezerotrusttunnelcloudflaredconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference interface {
	cdktf.ComplexObject
	Access() DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference
	CaPool() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableChunkedEncoding() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	Http2Origin() cdktf.IResolvable
	HttpHostHeader() *string
	InternalValue() *DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequest
	SetInternalValue(val *DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequest)
	KeepAliveConnections() *float64
	KeepAliveTimeout() *float64
	NoHappyEyeballs() cdktf.IResolvable
	NoTlsVerify() cdktf.IResolvable
	OriginServerName() *string
	ProxyType() *string
	TcpKeepAlive() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsTimeout() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference
type jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Access() DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference {
	var returns DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) CaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) DisableChunkedEncoding() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableChunkedEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Http2Origin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"http2Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) HttpHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InternalValue() *DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequest {
	var returns *DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoHappyEyeballs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"noHappyEyeballs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) NoTlsVerify() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"noTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) OriginServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TcpKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) TlsTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsTimeout",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredConfig.DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference_Override(d DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredConfig.DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetInternalValue(val *DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredConfigConfigOriginRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


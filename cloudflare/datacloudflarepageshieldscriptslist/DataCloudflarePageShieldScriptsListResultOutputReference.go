// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldscriptslist

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/datacloudflarepageshieldscriptslist/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflarePageShieldScriptsListResultOutputReference interface {
	cdktf.ComplexObject
	AddedAt() *string
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
	CryptominingScore() *float64
	DataflowScore() *float64
	DomainReportedMalicious() cdktf.IResolvable
	FetchedAt() *string
	FirstPageUrl() *string
	FirstSeenAt() *string
	// Experimental.
	Fqn() *string
	Hash() *string
	Host() *string
	Id() *string
	InternalValue() *DataCloudflarePageShieldScriptsListResult
	SetInternalValue(val *DataCloudflarePageShieldScriptsListResult)
	JsIntegrityScore() *float64
	LastSeenAt() *string
	MagecartScore() *float64
	MaliciousDomainCategories() *[]*string
	MaliciousUrlCategories() *[]*string
	MalwareScore() *float64
	ObfuscationScore() *float64
	PageUrls() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	UrlContainsCdnCgiPath() cdktf.IResolvable
	UrlReportedMalicious() cdktf.IResolvable
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

// The jsii proxy struct for DataCloudflarePageShieldScriptsListResultOutputReference
type jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) AddedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) CryptominingScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cryptominingScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) DataflowScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataflowScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) DomainReportedMalicious() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"domainReportedMalicious",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) FetchedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fetchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) FirstPageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) FirstSeenAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstSeenAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Hash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) InternalValue() *DataCloudflarePageShieldScriptsListResult {
	var returns *DataCloudflarePageShieldScriptsListResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) JsIntegrityScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsIntegrityScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) LastSeenAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSeenAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) MagecartScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"magecartScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) MaliciousDomainCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"maliciousDomainCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) MaliciousUrlCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"maliciousUrlCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) MalwareScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"malwareScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) ObfuscationScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"obfuscationScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) PageUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pageUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) UrlContainsCdnCgiPath() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"urlContainsCdnCgiPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) UrlReportedMalicious() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"urlReportedMalicious",
		&returns,
	)
	return returns
}


func NewDataCloudflarePageShieldScriptsListResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflarePageShieldScriptsListResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflarePageShieldScriptsListResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsListResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflarePageShieldScriptsListResultOutputReference_Override(d DataCloudflarePageShieldScriptsListResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldScriptsList.DataCloudflarePageShieldScriptsListResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference)SetInternalValue(val *DataCloudflarePageShieldScriptsListResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflarePageShieldScriptsListResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


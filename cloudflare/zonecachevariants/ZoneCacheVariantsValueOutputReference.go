// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonecachevariants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v13/zonecachevariants/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneCacheVariantsValueOutputReference interface {
	cdktf.ComplexObject
	Avif() *[]*string
	SetAvif(val *[]*string)
	AvifInput() *[]*string
	Bmp() *[]*string
	SetBmp(val *[]*string)
	BmpInput() *[]*string
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
	// Experimental.
	Fqn() *string
	Gif() *[]*string
	SetGif(val *[]*string)
	GifInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jp2() *[]*string
	SetJp2(val *[]*string)
	Jp2Input() *[]*string
	Jpeg() *[]*string
	SetJpeg(val *[]*string)
	JpegInput() *[]*string
	Jpg() *[]*string
	SetJpg(val *[]*string)
	Jpg2() *[]*string
	SetJpg2(val *[]*string)
	Jpg2Input() *[]*string
	JpgInput() *[]*string
	Png() *[]*string
	SetPng(val *[]*string)
	PngInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tif() *[]*string
	SetTif(val *[]*string)
	Tiff() *[]*string
	SetTiff(val *[]*string)
	TiffInput() *[]*string
	TifInput() *[]*string
	Webp() *[]*string
	SetWebp(val *[]*string)
	WebpInput() *[]*string
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
	ResetAvif()
	ResetBmp()
	ResetGif()
	ResetJp2()
	ResetJpeg()
	ResetJpg()
	ResetJpg2()
	ResetPng()
	ResetTif()
	ResetTiff()
	ResetWebp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneCacheVariantsValueOutputReference
type jsiiProxy_ZoneCacheVariantsValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Avif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"avif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) AvifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"avifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Bmp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bmp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) BmpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bmpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Gif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) GifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jp2() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jp2Input() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jpeg() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpeg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) JpegInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpegInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jpg() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jpg2() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Jpg2Input() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) JpgInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Png() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"png",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) PngInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pngInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Tif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Tiff() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tiff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) TiffInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tiffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) TifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) Webp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference) WebpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webpInput",
		&returns,
	)
	return returns
}


func NewZoneCacheVariantsValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZoneCacheVariantsValueOutputReference {
	_init_.Initialize()

	if err := validateNewZoneCacheVariantsValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneCacheVariantsValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneCacheVariants.ZoneCacheVariantsValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZoneCacheVariantsValueOutputReference_Override(z ZoneCacheVariantsValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneCacheVariants.ZoneCacheVariantsValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetAvif(val *[]*string) {
	if err := j.validateSetAvifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetBmp(val *[]*string) {
	if err := j.validateSetBmpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bmp",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetGif(val *[]*string) {
	if err := j.validateSetGifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetJp2(val *[]*string) {
	if err := j.validateSetJp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jp2",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetJpeg(val *[]*string) {
	if err := j.validateSetJpegParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpeg",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetJpg(val *[]*string) {
	if err := j.validateSetJpgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpg",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetJpg2(val *[]*string) {
	if err := j.validateSetJpg2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpg2",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetPng(val *[]*string) {
	if err := j.validateSetPngParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"png",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetTif(val *[]*string) {
	if err := j.validateSetTifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetTiff(val *[]*string) {
	if err := j.validateSetTiffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tiff",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariantsValueOutputReference)SetWebp(val *[]*string) {
	if err := j.validateSetWebpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webp",
		val,
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetAvif() {
	_jsii_.InvokeVoid(
		z,
		"resetAvif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetBmp() {
	_jsii_.InvokeVoid(
		z,
		"resetBmp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetGif() {
	_jsii_.InvokeVoid(
		z,
		"resetGif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetJp2() {
	_jsii_.InvokeVoid(
		z,
		"resetJp2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetJpeg() {
	_jsii_.InvokeVoid(
		z,
		"resetJpeg",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetJpg() {
	_jsii_.InvokeVoid(
		z,
		"resetJpg",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetJpg2() {
	_jsii_.InvokeVoid(
		z,
		"resetJpg2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetPng() {
	_jsii_.InvokeVoid(
		z,
		"resetPng",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetTif() {
	_jsii_.InvokeVoid(
		z,
		"resetTif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetTiff() {
	_jsii_.InvokeVoid(
		z,
		"resetTiff",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ResetWebp() {
	_jsii_.InvokeVoid(
		z,
		"resetWebp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZoneCacheVariantsValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


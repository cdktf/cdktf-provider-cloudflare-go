package imagevariant


type ImageVariantOptions struct {
	// The fit property describes how the width and height dimensions should be interpreted. Available values: "scale-down", "contain", "cover", "crop", "pad".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/image_variant#fit ImageVariant#fit}
	Fit *string `field:"required" json:"fit" yaml:"fit"`
	// Maximum height in image pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/image_variant#height ImageVariant#height}
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// What EXIF data should be preserved in the output image. Available values: "keep", "copyright", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/image_variant#metadata ImageVariant#metadata}
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// Maximum width in image pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/image_variant#width ImageVariant#width}
	Width *float64 `field:"required" json:"width" yaml:"width"`
}


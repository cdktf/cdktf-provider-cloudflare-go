// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonecachevariants


type ZoneCacheVariantsValue struct {
	// List of strings with the MIME types of all the variants that should be served for avif.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#avif ZoneCacheVariants#avif}
	Avif *[]*string `field:"optional" json:"avif" yaml:"avif"`
	// List of strings with the MIME types of all the variants that should be served for bmp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#bmp ZoneCacheVariants#bmp}
	Bmp *[]*string `field:"optional" json:"bmp" yaml:"bmp"`
	// List of strings with the MIME types of all the variants that should be served for gif.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#gif ZoneCacheVariants#gif}
	Gif *[]*string `field:"optional" json:"gif" yaml:"gif"`
	// List of strings with the MIME types of all the variants that should be served for jp2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#jp2 ZoneCacheVariants#jp2}
	Jp2 *[]*string `field:"optional" json:"jp2" yaml:"jp2"`
	// List of strings with the MIME types of all the variants that should be served for jpeg.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#jpeg ZoneCacheVariants#jpeg}
	Jpeg *[]*string `field:"optional" json:"jpeg" yaml:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served for jpg.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#jpg ZoneCacheVariants#jpg}
	Jpg *[]*string `field:"optional" json:"jpg" yaml:"jpg"`
	// List of strings with the MIME types of all the variants that should be served for jpg2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#jpg2 ZoneCacheVariants#jpg2}
	Jpg2 *[]*string `field:"optional" json:"jpg2" yaml:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served for png.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#png ZoneCacheVariants#png}
	Png *[]*string `field:"optional" json:"png" yaml:"png"`
	// List of strings with the MIME types of all the variants that should be served for tif.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#tif ZoneCacheVariants#tif}
	Tif *[]*string `field:"optional" json:"tif" yaml:"tif"`
	// List of strings with the MIME types of all the variants that should be served for tiff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#tiff ZoneCacheVariants#tiff}
	Tiff *[]*string `field:"optional" json:"tiff" yaml:"tiff"`
	// List of strings with the MIME types of all the variants that should be served for webp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zone_cache_variants#webp ZoneCacheVariants#webp}
	Webp *[]*string `field:"optional" json:"webp" yaml:"webp"`
}


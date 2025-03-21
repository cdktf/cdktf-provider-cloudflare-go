// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecord


type DnsRecordData struct {
	// Algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#algorithm DnsRecord#algorithm}
	Algorithm *float64 `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Altitude of location in meters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#altitude DnsRecord#altitude}
	Altitude *float64 `field:"optional" json:"altitude" yaml:"altitude"`
	// Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#certificate DnsRecord#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Digest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#digest DnsRecord#digest}
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Digest Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#digest_type DnsRecord#digest_type}
	DigestType *float64 `field:"optional" json:"digestType" yaml:"digestType"`
	// fingerprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#fingerprint DnsRecord#fingerprint}
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
	// Flags for the CAA record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#flags DnsRecord#flags}
	Flags *float64 `field:"optional" json:"flags" yaml:"flags"`
	// Key Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#key_tag DnsRecord#key_tag}
	KeyTag *float64 `field:"optional" json:"keyTag" yaml:"keyTag"`
	// Degrees of latitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#lat_degrees DnsRecord#lat_degrees}
	LatDegrees *float64 `field:"optional" json:"latDegrees" yaml:"latDegrees"`
	// Latitude direction. Available values: "N", "S".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#lat_direction DnsRecord#lat_direction}
	LatDirection *string `field:"optional" json:"latDirection" yaml:"latDirection"`
	// Minutes of latitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#lat_minutes DnsRecord#lat_minutes}
	LatMinutes *float64 `field:"optional" json:"latMinutes" yaml:"latMinutes"`
	// Seconds of latitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#lat_seconds DnsRecord#lat_seconds}
	LatSeconds *float64 `field:"optional" json:"latSeconds" yaml:"latSeconds"`
	// Degrees of longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#long_degrees DnsRecord#long_degrees}
	LongDegrees *float64 `field:"optional" json:"longDegrees" yaml:"longDegrees"`
	// Longitude direction. Available values: "E", "W".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#long_direction DnsRecord#long_direction}
	LongDirection *string `field:"optional" json:"longDirection" yaml:"longDirection"`
	// Minutes of longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#long_minutes DnsRecord#long_minutes}
	LongMinutes *float64 `field:"optional" json:"longMinutes" yaml:"longMinutes"`
	// Seconds of longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#long_seconds DnsRecord#long_seconds}
	LongSeconds *float64 `field:"optional" json:"longSeconds" yaml:"longSeconds"`
	// Matching Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#matching_type DnsRecord#matching_type}
	MatchingType *float64 `field:"optional" json:"matchingType" yaml:"matchingType"`
	// Order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#order DnsRecord#order}
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// The port of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#port DnsRecord#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Horizontal precision of location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#precision_horz DnsRecord#precision_horz}
	PrecisionHorz *float64 `field:"optional" json:"precisionHorz" yaml:"precisionHorz"`
	// Vertical precision of location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#precision_vert DnsRecord#precision_vert}
	PrecisionVert *float64 `field:"optional" json:"precisionVert" yaml:"precisionVert"`
	// Preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#preference DnsRecord#preference}
	Preference *float64 `field:"optional" json:"preference" yaml:"preference"`
	// priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#priority DnsRecord#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#protocol DnsRecord#protocol}
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// Public Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#public_key DnsRecord#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#regex DnsRecord#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Replacement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#replacement DnsRecord#replacement}
	Replacement *string `field:"optional" json:"replacement" yaml:"replacement"`
	// Selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#selector DnsRecord#selector}
	Selector *float64 `field:"optional" json:"selector" yaml:"selector"`
	// Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#service DnsRecord#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Size of location in meters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#size DnsRecord#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#tag DnsRecord#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#target DnsRecord#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#type DnsRecord#type}
	Type *float64 `field:"optional" json:"type" yaml:"type"`
	// Usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#usage DnsRecord#usage}
	Usage *float64 `field:"optional" json:"usage" yaml:"usage"`
	// Value of the record. This field's semantics depend on the chosen tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#value DnsRecord#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The record weight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/dns_record#weight DnsRecord#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RecordData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#algorithm Record#algorithm}.
	Algorithm *float64 `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#altitude Record#altitude}.
	Altitude *float64 `field:"optional" json:"altitude" yaml:"altitude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#certificate Record#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#content Record#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#digest Record#digest}.
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#digest_type Record#digest_type}.
	DigestType *float64 `field:"optional" json:"digestType" yaml:"digestType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#fingerprint Record#fingerprint}.
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#flags Record#flags}.
	Flags *string `field:"optional" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#key_tag Record#key_tag}.
	KeyTag *float64 `field:"optional" json:"keyTag" yaml:"keyTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#lat_degrees Record#lat_degrees}.
	LatDegrees *float64 `field:"optional" json:"latDegrees" yaml:"latDegrees"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#lat_direction Record#lat_direction}.
	LatDirection *string `field:"optional" json:"latDirection" yaml:"latDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#lat_minutes Record#lat_minutes}.
	LatMinutes *float64 `field:"optional" json:"latMinutes" yaml:"latMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#lat_seconds Record#lat_seconds}.
	LatSeconds *float64 `field:"optional" json:"latSeconds" yaml:"latSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#long_degrees Record#long_degrees}.
	LongDegrees *float64 `field:"optional" json:"longDegrees" yaml:"longDegrees"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#long_direction Record#long_direction}.
	LongDirection *string `field:"optional" json:"longDirection" yaml:"longDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#long_minutes Record#long_minutes}.
	LongMinutes *float64 `field:"optional" json:"longMinutes" yaml:"longMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#long_seconds Record#long_seconds}.
	LongSeconds *float64 `field:"optional" json:"longSeconds" yaml:"longSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#matching_type Record#matching_type}.
	MatchingType *float64 `field:"optional" json:"matchingType" yaml:"matchingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#name Record#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#order Record#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#port Record#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#precision_horz Record#precision_horz}.
	PrecisionHorz *float64 `field:"optional" json:"precisionHorz" yaml:"precisionHorz"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#precision_vert Record#precision_vert}.
	PrecisionVert *float64 `field:"optional" json:"precisionVert" yaml:"precisionVert"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#preference Record#preference}.
	Preference *float64 `field:"optional" json:"preference" yaml:"preference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#priority Record#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#proto Record#proto}.
	Proto *string `field:"optional" json:"proto" yaml:"proto"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#protocol Record#protocol}.
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#public_key Record#public_key}.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#regex Record#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#replacement Record#replacement}.
	Replacement *string `field:"optional" json:"replacement" yaml:"replacement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#selector Record#selector}.
	Selector *float64 `field:"optional" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#service Record#service}.
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#size Record#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#tag Record#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#target Record#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#type Record#type}.
	Type *float64 `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#usage Record#usage}.
	Usage *float64 `field:"optional" json:"usage" yaml:"usage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#value Record#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#weight Record#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


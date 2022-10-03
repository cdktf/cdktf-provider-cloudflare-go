package apishield


type ApiShieldAuthIdCharacteristics struct {
	// The name of the characteristic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_shield#name ApiShield#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of characteristic. Available values: `header`, `cookie`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_shield#type ApiShield#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}


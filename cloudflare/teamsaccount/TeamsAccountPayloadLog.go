package teamsaccount


type TeamsAccountPayloadLog struct {
	// Public key used to encrypt matched payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.9.0/docs/resources/teams_account#public_key TeamsAccount#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}


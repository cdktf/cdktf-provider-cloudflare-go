// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupRequire struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#any_valid_service_token ZeroTrustAccessGroup#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessGroupRequireAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#auth_context ZeroTrustAccessGroup#auth_context}.
	AuthContext *ZeroTrustAccessGroupRequireAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#auth_method ZeroTrustAccessGroup#auth_method}.
	AuthMethod *ZeroTrustAccessGroupRequireAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#azure_ad ZeroTrustAccessGroup#azure_ad}.
	AzureAd *ZeroTrustAccessGroupRequireAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#certificate ZeroTrustAccessGroup#certificate}.
	Certificate *ZeroTrustAccessGroupRequireCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#common_name ZeroTrustAccessGroup#common_name}.
	CommonName *ZeroTrustAccessGroupRequireCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#device_posture ZeroTrustAccessGroup#device_posture}.
	DevicePosture *ZeroTrustAccessGroupRequireDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#email ZeroTrustAccessGroup#email}.
	Email *ZeroTrustAccessGroupRequireEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#email_domain ZeroTrustAccessGroup#email_domain}.
	EmailDomain *ZeroTrustAccessGroupRequireEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#email_list ZeroTrustAccessGroup#email_list}.
	EmailList *ZeroTrustAccessGroupRequireEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#everyone ZeroTrustAccessGroup#everyone}
	Everyone *ZeroTrustAccessGroupRequireEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#external_evaluation ZeroTrustAccessGroup#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessGroupRequireExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#geo ZeroTrustAccessGroup#geo}.
	Geo *ZeroTrustAccessGroupRequireGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#github_organization ZeroTrustAccessGroup#github_organization}.
	GithubOrganization *ZeroTrustAccessGroupRequireGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#group ZeroTrustAccessGroup#group}.
	Group *ZeroTrustAccessGroupRequireGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#gsuite ZeroTrustAccessGroup#gsuite}.
	Gsuite *ZeroTrustAccessGroupRequireGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#ip ZeroTrustAccessGroup#ip}.
	Ip *ZeroTrustAccessGroupRequireIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#ip_list ZeroTrustAccessGroup#ip_list}.
	IpList *ZeroTrustAccessGroupRequireIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#login_method ZeroTrustAccessGroup#login_method}.
	LoginMethod *ZeroTrustAccessGroupRequireLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#oidc ZeroTrustAccessGroup#oidc}.
	Oidc *ZeroTrustAccessGroupRequireOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#okta ZeroTrustAccessGroup#okta}.
	Okta *ZeroTrustAccessGroupRequireOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#saml ZeroTrustAccessGroup#saml}.
	Saml *ZeroTrustAccessGroupRequireSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_group#service_token ZeroTrustAccessGroup#service_token}.
	ServiceToken *ZeroTrustAccessGroupRequireServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


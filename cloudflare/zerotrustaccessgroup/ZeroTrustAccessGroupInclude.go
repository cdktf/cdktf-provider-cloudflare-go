// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupInclude struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#any_valid_service_token ZeroTrustAccessGroup#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessGroupIncludeAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#auth_context ZeroTrustAccessGroup#auth_context}.
	AuthContext *ZeroTrustAccessGroupIncludeAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#auth_method ZeroTrustAccessGroup#auth_method}.
	AuthMethod *ZeroTrustAccessGroupIncludeAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#azure_ad ZeroTrustAccessGroup#azure_ad}.
	AzureAd *ZeroTrustAccessGroupIncludeAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#certificate ZeroTrustAccessGroup#certificate}.
	Certificate *ZeroTrustAccessGroupIncludeCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#common_name ZeroTrustAccessGroup#common_name}.
	CommonName *ZeroTrustAccessGroupIncludeCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#device_posture ZeroTrustAccessGroup#device_posture}.
	DevicePosture *ZeroTrustAccessGroupIncludeDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#email ZeroTrustAccessGroup#email}.
	Email *ZeroTrustAccessGroupIncludeEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#email_domain ZeroTrustAccessGroup#email_domain}.
	EmailDomain *ZeroTrustAccessGroupIncludeEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#email_list ZeroTrustAccessGroup#email_list}.
	EmailList *ZeroTrustAccessGroupIncludeEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#everyone ZeroTrustAccessGroup#everyone}
	Everyone *ZeroTrustAccessGroupIncludeEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#external_evaluation ZeroTrustAccessGroup#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessGroupIncludeExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#geo ZeroTrustAccessGroup#geo}.
	Geo *ZeroTrustAccessGroupIncludeGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#github_organization ZeroTrustAccessGroup#github_organization}.
	GithubOrganization *ZeroTrustAccessGroupIncludeGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#group ZeroTrustAccessGroup#group}.
	Group *ZeroTrustAccessGroupIncludeGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#gsuite ZeroTrustAccessGroup#gsuite}.
	Gsuite *ZeroTrustAccessGroupIncludeGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#ip ZeroTrustAccessGroup#ip}.
	Ip *ZeroTrustAccessGroupIncludeIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#ip_list ZeroTrustAccessGroup#ip_list}.
	IpList *ZeroTrustAccessGroupIncludeIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#linked_app_token ZeroTrustAccessGroup#linked_app_token}.
	LinkedAppToken *ZeroTrustAccessGroupIncludeLinkedAppToken `field:"optional" json:"linkedAppToken" yaml:"linkedAppToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#login_method ZeroTrustAccessGroup#login_method}.
	LoginMethod *ZeroTrustAccessGroupIncludeLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#oidc ZeroTrustAccessGroup#oidc}.
	Oidc *ZeroTrustAccessGroupIncludeOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#okta ZeroTrustAccessGroup#okta}.
	Okta *ZeroTrustAccessGroupIncludeOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#saml ZeroTrustAccessGroup#saml}.
	Saml *ZeroTrustAccessGroupIncludeSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_group#service_token ZeroTrustAccessGroup#service_token}.
	ServiceToken *ZeroTrustAccessGroupIncludeServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


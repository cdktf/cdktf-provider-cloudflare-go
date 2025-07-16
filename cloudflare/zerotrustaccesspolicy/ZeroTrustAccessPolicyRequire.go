// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyRequire struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#any_valid_service_token ZeroTrustAccessPolicy#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessPolicyRequireAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#auth_context ZeroTrustAccessPolicy#auth_context}.
	AuthContext *ZeroTrustAccessPolicyRequireAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#auth_method ZeroTrustAccessPolicy#auth_method}.
	AuthMethod *ZeroTrustAccessPolicyRequireAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#azure_ad ZeroTrustAccessPolicy#azure_ad}.
	AzureAd *ZeroTrustAccessPolicyRequireAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#certificate ZeroTrustAccessPolicy#certificate}.
	Certificate *ZeroTrustAccessPolicyRequireCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#common_name ZeroTrustAccessPolicy#common_name}.
	CommonName *ZeroTrustAccessPolicyRequireCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#device_posture ZeroTrustAccessPolicy#device_posture}.
	DevicePosture *ZeroTrustAccessPolicyRequireDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}.
	Email *ZeroTrustAccessPolicyRequireEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#email_domain ZeroTrustAccessPolicy#email_domain}.
	EmailDomain *ZeroTrustAccessPolicyRequireEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#email_list ZeroTrustAccessPolicy#email_list}.
	EmailList *ZeroTrustAccessPolicyRequireEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#everyone ZeroTrustAccessPolicy#everyone}
	Everyone *ZeroTrustAccessPolicyRequireEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#external_evaluation ZeroTrustAccessPolicy#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessPolicyRequireExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#geo ZeroTrustAccessPolicy#geo}.
	Geo *ZeroTrustAccessPolicyRequireGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#github_organization ZeroTrustAccessPolicy#github_organization}.
	GithubOrganization *ZeroTrustAccessPolicyRequireGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#group ZeroTrustAccessPolicy#group}.
	Group *ZeroTrustAccessPolicyRequireGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#gsuite ZeroTrustAccessPolicy#gsuite}.
	Gsuite *ZeroTrustAccessPolicyRequireGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#ip ZeroTrustAccessPolicy#ip}.
	Ip *ZeroTrustAccessPolicyRequireIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#ip_list ZeroTrustAccessPolicy#ip_list}.
	IpList *ZeroTrustAccessPolicyRequireIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#login_method ZeroTrustAccessPolicy#login_method}.
	LoginMethod *ZeroTrustAccessPolicyRequireLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#oidc ZeroTrustAccessPolicy#oidc}.
	Oidc *ZeroTrustAccessPolicyRequireOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#okta ZeroTrustAccessPolicy#okta}.
	Okta *ZeroTrustAccessPolicyRequireOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#saml ZeroTrustAccessPolicy#saml}.
	Saml *ZeroTrustAccessPolicyRequireSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy#service_token ZeroTrustAccessPolicy#service_token}.
	ServiceToken *ZeroTrustAccessPolicyRequireServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


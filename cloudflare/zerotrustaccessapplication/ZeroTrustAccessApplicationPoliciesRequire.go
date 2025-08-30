// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesRequire struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#any_valid_service_token ZeroTrustAccessApplication#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#auth_context ZeroTrustAccessApplication#auth_context}.
	AuthContext *ZeroTrustAccessApplicationPoliciesRequireAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#auth_method ZeroTrustAccessApplication#auth_method}.
	AuthMethod *ZeroTrustAccessApplicationPoliciesRequireAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#azure_ad ZeroTrustAccessApplication#azure_ad}.
	AzureAd *ZeroTrustAccessApplicationPoliciesRequireAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#certificate ZeroTrustAccessApplication#certificate}.
	Certificate *ZeroTrustAccessApplicationPoliciesRequireCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#common_name ZeroTrustAccessApplication#common_name}.
	CommonName *ZeroTrustAccessApplicationPoliciesRequireCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#device_posture ZeroTrustAccessApplication#device_posture}.
	DevicePosture *ZeroTrustAccessApplicationPoliciesRequireDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#email ZeroTrustAccessApplication#email}.
	Email *ZeroTrustAccessApplicationPoliciesRequireEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#email_domain ZeroTrustAccessApplication#email_domain}.
	EmailDomain *ZeroTrustAccessApplicationPoliciesRequireEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#email_list ZeroTrustAccessApplication#email_list}.
	EmailList *ZeroTrustAccessApplicationPoliciesRequireEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#everyone ZeroTrustAccessApplication#everyone}
	Everyone *ZeroTrustAccessApplicationPoliciesRequireEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#external_evaluation ZeroTrustAccessApplication#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessApplicationPoliciesRequireExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#geo ZeroTrustAccessApplication#geo}.
	Geo *ZeroTrustAccessApplicationPoliciesRequireGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#github_organization ZeroTrustAccessApplication#github_organization}.
	GithubOrganization *ZeroTrustAccessApplicationPoliciesRequireGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#group ZeroTrustAccessApplication#group}.
	Group *ZeroTrustAccessApplicationPoliciesRequireGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#gsuite ZeroTrustAccessApplication#gsuite}.
	Gsuite *ZeroTrustAccessApplicationPoliciesRequireGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#ip ZeroTrustAccessApplication#ip}.
	Ip *ZeroTrustAccessApplicationPoliciesRequireIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#ip_list ZeroTrustAccessApplication#ip_list}.
	IpList *ZeroTrustAccessApplicationPoliciesRequireIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#linked_app_token ZeroTrustAccessApplication#linked_app_token}.
	LinkedAppToken *ZeroTrustAccessApplicationPoliciesRequireLinkedAppToken `field:"optional" json:"linkedAppToken" yaml:"linkedAppToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#login_method ZeroTrustAccessApplication#login_method}.
	LoginMethod *ZeroTrustAccessApplicationPoliciesRequireLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#oidc ZeroTrustAccessApplication#oidc}.
	Oidc *ZeroTrustAccessApplicationPoliciesRequireOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#okta ZeroTrustAccessApplication#okta}.
	Okta *ZeroTrustAccessApplicationPoliciesRequireOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#saml ZeroTrustAccessApplication#saml}.
	Saml *ZeroTrustAccessApplicationPoliciesRequireSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_application#service_token ZeroTrustAccessApplication#service_token}.
	ServiceToken *ZeroTrustAccessApplicationPoliciesRequireServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


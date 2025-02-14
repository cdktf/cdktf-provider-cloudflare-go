// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesInclude struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#any_valid_service_token ZeroTrustAccessApplication#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#auth_context ZeroTrustAccessApplication#auth_context}.
	AuthContext *ZeroTrustAccessApplicationPoliciesIncludeAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#auth_method ZeroTrustAccessApplication#auth_method}.
	AuthMethod *ZeroTrustAccessApplicationPoliciesIncludeAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#azure_ad ZeroTrustAccessApplication#azure_ad}.
	AzureAd *ZeroTrustAccessApplicationPoliciesIncludeAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#certificate ZeroTrustAccessApplication#certificate}.
	Certificate *ZeroTrustAccessApplicationPoliciesIncludeCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#common_name ZeroTrustAccessApplication#common_name}.
	CommonName *ZeroTrustAccessApplicationPoliciesIncludeCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#device_posture ZeroTrustAccessApplication#device_posture}.
	DevicePosture *ZeroTrustAccessApplicationPoliciesIncludeDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#email ZeroTrustAccessApplication#email}.
	Email *ZeroTrustAccessApplicationPoliciesIncludeEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#email_domain ZeroTrustAccessApplication#email_domain}.
	EmailDomain *ZeroTrustAccessApplicationPoliciesIncludeEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#email_list ZeroTrustAccessApplication#email_list}.
	EmailList *ZeroTrustAccessApplicationPoliciesIncludeEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#everyone ZeroTrustAccessApplication#everyone}
	Everyone *ZeroTrustAccessApplicationPoliciesIncludeEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#external_evaluation ZeroTrustAccessApplication#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#geo ZeroTrustAccessApplication#geo}.
	Geo *ZeroTrustAccessApplicationPoliciesIncludeGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#github_organization ZeroTrustAccessApplication#github_organization}.
	GithubOrganization *ZeroTrustAccessApplicationPoliciesIncludeGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#group ZeroTrustAccessApplication#group}.
	Group *ZeroTrustAccessApplicationPoliciesIncludeGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#gsuite ZeroTrustAccessApplication#gsuite}.
	Gsuite *ZeroTrustAccessApplicationPoliciesIncludeGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#ip ZeroTrustAccessApplication#ip}.
	Ip *ZeroTrustAccessApplicationPoliciesIncludeIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#ip_list ZeroTrustAccessApplication#ip_list}.
	IpList *ZeroTrustAccessApplicationPoliciesIncludeIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#login_method ZeroTrustAccessApplication#login_method}.
	LoginMethod *ZeroTrustAccessApplicationPoliciesIncludeLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#okta ZeroTrustAccessApplication#okta}.
	Okta *ZeroTrustAccessApplicationPoliciesIncludeOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#saml ZeroTrustAccessApplication#saml}.
	Saml *ZeroTrustAccessApplicationPoliciesIncludeSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#service_token ZeroTrustAccessApplication#service_token}.
	ServiceToken *ZeroTrustAccessApplicationPoliciesIncludeServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyInclude struct {
	// An empty object which matches on all service tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#any_valid_service_token ZeroTrustAccessPolicy#any_valid_service_token}
	AnyValidServiceToken *ZeroTrustAccessPolicyIncludeAnyValidServiceToken `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#auth_context ZeroTrustAccessPolicy#auth_context}.
	AuthContext *ZeroTrustAccessPolicyIncludeAuthContext `field:"optional" json:"authContext" yaml:"authContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#auth_method ZeroTrustAccessPolicy#auth_method}.
	AuthMethod *ZeroTrustAccessPolicyIncludeAuthMethod `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#azure_ad ZeroTrustAccessPolicy#azure_ad}.
	AzureAd *ZeroTrustAccessPolicyIncludeAzureAd `field:"optional" json:"azureAd" yaml:"azureAd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#certificate ZeroTrustAccessPolicy#certificate}.
	Certificate *ZeroTrustAccessPolicyIncludeCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#common_name ZeroTrustAccessPolicy#common_name}.
	CommonName *ZeroTrustAccessPolicyIncludeCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#device_posture ZeroTrustAccessPolicy#device_posture}.
	DevicePosture *ZeroTrustAccessPolicyIncludeDevicePosture `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}.
	Email *ZeroTrustAccessPolicyIncludeEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#email_domain ZeroTrustAccessPolicy#email_domain}.
	EmailDomain *ZeroTrustAccessPolicyIncludeEmailDomain `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#email_list ZeroTrustAccessPolicy#email_list}.
	EmailList *ZeroTrustAccessPolicyIncludeEmailListStruct `field:"optional" json:"emailList" yaml:"emailList"`
	// An empty object which matches on all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#everyone ZeroTrustAccessPolicy#everyone}
	Everyone *ZeroTrustAccessPolicyIncludeEveryone `field:"optional" json:"everyone" yaml:"everyone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#external_evaluation ZeroTrustAccessPolicy#external_evaluation}.
	ExternalEvaluation *ZeroTrustAccessPolicyIncludeExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#geo ZeroTrustAccessPolicy#geo}.
	Geo *ZeroTrustAccessPolicyIncludeGeo `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#github_organization ZeroTrustAccessPolicy#github_organization}.
	GithubOrganization *ZeroTrustAccessPolicyIncludeGithubOrganization `field:"optional" json:"githubOrganization" yaml:"githubOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#group ZeroTrustAccessPolicy#group}.
	Group *ZeroTrustAccessPolicyIncludeGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#gsuite ZeroTrustAccessPolicy#gsuite}.
	Gsuite *ZeroTrustAccessPolicyIncludeGsuite `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#ip ZeroTrustAccessPolicy#ip}.
	Ip *ZeroTrustAccessPolicyIncludeIp `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#ip_list ZeroTrustAccessPolicy#ip_list}.
	IpList *ZeroTrustAccessPolicyIncludeIpListStruct `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#login_method ZeroTrustAccessPolicy#login_method}.
	LoginMethod *ZeroTrustAccessPolicyIncludeLoginMethod `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#okta ZeroTrustAccessPolicy#okta}.
	Okta *ZeroTrustAccessPolicyIncludeOkta `field:"optional" json:"okta" yaml:"okta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#saml ZeroTrustAccessPolicy#saml}.
	Saml *ZeroTrustAccessPolicyIncludeSaml `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#service_token ZeroTrustAccessPolicy#service_token}.
	ServiceToken *ZeroTrustAccessPolicyIncludeServiceToken `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


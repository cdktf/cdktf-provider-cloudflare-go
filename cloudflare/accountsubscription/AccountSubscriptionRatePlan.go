// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsubscription


type AccountSubscriptionRatePlan struct {
	// The currency applied to the rate plan subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#currency AccountSubscription#currency}
	Currency *string `field:"optional" json:"currency" yaml:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#externally_managed AccountSubscription#externally_managed}
	ExternallyManaged interface{} `field:"optional" json:"externallyManaged" yaml:"externallyManaged"`
	// The ID of the rate plan. Available values: "free", "lite", "pro", "pro_plus", "business", "enterprise", "partners_free", "partners_pro", "partners_business", "partners_enterprise".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#id AccountSubscription#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#is_contract AccountSubscription#is_contract}
	IsContract interface{} `field:"optional" json:"isContract" yaml:"isContract"`
	// The full name of the rate plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#public_name AccountSubscription#public_name}
	PublicName *string `field:"optional" json:"publicName" yaml:"publicName"`
	// The scope that this rate plan applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#scope AccountSubscription#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The list of sets this rate plan applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/account_subscription#sets AccountSubscription#sets}
	Sets *[]*string `field:"optional" json:"sets" yaml:"sets"`
}


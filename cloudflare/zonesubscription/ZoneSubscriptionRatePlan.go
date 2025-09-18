// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonesubscription


type ZoneSubscriptionRatePlan struct {
	// The currency applied to the rate plan subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#currency ZoneSubscription#currency}
	Currency *string `field:"optional" json:"currency" yaml:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#externally_managed ZoneSubscription#externally_managed}
	ExternallyManaged interface{} `field:"optional" json:"externallyManaged" yaml:"externallyManaged"`
	// The ID of the rate plan. Available values: "free", "lite", "pro", "pro_plus", "business", "enterprise", "partners_free", "partners_pro", "partners_business", "partners_enterprise".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#id ZoneSubscription#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#is_contract ZoneSubscription#is_contract}
	IsContract interface{} `field:"optional" json:"isContract" yaml:"isContract"`
	// The full name of the rate plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#public_name ZoneSubscription#public_name}
	PublicName *string `field:"optional" json:"publicName" yaml:"publicName"`
	// The scope that this rate plan applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#scope ZoneSubscription#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The list of sets this rate plan applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zone_subscription#sets ZoneSubscription#sets}
	Sets *[]*string `field:"optional" json:"sets" yaml:"sets"`
}


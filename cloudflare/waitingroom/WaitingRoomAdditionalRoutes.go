// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroom


type WaitingRoomAdditionalRoutes struct {
	// The hostname to which this waiting room will be applied (no wildcards).
	//
	// The hostname must be the primary domain, subdomain, or custom hostname (if using SSL for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/waiting_room#host WaitingRoom#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Sets the path within the host to enable the waiting room on.
	//
	// The waiting room will be enabled for all subpaths as well. If there are two waiting rooms on the same subpath, the waiting room for the most specific path will be chosen. Wildcards and query parameters are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/waiting_room#path WaitingRoom#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}


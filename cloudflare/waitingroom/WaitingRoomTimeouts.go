// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroom


type WaitingRoomTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/waiting_room#create WaitingRoom#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/waiting_room#update WaitingRoom#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


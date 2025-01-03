// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersServeStale struct {
	// Disable stale while updating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/ruleset#disable_stale_while_updating Ruleset#disable_stale_while_updating}
	DisableStaleWhileUpdating interface{} `field:"optional" json:"disableStaleWhileUpdating" yaml:"disableStaleWhileUpdating"`
}


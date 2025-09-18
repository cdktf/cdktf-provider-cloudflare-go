// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectSourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#deployments_enabled PagesProject#deployments_enabled}.
	DeploymentsEnabled interface{} `field:"optional" json:"deploymentsEnabled" yaml:"deploymentsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#owner PagesProject#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#path_excludes PagesProject#path_excludes}.
	PathExcludes *[]*string `field:"optional" json:"pathExcludes" yaml:"pathExcludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#path_includes PagesProject#path_includes}.
	PathIncludes *[]*string `field:"optional" json:"pathIncludes" yaml:"pathIncludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#pr_comments_enabled PagesProject#pr_comments_enabled}.
	PrCommentsEnabled interface{} `field:"optional" json:"prCommentsEnabled" yaml:"prCommentsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#preview_branch_excludes PagesProject#preview_branch_excludes}.
	PreviewBranchExcludes *[]*string `field:"optional" json:"previewBranchExcludes" yaml:"previewBranchExcludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#preview_branch_includes PagesProject#preview_branch_includes}.
	PreviewBranchIncludes *[]*string `field:"optional" json:"previewBranchIncludes" yaml:"previewBranchIncludes"`
	// Available values: "all", "none", "custom".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#preview_deployment_setting PagesProject#preview_deployment_setting}
	PreviewDeploymentSetting *string `field:"optional" json:"previewDeploymentSetting" yaml:"previewDeploymentSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#production_branch PagesProject#production_branch}.
	ProductionBranch *string `field:"optional" json:"productionBranch" yaml:"productionBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#production_deployments_enabled PagesProject#production_deployments_enabled}.
	ProductionDeploymentsEnabled interface{} `field:"optional" json:"productionDeploymentsEnabled" yaml:"productionDeploymentsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#repo_name PagesProject#repo_name}.
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}


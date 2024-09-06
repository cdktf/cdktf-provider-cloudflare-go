// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectBuildConfig struct {
	// Enable build caching for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#build_caching PagesProject#build_caching}
	BuildCaching interface{} `field:"optional" json:"buildCaching" yaml:"buildCaching"`
	// Command used to build project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#build_command PagesProject#build_command}
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// Output directory of the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#destination_dir PagesProject#destination_dir}
	DestinationDir *string `field:"optional" json:"destinationDir" yaml:"destinationDir"`
	// Your project's root directory, where Cloudflare runs the build command.
	//
	// If your site is not in a subdirectory, leave this path value empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#root_dir PagesProject#root_dir}
	RootDir *string `field:"optional" json:"rootDir" yaml:"rootDir"`
	// The classifying tag for analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#web_analytics_tag PagesProject#web_analytics_tag}
	WebAnalyticsTag *string `field:"optional" json:"webAnalyticsTag" yaml:"webAnalyticsTag"`
	// The auth token for analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/pages_project#web_analytics_token PagesProject#web_analytics_token}
	WebAnalyticsToken *string `field:"optional" json:"webAnalyticsToken" yaml:"webAnalyticsToken"`
}


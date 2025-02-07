// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pagesproject

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PagesProjectLatestDeploymentEnvVarsMap) validateGetParameters(key *string) error {
	return nil
}

func (p *jsiiProxy_PagesProjectLatestDeploymentEnvVarsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (p *jsiiProxy_PagesProjectLatestDeploymentEnvVarsMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PagesProjectLatestDeploymentEnvVarsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PagesProjectLatestDeploymentEnvVarsMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func validateNewPagesProjectLatestDeploymentEnvVarsMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}


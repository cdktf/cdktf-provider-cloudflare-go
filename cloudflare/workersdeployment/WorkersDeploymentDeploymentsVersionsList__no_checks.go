// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsVersionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersDeploymentDeploymentsVersionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


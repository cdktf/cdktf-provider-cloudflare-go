// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentVersionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentVersionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersDeploymentVersionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


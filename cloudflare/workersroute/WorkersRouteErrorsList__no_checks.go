// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersroute

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersRouteErrorsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersRouteErrorsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersRouteErrorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteErrorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteErrorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteErrorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersRouteErrorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


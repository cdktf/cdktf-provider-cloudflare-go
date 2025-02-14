// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersroute

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersRouteMessagesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersRouteMessagesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersRouteMessagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteMessagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteMessagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersRouteMessagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersRouteMessagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


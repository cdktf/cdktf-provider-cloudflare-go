// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersScriptNamedHandlersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptNamedHandlersList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptNamedHandlersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptNamedHandlersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptNamedHandlersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptNamedHandlersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersScriptNamedHandlersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersScriptBindingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptBindingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptBindingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersScriptBindingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


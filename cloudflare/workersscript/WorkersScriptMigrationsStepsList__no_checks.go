// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersScriptMigrationsStepsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptMigrationsStepsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptMigrationsStepsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersScriptMigrationsStepsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


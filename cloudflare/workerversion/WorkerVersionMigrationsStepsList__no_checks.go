// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workerversion

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerVersionMigrationsStepsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionMigrationsStepsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionMigrationsStepsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionMigrationsStepsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionMigrationsStepsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionMigrationsStepsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionMigrationsStepsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerVersionMigrationsStepsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workerversion

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerVersionBindingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionBindingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionBindingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerVersionBindingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


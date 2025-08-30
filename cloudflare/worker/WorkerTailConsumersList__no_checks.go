// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerTailConsumersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerTailConsumersList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerTailConsumersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerTailConsumersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerTailConsumersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerTailConsumersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerTailConsumersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerTailConsumersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


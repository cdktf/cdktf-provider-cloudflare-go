// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workerscrontrigger

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersCronTriggerSchedulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersCronTriggerSchedulesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersCronTriggerSchedulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersCronTriggerSchedulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersCronTriggerSchedulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersCronTriggerSchedulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersCronTriggerSchedulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersCronTriggerSchedulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


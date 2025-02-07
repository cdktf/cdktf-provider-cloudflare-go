// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queue

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueConsumersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumersList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueueConsumersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueueConsumersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueueConsumersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


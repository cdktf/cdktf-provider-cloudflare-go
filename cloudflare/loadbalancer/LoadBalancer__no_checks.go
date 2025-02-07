// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancer) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validatePutAdaptiveRoutingParameters(value *LoadBalancerAdaptiveRouting) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validatePutLocationStrategyParameters(value *LoadBalancerLocationStrategy) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validatePutRandomSteeringParameters(value *LoadBalancerRandomSteering) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validatePutRulesParameters(value interface{}) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validatePutSessionAffinityAttributesParameters(value *LoadBalancerSessionAffinityAttributes) error {
	return nil
}

func validateLoadBalancer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLoadBalancer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLoadBalancer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLoadBalancer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetCountryPoolsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetDefaultPoolsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetFallbackPoolParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetNetworksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetPopPoolsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetProxiedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetRegionPoolsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetSessionAffinityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetSessionAffinityTtlParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetSteeringPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetTtlParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_LoadBalancer) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewLoadBalancerParameters(scope constructs.Construct, id *string, config *LoadBalancerConfig) error {
	return nil
}


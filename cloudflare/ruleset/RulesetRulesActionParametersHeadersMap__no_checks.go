// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateGetParameters(key *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersHeadersMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func validateNewRulesetRulesActionParametersHeadersMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}


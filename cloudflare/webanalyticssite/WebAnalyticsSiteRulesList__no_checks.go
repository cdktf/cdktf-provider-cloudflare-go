// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package webanalyticssite

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebAnalyticsSiteRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WebAnalyticsSiteRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WebAnalyticsSiteRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WebAnalyticsSiteRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WebAnalyticsSiteRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WebAnalyticsSiteRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWebAnalyticsSiteRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package snippetrules

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnippetRulesRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnippetRulesRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnippetRulesRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnippetRulesRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnippetRulesRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


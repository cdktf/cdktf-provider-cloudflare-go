//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudflareProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CloudflareProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCloudflareProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudflareProvider) validateSetApiClientLoggingParameters(val interface{}) error {
	return nil
}

func validateNewCloudflareProviderParameters(scope constructs.Construct, id *string, config *CloudflareProviderConfig) error {
	return nil
}


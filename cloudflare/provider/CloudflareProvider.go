// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v12/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs cloudflare}.
type CloudflareProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	ApiUserServiceKey() *string
	SetApiUserServiceKey(val *string)
	ApiUserServiceKeyInput() *string
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UserAgentOperatorSuffix() *string
	SetUserAgentOperatorSuffix(val *string)
	UserAgentOperatorSuffixInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiKey()
	ResetApiToken()
	ResetApiUserServiceKey()
	ResetBaseUrl()
	ResetEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUserAgentOperatorSuffix()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudflareProvider
type jsiiProxy_CloudflareProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_CloudflareProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiUserServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUserServiceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ApiUserServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUserServiceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) UserAgentOperatorSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentOperatorSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudflareProvider) UserAgentOperatorSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentOperatorSuffixInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs cloudflare} Resource.
func NewCloudflareProvider(scope constructs.Construct, id *string, config *CloudflareProviderConfig) CloudflareProvider {
	_init_.Initialize()

	if err := validateNewCloudflareProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudflareProvider{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs cloudflare} Resource.
func NewCloudflareProvider_Override(c CloudflareProvider, scope constructs.Construct, id *string, config *CloudflareProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetApiUserServiceKey(val *string) {
	_jsii_.Set(
		j,
		"apiUserServiceKey",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_CloudflareProvider)SetUserAgentOperatorSuffix(val *string) {
	_jsii_.Set(
		j,
		"userAgentOperatorSuffix",
		val,
	)
}

// Generates CDKTF code for importing a CloudflareProvider resource upon running "cdktf plan <stack-name>".
func CloudflareProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudflareProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudflareProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudflareProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudflareProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudflareProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudflareProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudflareProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudflareProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.provider.CloudflareProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudflareProvider) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudflareProvider) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		c,
		"resetAlias",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetApiKey() {
	_jsii_.InvokeVoid(
		c,
		"resetApiKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetApiToken() {
	_jsii_.InvokeVoid(
		c,
		"resetApiToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetApiUserServiceKey() {
	_jsii_.InvokeVoid(
		c,
		"resetApiUserServiceKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) ResetUserAgentOperatorSuffix() {
	_jsii_.InvokeVoid(
		c,
		"resetUserAgentOperatorSuffix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudflareProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudflareProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudflareProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudflareProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudflareProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudflareProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


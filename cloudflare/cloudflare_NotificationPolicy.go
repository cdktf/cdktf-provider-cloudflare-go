// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy cloudflare_notification_policy}.
type NotificationPolicy interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AlertType() *string
	SetAlertType(val *string)
	AlertTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Created() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailIntegration() NotificationPolicyEmailIntegrationList
	EmailIntegrationInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Filters() NotificationPolicyFiltersOutputReference
	FiltersInput() *NotificationPolicyFilters
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Modified() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PagerdutyIntegration() NotificationPolicyPagerdutyIntegrationList
	PagerdutyIntegrationInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebhooksIntegration() NotificationPolicyWebhooksIntegrationList
	WebhooksIntegrationInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEmailIntegration(value interface{})
	PutFilters(value *NotificationPolicyFilters)
	PutPagerdutyIntegration(value interface{})
	PutWebhooksIntegration(value interface{})
	ResetDescription()
	ResetEmailIntegration()
	ResetFilters()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPagerdutyIntegration()
	ResetWebhooksIntegration()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NotificationPolicy
type jsiiProxy_NotificationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NotificationPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AlertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AlertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) EmailIntegration() NotificationPolicyEmailIntegrationList {
	var returns NotificationPolicyEmailIntegrationList
	_jsii_.Get(
		j,
		"emailIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) EmailIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Filters() NotificationPolicyFiltersOutputReference {
	var returns NotificationPolicyFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) FiltersInput() *NotificationPolicyFilters {
	var returns *NotificationPolicyFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Modified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) PagerdutyIntegration() NotificationPolicyPagerdutyIntegrationList {
	var returns NotificationPolicyPagerdutyIntegrationList
	_jsii_.Get(
		j,
		"pagerdutyIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) PagerdutyIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagerdutyIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) WebhooksIntegration() NotificationPolicyWebhooksIntegrationList {
	var returns NotificationPolicyWebhooksIntegrationList
	_jsii_.Get(
		j,
		"webhooksIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) WebhooksIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhooksIntegrationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy cloudflare_notification_policy} Resource.
func NewNotificationPolicy(scope constructs.Construct, id *string, config *NotificationPolicyConfig) NotificationPolicy {
	_init_.Initialize()

	j := jsiiProxy_NotificationPolicy{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.NotificationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy cloudflare_notification_policy} Resource.
func NewNotificationPolicy_Override(n NotificationPolicy, scope constructs.Construct, id *string, config *NotificationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.NotificationPolicy",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetAlertType(val *string) {
	_jsii_.Set(
		j,
		"alertType",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
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
func NotificationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.NotificationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NotificationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.NotificationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NotificationPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NotificationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutEmailIntegration(value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"putEmailIntegration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutFilters(value *NotificationPolicyFilters) {
	_jsii_.InvokeVoid(
		n,
		"putFilters",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutPagerdutyIntegration(value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"putPagerdutyIntegration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutWebhooksIntegration(value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"putWebhooksIntegration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetEmailIntegration() {
	_jsii_.InvokeVoid(
		n,
		"resetEmailIntegration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetFilters() {
	_jsii_.InvokeVoid(
		n,
		"resetFilters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetPagerdutyIntegration() {
	_jsii_.InvokeVoid(
		n,
		"resetPagerdutyIntegration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetWebhooksIntegration() {
	_jsii_.InvokeVoid(
		n,
		"resetWebhooksIntegration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyFiltersOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() *[]*string
	SetEnabled(val *[]*string)
	EnabledInput() *[]*string
	// Experimental.
	Fqn() *string
	HealthCheckId() *[]*string
	SetHealthCheckId(val *[]*string)
	HealthCheckIdInput() *[]*string
	InternalValue() *NotificationPolicyFilters
	SetInternalValue(val *NotificationPolicyFilters)
	Limit() *[]*string
	SetLimit(val *[]*string)
	LimitInput() *[]*string
	PoolId() *[]*string
	SetPoolId(val *[]*string)
	PoolIdInput() *[]*string
	Product() *[]*string
	SetProduct(val *[]*string)
	ProductInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
	Slo() *[]*string
	SetSlo(val *[]*string)
	SloInput() *[]*string
	Status() *[]*string
	SetStatus(val *[]*string)
	StatusInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEnabled()
	ResetHealthCheckId()
	ResetLimit()
	ResetPoolId()
	ResetProduct()
	ResetServices()
	ResetSlo()
	ResetStatus()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationPolicyFiltersOutputReference
type jsiiProxy_NotificationPolicyFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Enabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InternalValue() *NotificationPolicyFilters {
	var returns *NotificationPolicyFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Limit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) LimitInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Product() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProductInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Slo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SloInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sloInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Status() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) StatusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewNotificationPolicyFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationPolicyFiltersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NotificationPolicyFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationPolicyFiltersOutputReference_Override(n NotificationPolicyFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetEnabled(val *[]*string) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetHealthCheckId(val *[]*string) {
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetInternalValue(val *NotificationPolicyFilters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetLimit(val *[]*string) {
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetPoolId(val *[]*string) {
	_jsii_.Set(
		j,
		"poolId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetProduct(val *[]*string) {
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetServices(val *[]*string) {
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetSlo(val *[]*string) {
	_jsii_.Set(
		j,
		"slo",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetStatus(val *[]*string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		n,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		n,
		"resetLimit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPoolId() {
	_jsii_.InvokeVoid(
		n,
		"resetPoolId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProduct() {
	_jsii_.InvokeVoid(
		n,
		"resetProduct",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		n,
		"resetServices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetSlo() {
	_jsii_.InvokeVoid(
		n,
		"resetSlo",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		n,
		"resetStatus",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		n,
		"resetZones",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-cloudflare-go/cloudflare/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application cloudflare_spectrum_application}.
type SpectrumApplication interface {
	cdktf.TerraformResource
	ArgoSmartRouting() interface{}
	SetArgoSmartRouting(val interface{})
	ArgoSmartRoutingInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dns() SpectrumApplicationDnsOutputReference
	DnsInput() *SpectrumApplicationDns
	EdgeIpConnectivity() *string
	SetEdgeIpConnectivity(val *string)
	EdgeIpConnectivityInput() *string
	EdgeIps() *[]*string
	SetEdgeIps(val *[]*string)
	EdgeIpsInput() *[]*string
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
	IpFirewall() interface{}
	SetIpFirewall(val interface{})
	IpFirewallInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OriginDirect() *[]*string
	SetOriginDirect(val *[]*string)
	OriginDirectInput() *[]*string
	OriginDns() SpectrumApplicationOriginDnsOutputReference
	OriginDnsInput() *SpectrumApplicationOriginDns
	OriginPort() *float64
	SetOriginPort(val *float64)
	OriginPortInput() *float64
	OriginPortRange() SpectrumApplicationOriginPortRangeOutputReference
	OriginPortRangeInput() *SpectrumApplicationOriginPortRange
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyProtocol() *string
	SetProxyProtocol(val *string)
	ProxyProtocolInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tls() *string
	SetTls(val *string)
	TlsInput() *string
	TrafficType() *string
	SetTrafficType(val *string)
	TrafficTypeInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutDns(value *SpectrumApplicationDns)
	PutOriginDns(value *SpectrumApplicationOriginDns)
	PutOriginPortRange(value *SpectrumApplicationOriginPortRange)
	ResetArgoSmartRouting()
	ResetEdgeIpConnectivity()
	ResetEdgeIps()
	ResetId()
	ResetIpFirewall()
	ResetOriginDirect()
	ResetOriginDns()
	ResetOriginPort()
	ResetOriginPortRange()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProxyProtocol()
	ResetTls()
	ResetTrafficType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SpectrumApplication
type jsiiProxy_SpectrumApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpectrumApplication) ArgoSmartRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argoSmartRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ArgoSmartRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argoSmartRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Dns() SpectrumApplicationDnsOutputReference {
	var returns SpectrumApplicationDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) DnsInput() *SpectrumApplicationDns {
	var returns *SpectrumApplicationDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIpConnectivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeIpConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIpConnectivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeIpConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IpFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IpFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDirect() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDirectInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDns() SpectrumApplicationOriginDnsOutputReference {
	var returns SpectrumApplicationOriginDnsOutputReference
	_jsii_.Get(
		j,
		"originDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDnsInput() *SpectrumApplicationOriginDns {
	var returns *SpectrumApplicationOriginDns
	_jsii_.Get(
		j,
		"originDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortRange() SpectrumApplicationOriginPortRangeOutputReference {
	var returns SpectrumApplicationOriginPortRangeOutputReference
	_jsii_.Get(
		j,
		"originPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortRangeInput() *SpectrumApplicationOriginPortRange {
	var returns *SpectrumApplicationOriginPortRange
	_jsii_.Get(
		j,
		"originPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProxyProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProxyProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Tls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TrafficType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TrafficTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application cloudflare_spectrum_application} Resource.
func NewSpectrumApplication(scope constructs.Construct, id *string, config *SpectrumApplicationConfig) SpectrumApplication {
	_init_.Initialize()

	j := jsiiProxy_SpectrumApplication{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.SpectrumApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application cloudflare_spectrum_application} Resource.
func NewSpectrumApplication_Override(s SpectrumApplication, scope constructs.Construct, id *string, config *SpectrumApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.SpectrumApplication",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetArgoSmartRouting(val interface{}) {
	_jsii_.Set(
		j,
		"argoSmartRouting",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetEdgeIpConnectivity(val *string) {
	_jsii_.Set(
		j,
		"edgeIpConnectivity",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetEdgeIps(val *[]*string) {
	_jsii_.Set(
		j,
		"edgeIps",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetIpFirewall(val interface{}) {
	_jsii_.Set(
		j,
		"ipFirewall",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetOriginDirect(val *[]*string) {
	_jsii_.Set(
		j,
		"originDirect",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetOriginPort(val *float64) {
	_jsii_.Set(
		j,
		"originPort",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetProxyProtocol(val *string) {
	_jsii_.Set(
		j,
		"proxyProtocol",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetTls(val *string) {
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetTrafficType(val *string) {
	_jsii_.Set(
		j,
		"trafficType",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication) SetZoneId(val *string) {
	_jsii_.Set(
		j,
		"zoneId",
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
func SpectrumApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.SpectrumApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpectrumApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.SpectrumApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpectrumApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpectrumApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutDns(value *SpectrumApplicationDns) {
	_jsii_.InvokeVoid(
		s,
		"putDns",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutOriginDns(value *SpectrumApplicationOriginDns) {
	_jsii_.InvokeVoid(
		s,
		"putOriginDns",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutOriginPortRange(value *SpectrumApplicationOriginPortRange) {
	_jsii_.InvokeVoid(
		s,
		"putOriginPortRange",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetArgoSmartRouting() {
	_jsii_.InvokeVoid(
		s,
		"resetArgoSmartRouting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetEdgeIpConnectivity() {
	_jsii_.InvokeVoid(
		s,
		"resetEdgeIpConnectivity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetEdgeIps() {
	_jsii_.InvokeVoid(
		s,
		"resetEdgeIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetIpFirewall() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFirewall",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginDirect() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginDirect",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginDns() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginDns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginPort() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginPortRange() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginPortRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetProxyProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProxyProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetTls() {
	_jsii_.InvokeVoid(
		s,
		"resetTls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetTrafficType() {
	_jsii_.InvokeVoid(
		s,
		"resetTrafficType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

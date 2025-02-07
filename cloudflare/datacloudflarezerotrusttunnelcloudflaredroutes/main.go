// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflaredroutes

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutes",
		reflect.TypeOf((*DataCloudflareZeroTrustTunnelCloudflaredRoutes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "existedAt", GoGetter: "ExistedAt"},
			_jsii_.MemberProperty{JsiiProperty: "existedAtInput", GoGetter: "ExistedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isDeleted", GoGetter: "IsDeleted"},
			_jsii_.MemberProperty{JsiiProperty: "isDeletedInput", GoGetter: "IsDeletedInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxItems", GoGetter: "MaxItems"},
			_jsii_.MemberProperty{JsiiProperty: "maxItemsInput", GoGetter: "MaxItemsInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSubset", GoGetter: "NetworkSubset"},
			_jsii_.MemberProperty{JsiiProperty: "networkSubsetInput", GoGetter: "NetworkSubsetInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSuperset", GoGetter: "NetworkSuperset"},
			_jsii_.MemberProperty{JsiiProperty: "networkSupersetInput", GoGetter: "NetworkSupersetInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetExistedAt", GoMethod: "ResetExistedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsDeleted", GoMethod: "ResetIsDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxItems", GoMethod: "ResetMaxItems"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSubset", GoMethod: "ResetNetworkSubset"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSuperset", GoMethod: "ResetNetworkSuperset"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteId", GoMethod: "ResetRouteId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelId", GoMethod: "ResetTunnelId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunTypes", GoMethod: "ResetTunTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualNetworkId", GoMethod: "ResetVirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberProperty{JsiiProperty: "routeId", GoGetter: "RouteId"},
			_jsii_.MemberProperty{JsiiProperty: "routeIdInput", GoGetter: "RouteIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelId", GoGetter: "TunnelId"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelIdInput", GoGetter: "TunnelIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunTypes", GoGetter: "TunTypes"},
			_jsii_.MemberProperty{JsiiProperty: "tunTypesInput", GoGetter: "TunTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkIdInput", GoGetter: "VirtualNetworkIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutes{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutesConfig",
		reflect.TypeOf((*DataCloudflareZeroTrustTunnelCloudflaredRoutesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutesResult",
		reflect.TypeOf((*DataCloudflareZeroTrustTunnelCloudflaredRoutesResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList",
		reflect.TypeOf((*DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutesResultList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredRoutes.DataCloudflareZeroTrustTunnelCloudflaredRoutesResultOutputReference",
		reflect.TypeOf((*DataCloudflareZeroTrustTunnelCloudflaredRoutesResultOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deletedAt", GoGetter: "DeletedAt"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "network", GoGetter: "Network"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelId", GoGetter: "TunnelId"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelName", GoGetter: "TunnelName"},
			_jsii_.MemberProperty{JsiiProperty: "tunType", GoGetter: "TunType"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkName", GoGetter: "VirtualNetworkName"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredRoutesResultOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

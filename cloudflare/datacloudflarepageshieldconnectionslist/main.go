// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarepageshieldconnectionslist

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldConnectionsList.DataCloudflarePageShieldConnectionsList",
		reflect.TypeOf((*DataCloudflarePageShieldConnectionsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "direction", GoGetter: "Direction"},
			_jsii_.MemberProperty{JsiiProperty: "directionInput", GoGetter: "DirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCdnCgi", GoGetter: "ExcludeCdnCgi"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCdnCgiInput", GoGetter: "ExcludeCdnCgiInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeUrls", GoGetter: "ExcludeUrls"},
			_jsii_.MemberProperty{JsiiProperty: "excludeUrlsInput", GoGetter: "ExcludeUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "export", GoGetter: "Export"},
			_jsii_.MemberProperty{JsiiProperty: "exportInput", GoGetter: "ExportInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberProperty{JsiiProperty: "hostsInput", GoGetter: "HostsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxItems", GoGetter: "MaxItems"},
			_jsii_.MemberProperty{JsiiProperty: "maxItemsInput", GoGetter: "MaxItemsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderBy", GoGetter: "OrderBy"},
			_jsii_.MemberProperty{JsiiProperty: "orderByInput", GoGetter: "OrderByInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "page", GoGetter: "Page"},
			_jsii_.MemberProperty{JsiiProperty: "pageInput", GoGetter: "PageInput"},
			_jsii_.MemberProperty{JsiiProperty: "pageUrl", GoGetter: "PageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "pageUrlInput", GoGetter: "PageUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "perPage", GoGetter: "PerPage"},
			_jsii_.MemberProperty{JsiiProperty: "perPageInput", GoGetter: "PerPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "prioritizeMalicious", GoGetter: "PrioritizeMalicious"},
			_jsii_.MemberProperty{JsiiProperty: "prioritizeMaliciousInput", GoGetter: "PrioritizeMaliciousInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirection", GoMethod: "ResetDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeCdnCgi", GoMethod: "ResetExcludeCdnCgi"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeUrls", GoMethod: "ResetExcludeUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetExport", GoMethod: "ResetExport"},
			_jsii_.MemberMethod{JsiiMethod: "resetHosts", GoMethod: "ResetHosts"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxItems", GoMethod: "ResetMaxItems"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderBy", GoMethod: "ResetOrderBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPage", GoMethod: "ResetPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetPageUrl", GoMethod: "ResetPageUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerPage", GoMethod: "ResetPerPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrioritizeMalicious", GoMethod: "ResetPrioritizeMalicious"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrls", GoMethod: "ResetUrls"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "urls", GoGetter: "Urls"},
			_jsii_.MemberProperty{JsiiProperty: "urlsInput", GoGetter: "UrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneId", GoGetter: "ZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "zoneIdInput", GoGetter: "ZoneIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflarePageShieldConnectionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldConnectionsList.DataCloudflarePageShieldConnectionsListConfig",
		reflect.TypeOf((*DataCloudflarePageShieldConnectionsListConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldConnectionsList.DataCloudflarePageShieldConnectionsListResult",
		reflect.TypeOf((*DataCloudflarePageShieldConnectionsListResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldConnectionsList.DataCloudflarePageShieldConnectionsListResultList",
		reflect.TypeOf((*DataCloudflarePageShieldConnectionsListResultList)(nil)).Elem(),
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
			j := jsiiProxy_DataCloudflarePageShieldConnectionsListResultList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflarePageShieldConnectionsList.DataCloudflarePageShieldConnectionsListResultOutputReference",
		reflect.TypeOf((*DataCloudflarePageShieldConnectionsListResultOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addedAt", GoGetter: "AddedAt"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainReportedMalicious", GoGetter: "DomainReportedMalicious"},
			_jsii_.MemberProperty{JsiiProperty: "firstPageUrl", GoGetter: "FirstPageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "firstSeenAt", GoGetter: "FirstSeenAt"},
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
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastSeenAt", GoGetter: "LastSeenAt"},
			_jsii_.MemberProperty{JsiiProperty: "maliciousDomainCategories", GoGetter: "MaliciousDomainCategories"},
			_jsii_.MemberProperty{JsiiProperty: "maliciousUrlCategories", GoGetter: "MaliciousUrlCategories"},
			_jsii_.MemberProperty{JsiiProperty: "pageUrls", GoGetter: "PageUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlContainsCdnCgiPath", GoGetter: "UrlContainsCdnCgiPath"},
			_jsii_.MemberProperty{JsiiProperty: "urlReportedMalicious", GoGetter: "UrlReportedMalicious"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflarePageShieldConnectionsListResultOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

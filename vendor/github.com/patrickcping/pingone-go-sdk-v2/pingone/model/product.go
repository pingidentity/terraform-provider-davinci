package model

import (
	"fmt"
	"sort"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"golang.org/x/exp/slices"
)

type ProductMapping struct {
	ProductCode  string
	APICode      management.EnumProductType
	Default      bool
	Selectable   bool
	Deprecated   bool
	DisplayOrder int
}

var productMappingList []ProductMapping

func init() {
	productMappingList = []ProductMapping{
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_BASE,
			ProductCode: "SSO",
			Default:     true,
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_PROVISIONING,
			ProductCode: "Provisioning",
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_DAVINCI,
			ProductCode: "DaVinci",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_ORCHESTRATE,
			ProductCode: "DaVinciLegacy",
			Selectable:  false,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_MFA,
			ProductCode: "MFA",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_RISK,
			ProductCode: "Risk",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_VERIFY,
			ProductCode: "Verify",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_CREDENTIALS,
			ProductCode: "Credentials",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_INTELLIGENCE,
			ProductCode: "APIIntelligence",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_AUTHORIZE,
			ProductCode: "Authorize",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_FRAUD,
			ProductCode: "Fraud",
			Selectable:  false,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ID,
			ProductCode: "PingID",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_FEDERATE,
			ProductCode: "PingFederate",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ACCESS,
			ProductCode: "PingAccess",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_DIRECTORY,
			ProductCode: "PingDirectory",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_AUTHORIZE,
			ProductCode: "PingAuthorize",
			Selectable:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_CENTRAL,
			ProductCode: "PingCentral",
			Selectable:  true,
		},

		{
			APICode:     management.ENUMPRODUCTTYPE_DATA_SYNC,
			ProductCode: "PingDataSync",
			Selectable:  false,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_DATA_GOVERNANCE,
			ProductCode: "PingDataGovernance",
			Selectable:  false,
			Deprecated:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_FOR_ENTERPRISE,
			ProductCode: "PingOneEnterprise",
			Selectable:  false,
			Deprecated:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ID_SDK,
			ProductCode: "PingIDSDK",
			Selectable:  false,
			Deprecated:  true,
		},
		{
			APICode:     management.ENUMPRODUCTTYPE_ONE_FOR_SAAS,
			ProductCode: "PingOneForSaaS",
			Selectable:  false,
			Deprecated:  true,
		},
	}
}

func FindProductByName(product string) (ProductMapping, error) {

	idx := slices.IndexFunc(productMappingList, func(c ProductMapping) bool { return c.ProductCode == product })

	if idx < 0 {
		return ProductMapping{}, fmt.Errorf("Cannot find product from name: %s", product)
	}

	return productMappingList[idx], nil

}

func FindProductByAPICode(apiCode management.EnumProductType) (ProductMapping, error) {

	idx := slices.IndexFunc(productMappingList, func(c ProductMapping) bool { return c.APICode == apiCode })

	if idx < 0 {

		return ProductMapping{}, fmt.Errorf("Cannot find product from api code: %v", apiCode)
	}

	return productMappingList[idx], nil

}

func ProductsSelectableList() []string {

	productList := make([]string, 0)

	for _, product := range productMappingList {

		if product.Selectable {
			productList = append(productList, product.ProductCode)
		}

	}

	sort.Strings(productList)

	return productList
}

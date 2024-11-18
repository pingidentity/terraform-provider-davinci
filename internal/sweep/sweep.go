package sweep

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
)

var (
	EnvironmentNamePrefix = "tf-testacc-dv-"
)

func SweepClient(ctx context.Context) (*Client, error) {

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
		ForceDelete:   true,
	}

	return config.APIClient(ctx, getProviderTestingVersion())

}

func getProviderTestingVersion() string {
	returnVar := "dev"
	if v := os.Getenv("PINGONE_TESTING_PROVIDER_VERSION"); v != "" {
		returnVar = v
	}
	return returnVar
}

func FetchTaggedEnvironments(ctx context.Context, apiClient *management.APIClient) ([]management.Environment, error) {
	return FetchTaggedEnvironmentsByPrefix(ctx, apiClient, EnvironmentNamePrefix)
}

func FetchTaggedEnvironmentsByPrefix(ctx context.Context, apiClient *management.APIClient, prefix string) ([]management.Environment, error) {

	filter := fmt.Sprintf("name sw \"%s\"", prefix)

	resp, diags := ParseResponse(
		ctx,
		func() (any, *http.Response, error) {
			pagedIterator := apiClient.EnvironmentsApi.ReadAllEnvironments(ctx).Filter(filter).Execute()

			returnEnvironments := make([]management.Environment, 0)

			var initialHttpResponse *http.Response

			for pageCursor, err := range pagedIterator {
				if err != nil {
					return nil, pageCursor.HTTPResponse, err
				}

				if initialHttpResponse == nil {
					initialHttpResponse = pageCursor.HTTPResponse
				}

				if environments, ok := pageCursor.EntityArray.Embedded.GetEnvironmentsOk(); ok {

					for _, environment := range environments {
						if environment.GetName() == "Administrators" {
							return nil, nil, fmt.Errorf("Unsafe filter, Administrators environment present: %s", filter)
						}
					}

					returnEnvironments = append(returnEnvironments, environments...)
				}
			}

			return returnEnvironments, initialHttpResponse, nil
		},
		"ReadAllEnvironments",
		CustomErrorResourceNotFoundWarning,
		func(ctx context.Context, r *http.Response, p1error *model.P1Error) bool {

			if p1error != nil {
				var err error

				// Permissions may not have propagated by this point
				m, err := regexp.MatchString("^The request could not be completed. You do not have access to this resource.", p1error.GetMessage())
				if err == nil && m {
					tflog.Warn(ctx, "Insufficient PingOne privileges detected")
					return true
				}
				if err != nil {
					tflog.Warn(ctx, "Cannot match error string for retry")
					return false
				}

			}

			return false
		},
	)
	if diags.HasError() {
		return nil, fmt.Errorf("Error getting environments for sweep")
	}

	respList := resp.([]management.Environment)

	return respList, nil
}

func CreateTestEnvironment(ctx context.Context, apiClient *management.APIClient, region management.EnumRegionCode, index string) error {

	environmentLicense := os.Getenv("PINGONE_LICENSE_ID")

	environment := *management.NewEnvironment(
		*management.NewEnvironmentLicense(environmentLicense),
		fmt.Sprintf("%sdynamic-%s", EnvironmentNamePrefix, index),
		management.EnvironmentRegion{
			EnumRegionCode: &region,
		},
		management.ENUMENVIRONMENTTYPE_SANDBOX,
	)

	productBOMItems := make([]management.BillOfMaterialsProductsInner, 0)

	productBOMItems = append(productBOMItems, *management.NewBillOfMaterialsProductsInner(management.ENUMPRODUCTTYPE_ONE_BASE))
	productBOMItems = append(productBOMItems, *management.NewBillOfMaterialsProductsInner(management.ENUMPRODUCTTYPE_ONE_MFA))
	productBOMItems = append(productBOMItems, *management.NewBillOfMaterialsProductsInner(management.ENUMPRODUCTTYPE_ONE_RISK))
	productBOMItems = append(productBOMItems, *management.NewBillOfMaterialsProductsInner(management.ENUMPRODUCTTYPE_ONE_AUTHORIZE))

	environment.SetBillOfMaterials(*management.NewBillOfMaterials(productBOMItems))

	resp, diags := ParseResponse(
		ctx,

		func() (interface{}, *http.Response, error) {
			return apiClient.EnvironmentsApi.CreateEnvironmentActiveLicense(ctx).Environment(environment).Execute()
		},
		"CreateEnvironmentActiveLicense",
		nil,
		DefaultRetryable,
	)
	if diags.HasError() {
		return fmt.Errorf("Cannot create environment `%s`", environment.GetName())
	}

	environmentID := resp.(*management.Environment).GetId()

	// A population, because we must have one

	population := *management.NewPopulation("Default")

	_, diags = ParseResponse(
		ctx,

		func() (interface{}, *http.Response, error) {
			return apiClient.PopulationsApi.CreatePopulation(ctx, environmentID).Population(population).Execute()
		},
		"CreatePopulation",
		DefaultCustomError,
		DefaultCreateReadRetryable,
	)
	if diags.HasError() {
		return fmt.Errorf("Cannot create population for environment `%s`", environment.GetName())
	}

	return nil

}

// from tf-provider-pingone internal/service/base/sweep.go

func init() {
	fmt.Println("registering sweepers")
	resource.AddTestSweepers("pingone_environment", &resource.Sweeper{
		Name:         "pingone_environment",
		F:            sweepEnvironments,
		Dependencies: []string{
			// "pingone_group",
			// "pingone_population",
		},
	})
}

func sweepEnvironments(region string) error {
	var ctx = context.Background()

	p1Client, err := SweepClient(ctx)

	if err != nil {
		return err
	}

	apiClient := p1Client.API.ManagementAPIClient
	ctx = context.WithValue(ctx, management.ContextServerVariables, map[string]string{
		"suffix": p1Client.API.Region.URLSuffix,
	})

	err = CreateTestEnvironment(ctx, apiClient, p1Client.API.Region.APICode, "general-test")
	if err != nil {
		log.Printf("Error creating environment `general-test` during sweep: %s", err)
	}

	environments, err := FetchTaggedEnvironmentsByPrefix(ctx, apiClient, fmt.Sprintf("%sdynamic-", EnvironmentNamePrefix))
	if err != nil {
		return err
	}

	for _, environment := range environments {
		fmt.Printf("Destroying environment %s\n", environment.GetName())

		// Reset back to sandbox
		if environment.GetType() == "PRODUCTION" {
			updateEnvironmentTypeRequest := *management.NewUpdateEnvironmentTypeRequest()
			updateEnvironmentTypeRequest.SetType("SANDBOX")
			_, _, err := apiClient.EnvironmentsApi.UpdateEnvironmentType(ctx, environment.GetId()).UpdateEnvironmentTypeRequest(updateEnvironmentTypeRequest).Execute()

			if err != nil {
				log.Printf("Error setting environment %s of type PRODUCTION to SANDBOX during sweep: %s", environment.GetName(), err)
			}
		}

		// Delete the environment
		_, err := apiClient.EnvironmentsApi.DeleteEnvironment(ctx, environment.GetId()).Execute()

		if err != nil {
			log.Printf("Error destroying environment %s during sweep: %s", environment.GetName(), err)
		}

	}

	return nil

}

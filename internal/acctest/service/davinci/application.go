package davinci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func Application_CheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var ctx = context.Background()

		c, err := acctest.TestClient()
		if err != nil {
			return err
		}

		p1Client, err := acctest.PingOneTestClient(ctx)
		if err != nil {
			return err
		}

		apiClient := p1Client.API.ManagementAPIClient

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "davinci_application" {
				continue
			}

			shouldContinue, err := acctest.CheckParentEnvironmentDestroy(ctx, apiClient, rs.Primary.Attributes["environment_id"])
			if err != nil {
				return err
			}

			if shouldContinue {
				continue
			}

			companyId := rs.Primary.Attributes["environment_id"]

			_, err = sdk.DoRetryable(
				ctx,
				c,
				companyId,
				func() (interface{}, *http.Response, error) {
					return c.ReadApplicationWithResponse(companyId, rs.Primary.ID)
				},
			)

			if err != nil {
				if dvError, ok := err.(dv.ErrorResponse); ok {
					if dvError.HttpResponseCode == http.StatusNotFound {
						shouldContinue = true
					}
				}
			}

			if shouldContinue {
				continue
			}

			return fmt.Errorf("DaVinci Application Instance %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func Application_GetIDs(resourceName string, environmentID, applicationID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceName)
		}

		*applicationID = rs.Primary.ID
		*environmentID = rs.Primary.Attributes["environment_id"]

		return nil
	}
}

func Application_RemovalDrift_PreConfig(t *testing.T, environmentID, applicationID string) {
	c, err := acctest.TestClient()
	if err != nil {
		t.Fatalf("Failed to get API client: %v", err)
	}

	if environmentID == "" || applicationID == "" {
		t.Fatalf("One of environment ID or application ID cannot be determined. Environment ID: %s, Application ID: %s", environmentID, applicationID)
	}

	var ctx = context.Background()

	_, err = sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteApplicationWithResponse(environmentID, applicationID)
		},
	)
	if err != nil {
		t.Fatalf("Failed to delete application: %v", err)
	}
}

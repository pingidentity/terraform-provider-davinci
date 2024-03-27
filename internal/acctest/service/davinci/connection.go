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

func Connection_CheckDestroy() resource.TestCheckFunc {
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
			if rs.Type != "davinci_connection" {
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
					return c.ReadConnectionWithResponse(companyId, rs.Primary.ID)
				},
			)

			if err != nil {
				if dvError, ok := err.(dv.ErrorResponse); ok {
					if dvError.HttpResponseCode == http.StatusNotFound || dvError.Code == dv.DV_ERROR_CODE_CONNECTION_NOT_FOUND {
						shouldContinue = true
					}
				}
			}

			if shouldContinue {
				continue
			}

			return fmt.Errorf("DaVinci Connection Instance %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func Connection_GetIDs(resourceName string, environmentID, connectionID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceName)
		}

		*connectionID = rs.Primary.ID
		*environmentID = rs.Primary.Attributes["environment_id"]

		return nil
	}
}

func Connection_RemovalDrift_PreConfig(t *testing.T, environmentID, connectionID string) {
	c, err := acctest.TestClient()
	if err != nil {
		t.Fatalf("Failed to get API client: %v", err)
	}

	if environmentID == "" || connectionID == "" {
		t.Fatalf("One of environment ID or connection ID cannot be determined. Environment ID: %s, Connection ID: %s", environmentID, connectionID)
	}

	var ctx = context.Background()

	_, err = sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteConnectionWithResponse(environmentID, connectionID)
		},
	)
	if err != nil {
		t.Fatalf("Failed to delete connection: %v", err)
	}
}

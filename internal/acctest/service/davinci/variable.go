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

func Variable_CheckDestroy() resource.TestCheckFunc {
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
			if rs.Type != "davinci_variable" {
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
					return c.ReadVariableWithResponse(&companyId, rs.Primary.ID)
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

			return fmt.Errorf("DaVinci Variable Instance %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func Variable_GetIDs(resourceName string, environmentID, variableID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceName)
		}

		*variableID = rs.Primary.ID
		*environmentID = rs.Primary.Attributes["environment_id"]

		return nil
	}
}

func Variable_RemovalDrift_PreConfig(t *testing.T, environmentID, variableID string) {
	c, err := acctest.TestClient()
	if err != nil {
		t.Fatalf("Failed to get API client: %v", err)
	}

	if environmentID == "" || variableID == "" {
		t.Fatalf("One of environment ID or variable ID cannot be determined. Environment ID: %s, Variable ID: %s", environmentID, variableID)
	}

	var ctx = context.Background()

	_, err = sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteVariableWithResponse(&environmentID, variableID)
		},
	)
	if err != nil {
		t.Fatalf("Failed to delete variable: %v", err)
	}
}

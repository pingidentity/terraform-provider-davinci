package davinci

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/samir-gandhi/davinci-client-go/davinci"
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

			_, err = c.ReadConnection(&companyId, rs.Primary.ID)

			if err != nil {
				// currently a 400 (rather than 404) is returned if the connection is not found.
				// The comparison is made to match the entire error message to avoid false positives
				if strings.Contains(err.Error(), "status: 400, body: {\"cause\":null,\"logLevel\":\"error\",\"serviceName\":null,\"message\":\"Error retrieving connectors\",\"errorMessage\":\"Error retrieving connectors\",\"success\":false,\"httpResponseCode\":400,\"code\":7005}") {
					shouldContinue = true
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

func Connection_RemovalDrift_PreConfig(c *davinci.APIClient, t *testing.T, environmentID, connectionID string) {
	if environmentID == "" || connectionID == "" {
		t.Fatalf("One of environment ID or connection ID cannot be determined. Environment ID: %s, Connection ID: %s", environmentID, connectionID)
	}

	_, err := c.DeleteConnection(&environmentID, connectionID)
	if err != nil {
		t.Fatalf("Failed to delete connection: %v", err)
	}
}

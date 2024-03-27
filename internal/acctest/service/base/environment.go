package base

import (
	"context"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

func Environment_RemovalDrift_PreConfig(ctx context.Context, apiClient *management.APIClient, t *testing.T, environmentID string) {
	if environmentID == "" {
		t.Fatalf("Environment ID cannot be determined. Environment ID: %s", environmentID)
	}

	_, err := apiClient.EnvironmentsApi.DeleteEnvironment(ctx, environmentID).Execute()
	if err != nil {
		t.Fatalf("Failed to delete environment: %v", err)
	}
}

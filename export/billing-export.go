package export

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-10-01/costmanagement"
	"github.com/Azure/go-autorest/autorest"
	"github.com/pjirsa/azure-costmanagement-samples/internal/config"
	"github.com/pjirsa/azure-costmanagement-samples/internal/iam"
)

func getCostClient() costmanagement.ExportsClient {
	costClient := costmanagement.NewExportsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	costClient.Authorizer = a
	costClient.AddToUserAgent(config.UserAgent())
	return costClient
}

func List(ctx context.Context, scope string) (costmanagement.ExportListResult, error) {
	costClient := getCostClient()
	result, err := costClient.List(ctx, scope)

	if err != nil {
		return result, fmt.Errorf("An error has happened")
	}

	return result, err
}

func Get(ctx context.Context, scope string, exportName string) (costmanagement.Export, error) {
	costClient := getCostClient()

	result, err := costClient.Get(ctx, scope, exportName)

	if err != nil {
		return result, fmt.Errorf("Error getting export")
	}

	return result, err
}

func CreateOrUpdate(ctx context.Context, scope string, exportName string, parameters costmanagement.Export) (costmanagement.Export, error) {
	costClient := getCostClient()

	result, err := costClient.CreateOrUpdate(ctx, scope, exportName, parameters)

	if err != nil {
		return result, fmt.Errorf("Error creating export")
	}

	return result, err
}

func Delete(ctx context.Context, scope string, exportName string) (autorest.Response, error) {
	costClient := getCostClient()

	result, err := costClient.Delete(ctx, scope, exportName)

	if err != nil {
		return result, fmt.Errorf("Error deleting export")
	}

	return result, err
}

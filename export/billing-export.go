package export

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-10-01/costmanagement"
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

func List(ctx context.Context, scope string) (export costmanagement.ExportListResult, err error) {
	costClient := getCostClient()
	result, err := costClient.List(ctx, scope)

	if err != nil {
		return export, fmt.Errorf("An error has happened")
	}

	return result, err
}

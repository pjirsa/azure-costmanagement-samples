package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pjirsa/azure-costmanagement-samples/export"
	"github.com/pjirsa/azure-costmanagement-samples/internal/config"
)

func main() {
	err := config.ParseEnvironment()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()

	ListExports(ctx)
}

func ListExports(ctx context.Context) {
	//scope := "subscriptions/" + config.SubscriptionID()
	scope := fmt.Sprintf("/providers/Microsoft.Management/managementGroups/%s/", config.ManagementGroupId())

	resp, err := export.List(ctx, scope)

	if err != nil {
		fmt.Println("An error happened")
	}

	result, _ := json.Marshal(resp.Value)

	fmt.Println("List of configured exports:")
	fmt.Println(string(result))
}

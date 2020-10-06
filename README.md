---
languages:
- go
products:
- azure
page_type: sample
description: "A collection of samples showing how to use the Azure SDK for Go."
---

# Azure SDK for Go Samples - Consumption Management Edition 🐱‍💻

azure-sdk-for-go-samples is a collection of sample usages of the [Azure/azure-sdk-for-go][].

For general SDK help start with the [main SDK README][].

## To run tests

1. set up authentication (see following)
1. `go run ./test-exports.go`

To use service principal authentication, create a principal by running `az ad sp create-for-rbac -n "<yourAppName>"` and set the following environment variables. You can copy `.env.tpl` to a `.env` file in each package for ease of use.

```bash
export AZURE_SUBSCRIPTION_ID=
export AZURE_TENANT_ID=
export AZURE_CLIENT_ID=
export AZURE_CLIENT_SECRET=

export AZURE_LOCATION_DEFAULT=westus2
export AZURE_BASE_GROUP_NAME=azure-samples-go
export AZURE_KEEP_SAMPLE_RESOURCES=0
```

## Resources

- SDK code is at [Azure/azure-sdk-for-go][].
- SDK docs are at [godoc.org](https://godoc.org/github.com/Azure/azure-sdk-for-go/).
- SDK notifications are published via the [Azure update feed][].
- Azure API docs are at [docs.microsoft.com/rest/api](https://docs.microsoft.com/rest/api/).
- General Azure docs are at [docs.microsoft.com/azure](https://docs.microsoft.com/azure).

## License

This code is provided under the MIT license. See [LICENSE][] for details.
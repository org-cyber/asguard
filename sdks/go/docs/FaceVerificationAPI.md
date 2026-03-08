# \FaceVerificationAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeFace**](FaceVerificationAPI.md#AnalyzeFace) | **Post** /v1/analyze | Analyze face image quality and extract embedding
[**CompareFaces**](FaceVerificationAPI.md#CompareFaces) | **Post** /v1/compare | Compare a probe face image with a reference embedding



## AnalyzeFace

> AnalyzeResponse AnalyzeFace(ctx).AnalyzeFaceRequest(analyzeFaceRequest).Execute()

Analyze face image quality and extract embedding

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	analyzeFaceRequest := *openapiclient.NewAnalyzeFaceRequest("Image_example") // AnalyzeFaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FaceVerificationAPI.AnalyzeFace(context.Background()).AnalyzeFaceRequest(analyzeFaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaceVerificationAPI.AnalyzeFace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyzeFace`: AnalyzeResponse
	fmt.Fprintf(os.Stdout, "Response from `FaceVerificationAPI.AnalyzeFace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeFaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyzeFaceRequest** | [**AnalyzeFaceRequest**](AnalyzeFaceRequest.md) |  | 

### Return type

[**AnalyzeResponse**](AnalyzeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareFaces

> CompareResponse CompareFaces(ctx).CompareFacesRequest(compareFacesRequest).Execute()

Compare a probe face image with a reference embedding

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	compareFacesRequest := *openapiclient.NewCompareFacesRequest("ProbeImage_example", []float32{float32(123)}) // CompareFacesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FaceVerificationAPI.CompareFaces(context.Background()).CompareFacesRequest(compareFacesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaceVerificationAPI.CompareFaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompareFaces`: CompareResponse
	fmt.Fprintf(os.Stdout, "Response from `FaceVerificationAPI.CompareFaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompareFacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compareFacesRequest** | [**CompareFacesRequest**](CompareFacesRequest.md) |  | 

### Return type

[**CompareResponse**](CompareResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


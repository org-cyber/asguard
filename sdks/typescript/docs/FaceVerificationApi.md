# FaceVerificationApi

All URIs are relative to *http://localhost:8081*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**analyzeFace**](#analyzeface) | **POST** /v1/analyze | Analyze face image quality and extract embedding|
|[**compareFaces**](#comparefaces) | **POST** /v1/compare | Compare a probe face image with a reference embedding|

# **analyzeFace**
> AnalyzeResponse analyzeFace(analyzeFaceRequest)


### Example

```typescript
import {
    FaceVerificationApi,
    Configuration,
    AnalyzeFaceRequest
} from 'asguard';

const configuration = new Configuration();
const apiInstance = new FaceVerificationApi(configuration);

let analyzeFaceRequest: AnalyzeFaceRequest; //

const { status, data } = await apiInstance.analyzeFace(
    analyzeFaceRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **analyzeFaceRequest** | **AnalyzeFaceRequest**|  | |


### Return type

**AnalyzeResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Face analyzed successfully |  -  |
|**400** | Invalid image or no face detected |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **compareFaces**
> CompareResponse compareFaces(compareFacesRequest)


### Example

```typescript
import {
    FaceVerificationApi,
    Configuration,
    CompareFacesRequest
} from 'asguard';

const configuration = new Configuration();
const apiInstance = new FaceVerificationApi(configuration);

let compareFacesRequest: CompareFacesRequest; //

const { status, data } = await apiInstance.compareFaces(
    compareFacesRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **compareFacesRequest** | **CompareFacesRequest**|  | |


### Return type

**CompareResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Comparison completed |  -  |
|**400** | Invalid input |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


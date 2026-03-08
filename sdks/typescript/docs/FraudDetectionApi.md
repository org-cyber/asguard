# FraudDetectionApi

All URIs are relative to *http://localhost:8081*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**checkFraud**](#checkfraud) | **POST** /analyze | Check transaction for fraud risk|

# **checkFraud**
> FraudCheckResponse checkFraud(fraudCheckRequest)


### Example

```typescript
import {
    FraudDetectionApi,
    Configuration,
    FraudCheckRequest
} from 'asguard';

const configuration = new Configuration();
const apiInstance = new FraudDetectionApi(configuration);

let fraudCheckRequest: FraudCheckRequest; //

const { status, data } = await apiInstance.checkFraud(
    fraudCheckRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **fraudCheckRequest** | **FraudCheckRequest**|  | |


### Return type

**FraudCheckResponse**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Risk assessment completed |  -  |
|**400** | Invalid request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# SystemApi

All URIs are relative to *http://localhost:8081*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**healthCheckFace**](#healthcheckface) | **GET** /face/health | Face service health check|
|[**healthCheckFraud**](#healthcheckfraud) | **GET** /fraud/health | Fraud service health check|

# **healthCheckFace**
> HealthResponse healthCheckFace()


### Example

```typescript
import {
    SystemApi,
    Configuration
} from 'asguard';

const configuration = new Configuration();
const apiInstance = new SystemApi(configuration);

const { status, data } = await apiInstance.healthCheckFace();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**HealthResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Service status |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **healthCheckFraud**
> HealthResponse healthCheckFraud()


### Example

```typescript
import {
    SystemApi,
    Configuration
} from 'asguard';

const configuration = new Configuration();
const apiInstance = new SystemApi(configuration);

const { status, data } = await apiInstance.healthCheckFraud();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**HealthResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Service status |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


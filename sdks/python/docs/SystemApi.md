# asguard.SystemApi

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**health_check_face**](SystemApi.md#health_check_face) | **GET** /face/health | Face service health check
[**health_check_fraud**](SystemApi.md#health_check_fraud) | **GET** /fraud/health | Fraud service health check


# **health_check_face**
> HealthResponse health_check_face()

Face service health check

### Example


```python
import asguard
from asguard.models.health_response import HealthResponse
from asguard.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8081
# See configuration.py for a list of all supported configuration parameters.
configuration = asguard.Configuration(
    host = "http://localhost:8081"
)


# Enter a context with an instance of the API client
with asguard.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = asguard.SystemApi(api_client)

    try:
        # Face service health check
        api_response = api_instance.health_check_face()
        print("The response of SystemApi->health_check_face:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SystemApi->health_check_face: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Service status |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **health_check_fraud**
> HealthResponse health_check_fraud()

Fraud service health check

### Example


```python
import asguard
from asguard.models.health_response import HealthResponse
from asguard.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8081
# See configuration.py for a list of all supported configuration parameters.
configuration = asguard.Configuration(
    host = "http://localhost:8081"
)


# Enter a context with an instance of the API client
with asguard.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = asguard.SystemApi(api_client)

    try:
        # Fraud service health check
        api_response = api_instance.health_check_fraud()
        print("The response of SystemApi->health_check_fraud:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SystemApi->health_check_fraud: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Service status |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


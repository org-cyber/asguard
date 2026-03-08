# asguard.FaceVerificationApi

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**analyze_face**](FaceVerificationApi.md#analyze_face) | **POST** /v1/analyze | Analyze face image quality and extract embedding
[**compare_faces**](FaceVerificationApi.md#compare_faces) | **POST** /v1/compare | Compare a probe face image with a reference embedding


# **analyze_face**
> AnalyzeResponse analyze_face(analyze_face_request)

Analyze face image quality and extract embedding

### Example

* Bearer (API Key or JWT) Authentication (BearerAuth):

```python
import asguard
from asguard.models.analyze_face_request import AnalyzeFaceRequest
from asguard.models.analyze_response import AnalyzeResponse
from asguard.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8081
# See configuration.py for a list of all supported configuration parameters.
configuration = asguard.Configuration(
    host = "http://localhost:8081"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (API Key or JWT): BearerAuth
configuration = asguard.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with asguard.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = asguard.FaceVerificationApi(api_client)
    analyze_face_request = asguard.AnalyzeFaceRequest() # AnalyzeFaceRequest | 

    try:
        # Analyze face image quality and extract embedding
        api_response = api_instance.analyze_face(analyze_face_request)
        print("The response of FaceVerificationApi->analyze_face:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FaceVerificationApi->analyze_face: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyze_face_request** | [**AnalyzeFaceRequest**](AnalyzeFaceRequest.md)|  | 

### Return type

[**AnalyzeResponse**](AnalyzeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Face analyzed successfully |  -  |
**400** | Invalid image or no face detected |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **compare_faces**
> CompareResponse compare_faces(compare_faces_request)

Compare a probe face image with a reference embedding

### Example

* Bearer (API Key or JWT) Authentication (BearerAuth):

```python
import asguard
from asguard.models.compare_faces_request import CompareFacesRequest
from asguard.models.compare_response import CompareResponse
from asguard.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8081
# See configuration.py for a list of all supported configuration parameters.
configuration = asguard.Configuration(
    host = "http://localhost:8081"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (API Key or JWT): BearerAuth
configuration = asguard.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with asguard.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = asguard.FaceVerificationApi(api_client)
    compare_faces_request = asguard.CompareFacesRequest() # CompareFacesRequest | 

    try:
        # Compare a probe face image with a reference embedding
        api_response = api_instance.compare_faces(compare_faces_request)
        print("The response of FaceVerificationApi->compare_faces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FaceVerificationApi->compare_faces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compare_faces_request** | [**CompareFacesRequest**](CompareFacesRequest.md)|  | 

### Return type

[**CompareResponse**](CompareResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Comparison completed |  -  |
**400** | Invalid input |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


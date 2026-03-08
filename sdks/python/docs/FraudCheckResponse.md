# FraudCheckResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**transaction_id** | **str** |  | 
**risk_score** | **float** |  | 
**risk_level** | **str** |  | 
**reasons** | **List[str]** |  | [optional] 
**ai_triggered** | **bool** |  | [optional] 
**ai_confidence** | **float** |  | [optional] 
**ai_recommendation** | **str** |  | [optional] 
**ai_fraud_probability** | **float** |  | [optional] 
**ai_summary** | **str** |  | [optional] 
**message** | **str** |  | [optional] 
**processing_time_ms** | **int** | Request processing time in milliseconds | [optional] 

## Example

```python
from asguard.models.fraud_check_response import FraudCheckResponse

# TODO update the JSON string below
json = "{}"
# create an instance of FraudCheckResponse from a JSON string
fraud_check_response_instance = FraudCheckResponse.from_json(json)
# print the JSON string representation of the object
print(FraudCheckResponse.to_json())

# convert the object into a dict
fraud_check_response_dict = fraud_check_response_instance.to_dict()
# create an instance of FraudCheckResponse from a dict
fraud_check_response_from_dict = FraudCheckResponse.from_dict(fraud_check_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



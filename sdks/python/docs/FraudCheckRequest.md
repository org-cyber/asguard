# FraudCheckRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**transaction_id** | **str** | Unique transaction identifier | 
**amount** | **float** | Transaction amount | 
**currency** | **str** | ISO 4217 currency code (USD, EUR, etc.) | 
**user_id** | **str** | Customer identifier | 
**device_id** | **str** | Device fingerprint | [optional] 
**ip_address** | **str** | Customer IP address | [optional] 
**location** | **str** | Geo location (city, country) | [optional] 
**timestamp** | **datetime** | Transaction timestamp | [optional] 

## Example

```python
from asguard.models.fraud_check_request import FraudCheckRequest

# TODO update the JSON string below
json = "{}"
# create an instance of FraudCheckRequest from a JSON string
fraud_check_request_instance = FraudCheckRequest.from_json(json)
# print the JSON string representation of the object
print(FraudCheckRequest.to_json())

# convert the object into a dict
fraud_check_request_dict = fraud_check_request_instance.to_dict()
# create an instance of FraudCheckRequest from a dict
fraud_check_request_from_dict = FraudCheckRequest.from_dict(fraud_check_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



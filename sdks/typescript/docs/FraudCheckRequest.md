# FraudCheckRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**transaction_id** | **string** | Unique transaction identifier | [default to undefined]
**amount** | **number** | Transaction amount | [default to undefined]
**currency** | **string** | ISO 4217 currency code (USD, EUR, etc.) | [default to undefined]
**user_id** | **string** | Customer identifier | [default to undefined]
**device_id** | **string** | Device fingerprint | [optional] [default to undefined]
**ip_address** | **string** | Customer IP address | [optional] [default to undefined]
**location** | **string** | Geo location (city, country) | [optional] [default to undefined]
**timestamp** | **string** | Transaction timestamp | [optional] [default to undefined]

## Example

```typescript
import { FraudCheckRequest } from 'asguard';

const instance: FraudCheckRequest = {
    transaction_id,
    amount,
    currency,
    user_id,
    device_id,
    ip_address,
    location,
    timestamp,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

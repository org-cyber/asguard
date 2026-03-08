# FraudCheckResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**transaction_id** | **string** |  | [default to undefined]
**risk_score** | **number** |  | [default to undefined]
**risk_level** | **string** |  | [default to undefined]
**reasons** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**ai_triggered** | **boolean** |  | [optional] [default to undefined]
**ai_confidence** | **number** |  | [optional] [default to undefined]
**ai_recommendation** | **string** |  | [optional] [default to undefined]
**ai_fraud_probability** | **number** |  | [optional] [default to undefined]
**ai_summary** | **string** |  | [optional] [default to undefined]
**message** | **string** |  | [optional] [default to undefined]
**processing_time_ms** | **number** | Request processing time in milliseconds | [optional] [default to undefined]

## Example

```typescript
import { FraudCheckResponse } from 'asguard';

const instance: FraudCheckResponse = {
    transaction_id,
    risk_score,
    risk_level,
    reasons,
    ai_triggered,
    ai_confidence,
    ai_recommendation,
    ai_fraud_probability,
    ai_summary,
    message,
    processing_time_ms,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

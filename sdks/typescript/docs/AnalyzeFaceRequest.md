# AnalyzeFaceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image** | **string** | Base64-encoded image (with optional data URI prefix) | [default to undefined]
**quality_checks** | **boolean** | Enable quality checks (sharpness, brightness, face size) | [optional] [default to false]

## Example

```typescript
import { AnalyzeFaceRequest } from 'asguard';

const instance: AnalyzeFaceRequest = {
    image,
    quality_checks,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# CompareResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **boolean** |  | [default to undefined]
**match** | **boolean** | Whether faces match (distance &lt; threshold_used) | [default to undefined]
**confidence** | **number** | Match confidence score (0-1) | [default to undefined]
**distance** | **number** | Euclidean distance between embeddings | [default to undefined]
**threshold_used** | **number** | The distance threshold used for this comparison | [default to undefined]
**probe_quality** | **number** | Quality score of the probe image (0-1) | [optional] [default to undefined]
**processing_time_ms** | **number** |  | [optional] [default to undefined]
**error** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { CompareResponse } from 'asguard';

const instance: CompareResponse = {
    success,
    match,
    confidence,
    distance,
    threshold_used,
    probe_quality,
    processing_time_ms,
    error,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# AnalyzeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **boolean** | Whether analysis succeeded | [default to undefined]
**face_detected** | **boolean** | Whether a face was found | [default to undefined]
**embedding** | **Array&lt;number&gt;** | 128-dimensional face embedding vector (if face detected) | [optional] [default to undefined]
**quality_score** | **number** | Overall quality score (0-1) | [optional] [default to undefined]
**sharpness** | **number** | Normalized sharpness score (0-1) | [optional] [default to undefined]
**brightness** | **number** | Normalized brightness score (0-1) | [optional] [default to undefined]
**face_size_ratio** | **number** | Face area relative to image (0-1) | [optional] [default to undefined]
**warnings** | **Array&lt;string&gt;** | List of quality warnings (e.g., \&quot;too_dark\&quot;, \&quot;too_blurry\&quot;) | [optional] [default to undefined]
**processing_time_ms** | **number** | Request processing time in milliseconds | [optional] [default to undefined]
**error** | **string** | Error message if success&#x3D;false | [optional] [default to undefined]

## Example

```typescript
import { AnalyzeResponse } from 'asguard';

const instance: AnalyzeResponse = {
    success,
    face_detected,
    embedding,
    quality_score,
    sharpness,
    brightness,
    face_size_ratio,
    warnings,
    processing_time_ms,
    error,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

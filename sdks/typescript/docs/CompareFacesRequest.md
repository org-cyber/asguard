# CompareFacesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**probe_image** | **string** | Base64-encoded image (with optional data URI prefix) | [default to undefined]
**reference_embedding** | **Array&lt;number&gt;** | 128-dimensional embedding of the reference face | [default to undefined]
**threshold** | **number** | Distance threshold for matching (lower &#x3D; stricter) | [optional] [default to 0.6]

## Example

```typescript
import { CompareFacesRequest } from 'asguard';

const instance: CompareFacesRequest = {
    probe_image,
    reference_embedding,
    threshold,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# CompareFacesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**probe_image** | **str** | Base64-encoded image (with optional data URI prefix) | 
**reference_embedding** | **List[float]** | 128-dimensional embedding of the reference face | 
**threshold** | **float** | Distance threshold for matching (lower &#x3D; stricter) | [optional] [default to 0.6]

## Example

```python
from asguard.models.compare_faces_request import CompareFacesRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CompareFacesRequest from a JSON string
compare_faces_request_instance = CompareFacesRequest.from_json(json)
# print the JSON string representation of the object
print(CompareFacesRequest.to_json())

# convert the object into a dict
compare_faces_request_dict = compare_faces_request_instance.to_dict()
# create an instance of CompareFacesRequest from a dict
compare_faces_request_from_dict = CompareFacesRequest.from_dict(compare_faces_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



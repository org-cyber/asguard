# AnalyzeFaceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image** | **str** | Base64-encoded image (with optional data URI prefix) | 
**quality_checks** | **bool** | Enable quality checks (sharpness, brightness, face size) | [optional] [default to False]

## Example

```python
from asguard.models.analyze_face_request import AnalyzeFaceRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AnalyzeFaceRequest from a JSON string
analyze_face_request_instance = AnalyzeFaceRequest.from_json(json)
# print the JSON string representation of the object
print(AnalyzeFaceRequest.to_json())

# convert the object into a dict
analyze_face_request_dict = analyze_face_request_instance.to_dict()
# create an instance of AnalyzeFaceRequest from a dict
analyze_face_request_from_dict = AnalyzeFaceRequest.from_dict(analyze_face_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



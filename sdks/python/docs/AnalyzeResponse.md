# AnalyzeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** | Whether analysis succeeded | 
**face_detected** | **bool** | Whether a face was found | 
**embedding** | **List[float]** | 128-dimensional face embedding vector (if face detected) | [optional] 
**quality_score** | **float** | Overall quality score (0-1) | [optional] 
**sharpness** | **float** | Normalized sharpness score (0-1) | [optional] 
**brightness** | **float** | Normalized brightness score (0-1) | [optional] 
**face_size_ratio** | **float** | Face area relative to image (0-1) | [optional] 
**warnings** | **List[str]** | List of quality warnings (e.g., \&quot;too_dark\&quot;, \&quot;too_blurry\&quot;) | [optional] 
**processing_time_ms** | **int** | Request processing time in milliseconds | [optional] 
**error** | **str** | Error message if success&#x3D;false | [optional] 

## Example

```python
from asguard.models.analyze_response import AnalyzeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AnalyzeResponse from a JSON string
analyze_response_instance = AnalyzeResponse.from_json(json)
# print the JSON string representation of the object
print(AnalyzeResponse.to_json())

# convert the object into a dict
analyze_response_dict = analyze_response_instance.to_dict()
# create an instance of AnalyzeResponse from a dict
analyze_response_from_dict = AnalyzeResponse.from_dict(analyze_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



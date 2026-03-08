# CompareResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | 
**match** | **bool** | Whether faces match (distance &lt; threshold_used) | 
**confidence** | **float** | Match confidence score (0-1) | 
**distance** | **float** | Euclidean distance between embeddings | 
**threshold_used** | **float** | The distance threshold used for this comparison | 
**probe_quality** | **float** | Quality score of the probe image (0-1) | [optional] 
**processing_time_ms** | **int** |  | [optional] 
**error** | **str** |  | [optional] 

## Example

```python
from asguard.models.compare_response import CompareResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CompareResponse from a JSON string
compare_response_instance = CompareResponse.from_json(json)
# print the JSON string representation of the object
print(CompareResponse.to_json())

# convert the object into a dict
compare_response_dict = compare_response_instance.to_dict()
# create an instance of CompareResponse from a dict
compare_response_from_dict = CompareResponse.from_dict(compare_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



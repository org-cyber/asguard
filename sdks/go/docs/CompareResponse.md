# CompareResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Match** | **bool** | Whether faces match (distance &lt; threshold_used) | 
**Confidence** | **float32** | Match confidence score (0-1) | 
**Distance** | **float32** | Euclidean distance between embeddings | 
**ThresholdUsed** | **float32** | The distance threshold used for this comparison | 
**ProbeQuality** | Pointer to **float32** | Quality score of the probe image (0-1) | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewCompareResponse

`func NewCompareResponse(success bool, match bool, confidence float32, distance float32, thresholdUsed float32, ) *CompareResponse`

NewCompareResponse instantiates a new CompareResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareResponseWithDefaults

`func NewCompareResponseWithDefaults() *CompareResponse`

NewCompareResponseWithDefaults instantiates a new CompareResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CompareResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CompareResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CompareResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMatch

`func (o *CompareResponse) GetMatch() bool`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CompareResponse) GetMatchOk() (*bool, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CompareResponse) SetMatch(v bool)`

SetMatch sets Match field to given value.


### GetConfidence

`func (o *CompareResponse) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CompareResponse) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CompareResponse) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetDistance

`func (o *CompareResponse) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CompareResponse) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CompareResponse) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetThresholdUsed

`func (o *CompareResponse) GetThresholdUsed() float32`

GetThresholdUsed returns the ThresholdUsed field if non-nil, zero value otherwise.

### GetThresholdUsedOk

`func (o *CompareResponse) GetThresholdUsedOk() (*float32, bool)`

GetThresholdUsedOk returns a tuple with the ThresholdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUsed

`func (o *CompareResponse) SetThresholdUsed(v float32)`

SetThresholdUsed sets ThresholdUsed field to given value.


### GetProbeQuality

`func (o *CompareResponse) GetProbeQuality() float32`

GetProbeQuality returns the ProbeQuality field if non-nil, zero value otherwise.

### GetProbeQualityOk

`func (o *CompareResponse) GetProbeQualityOk() (*float32, bool)`

GetProbeQualityOk returns a tuple with the ProbeQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeQuality

`func (o *CompareResponse) SetProbeQuality(v float32)`

SetProbeQuality sets ProbeQuality field to given value.

### HasProbeQuality

`func (o *CompareResponse) HasProbeQuality() bool`

HasProbeQuality returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *CompareResponse) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *CompareResponse) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *CompareResponse) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *CompareResponse) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetError

`func (o *CompareResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CompareResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CompareResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CompareResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AnalyzeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether analysis succeeded | 
**FaceDetected** | **bool** | Whether a face was found | 
**Embedding** | Pointer to **[]float32** | 128-dimensional face embedding vector (if face detected) | [optional] 
**QualityScore** | Pointer to **float32** | Overall quality score (0-1) | [optional] 
**Sharpness** | Pointer to **float32** | Normalized sharpness score (0-1) | [optional] 
**Brightness** | Pointer to **float32** | Normalized brightness score (0-1) | [optional] 
**FaceSizeRatio** | Pointer to **float32** | Face area relative to image (0-1) | [optional] 
**Warnings** | Pointer to **[]string** | List of quality warnings (e.g., \&quot;too_dark\&quot;, \&quot;too_blurry\&quot;) | [optional] 
**ProcessingTimeMs** | Pointer to **int32** | Request processing time in milliseconds | [optional] 
**Error** | Pointer to **string** | Error message if success&#x3D;false | [optional] 

## Methods

### NewAnalyzeResponse

`func NewAnalyzeResponse(success bool, faceDetected bool, ) *AnalyzeResponse`

NewAnalyzeResponse instantiates a new AnalyzeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeResponseWithDefaults

`func NewAnalyzeResponseWithDefaults() *AnalyzeResponse`

NewAnalyzeResponseWithDefaults instantiates a new AnalyzeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AnalyzeResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AnalyzeResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AnalyzeResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetFaceDetected

`func (o *AnalyzeResponse) GetFaceDetected() bool`

GetFaceDetected returns the FaceDetected field if non-nil, zero value otherwise.

### GetFaceDetectedOk

`func (o *AnalyzeResponse) GetFaceDetectedOk() (*bool, bool)`

GetFaceDetectedOk returns a tuple with the FaceDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceDetected

`func (o *AnalyzeResponse) SetFaceDetected(v bool)`

SetFaceDetected sets FaceDetected field to given value.


### GetEmbedding

`func (o *AnalyzeResponse) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *AnalyzeResponse) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *AnalyzeResponse) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.

### HasEmbedding

`func (o *AnalyzeResponse) HasEmbedding() bool`

HasEmbedding returns a boolean if a field has been set.

### GetQualityScore

`func (o *AnalyzeResponse) GetQualityScore() float32`

GetQualityScore returns the QualityScore field if non-nil, zero value otherwise.

### GetQualityScoreOk

`func (o *AnalyzeResponse) GetQualityScoreOk() (*float32, bool)`

GetQualityScoreOk returns a tuple with the QualityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScore

`func (o *AnalyzeResponse) SetQualityScore(v float32)`

SetQualityScore sets QualityScore field to given value.

### HasQualityScore

`func (o *AnalyzeResponse) HasQualityScore() bool`

HasQualityScore returns a boolean if a field has been set.

### GetSharpness

`func (o *AnalyzeResponse) GetSharpness() float32`

GetSharpness returns the Sharpness field if non-nil, zero value otherwise.

### GetSharpnessOk

`func (o *AnalyzeResponse) GetSharpnessOk() (*float32, bool)`

GetSharpnessOk returns a tuple with the Sharpness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharpness

`func (o *AnalyzeResponse) SetSharpness(v float32)`

SetSharpness sets Sharpness field to given value.

### HasSharpness

`func (o *AnalyzeResponse) HasSharpness() bool`

HasSharpness returns a boolean if a field has been set.

### GetBrightness

`func (o *AnalyzeResponse) GetBrightness() float32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *AnalyzeResponse) GetBrightnessOk() (*float32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *AnalyzeResponse) SetBrightness(v float32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *AnalyzeResponse) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### GetFaceSizeRatio

`func (o *AnalyzeResponse) GetFaceSizeRatio() float32`

GetFaceSizeRatio returns the FaceSizeRatio field if non-nil, zero value otherwise.

### GetFaceSizeRatioOk

`func (o *AnalyzeResponse) GetFaceSizeRatioOk() (*float32, bool)`

GetFaceSizeRatioOk returns a tuple with the FaceSizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceSizeRatio

`func (o *AnalyzeResponse) SetFaceSizeRatio(v float32)`

SetFaceSizeRatio sets FaceSizeRatio field to given value.

### HasFaceSizeRatio

`func (o *AnalyzeResponse) HasFaceSizeRatio() bool`

HasFaceSizeRatio returns a boolean if a field has been set.

### GetWarnings

`func (o *AnalyzeResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AnalyzeResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AnalyzeResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AnalyzeResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *AnalyzeResponse) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *AnalyzeResponse) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *AnalyzeResponse) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *AnalyzeResponse) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetError

`func (o *AnalyzeResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalyzeResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalyzeResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AnalyzeResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



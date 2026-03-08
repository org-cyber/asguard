# AnalyzeFaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Base64-encoded image (with optional data URI prefix) | 
**QualityChecks** | Pointer to **bool** | Enable quality checks (sharpness, brightness, face size) | [optional] [default to false]

## Methods

### NewAnalyzeFaceRequest

`func NewAnalyzeFaceRequest(image string, ) *AnalyzeFaceRequest`

NewAnalyzeFaceRequest instantiates a new AnalyzeFaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeFaceRequestWithDefaults

`func NewAnalyzeFaceRequestWithDefaults() *AnalyzeFaceRequest`

NewAnalyzeFaceRequestWithDefaults instantiates a new AnalyzeFaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *AnalyzeFaceRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AnalyzeFaceRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AnalyzeFaceRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetQualityChecks

`func (o *AnalyzeFaceRequest) GetQualityChecks() bool`

GetQualityChecks returns the QualityChecks field if non-nil, zero value otherwise.

### GetQualityChecksOk

`func (o *AnalyzeFaceRequest) GetQualityChecksOk() (*bool, bool)`

GetQualityChecksOk returns a tuple with the QualityChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityChecks

`func (o *AnalyzeFaceRequest) SetQualityChecks(v bool)`

SetQualityChecks sets QualityChecks field to given value.

### HasQualityChecks

`func (o *AnalyzeFaceRequest) HasQualityChecks() bool`

HasQualityChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



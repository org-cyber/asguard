# CompareFacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProbeImage** | **string** | Base64-encoded image (with optional data URI prefix) | 
**ReferenceEmbedding** | **[]float32** | 128-dimensional embedding of the reference face | 
**Threshold** | Pointer to **float32** | Distance threshold for matching (lower &#x3D; stricter) | [optional] [default to 0.6]

## Methods

### NewCompareFacesRequest

`func NewCompareFacesRequest(probeImage string, referenceEmbedding []float32, ) *CompareFacesRequest`

NewCompareFacesRequest instantiates a new CompareFacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareFacesRequestWithDefaults

`func NewCompareFacesRequestWithDefaults() *CompareFacesRequest`

NewCompareFacesRequestWithDefaults instantiates a new CompareFacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProbeImage

`func (o *CompareFacesRequest) GetProbeImage() string`

GetProbeImage returns the ProbeImage field if non-nil, zero value otherwise.

### GetProbeImageOk

`func (o *CompareFacesRequest) GetProbeImageOk() (*string, bool)`

GetProbeImageOk returns a tuple with the ProbeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeImage

`func (o *CompareFacesRequest) SetProbeImage(v string)`

SetProbeImage sets ProbeImage field to given value.


### GetReferenceEmbedding

`func (o *CompareFacesRequest) GetReferenceEmbedding() []float32`

GetReferenceEmbedding returns the ReferenceEmbedding field if non-nil, zero value otherwise.

### GetReferenceEmbeddingOk

`func (o *CompareFacesRequest) GetReferenceEmbeddingOk() (*[]float32, bool)`

GetReferenceEmbeddingOk returns a tuple with the ReferenceEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceEmbedding

`func (o *CompareFacesRequest) SetReferenceEmbedding(v []float32)`

SetReferenceEmbedding sets ReferenceEmbedding field to given value.


### GetThreshold

`func (o *CompareFacesRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CompareFacesRequest) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CompareFacesRequest) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CompareFacesRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



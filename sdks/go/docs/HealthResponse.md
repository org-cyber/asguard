# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ModelsLoaded** | Pointer to **bool** | Whether ML models are loaded (relevant for face service) | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthResponse

`func NewHealthResponse(status string, ) *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetModelsLoaded

`func (o *HealthResponse) GetModelsLoaded() bool`

GetModelsLoaded returns the ModelsLoaded field if non-nil, zero value otherwise.

### GetModelsLoadedOk

`func (o *HealthResponse) GetModelsLoadedOk() (*bool, bool)`

GetModelsLoadedOk returns a tuple with the ModelsLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelsLoaded

`func (o *HealthResponse) SetModelsLoaded(v bool)`

SetModelsLoaded sets ModelsLoaded field to given value.

### HasModelsLoaded

`func (o *HealthResponse) HasModelsLoaded() bool`

HasModelsLoaded returns a boolean if a field has been set.

### GetVersion

`func (o *HealthResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FraudCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | Unique transaction identifier | 
**Amount** | **float32** | Transaction amount | 
**Currency** | **string** | ISO 4217 currency code (USD, EUR, etc.) | 
**UserId** | **string** | Customer identifier | 
**DeviceId** | Pointer to **string** | Device fingerprint | [optional] 
**IpAddress** | Pointer to **string** | Customer IP address | [optional] 
**Location** | Pointer to **string** | Geo location (city, country) | [optional] 
**Timestamp** | Pointer to **time.Time** | Transaction timestamp | [optional] 

## Methods

### NewFraudCheckRequest

`func NewFraudCheckRequest(transactionId string, amount float32, currency string, userId string, ) *FraudCheckRequest`

NewFraudCheckRequest instantiates a new FraudCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraudCheckRequestWithDefaults

`func NewFraudCheckRequestWithDefaults() *FraudCheckRequest`

NewFraudCheckRequestWithDefaults instantiates a new FraudCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *FraudCheckRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FraudCheckRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FraudCheckRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *FraudCheckRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FraudCheckRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FraudCheckRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *FraudCheckRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FraudCheckRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FraudCheckRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUserId

`func (o *FraudCheckRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FraudCheckRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FraudCheckRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDeviceId

`func (o *FraudCheckRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *FraudCheckRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *FraudCheckRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *FraudCheckRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIpAddress

`func (o *FraudCheckRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FraudCheckRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FraudCheckRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *FraudCheckRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLocation

`func (o *FraudCheckRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FraudCheckRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FraudCheckRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FraudCheckRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTimestamp

`func (o *FraudCheckRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FraudCheckRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FraudCheckRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FraudCheckRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



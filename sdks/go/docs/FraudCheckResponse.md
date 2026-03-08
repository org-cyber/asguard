# FraudCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**RiskScore** | **float32** |  | 
**RiskLevel** | **string** |  | 
**Reasons** | Pointer to **[]string** |  | [optional] 
**AiTriggered** | Pointer to **bool** |  | [optional] 
**AiConfidence** | Pointer to **float32** |  | [optional] 
**AiRecommendation** | Pointer to **string** |  | [optional] 
**AiFraudProbability** | Pointer to **float32** |  | [optional] 
**AiSummary** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** | Request processing time in milliseconds | [optional] 

## Methods

### NewFraudCheckResponse

`func NewFraudCheckResponse(transactionId string, riskScore float32, riskLevel string, ) *FraudCheckResponse`

NewFraudCheckResponse instantiates a new FraudCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraudCheckResponseWithDefaults

`func NewFraudCheckResponseWithDefaults() *FraudCheckResponse`

NewFraudCheckResponseWithDefaults instantiates a new FraudCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *FraudCheckResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FraudCheckResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FraudCheckResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetRiskScore

`func (o *FraudCheckResponse) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *FraudCheckResponse) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *FraudCheckResponse) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskLevel

`func (o *FraudCheckResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *FraudCheckResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *FraudCheckResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.


### GetReasons

`func (o *FraudCheckResponse) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *FraudCheckResponse) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *FraudCheckResponse) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *FraudCheckResponse) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetAiTriggered

`func (o *FraudCheckResponse) GetAiTriggered() bool`

GetAiTriggered returns the AiTriggered field if non-nil, zero value otherwise.

### GetAiTriggeredOk

`func (o *FraudCheckResponse) GetAiTriggeredOk() (*bool, bool)`

GetAiTriggeredOk returns a tuple with the AiTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTriggered

`func (o *FraudCheckResponse) SetAiTriggered(v bool)`

SetAiTriggered sets AiTriggered field to given value.

### HasAiTriggered

`func (o *FraudCheckResponse) HasAiTriggered() bool`

HasAiTriggered returns a boolean if a field has been set.

### GetAiConfidence

`func (o *FraudCheckResponse) GetAiConfidence() float32`

GetAiConfidence returns the AiConfidence field if non-nil, zero value otherwise.

### GetAiConfidenceOk

`func (o *FraudCheckResponse) GetAiConfidenceOk() (*float32, bool)`

GetAiConfidenceOk returns a tuple with the AiConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfidence

`func (o *FraudCheckResponse) SetAiConfidence(v float32)`

SetAiConfidence sets AiConfidence field to given value.

### HasAiConfidence

`func (o *FraudCheckResponse) HasAiConfidence() bool`

HasAiConfidence returns a boolean if a field has been set.

### GetAiRecommendation

`func (o *FraudCheckResponse) GetAiRecommendation() string`

GetAiRecommendation returns the AiRecommendation field if non-nil, zero value otherwise.

### GetAiRecommendationOk

`func (o *FraudCheckResponse) GetAiRecommendationOk() (*string, bool)`

GetAiRecommendationOk returns a tuple with the AiRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRecommendation

`func (o *FraudCheckResponse) SetAiRecommendation(v string)`

SetAiRecommendation sets AiRecommendation field to given value.

### HasAiRecommendation

`func (o *FraudCheckResponse) HasAiRecommendation() bool`

HasAiRecommendation returns a boolean if a field has been set.

### GetAiFraudProbability

`func (o *FraudCheckResponse) GetAiFraudProbability() float32`

GetAiFraudProbability returns the AiFraudProbability field if non-nil, zero value otherwise.

### GetAiFraudProbabilityOk

`func (o *FraudCheckResponse) GetAiFraudProbabilityOk() (*float32, bool)`

GetAiFraudProbabilityOk returns a tuple with the AiFraudProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiFraudProbability

`func (o *FraudCheckResponse) SetAiFraudProbability(v float32)`

SetAiFraudProbability sets AiFraudProbability field to given value.

### HasAiFraudProbability

`func (o *FraudCheckResponse) HasAiFraudProbability() bool`

HasAiFraudProbability returns a boolean if a field has been set.

### GetAiSummary

`func (o *FraudCheckResponse) GetAiSummary() string`

GetAiSummary returns the AiSummary field if non-nil, zero value otherwise.

### GetAiSummaryOk

`func (o *FraudCheckResponse) GetAiSummaryOk() (*string, bool)`

GetAiSummaryOk returns a tuple with the AiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSummary

`func (o *FraudCheckResponse) SetAiSummary(v string)`

SetAiSummary sets AiSummary field to given value.

### HasAiSummary

`func (o *FraudCheckResponse) HasAiSummary() bool`

HasAiSummary returns a boolean if a field has been set.

### GetMessage

`func (o *FraudCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FraudCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FraudCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FraudCheckResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *FraudCheckResponse) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *FraudCheckResponse) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *FraudCheckResponse) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *FraudCheckResponse) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExternalGateDtoReviewSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredReview** | **bool** |  | 
**AllowedReviewers** | [**[]ExternalGateDtoReviewSettingsAllowedReviewersInner**](ExternalGateDtoReviewSettingsAllowedReviewersInner.md) |  | 

## Methods

### NewExternalGateDtoReviewSettings

`func NewExternalGateDtoReviewSettings(requiredReview bool, allowedReviewers []ExternalGateDtoReviewSettingsAllowedReviewersInner, ) *ExternalGateDtoReviewSettings`

NewExternalGateDtoReviewSettings instantiates a new ExternalGateDtoReviewSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGateDtoReviewSettingsWithDefaults

`func NewExternalGateDtoReviewSettingsWithDefaults() *ExternalGateDtoReviewSettings`

NewExternalGateDtoReviewSettingsWithDefaults instantiates a new ExternalGateDtoReviewSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredReview

`func (o *ExternalGateDtoReviewSettings) GetRequiredReview() bool`

GetRequiredReview returns the RequiredReview field if non-nil, zero value otherwise.

### GetRequiredReviewOk

`func (o *ExternalGateDtoReviewSettings) GetRequiredReviewOk() (*bool, bool)`

GetRequiredReviewOk returns a tuple with the RequiredReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReview

`func (o *ExternalGateDtoReviewSettings) SetRequiredReview(v bool)`

SetRequiredReview sets RequiredReview field to given value.


### GetAllowedReviewers

`func (o *ExternalGateDtoReviewSettings) GetAllowedReviewers() []ExternalGateDtoReviewSettingsAllowedReviewersInner`

GetAllowedReviewers returns the AllowedReviewers field if non-nil, zero value otherwise.

### GetAllowedReviewersOk

`func (o *ExternalGateDtoReviewSettings) GetAllowedReviewersOk() (*[]ExternalGateDtoReviewSettingsAllowedReviewersInner, bool)`

GetAllowedReviewersOk returns a tuple with the AllowedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedReviewers

`func (o *ExternalGateDtoReviewSettings) SetAllowedReviewers(v []ExternalGateDtoReviewSettingsAllowedReviewersInner)`

SetAllowedReviewers sets AllowedReviewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



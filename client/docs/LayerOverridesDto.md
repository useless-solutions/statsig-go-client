# LayerOverridesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionalOverrides** | [**[]LayerOverridesDtoConditionalOverridesInner**](LayerOverridesDtoConditionalOverridesInner.md) |  | 
**IdOverrides** | [**[]LayerOverridesDtoIdOverridesInner**](LayerOverridesDtoIdOverridesInner.md) |  | 

## Methods

### NewLayerOverridesDto

`func NewLayerOverridesDto(conditionalOverrides []LayerOverridesDtoConditionalOverridesInner, idOverrides []LayerOverridesDtoIdOverridesInner, ) *LayerOverridesDto`

NewLayerOverridesDto instantiates a new LayerOverridesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerOverridesDtoWithDefaults

`func NewLayerOverridesDtoWithDefaults() *LayerOverridesDto`

NewLayerOverridesDtoWithDefaults instantiates a new LayerOverridesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionalOverrides

`func (o *LayerOverridesDto) GetConditionalOverrides() []LayerOverridesDtoConditionalOverridesInner`

GetConditionalOverrides returns the ConditionalOverrides field if non-nil, zero value otherwise.

### GetConditionalOverridesOk

`func (o *LayerOverridesDto) GetConditionalOverridesOk() (*[]LayerOverridesDtoConditionalOverridesInner, bool)`

GetConditionalOverridesOk returns a tuple with the ConditionalOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalOverrides

`func (o *LayerOverridesDto) SetConditionalOverrides(v []LayerOverridesDtoConditionalOverridesInner)`

SetConditionalOverrides sets ConditionalOverrides field to given value.


### GetIdOverrides

`func (o *LayerOverridesDto) GetIdOverrides() []LayerOverridesDtoIdOverridesInner`

GetIdOverrides returns the IdOverrides field if non-nil, zero value otherwise.

### GetIdOverridesOk

`func (o *LayerOverridesDto) GetIdOverridesOk() (*[]LayerOverridesDtoIdOverridesInner, bool)`

GetIdOverridesOk returns a tuple with the IdOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOverrides

`func (o *LayerOverridesDto) SetIdOverrides(v []LayerOverridesDtoIdOverridesInner)`

SetIdOverrides sets IdOverrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AutotuneExperimentDtoVariantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Variant name | 
**Json** | **interface{}** | Variant JSON data | 
**Id** | **string** | The name that was originally given to the autotune on creation but formatted as an ID (\&quot;A Autotune\&quot; -&gt; \&quot;a_autotune\&quot;). | 

## Methods

### NewAutotuneExperimentDtoVariantsInner

`func NewAutotuneExperimentDtoVariantsInner(name string, json interface{}, id string, ) *AutotuneExperimentDtoVariantsInner`

NewAutotuneExperimentDtoVariantsInner instantiates a new AutotuneExperimentDtoVariantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotuneExperimentDtoVariantsInnerWithDefaults

`func NewAutotuneExperimentDtoVariantsInnerWithDefaults() *AutotuneExperimentDtoVariantsInner`

NewAutotuneExperimentDtoVariantsInnerWithDefaults instantiates a new AutotuneExperimentDtoVariantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutotuneExperimentDtoVariantsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotuneExperimentDtoVariantsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotuneExperimentDtoVariantsInner) SetName(v string)`

SetName sets Name field to given value.


### GetJson

`func (o *AutotuneExperimentDtoVariantsInner) GetJson() interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AutotuneExperimentDtoVariantsInner) GetJsonOk() (*interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AutotuneExperimentDtoVariantsInner) SetJson(v interface{})`

SetJson sets Json field to given value.


### SetJsonNil

`func (o *AutotuneExperimentDtoVariantsInner) SetJsonNil(b bool)`

 SetJsonNil sets the value for Json to be an explicit nil

### UnsetJson
`func (o *AutotuneExperimentDtoVariantsInner) UnsetJson()`

UnsetJson ensures that no value is present for Json, not even an explicit nil
### GetId

`func (o *AutotuneExperimentDtoVariantsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotuneExperimentDtoVariantsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotuneExperimentDtoVariantsInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



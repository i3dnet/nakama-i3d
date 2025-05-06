# ApplicationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Application type ID | [readonly] 
**TypeDescription** | **string** | Application type description | [readonly] 

## Methods

### NewApplicationType

`func NewApplicationType(typeId int32, typeDescription string, ) *ApplicationType`

NewApplicationType instantiates a new ApplicationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTypeWithDefaults

`func NewApplicationTypeWithDefaults() *ApplicationType`

NewApplicationTypeWithDefaults instantiates a new ApplicationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *ApplicationType) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ApplicationType) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ApplicationType) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetTypeDescription

`func (o *ApplicationType) GetTypeDescription() string`

GetTypeDescription returns the TypeDescription field if non-nil, zero value otherwise.

### GetTypeDescriptionOk

`func (o *ApplicationType) GetTypeDescriptionOk() (*string, bool)`

GetTypeDescriptionOk returns a tuple with the TypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDescription

`func (o *ApplicationType) SetTypeDescription(v string)`

SetTypeDescription sets TypeDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



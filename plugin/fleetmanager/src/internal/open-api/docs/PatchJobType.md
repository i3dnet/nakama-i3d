# PatchJobType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Patch job type ID | [readonly] 
**TypeDescription** | **string** | Patch job type description | [readonly] 

## Methods

### NewPatchJobType

`func NewPatchJobType(typeId int32, typeDescription string, ) *PatchJobType`

NewPatchJobType instantiates a new PatchJobType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobTypeWithDefaults

`func NewPatchJobTypeWithDefaults() *PatchJobType`

NewPatchJobTypeWithDefaults instantiates a new PatchJobType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *PatchJobType) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PatchJobType) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PatchJobType) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetTypeDescription

`func (o *PatchJobType) GetTypeDescription() string`

GetTypeDescription returns the TypeDescription field if non-nil, zero value otherwise.

### GetTypeDescriptionOk

`func (o *PatchJobType) GetTypeDescriptionOk() (*string, bool)`

GetTypeDescriptionOk returns a tuple with the TypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDescription

`func (o *PatchJobType) SetTypeDescription(v string)`

SetTypeDescription sets TypeDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



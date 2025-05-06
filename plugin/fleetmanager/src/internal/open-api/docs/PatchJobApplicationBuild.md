# PatchJobApplicationBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a patch job application build | [readonly] 
**OldApplicationBuildId** | **string** | ID of an old application build | 
**NewApplicationBuildId** | **string** | ID of a new application build | 

## Methods

### NewPatchJobApplicationBuild

`func NewPatchJobApplicationBuild(id string, oldApplicationBuildId string, newApplicationBuildId string, ) *PatchJobApplicationBuild`

NewPatchJobApplicationBuild instantiates a new PatchJobApplicationBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobApplicationBuildWithDefaults

`func NewPatchJobApplicationBuildWithDefaults() *PatchJobApplicationBuild`

NewPatchJobApplicationBuildWithDefaults instantiates a new PatchJobApplicationBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJobApplicationBuild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJobApplicationBuild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJobApplicationBuild) SetId(v string)`

SetId sets Id field to given value.


### GetOldApplicationBuildId

`func (o *PatchJobApplicationBuild) GetOldApplicationBuildId() string`

GetOldApplicationBuildId returns the OldApplicationBuildId field if non-nil, zero value otherwise.

### GetOldApplicationBuildIdOk

`func (o *PatchJobApplicationBuild) GetOldApplicationBuildIdOk() (*string, bool)`

GetOldApplicationBuildIdOk returns a tuple with the OldApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldApplicationBuildId

`func (o *PatchJobApplicationBuild) SetOldApplicationBuildId(v string)`

SetOldApplicationBuildId sets OldApplicationBuildId field to given value.


### GetNewApplicationBuildId

`func (o *PatchJobApplicationBuild) GetNewApplicationBuildId() string`

GetNewApplicationBuildId returns the NewApplicationBuildId field if non-nil, zero value otherwise.

### GetNewApplicationBuildIdOk

`func (o *PatchJobApplicationBuild) GetNewApplicationBuildIdOk() (*string, bool)`

GetNewApplicationBuildIdOk returns a tuple with the NewApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplicationBuildId

`func (o *PatchJobApplicationBuild) SetNewApplicationBuildId(v string)`

SetNewApplicationBuildId sets NewApplicationBuildId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



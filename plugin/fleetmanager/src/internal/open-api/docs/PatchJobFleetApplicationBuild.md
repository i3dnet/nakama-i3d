# PatchJobFleetApplicationBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the fleet | [readonly] 
**Name** | **string** | Name of the fleet | [readonly] 
**ApplicationBuild** | [**[]PatchJobFleetApplicationBuildDetail**](PatchJobFleetApplicationBuildDetail.md) | Array of the application build belonging to the fleet | [readonly] 

## Methods

### NewPatchJobFleetApplicationBuild

`func NewPatchJobFleetApplicationBuild(id string, name string, applicationBuild []PatchJobFleetApplicationBuildDetail, ) *PatchJobFleetApplicationBuild`

NewPatchJobFleetApplicationBuild instantiates a new PatchJobFleetApplicationBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobFleetApplicationBuildWithDefaults

`func NewPatchJobFleetApplicationBuildWithDefaults() *PatchJobFleetApplicationBuild`

NewPatchJobFleetApplicationBuildWithDefaults instantiates a new PatchJobFleetApplicationBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJobFleetApplicationBuild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJobFleetApplicationBuild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJobFleetApplicationBuild) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchJobFleetApplicationBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchJobFleetApplicationBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchJobFleetApplicationBuild) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationBuild

`func (o *PatchJobFleetApplicationBuild) GetApplicationBuild() []PatchJobFleetApplicationBuildDetail`

GetApplicationBuild returns the ApplicationBuild field if non-nil, zero value otherwise.

### GetApplicationBuildOk

`func (o *PatchJobFleetApplicationBuild) GetApplicationBuildOk() (*[]PatchJobFleetApplicationBuildDetail, bool)`

GetApplicationBuildOk returns a tuple with the ApplicationBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuild

`func (o *PatchJobFleetApplicationBuild) SetApplicationBuild(v []PatchJobFleetApplicationBuildDetail)`

SetApplicationBuild sets ApplicationBuild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



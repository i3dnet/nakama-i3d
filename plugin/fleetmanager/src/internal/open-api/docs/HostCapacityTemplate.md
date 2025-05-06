# HostCapacityTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Host Capacity Template | [readonly] 
**FleetIds** | **[]string** | IDs of the Fleets that use this Host Capacity Template | [readonly] 
**ApplicationBuildIds** | **[]string** | IDs of the Application Builds that use this Host Capacity Template | [readonly] 
**Name** | **string** | The public name of your Host Capacity Template | 
**InUse** | **int32** | 1 if the Host Capacity Template is in use by a Fleet or Application Build, 0 if it is not. | 
**CreatedAt** | **int32** | The Unix timestamp at which this element is created | [readonly] 

## Methods

### NewHostCapacityTemplate

`func NewHostCapacityTemplate(id string, fleetIds []string, applicationBuildIds []string, name string, inUse int32, createdAt int32, ) *HostCapacityTemplate`

NewHostCapacityTemplate instantiates a new HostCapacityTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCapacityTemplateWithDefaults

`func NewHostCapacityTemplateWithDefaults() *HostCapacityTemplate`

NewHostCapacityTemplateWithDefaults instantiates a new HostCapacityTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostCapacityTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostCapacityTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostCapacityTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetFleetIds

`func (o *HostCapacityTemplate) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *HostCapacityTemplate) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *HostCapacityTemplate) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetApplicationBuildIds

`func (o *HostCapacityTemplate) GetApplicationBuildIds() []string`

GetApplicationBuildIds returns the ApplicationBuildIds field if non-nil, zero value otherwise.

### GetApplicationBuildIdsOk

`func (o *HostCapacityTemplate) GetApplicationBuildIdsOk() (*[]string, bool)`

GetApplicationBuildIdsOk returns a tuple with the ApplicationBuildIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildIds

`func (o *HostCapacityTemplate) SetApplicationBuildIds(v []string)`

SetApplicationBuildIds sets ApplicationBuildIds field to given value.


### GetName

`func (o *HostCapacityTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostCapacityTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostCapacityTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetInUse

`func (o *HostCapacityTemplate) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *HostCapacityTemplate) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *HostCapacityTemplate) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetCreatedAt

`func (o *HostCapacityTemplate) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostCapacityTemplate) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostCapacityTemplate) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


